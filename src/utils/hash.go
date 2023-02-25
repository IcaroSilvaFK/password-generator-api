package methods

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(s string) (string, error) {

	hashMaker := md5.New()

	defer hashMaker.Reset()

	_, err := hashMaker.Write([]byte(s))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hashMaker.Sum(nil)), nil
}
