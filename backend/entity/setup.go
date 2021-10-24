package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connent database")
	}

	//migrate the schema
	database.AutoMigrate(
		&Patient{},
		&User{},
		&Role{},
		&Appointment{},
		&Medical_product{},
		&Screening_record{},
	)
	db = database
}

/*package entity
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	database.AutoMigrate(
		&Video{}, &User{}, &Playlist{}, &Resolution{}, &WatchVideo{},
	)

	db = database
}
/* Example
package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect satabase")
	}

	//Migrate the schma
	database.AutoMigrate(&User{})

	db = database
} */
