package utils

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil
	}
	defer file.Close()
	d, _ := io.ReadAll(file)
	return string(d), nil
}

func WriteFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(content)
	return nil
}

func FormatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
