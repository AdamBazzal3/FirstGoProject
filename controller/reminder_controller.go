package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reminders.com/m/Models"
	"reminders.com/m/Repository"
	"strconv"
)

type ReminderController struct {
	ReminderService *Repository.RemindersProviderRepository
}

func (receiver *ReminderController) PostCreateReminder(c echo.Context) error {
	reminder := &Models.Reminder{}

	if err := c.Bind(reminder); err != nil {
		return err
	}

	receiver.ReminderService.CreateReminder(reminder)
	return c.JSON(http.StatusCreated, reminder)
}

func (receiver *ReminderController) GetReminder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	var reminder Models.Reminder

	receiver.ReminderService.GetReminder(id, &reminder)

	return c.JSON(http.StatusOK, reminder)
}

func (receiver *ReminderController) GetAllReminders(c echo.Context) error {
	var reminders []Models.Reminder

	receiver.ReminderService.GetAllReminders(&reminders)

	return c.JSON(http.StatusOK, reminders)
}

func (receiver *ReminderController) PutUpdateReminder(c echo.Context) error {
	r := new(Models.Reminder)

	if err := c.Bind(r); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	r.Id = id
	receiver.ReminderService.UpdateReminder(*r)

	return c.JSON(http.StatusOK, r)
}

func (receiver *ReminderController) DeleteReminder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	receiver.ReminderService.DeleteReminder(id)
	return c.NoContent(http.StatusNoContent)
}
