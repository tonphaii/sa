package entity

import (
	"time"

	"gorm.io/gorm"
)

// Entity หลัก บันทึกเวลาขาออก
type RecordTimeOut struct {
	gorm.Model

	RecordTimeOutID string
	TimeOUT         time.Time
	ODO_Meter       int

	EmployeeID *uint
	Employee   Employee

	CarID *uint
	Car   Car

	CaseID *uint
	Case   Case
}

type Car struct {
	gorm.Model

	Car_ID  string
	Name    string
	Company string
	TypeCar string

	RecordTimeOut     []RecordTimeOut     `gorm:"foreignKey:CarID"`
	VehicleInspection []VehicleInspection `gorm:"foreignKey:CarID"`
}

type Case struct {
	gorm.Model

	Case_ID       string
	Case_Name     string
	TypeCase      string
	Location      string
	Whistleblower string

	RecordTimeOut []RecordTimeOut `gorm:"foreignKey:CaseID"`
}

// main Table to link to another Feature
type Employee struct {
	gorm.Model
	Name    string
	Surname string

	//For Enter Relation

	RecordTimeOut     []RecordTimeOut     `gorm:"foreignKey:EmployeeID"`
	VehicleInspection []VehicleInspection `gorm:"foreignKey:EmployeeID"`
}
