package database

import (
	"fmt"

	"github.com/Infintive/predictive-go/app/models"
)

var memoryDB = make([]models.User, 5)
var index = 0

func InitMemoryDB() {
	memoryDB[0] = models.NewUser("rajesh@yahoo.com", "admin123", "Rajesh", 100)
	memoryDB[1] = models.NewUser("ralph@yahoo.com", "admin123", "Ralph", 101)
	memoryDB[2] = models.NewUser("zach@yahoo.com", "admin123", "Zach", 101)
	index = 2

}

func AddUser(u models.User) bool {

	for _, usr := range memoryDB {
		if usr.Email == u.Email {
			return false
		}
	}
	index++
	memoryDB[index] = u
	printUsers()
	return true
}

func printUsers() {
	fmt.Println("Current users in the memory database")
	for _, user := range memoryDB {
		fmt.Printf("User: %+v\n", user)
	}
}

func GetUserByEmail(email string) *models.User {
	for _, user := range memoryDB {
		if email == user.Email {
			return &user
		}
	}
	return nil
}
