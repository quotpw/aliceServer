package helpers

import (
	"aliceServer/alice"
	"aliceServer/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initFunctionNames(db *gorm.DB) {
	var funcNames []string

	for funcName, _ := range alice.Functions {
		funcNames = append(funcNames, funcName)
	}

	// Delete all function names that are not in the list
	var functionsInDB []entities.Function
	db.Find(&functionsInDB)
	for _, dbFunc := range functionsInDB {
		needDelete := true
		for _, funcName := range funcNames {
			if funcName == dbFunc.Name {
				needDelete = false
				break
			}
		}
		if needDelete {
			db.Delete(&dbFunc)
		}
	}

	// Add all function names that are not in the database
	for _, funcName := range funcNames {
		var funcInDB entities.Function
		db.Where("name = ?", funcName).First(&funcInDB)
		if funcInDB.ID == 0 {
			db.Create(&entities.Function{Name: funcName})
		}
	}
}

func InitDatabase(database, host, port, username, password string) *gorm.DB {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&entities.Query{}, &entities.Function{})
	_ = db.AutoMigrate(&entities.Device{}, &entities.DeviceName{})

	initFunctionNames(db)

	return db
}
