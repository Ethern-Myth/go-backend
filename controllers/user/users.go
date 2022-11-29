package user

import (
	"backend_sqlite/data"
	"backend_sqlite/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetUsers(u *gin.Context){
	var users[] models.User
	data.DB.Find(&users)
	u.JSON(http.StatusOK, users)
}

func GetUser(u *gin.Context){
	var user models.User
	if err:= data.DB.Where("id = ?", u.Param("id")).First(&user).Error; err != nil{
		u.JSON(http.StatusBadRequest, gin.H{"error":"User not found"})
		return
	}
	u.JSON(http.StatusOK, user)
}

func CreateUser(u *gin.Context){
	var newUser models.CreateUser
	if err:=u.ShouldBindJSON(&newUser); err != nil{
		u.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	user := models.User{
		Name: newUser.Name,
		Surname: newUser.Surname,
		Email: newUser.Email,
		Password: newUser.Password,
	}
	data.DB.Create(&user)
	u.JSON(http.StatusOK, user)
}

func UpdateUser(u *gin.Context){
	var user models.User
	if err:= data.DB.Where("id = ?", u.Param("id")).First(&user).Error; err != nil{
		u.JSON(http.StatusBadRequest, gin.H{"error":"User not found to update"})
		return
	}
	var updateData models.UpdateUser
	if err:= u.ShouldBindJSON(&updateData); err!=nil{
		u.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	data.DB.Model(&user).Update(updateData)
	u.JSON(http.StatusOK, updateData)
}

func DeleteUser(u * gin.Context){
	var user models.User
	if err:= data.DB.Where("id = ?", u.Param("id")).First(&user).Error; err !=nil{
		u.JSON(http.StatusBadRequest, gin.H{"error":"User not found to delete"})
		return
	}
	data.DB.Delete(&user)
	u.JSON(http.StatusOK, gin.H{"deleted":true})
}