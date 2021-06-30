package domain

import "time"

type FullName struct {
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName"`
}

type PolicyHolder struct {
	//gorm.models
	AadhaarCardNo string    `json:"aadhaarCardNo" gorm:"primary_key"`
	Name          FullName  `gorm:"embedded"`
	DOB           time.Time `json:"dob"`
	Policies      []Policy  `json:"policies" gorm:"foreignkey:AadhaarCardNo"`
}

type Policy struct {
	PolicyNo      int64     `json:"policyNo" gorm:"primary_key"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	InsuredAmount int32     `json:"insuredAmount"`
	AadhaarCardNo string    `json:"-"`
	//VehicleInstance Vehicle   `json:"vehicleInstance" gorm:"foreignkey:PolicyNo"`
}

/*
type Vehicle struct {
	RegNo     string    `json:"regNo" gorm:"primary_key"`
	Brand     string    `json:"brand"`
	ModelName string    `json:"modelName"`
	DOP       time.Time `json:"dop"`
	PolicyNo  int64     `json:"-"`
}
*/
