package user_repo

import (
	"aisale/database"
	userModel "aisale/database/models/user_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	db := database.GetDB()

	var user userModel.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user_model"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	db := database.GetDB()

	var users []userModel.User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	db := database.GetDB()

	id := c.Param("id")
	var user userModel.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()

	id := c.Param("id")
	var user userModel.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user_model"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()

	id := c.Param("id")
	if err := db.Delete(&userModel.User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

//func CreateUser(name, email string) (*User, error) {
//	var collection = database.GetCollection()
//
//	user_model := User{Name: name, Email: email}
//	result, err := collection.InsertOne(context.TODO(), user_model)
//	if err != nil {
//		return nil, err
//	}
//	user_model.ID = result.InsertedID.(primitive.ObjectID)
//	return &user_model, nil
//}
//
//func GetUserByID(id primitive.ObjectID) (*User, error) {
//	var collection = database.GetCollection()
//
//	var user_model User
//	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user_model)
//
//	if err != nil {
//		return nil, err
//	}
//	return &user_model, nil
//}
//
//func UpdateUser(id primitive.ObjectID, name, email string) error {
//	var collection = database.GetCollection()
//
//	_, err := collection.UpdateOne(
//		context.TODO(),
//		bson.M{"_id": id},
//		bson.M{"$set": bson.M{"name": name, "email": email}},
//	)
//	return err
//}
//
//func DeleteUser(id primitive.ObjectID) error {
//	var collection = database.GetCollection()
//
//	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
//	return err
//}
