package controller

import "gorm.io/gorm"

type AuthenticationController struct {
	Database *gorm.DB
}

func (a AuthenticationController) SignIn() {

}
