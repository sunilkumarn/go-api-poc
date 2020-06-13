package models

import (
	"fmt"
	"leads_module/config"

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

	database.AutoMigrate(&LeadDetail{}, &LeadAddress{}, &LeadSubDetail{})
	database.Model(&LeadAddress{}).AddForeignKey("lead_detail_id", "lead_details(id)", "CASCADE", "CASCADE")
	database.Model(&LeadSubDetail{}).AddForeignKey("lead_detail_id", "lead_details(id)", "CASCADE", "CASCADE")
	database.Model(&LeadDetail{}).AddIndex("leaddetails_converted_leadstatus_idx", "converted", "lead_status")

	DB = database
}
