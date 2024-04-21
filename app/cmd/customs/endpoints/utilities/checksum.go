package utilities

import (
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/pkg/errors"
)

func MD5(file io.Reader) (string, error) {
	hash := md5.New()
	var _, err = io.Copy(hash, file)
	if err != nil {
		return "", errors.Wrap(err, "Error calculating MD5 checksum")
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
