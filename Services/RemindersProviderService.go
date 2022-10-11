package Services

import (
	"gorm.io/gorm"
	"reminders.com/m/Models"
)

func New(database *gorm.DB) *RemindersProviderService {
	return &RemindersProviderService{database: database}
}

type RemindersProviderService struct {
	database *gorm.DB
}

func (receiver RemindersProviderService) CreateReminder(reminder *Models.Reminder) {
	receiver.database.Create(reminder)
}

func (receiver RemindersProviderService) GetReminder(id int, reminder *Models.Reminder) *Models.Reminder {
	receiver.database.First(reminder, id)

	return reminder
}

func (receiver RemindersProviderService) GetAllReminders(reminders []Models.Reminder) []Models.Reminder {

	receiver.database.Find(&reminders)

	return reminders
}

func (receiver RemindersProviderService) UpdateReminder(r Models.Reminder) {
	receiver.database.Model(&r).Updates(r)
}

func (receiver RemindersProviderService) DeleteReminder(id int) {
	receiver.database.Delete(&Models.Reminder{}, id)
}
