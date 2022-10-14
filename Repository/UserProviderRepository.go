package Repository

import (
	"gorm.io/gorm"
	"reminders.com/m/Models"
)

type UserProviderRepository struct {
	Database *gorm.DB
}

func (receiver *UserProviderRepository) CreateUser(reminder *Models.User) error {
	result := receiver.Database.Create(reminder)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (receiver *UserProviderRepository) GetUserById(id int, reminder *Models.User) *Models.User {
	receiver.Database.First(reminder, id)

	return reminder
}

func (receiver *UserProviderRepository) GetUserByUsername(username string, reminder *Models.User) error {
	result := receiver.Database.First(reminder, "username = ?", username)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (receiver *UserProviderRepository) GetAllUsers(reminders []Models.User) []Models.User {

	receiver.Database.Find(&reminders)

	return reminders
}

func (receiver *UserProviderRepository) UpdateUser(r Models.User) {
	receiver.Database.Model(&r).Updates(r)
}

func (receiver *UserProviderRepository) DeleteUser(id int) {
	receiver.Database.Delete(&Models.User{}, id)
}
