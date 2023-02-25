package services

import (
	"log"

	"github.com/IcaroSilvaFK/password-generator-api/src/repository"
	methods "github.com/IcaroSilvaFK/password-generator-api/src/utils"
	"github.com/google/uuid"
)

type PassService struct{}

type CreateInBatchStructure struct {
	ID   string
	PASS string
}

func NewPassService() *PassService {
	return &PassService{}
}

func (p PassService) NewPass(pass string) (repository.Model, error) {

	hash, err := methods.Hash(pass)

	passCreated := repository.NewPass(hash)

	if err != nil {
		return repository.Model{}, err
	}

	return passCreated, nil
}

func (p PassService) FindAllPass() []repository.Model {
	res := repository.FindAll()

	return res
}

func (p PassService) RandomPass(quantity int) ([]repository.CreateInBatchPassType, error) {

	var hashArr []repository.CreateInBatchPassType

	for i := 0; i < quantity; i++ {
		uu := uuid.NewString()
		hash, err := methods.Hash(uu)

		if err != nil {
			log.Fatal(err)
		}

		hashArr = append(hashArr, repository.CreateInBatchPassType{
			ID:   uu,
			PASS: hash,
		})
	}

	repository.CreateInBatchPass(hashArr, quantity)

	return hashArr, nil
}
