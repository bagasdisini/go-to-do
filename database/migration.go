package database

import (
	"fmt"
	"go-to-do/models"
	"go-to-do/pkg"
)

func RunMigration() {
	err := pkg.DB.AutoMigrate(&models.Activity{}, &models.Todo{})

	if err != nil {
		fmt.Println(err)
		panic("Database Migration Failed!")
	}

	fmt.Println("Database Migration Successful!")
}
