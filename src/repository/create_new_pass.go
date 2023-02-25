package repository

import (
	"time"

	infra "github.com/IcaroSilvaFK/password-generator-api/src/infra/database"
	"github.com/google/uuid"
)

type Model struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	PASS      string    `json:"pass"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPass(pass string) {
	passCreated := Model{
		PASS:      pass,
		CreatedAt: time.Now(),
		ID:        uuid.NewString(),
	}
	passCreated.Save()
}

func (p Model) Save() {

	db := infra.NewDatabase()

	db.Create(&p)

}
