package database

import (
	"github.com/prabhat27m/react-golang-admin-project/go_admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	database, err := gorm.Open(mysql.Open("root@/go_admin"), &gorm.Config{})
	if err != nil {
		panic("Connection could not be established")
	}
	database.AutoMigrate(&models.User{})
	
}
