package repository

import infra "github.com/IcaroSilvaFK/password-generator-api/src/infra/database"

func FindAll() []Model {

	var allPass []Model

	db := infra.NewDatabase()

	db.Debug().Table("pass").Select("*").Find(&allPass)

	return allPass
}
