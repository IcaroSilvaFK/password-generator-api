package repository

import infra "github.com/IcaroSilvaFK/password-generator-api/src/infra/database"

type CreateInBatchPassType struct {
	ID   string `json:"id"`
	PASS string `json:"pass"`
}

func CreateInBatchPass(pass []CreateInBatchPassType, limit int) {

	db := infra.NewDatabase()

	db.Table("pass").CreateInBatches(pass, limit)
}
