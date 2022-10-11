package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"reminders.com/m/Models"
	"reminders.com/m/Services"
	"reminders.com/m/controller"
)

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func handlePage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(request.Body).Decode(&message)
	if err != nil {
		return
	}
	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		return
	}
}

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

	service := Services.New(database)
	reminderController := controller.New(service)

	// Routes
	e.GET("/reminders", reminderController.GetAllReminders)
	e.POST("/reminders", reminderController.CreateReminder)
	e.GET("/reminders/:id", reminderController.GetReminder)
	e.PUT("/reminders/:id", reminderController.UpdateReminder)
	e.DELETE("/reminders/:id", reminderController.DeleteReminder)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
	//http.HandleFunc("/home", handlePage)
	//err = http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	log.Println("There was an error listening on port :8080", err)
	//}
}
