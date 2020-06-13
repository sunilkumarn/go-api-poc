package models

import (
	"farm_tales/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDataBase() {
	config, err := config.AppConfig()

	if err != nil {
		fmt.Println(err)
		panic("Failed to read configuration!")
	}

	dbConfig := config.Database

	connectorUrl := dbConfig.Username + ":" + dbConfig.Password + "@(" + dbConfig.Host + ")/" + dbConfig.Name + "?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(dbConfig.Adapter, connectorUrl)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Farm{}, &FarmAddress{})
	database.Model(&FarmAddress{}).AddForeignKey("farm_id", "farms(id)", "CASCADE", "CASCADE")
	database.Model(&Farm{}).AddIndex("farm_website_idx", "website")

	DB = database
}
