package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonphaii/sa-65-example/entity"
)

// GET /user/:id

func GetCase(c *gin.Context) {

	var cases entity.Case

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM cases WHERE id = ?", id).Scan(&cases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cases})

}

// GET /users

func ListCases(c *gin.Context) {

	var cases []entity.Case

	if err := entity.DB().Raw("SELECT * FROM cases").Scan(&cases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cases})

}

// DELETE /users/:id

func DeleteCase(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM cases WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "case not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /users

func UpdateCase(c *gin.Context) {

	var cases entity.Case

	if err := c.ShouldBindJSON(&cases); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", cases.ID).First(&cases); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "case not found"})

		return

	}

	if err := entity.DB().Save(&cases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cases})

}

// POST /users

func CreateCase(c *gin.Context) {

	var cases entity.Case

	if err := c.ShouldBindJSON(&cases); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&cases).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": cases})

}
