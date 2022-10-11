package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"reminders.com/m/Models"
	"strconv"
)

func New(db *gorm.DB) *ReminderController {
	return &ReminderController{db} // <- 33: a very sensible default value
}

type ReminderController struct {
	database *gorm.DB
}

func (receiver ReminderController) CreateReminder(c echo.Context) error {
	reminder := &Models.Reminder{}

	if err := c.Bind(reminder); err != nil {
		return err
	}

	receiver.database.Create(&reminder)
	return c.JSON(http.StatusCreated, reminder)
}

func (receiver ReminderController) GetReminder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var reminder Models.Reminder

	receiver.database.First(&reminder, id)

	return c.JSON(http.StatusOK, reminder)
}

func (receiver ReminderController) GetAllReminders(c echo.Context) error {
	var reminders []Models.Reminder

	receiver.database.Find(&reminders)

	return c.JSON(http.StatusOK, reminders)
}

func (receiver ReminderController) UpdateReminder(c echo.Context) error {
	r := new(Models.Reminder)
	var reminder Models.Reminder

	if err := c.Bind(r); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	receiver.database.First(&reminder, id)
	receiver.database.Model(&reminder).Where("id = ?", id).Updates(r)
	return c.JSON(http.StatusOK, reminder)
}

func (receiver ReminderController) DeleteReminder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	receiver.database.Delete(&Models.Reminder{}, id)
	return c.NoContent(http.StatusNoContent)
}
