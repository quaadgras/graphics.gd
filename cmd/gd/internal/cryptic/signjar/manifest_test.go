package signjar

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()
	t.Run("Full", func(t *testing.T) {
		const manifest = `Manifest-Version: 1.0
Built-By: nobody
Long-Header-Line: 0123456789abcdef0123456789abcdef0123456789abcdef
 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef

Name: foo
Ham: spam
Eggs: bacon

`
		main := http.Header{
			"Manifest-Version": []string{"1.0"},
			"Built-By":         []string{"nobody"},
			"Long-Header-Line": []string{
				"0123456789abcdef0123456789abcdef" +
					"0123456789abcdef0123456789abcdef0123456789abcdef" +
					"0123456789abcdef0123456789abcdef0123456789abcdef" +
					"0123456789abcdef0123456789abcdef0123456789abcdef"},
		}
		file := http.Header{
			"Name": []string{"foo"},
			"Ham":  []string{"spam"},
			"Eggs": []string{"bacon"},
		}
		expected := &FilesMap{
			Main:  main,
			Files: map[string]http.Header{"foo": file},
			Order: []string{"foo"},
		}
		parsed, err := ParseManifest([]byte(manifest))
		require.NoError(t, err)
		assert.Equal(t, expected, parsed)
		crlfManifest := []byte(strings.ReplaceAll(manifest, "\n", "\r\n"))
		parsed, err = ParseManifest(crlfManifest)
		require.NoError(t, err)
		assert.Equal(t, expected, parsed)
	})
	t.Run("TruncatedMain", func(t *testing.T) {
		const manifest = "Manifest-Version: 1.0\n"
		expected := &FilesMap{
			Main: http.Header{
				"Manifest-Version": []string{"1.0"},
			},
			Order: []string{},
			Files: map[string]http.Header{},
		}
		_, err := ParseManifest([]byte(manifest))
		require.ErrorIs(t, err, ErrManifestLineEndings)
		parsed, malformed, err := parseManifest([]byte(manifest))
		require.NoError(t, err)
		assert.True(t, malformed)
		assert.Equal(t, expected, parsed)
	})
	t.Run("TruncatedFile", func(t *testing.T) {
		const manifest = "Manifest-Version: 1.0\n\nName: foo\n"
		file := http.Header{
			"Name": []string{"foo"},
		}
		expected := &FilesMap{
			Main: http.Header{
				"Manifest-Version": []string{"1.0"},
			},
			Order: []string{"foo"},
			Files: map[string]http.Header{"foo": file},
		}
		_, err := ParseManifest([]byte(manifest))
		require.ErrorIs(t, err, ErrManifestLineEndings)
		parsed, malformed, err := parseManifest([]byte(manifest))
		require.NoError(t, err)
		assert.True(t, malformed)
		assert.Equal(t, expected, parsed)
	})
	t.Run("TrailingWhitespace", func(t *testing.T) {
		const manifest = "Manifest-Version: 1.0\n\nName: foo\n\n\n"
		file := http.Header{
			"Name": []string{"foo"},
		}
		expected := &FilesMap{
			Main: http.Header{
				"Manifest-Version": []string{"1.0"},
			},
			Order: []string{"foo"},
			Files: map[string]http.Header{"foo": file},
		}
		_, err := ParseManifest([]byte(manifest))
		require.ErrorIs(t, err, ErrManifestLineEndings)
		parsed, malformed, err := parseManifest([]byte(manifest))
		require.NoError(t, err)
		assert.True(t, malformed)
		assert.Equal(t, expected, parsed)
	})
	t.Run("InvalidNoName", func(t *testing.T) {
		const manifest = "Manifest-Version: 1.0\n\nFoo: bar\n\n"
		_, err := ParseManifest([]byte(manifest))
		require.Error(t, err)
	})
}

func TestDump(t *testing.T) {
	manifest := &FilesMap{
		Main: http.Header{
			"Manifest-Version": []string{"1.0"},
			"D":                []string{"D"},
			"C":                []string{"C"},
			"B":                []string{"B"},
			"A":                []string{"A"},
			"Long-Header":      []string{strings.Repeat("0123456789abcdef", 10)},
		},
		Files: map[string]http.Header{
			"foo": {"Name": []string{"foo"}, "Foo": []string{"bar"}},
			"bar": {"Name": []string{"bar"}, "Foo": []string{"bar"}},
		},
		Order: []string{"bar", "foo"},
	}
	result := string(manifest.Dump())
	expected := `Manifest-Version: 1.0
A: A
B: B
C: C
D: D
Long-Header: 0123456789abcdef0123456789abcdef0123456789abcdef012345678
 9abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcd
 ef0123456789abcdef0123456789abcdef

Name: bar
Foo: bar

Name: foo
Foo: bar

`
	expected = strings.ReplaceAll(expected, "\n", "\r\n")
	assert.Equal(t, expected, result)
}
