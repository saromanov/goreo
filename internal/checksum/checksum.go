package checksum

import (
	"crypto/sha256"
	"io"
	"os"
)

// Run at this moment provides calculation based on sha 256 check sum
func Run(algorithm, path string) ([]byte, error) {
	return calcSha256(path)
}

func calcSha256(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
