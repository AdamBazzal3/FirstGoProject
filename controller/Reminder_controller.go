package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reminders.com/m/Models"
	"reminders.com/m/Services"
	"strconv"
)

func New(service *Services.RemindersProviderService) *ReminderController {
	return &ReminderController{service}
}

type ReminderController struct {
	reminderService *Services.RemindersProviderService
}

func (receiver ReminderController) CreateReminder(c echo.Context) error {
	reminder := &Models.Reminder{}

	if err := c.Bind(reminder); err != nil {
		return err
	}

	receiver.reminderService.CreateReminder(reminder)
	return c.JSON(http.StatusCreated, reminder)
}

func (receiver ReminderController) GetReminder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var reminder Models.Reminder

	receiver.reminderService.GetReminder(id, &reminder)

	return c.JSON(http.StatusOK, reminder)
}

func (receiver ReminderController) GetAllReminders(c echo.Context) error {
	var reminders []Models.Reminder

	receiver.reminderService.GetAllReminders(reminders)

	return c.JSON(http.StatusOK, reminders)
}

func (receiver ReminderController) UpdateReminder(c echo.Context) error {
	r := new(Models.Reminder)

	if err := c.Bind(r); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	r.Id = id
	receiver.reminderService.UpdateReminder(*r)

	return c.JSON(http.StatusOK, r)
}

func (receiver ReminderController) DeleteReminder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	receiver.reminderService.DeleteReminder(id)
	return c.NoContent(http.StatusNoContent)
}
