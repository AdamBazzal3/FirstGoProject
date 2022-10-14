package Services

import (
	"gorm.io/gorm"
	"reminders.com/m/Models"
)

func NewAuthService(database *gorm.DB) *RemindersProviderService {
	return &RemindersProviderService{database: database}
}

type UserProviderService struct {
	database *gorm.DB
}

func (receiver UserProviderService) CreateUser(reminder *Models.User) {
	receiver.database.Create(reminder)
}

func (receiver UserProviderService) GetUser(id int, reminder *Models.User) *Models.User {
	receiver.database.First(reminder, id)

	return reminder
}

func (receiver UserProviderService) GetAllUsers(reminders []Models.User) []Models.User {

	receiver.database.Find(&reminders)

	return reminders
}

func (receiver UserProviderService) UpdateUser(r Models.User) {
	receiver.database.Model(&r).Updates(r)
}

func (receiver UserProviderService) DeleteUser(id int) {
	receiver.database.Delete(&Models.User{}, id)
}
