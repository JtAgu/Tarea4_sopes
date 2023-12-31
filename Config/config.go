package Config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var Uri = "root:secret@tcp(database:3306)/tarea3?charset=utf8"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(Uri), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}

	//err = Database.AutoMigrate(&Entities.Music{})
	if err != nil {
		return err
	}

	return nil
}
