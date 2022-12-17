package entity

import (
	"time"

	"gorm.io/gorm"
)

// Entity หลัก บันทึกเวลาขาออก
type VehicleInspection struct {
	gorm.Model

	VehicleInspectionID string
	Case_Time           time.Time
	Fail                string
	Damaged_Area        string

	EmployeeID *uint
	Employee   Employee

	CarID *uint
	Car   Car

	StatusID *uint
	Status   Status
}

type Status struct {
	gorm.Model

	Status_ID int
	Name      string

	VehicleInspection []VehicleInspection `gorm:"foreignKey:StatusID"`
}
