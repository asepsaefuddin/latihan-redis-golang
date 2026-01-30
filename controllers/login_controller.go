package controllers

import (
	"gin-quickstart/config"
	"gin-quickstart/helper"
	"gin-quickstart/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// melakukan pencarian email
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email/password salah"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email/password salah"})
		return
	}
	// buat jwt / kalau berhasil
	token, _ := helper.GenerateJWT(user.ID)
	c.JSON(http.StatusOK, gin.H{"message": "berhasil login", "token": token})
}
