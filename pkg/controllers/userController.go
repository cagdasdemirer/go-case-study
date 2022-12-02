package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-case-study/pkg/initializers"
	"go-case-study/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func UserCreate(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong input !"})
		return
	}

	// Hashing the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return
	}
	user.Password = string(hash)

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "User not created."})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var user models.User

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong input !"})
		return
	}

	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid email."})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid password."})
		return
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":     user.ID,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Error while generating token."})
		return
	}

	// Return token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AuthToken", tokenString, 3600*24, "", "", false, true)
	// Secure should be true in production

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Token successfully generated."})
}

func UserValidate(c *gin.Context) {
	user, err := c.Get("user")
	if !err {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Token successfully generated."})
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "Welcome Admin !"})
}
