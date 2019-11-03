package checksum

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// Run at this moment provides calculation based on sha 256 check sum
func Run(algorithm, path string) (string, error) {
	if algorithm == "" {
		return "", nil
	}
	return calcSha256(path)
}

func calcSha256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
