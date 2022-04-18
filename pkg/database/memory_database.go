package database

import (
	"github.com/Infintive/predictive-go/app/models"
)

func InitMemoryDB() {
	var memoryDB [3]models.User

	memoryDB[0] = models.NewUser("rajesh@yahoo.com", "admin123", "Rajesh", 100)
	memoryDB[1] = models.NewUser("ralph@yahoo.com", "admin123", "Ralph", 101)
	memoryDB[2] = models.NewUser("zach@yahoo.com", "admin123", "Zach", 101)

}
