package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	UserID int    `json:"userID"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (todo *Todo) Validate() (string, bool) {

	if todo.UserID < 1 {
		return "User is not recognized", false
	}

	if todo.Title == "" {
		return "Title is empty", false
	}

	if todo.Body == "" {
		return "Body is empty", false
	}

	//All the required parameters are present
	return "", true
}

func (todo *Todo) Create() *Todo {

	if resp, ok := todo.Validate(); !ok {
	}

	GetDB().Create(todo)
	return todo
}

func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {

	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
