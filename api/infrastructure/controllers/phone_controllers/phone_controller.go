package phone_controllers

import (
	"aisale/database/models/phone_model/phone_repo"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	phone_repo.Create(c)
}

func GetById(c *gin.Context) {
	phone_repo.GetById(c)
}

func UpdateById(c *gin.Context) {
	phone_repo.UpdateById(c)
}

func DeleteById(c *gin.Context) {
	phone_repo.DeleteById(c)
}

func GetAll(c *gin.Context) {
	phone_repo.GetAll(c)
}

func FindPhone(c *gin.Context) {
	phone_repo.Search(c)
}
