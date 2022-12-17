package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonphaii/sa-65-example/entity"
)

// POST /users

func CreateRecordTimeOut(c *gin.Context) {

	var recordtimeout entity.RecordTimeOut
	var car entity.Car
	var casess entity.Case
	var employees entity.Employee

	if err := c.ShouldBindJSON(&recordtimeout); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}
	if tx := entity.DB().Where("id = ?", recordtimeout.CarID).First(&car); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "car not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", recordtimeout.CaseID).First(&casess); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "case not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", recordtimeout.EmployeeID).First(&employees); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	rec := entity.RecordTimeOut{
		RecordTimeOutID: recordtimeout.RecordTimeOutID,
		ODO_Meter:       recordtimeout.ODO_Meter, //กรอกเลขไมล์
		Case:            casess,                  //โยง คสพ Entity Case
		Car:             car,                     //โยง คสพ Entity Car
		Employee:        employees,               //โยง คสพ Entity Employee
		TimeOUT:         recordtimeout.TimeOUT,   // field DateTime
	}

	//บันทึก

	if err := entity.DB().Create(&rec).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": rec})

}

// GET /user/:id

func GetRecordTimeOut(c *gin.Context) {

	var recordtimeout entity.RecordTimeOut

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM recordtimeouts WHERE id = ?", id).Scan(&recordtimeout).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recordtimeout})

}

// GET /users

func ListRecordTimeOuts(c *gin.Context) {

	var recordtimeouts []entity.RecordTimeOut

	if err := entity.DB().Raw("SELECT * FROM recordtimeouts").Scan(&recordtimeouts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": recordtimeouts})

}

// DELETE /users/:id

func DeleteRecordTimeOut(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM recordtimeouts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recordtimeout not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateRecordTimeOut(c *gin.Context) {

	var recordtimeout entity.RecordTimeOut

	if err := c.ShouldBindJSON(&recordtimeout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	if tx := entity.DB().Where("id = ?", recordtimeout.ID).First(&recordtimeout); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recordtimeout not found"})
		return

	}

	if err := entity.DB().Save(&recordtimeout).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": recordtimeout})

}
