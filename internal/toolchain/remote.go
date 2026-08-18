package toolchain

import (
	"bytes"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"crypto/ed25519"

	"golang.org/x/crypto/ssh"
)

// Remote is a Kit that drives another machine's project checkout over
// SSH: file operations become cat, builds run gd there. This is how a
// device without exec (the iPhone) gets a full build loop today — the
// remote builds and serves the install, and OnURL hands the resulting
// sidestore:// link back to the caller to open on-device.
type Remote struct {
	Config RemoteConfig
	// OnURL, if set, receives the first sidestore:// URL seen in GD output.
	OnURL func(url string)

	stateDir string
	out      io.Writer
	signer   ssh.Signer
	client   *ssh.Client
}

type RemoteConfig struct {
	User string `json:"user"`
	Host string `json:"host"`
	Port string `json:"port"`
	Dir  string `json:"dir"`
}

func remoteConfigFile(stateDir string) string { return filepath.Join(stateDir, "remote") }

// NewRemote parses "user@host[:port] /path/to/project", ensures a key
// pair exists under stateDir, and saves the configuration. Connection
// is lazy. The project path must not contain spaces.
func NewRemote(spec string, stateDir string, out io.Writer) (*Remote, error) {
	target, dir, ok := strings.Cut(strings.TrimSpace(spec), " ")
	if !ok {
		return nil, fmt.Errorf("usage: user@host[:port] /path/to/project")
	}
	user, addr, ok := strings.Cut(target, "@")
	if !ok {
		return nil, fmt.Errorf("usage: user@host[:port] /path/to/project")
	}
	host, port, ok := strings.Cut(addr, ":")
	if !ok {
		port = "22"
	}
	cfg := RemoteConfig{User: user, Host: host, Port: port, Dir: strings.TrimSpace(dir)}
	data, _ := json.Marshal(cfg)
	if err := os.MkdirAll(stateDir, 0o700); err != nil {
		return nil, err
	}
	if err := os.WriteFile(remoteConfigFile(stateDir), data, 0o600); err != nil {
		return nil, err
	}
	return openRemote(cfg, stateDir, out)
}

// LoadRemote restores a previously configured remote, if any.
func LoadRemote(stateDir string, out io.Writer) (*Remote, bool) {
	data, err := os.ReadFile(remoteConfigFile(stateDir))
	if err != nil {
		return nil, false
	}
	var cfg RemoteConfig
	if json.Unmarshal(data, &cfg) != nil {
		return nil, false
	}
	r, err := openRemote(cfg, stateDir, out)
	if err != nil {
		fmt.Fprintf(out, "remote: %v\n", err)
		return nil, false
	}
	return r, true
}

// ForgetRemote removes the saved remote configuration.
func ForgetRemote(stateDir string) { os.Remove(remoteConfigFile(stateDir)) }

func openRemote(cfg RemoteConfig, stateDir string, out io.Writer) (*Remote, error) {
	r := &Remote{Config: cfg, stateDir: stateDir, out: out}
	signer, err := r.ensureKey()
	if err != nil {
		return nil, err
	}
	r.signer = signer
	return r, nil
}

// PublicKey is the authorized_keys line to install on the remote.
func (r *Remote) PublicKey() string {
	return strings.TrimSpace(string(ssh.MarshalAuthorizedKey(r.signer.PublicKey()))) + " gd-harness"
}

func (r *Remote) ensureKey() (ssh.Signer, error) {
	keyFile := filepath.Join(r.stateDir, "id_ed25519")
	if data, err := os.ReadFile(keyFile); err == nil {
		return ssh.ParsePrivateKey(data)
	}
	_, private, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}
	block, err := ssh.MarshalPrivateKey(private, "gd-harness")
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(r.stateDir, 0o700); err != nil {
		return nil, err
	}
	if err := os.WriteFile(keyFile, pem.EncodeToMemory(block), 0o600); err != nil {
		return nil, err
	}
	return ssh.NewSignerFromKey(private)
}

// hostKey trusts the first key a host presents and pins it after.
func (r *Remote) hostKey(hostname string, _ net.Addr, key ssh.PublicKey) error {
	line := bytes.TrimSpace(ssh.MarshalAuthorizedKey(key))
	file := filepath.Join(r.stateDir, "hostkey-"+r.Config.Host)
	saved, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(r.out, "remote: trusting %s on first connection (%s)\n", hostname, ssh.FingerprintSHA256(key))
		return os.WriteFile(file, append(line, '\n'), 0o600)
	}
	if !bytes.Equal(bytes.TrimSpace(saved), line) {
		return fmt.Errorf("host key for %s CHANGED (%s) — refusing; delete %s if this is expected", hostname, ssh.FingerprintSHA256(key), file)
	}
	return nil
}

