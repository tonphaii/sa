package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonphaii/sa-65-example/entity"
)

// POST /users

// VehicleInspection
func CreateVehicleInspection(c *gin.Context) {

	var vehicleinspection entity.VehicleInspection
	var car entity.Car
	var status entity.Status
	var employees entity.Employee

	if err := c.ShouldBindJSON(&vehicleinspection); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}
	if tx := entity.DB().Where("id = ?", vehicleinspection.CarID).First(&car); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "car not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", vehicleinspection.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", vehicleinspection.EmployeeID).First(&employees); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	veh := entity.VehicleInspection{
		VehicleInspectionID: vehicleinspection.VehicleInspectionID,
		Fail:                vehicleinspection.Fail, //กรอกเลขปัญหา
		Damaged_Area:        vehicleinspection.Damaged_Area,
		Status:              status,                      //โยง คสพ Entity Case
		Car:                 car,                         //โยง คสพ Entity Car
		Employee:            employees,                   //โยง คสพ Entity Employee
		Case_Time:           vehicleinspection.Case_Time, // field DateTime
	}

	//บันทึก

	if err := entity.DB().Create(&veh).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": veh})

}

// GET /user/:id

func GetVehicleInspection(c *gin.Context) {

	var vehicleinspection entity.VehicleInspection

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM vehicleinspections WHERE id = ?", id).Scan(&vehicleinspection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vehicleinspection})

}

// GET /users

func ListVehicleInspection(c *gin.Context) {

	var vehicleinspections []entity.VehicleInspection

	if err := entity.DB().Raw("SELECT * FROM vehicleinspections").Scan(&vehicleinspections).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": vehicleinspections})

}

// DELETE /users/:id

func DeleteVehicleInspection(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM vehicleinspections WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vehicleinspection not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateVehicleInspection(c *gin.Context) {

	var vehicleinspection entity.VehicleInspection

	if err := c.ShouldBindJSON(&vehicleinspection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	if tx := entity.DB().Where("id = ?", vehicleinspection.ID).First(&vehicleinspection); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vehicleinspection not found"})
		return

	}

	if err := entity.DB().Save(&vehicleinspection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": vehicleinspection})

}
