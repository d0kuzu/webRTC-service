package phone_repo

import (
	"aisale/database"
	"aisale/database/models/phone_model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	db := database.GetDB()

	var phone phone_model.Phone
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&phone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create object"})
		return
	}

	c.JSON(http.StatusCreated, phone)
}

func GetAll(c *gin.Context) {
	db := database.GetDB()

	var phone []phone_model.Phone
	if err := db.Find(&phone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find objects"})
		return
	}

	c.JSON(http.StatusOK, phone)
}

func GetById(c *gin.Context) {
	db := database.GetDB()

	var phone phone_model.Phone
	if err := db.First(&phone, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find object"})
		return
	}

	c.JSON(http.StatusOK, phone)
}

func UpdateById(c *gin.Context) {
	db := database.GetDB()

	var phone phone_model.Phone
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&phone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update object"})
		return
	}

	c.JSON(http.StatusOK, phone)
}

func DeleteById(c *gin.Context) {
	db := database.GetDB()

	var phone phone_model.Phone
	if err := db.First(&phone, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find object"})
		return
	}

	if err := db.Delete(&phone).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete object"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Search(c *gin.Context) {
	db := database.GetDB()

	param := c.Param("param")

	var phone phone_model.Phone
	if _, err := strconv.Atoi(param); err == nil {
		if err := db.First(&phone, "phone = ?", param).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
			return
		}
	} else {
		if err := db.First(&phone, "name = ?", param).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
			return
		}
	}

	c.JSON(http.StatusOK, phone)
}