func (r *Remote) connect() (*ssh.Client, error) {
	if r.client != nil {
		return r.client, nil
	}
	client, err := ssh.Dial("tcp", net.JoinHostPort(r.Config.Host, r.Config.Port), &ssh.ClientConfig{
		User:            r.Config.User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(r.signer)},
		HostKeyCallback: r.hostKey,
		Timeout:         10 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("ssh %s@%s:%s: %w (is the harness key in authorized_keys?)", r.Config.User, r.Config.Host, r.Config.Port, err)
	}
	r.client = client
	return client, nil
}

// session runs one command, retrying once through a fresh connection
// if a cached one has gone stale.
func (r *Remote) session(run func(*ssh.Session) error) error {
	for attempt := range 2 {
		client, err := r.connect()
		if err != nil {
			return err
		}
		session, err := client.NewSession()
		if err != nil {
			r.client = nil
			client.Close()
			if attempt == 0 {
				continue
			}
			return err
		}
		defer session.Close()
		return run(session)
	}
	return nil
}

func (r *Remote) Describe() string {
	return r.Config.User + "@" + r.Config.Host + ":" + r.Config.Dir
}

func (r *Remote) resolve(p string) (string, error) {
	if p == "" {
		p = "."
	}
	joined := path.Clean(path.Join(r.Config.Dir, p))
	if joined != r.Config.Dir && !strings.HasPrefix(joined, r.Config.Dir+"/") {
		return "", fmt.Errorf("path escapes the project: %s", p)
	}
	return joined, nil
}

func quote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}

func (r *Remote) Shell(command string, timeout time.Duration) (string, error) {
	if timeout <= 0 {
		timeout = 5 * time.Minute
	}
	var buf bytes.Buffer
	err := r.session(func(s *ssh.Session) error {
		s.Stdout = &buf
		s.Stderr = &buf
		timer := time.AfterFunc(timeout, func() { s.Close() })
		defer timer.Stop()
		return s.Run("cd " + quote(r.Config.Dir) + " && (" + command + ")")
	})
	const cap = 100 << 10
	out := buf.Bytes()
	if len(out) > cap {
		out = append([]byte("(output truncated)\n"), out[len(out)-cap:]...)
	}
	return string(out), err
}

func (r *Remote) GD(verb string, goos string) error {
	command := "cd " + quote(r.Config.Dir) + " && "
	if goos != "" {
		command += "GOOS=" + goos + " "
	}
	command += "gd " + verb
	fmt.Fprintf(r.out, "$ %s (on %s)\n", command, r.Config.Host)
	return r.session(func(s *ssh.Session) error {
		w := io.Writer(r.out)
		if r.OnURL != nil {
			w = &urlScanner{out: r.out, onURL: r.OnURL}
		}
		s.Stdout = w
		s.Stderr = w
		return s.Run(command)
	})
}

func (r *Remote) ReadFile(p string) ([]byte, error) {
	resolved, err := r.resolve(p)
	if err != nil {
		return nil, err
	}
	var buf, errs bytes.Buffer
	err = r.session(func(s *ssh.Session) error {
		s.Stdout = &buf
		s.Stderr = &errs
		return s.Run("cat " + quote(resolved))
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, strings.TrimSpace(errs.String()))
	}
	return buf.Bytes(), nil
}

func (r *Remote) WriteFile(p string, data []byte) error {
	resolved, err := r.resolve(p)
	if err != nil {
		return err
	}
	var errs bytes.Buffer
	err = r.session(func(s *ssh.Session) error {
		s.Stdin = bytes.NewReader(data)
		s.Stderr = &errs
		return s.Run("mkdir -p " + quote(path.Dir(resolved)) + " && cat > " + quote(resolved))
	})
	if err != nil {
		return fmt.Errorf("%w: %s", err, strings.TrimSpace(errs.String()))
	}
	return nil
}

func (r *Remote) List(p string) (string, error) {
	resolved, err := r.resolve(p)
	if err != nil {
		return "", err
	}
	var buf, errs bytes.Buffer
	err = r.session(func(s *ssh.Session) error {
		s.Stdout = &buf
		s.Stderr = &errs
		return s.Run("ls -1p " + quote(resolved))
	})
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, strings.TrimSpace(errs.String()))
	}
	return buf.String(), nil
}

// urlScanner tees output while watching for the first sidestore:// URL.
type urlScanner struct {
	out   io.Writer
	onURL func(string)
	line  bytes.Buffer
	done  bool
}

func (u *urlScanner) Write(p []byte) (int, error) {
	if !u.done {
		for _, b := range p {
			if b == '\n' {
				u.scan(u.line.String())
				u.line.Reset()
			} else {
				u.line.WriteByte(b)
			}
		}
	}
	return u.out.Write(p)
}

func (u *urlScanner) scan(line string) {
	const scheme = "sidestore://"
	i := strings.Index(line, scheme)
	if i < 0 {
		return
	}
	url := line[i:]
	if end := strings.IndexAny(url, " \t)"); end >= 0 {
		url = url[:end]
	}
	u.done = true
	u.onURL(url)
}
