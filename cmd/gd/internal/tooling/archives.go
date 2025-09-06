package tooling

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ulikunitz/xz"
)

// detectTopDirForTar scans a tar archive to detect a single top-level directory
func detectTopDirForTar(src, archiveType string) (string, error) {
	file, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("failed to open source file: %w", err)
	}
	defer file.Close()

	var tr *tar.Reader
	switch archiveType {
	case "tar.gz":
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gz.Close()
		tr = tar.NewReader(gz)
	case "tar.xz":
		xzReader, err := xz.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("failed to create xz reader: %w", err)
		}
		tr = tar.NewReader(xzReader)
	default:
		return "", fmt.Errorf("unsupported tar compression: %s", archiveType)
	}

	var topDir string
	first := true
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read tar archive: %w", err)
		}

		name := header.Name
		if first {
			firstSlash := strings.Index(name, string(filepath.Separator))
			if firstSlash > 0 {
				topDir = name[:firstSlash]
			}
			first = false
		} else if topDir != "" {
			if !strings.HasPrefix(name, topDir+string(filepath.Separator)) && name != topDir {
				topDir = ""
			}
		}
	}

	return topDir, nil
}

// detectTopDirForZip scans a zip archive to detect a single top-level directory
func detectTopDirForZip(src string) (string, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return "", fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	var topDir string
	first := true
	for _, f := range r.File {
		name := f.Name
		if first {
			firstSlash := strings.Index(name, string(filepath.Separator))
			if firstSlash > 0 {
				topDir = name[:firstSlash]
			}
			first = false
		} else if topDir != "" {
			if !strings.HasPrefix(name, topDir+string(filepath.Separator)) && name != topDir {
				topDir = ""
			}
		}
	}
	return topDir, nil
}

// extractZip extracts a .zip file to the destination directory
func extractZip(src, dest, topDir string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	// Ensure destination directory exists
	if err := os.MkdirAll(dest, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	for _, f := range r.File {
		// Adjust path if stripping top-level directory
		targetName := f.Name
		if topDir != "" && targetName != topDir {
			targetName = strings.TrimPrefix(targetName, topDir+string(filepath.Separator))
		}
		if targetName == "" {
			continue // Skip top-level directory itself
		}

		target := filepath.Join(dest, targetName)

		// Check for ZipSlip vulnerability
		if !filepath.HasPrefix(target, filepath.Clean(dest)) {
			return fmt.Errorf("invalid file path: %s", target)
		}

		if f.FileInfo().IsDir() {
			// Create directory
			if err := os.MkdirAll(target, f.Mode()); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", target, err)
			}
			continue
		}

		// Create parent directory
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return fmt.Errorf("failed to create parent directory for %s: %w", target, err)
		}

		// Create and write to file
		dstFile, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", target, err)
		}

		srcFile, err := f.Open()
		if err != nil {
			dstFile.Close()
			return fmt.Errorf("failed to open zip entry %s: %w", f.Name, err)
		}

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			srcFile.Close()
			dstFile.Close()
			return fmt.Errorf("failed to write to file %s: %w", target, err)
		}

		srcFile.Close()
		dstFile.Close()
	}

	return nil
}

// extractTar extracts a tar-based archive (.tar.gz or .tar.xz) to the destination directory
func extractTar(src, dest, archiveType, topDir string) error {
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer file.Close()

	var tr *tar.Reader
	switch archiveType {
	case "tar.gz":
		gz, err := gzip.NewReader(file)
		if err != nil {
			return fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gz.Close()
		tr = tar.NewReader(gz)
	case "tar.xz":
		xzReader, err := xz.NewReader(file)
		if err != nil {
			return fmt.Errorf("failed to create xz reader: %w", err)
		}
		tr = tar.NewReader(xzReader)
	default:
		return fmt.Errorf("unsupported tar compression: %s", archiveType)
	}

	// Ensure destination directory exists
	if err := os.MkdirAll(dest, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar archive: %w", err)
		}

		// Adjust path if stripping top-level directory
		targetName := header.Name
		if topDir != "" && targetName != topDir {
			targetName = strings.TrimPrefix(targetName, topDir+string(filepath.Separator))
		}
		if targetName == "" {
			continue // Skip top-level directory itself
		}

		target := filepath.Join(dest, targetName)

		// Check for path traversal vulnerability
		if !filepath.HasPrefix(target, filepath.Clean(dest)) {
			return fmt.Errorf("invalid file path: %s", target)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			// Create directory
			if err := os.MkdirAll(target, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", target, err)
			}

		case tar.TypeReg:
			// Create parent directory
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory for %s: %w", target, err)
			}

			// Create and write to file
			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", target, err)
			}

			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return fmt.Errorf("failed to write to file %s: %w", target, err)
			}
			f.Close()

		default:
			return fmt.Errorf("unsupported file type %v in tar archive: %s", header.Typeflag, header.Name)
		}
	}

	return nil
}

// ExtractArchive is the main entry point for extracting archives
func ExtractArchive(src, dest, archiveType string, stripTopDir bool) error {
	var topDir string
	if stripTopDir {
		// Detect top-level directory
		switch archiveType {
		case "zip":
			var err error
			topDir, err = detectTopDirForZip(src)
			if err != nil {
				return err
			}
		case "tar.gz", "tar.xz":
			var err error
			topDir, err = detectTopDirForTar(src, archiveType)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported archive type: %s", archiveType)
		}
	}

	// Extract based on archive type
	switch archiveType {
	case "zip":
		return extractZip(src, dest, topDir)
	case "tar.gz", "tar.xz":
		return extractTar(src, dest, archiveType, topDir)
	default:
		return fmt.Errorf("unsupported archive type: %s", archiveType)
	}
}
