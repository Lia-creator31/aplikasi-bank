package models

import (
	"create-api/helpers"
	"github.com/luthfikw/govalidator"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"not null" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required"`
	Photos    []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	CreatedAt *time.Time `json:"-,omitempty"`
	UpdatedAt *time.Time `json:"-,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	//return
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
