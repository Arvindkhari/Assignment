package entity

//Person object for REST(CRUD)
type Student struct {
	Id            int64  `gorm:"primary_key;auto_increment;not_null"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	MobileNumber  int64  `json:"mobileNumber" gorm:"unique"`
	YearOfJoining int64  `json:"yearOfJoining"`
	EmailId       string `json:"emailId"`
}
