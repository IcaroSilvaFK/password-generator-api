package services

import (
	"log"

	"github.com/IcaroSilvaFK/password-generator-api/src/repository"
	methods "github.com/IcaroSilvaFK/password-generator-api/src/utils"
	"github.com/google/uuid"
)

func NewPass(pass string) (string, error) {

	hash, err := methods.Hash(pass)

	repository.NewPass(pass)

	if err != nil {
		return "", err
	}

	return hash, nil
}

func RandomPass(quantity int) ([]string, error) {

	var hashArr []string

	for i := 0; i < quantity; i++ {
		uu := uuid.NewString()

		hash, err := NewPass(uu)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		hashArr = append(hashArr, hash)
	}

	return hashArr, nil
}
