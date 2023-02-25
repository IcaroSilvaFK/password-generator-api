package repository

import (
	infra "github.com/IcaroSilvaFK/password-generator-api/src/infra/database"
	"github.com/google/uuid"
)

type Model struct {
	ID   string `json:"id" gorm:"primaryKey"`
	PASS string `json:"pass"`
}

func NewPass(pass string) Model {
	passCreated := Model{
		ID:   uuid.NewString(),
		PASS: pass,
	}
	passCreated.Save()

	return passCreated
}

func (p Model) Save() {

	db := infra.NewDatabase()

	db.Debug().Table("pass").Select("id", "pass").Create(&p)
}
