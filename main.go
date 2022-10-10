package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reminders.com/m/Models"
	"reminders.com/m/controller"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//database testing
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = database.AutoMigrate(&Models.Reminder{})
	if err != nil {
		return
	}
	reminderController := controller.New(database)

	// Routes
	e.GET("/reminders", reminderController.GetAllReminders)
	e.POST("/reminders", reminderController.CreateReminder)
	e.GET("/reminders/:id", reminderController.GetReminder)
	e.PUT("/reminders/:id", reminderController.UpdateReminder)
	e.DELETE("/reminders/:id", reminderController.DeleteReminder)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
