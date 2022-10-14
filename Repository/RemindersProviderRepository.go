package Repository

import (
	"gorm.io/gorm"
	"reminders.com/m/Models"
)

type RemindersProviderRepository struct {
	Database *gorm.DB
}

func (receiver *RemindersProviderRepository) CreateReminder(reminder *Models.Reminder) {
	receiver.Database.Create(reminder)
}

func (receiver *RemindersProviderRepository) GetReminder(id int, reminder *Models.Reminder) *Models.Reminder {
	receiver.Database.First(reminder, id)

	return reminder
}

func (receiver *RemindersProviderRepository) GetAllReminders(reminders *[]Models.Reminder) {
	receiver.Database.Find(reminders)
}

func (receiver *RemindersProviderRepository) UpdateReminder(r Models.Reminder) {
	receiver.Database.Model(&r).Updates(r)
}

func (receiver *RemindersProviderRepository) DeleteReminder(id int) {
	receiver.Database.Delete(&Models.Reminder{}, id)
}
