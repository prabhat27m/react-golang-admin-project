package database

import (
	"fmt"
	"github.com/prabhat27m/react-golang-admin-project/go_admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	
	database, err := gorm.Open(mysql.Open("root@/go_admin"), &gorm.Config{})
	if err != nil {
		panic("Connection could not be established")
	}
	DB = database

	fmt.Print("Creating table User...")
	database.AutoMigrate(&models.User{})

}
