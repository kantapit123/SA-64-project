package entity

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name              string
	Sex               string
	Age               *uint
	Allergy           string
	ID_card           string
	Time              time.Time
	Tel               string
	Appointments      []Appointment      `gorm:"foreignKey:Patient_appointmentID"`
	Screening_records []Screening_record `gorm:"foreignKey:Scr_patientID"`
}

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"uniqueIndex"`
	Password string
	//Role FK
	User_roleID *uint
	User_role   Role

	Screening_records []Screening_record `gorm:"foreignKey:Scr_userID"`
}
type Role struct {
	gorm.Model
	role_name string
	Users     []User `gorm:"foreignKey:User_roleID"`
}

type Appointment struct {
	gorm.Model
	date time.Time
	todo string

	Patient_appointmentID *uint
	Patient_appoint       Patient
}

type Medical_product struct {
	gorm.Model
	Medication_name   string
	Screening_records []Screening_record `gorm:"foreignKey:Scr_medical_productID"`
}

type Screening_record struct {
	gorm.Model
	Illnesses string
	Detail    string
	Queue     string

	//Patient	FK
	Scr_patientID *uint
	Scr_patient   Patient

	//User		FK
	Scr_userID *uint
	Scr_user   User

	//Medical_product	FK
	Scr_medical_productID *uint
	Scr_medical_product   Medical_product
}

/* ระบบคัดกรองผู้ป่วยก่อนเข้ารับการรักษา
package entity

import (
	"time"
	"gorm.io/gorm"
)

type Dentist struct{
	gorm.Model
	Name 		string
	username 	string							`gorm:"uniqueIndex"`
	password 	string
	// 1 user เป็นเจ้าของได้หลาย screening
	Screening_records []Screening_record 		`gorm:"foreignKey:dentist_ownerID"`

	Patient_infomations []Patient_infomation 	`gorm:"foreignKey:dentist_ownerID"`

}

type Patient_infomation struct{
	gorm.Model
	Name 				string
	age 				*uint
	//dentist_ownerID ทำหน้าที่เป็น FK
	dentist_ownerID		*uint
	//เป็นข้อมูล dentist ใช้เพื่อให้ join ง่ายขึ้น
	dentist_owner		Dentist
	// 1 patient มีได้ 1 screening record
	Screening_records []Screening_record		`gorm:"foreignKey:patient_ownerID"`
	// 1 patient มีได้ หลาย Appointment
	Appointments []Appointment					`gorm:"goreignKey:patient_ownerID"`


}

type Appointment struct{
	gorm.Model
	date 				time.Time
	type_id				*uint
	type_name			string



}

type Screening_record struct{
	gorm.Model
}


/* ---------------------------------------
package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string     `gorm:"uniqueIndex"`
	Videos    []Video    `gorm:"foreignKey:OwnerID"`
	Playlists []Playlist `gorm:"foreignKey:OwnerID"`
}

type Video struct {
	gorm.Model
	Name        string
	Url         string `gorm:"uniqueIndex"`
	OwnerID     *uint
	Owner       User
	WatchVideos []WatchVideo `gorm:"foreignKey:VideoID"`
}

type Playlist struct {
	gorm.Model
	Title       string
	OwnerID     *uint
	Owner       User
	WatchVideos []WatchVideo `gorm:"foreignKey:PlaylistID"`
}

type Resolution struct {
	gorm.Model
	Value       string
	WatchVideos []WatchVideo `gorm:"foreignKey:ResolutionID"`
}

type WatchVideo struct {
	gorm.Model
	WatchedTime time.Time

	//Video FK
	VideoID *uint
	Video   Video
	//Playlist	FK
	PlaylistID *uint
	Playlist   Playlist
	//Resolution FK
	ResolutionID *uint
	Resolution   Resolution
}

/*

-----------------------------------------
/* Example
package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string
	Age       uint8
	BirthDay  time.Time
} */
