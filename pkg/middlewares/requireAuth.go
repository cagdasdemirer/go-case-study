package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-case-study/pkg/initializers"
	"go-case-study/pkg/models"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("Auth controlling...")
	tokenString, err := c.Cookie("AuthToken")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized."})
	}

	//Decode and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Expiration Control
		if time.Now().Unix() > claims["expiration"].(int64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized."})
		}

		var user models.User
		initializers.DB.First(&user, claims["userid"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not a User anymore :("})
		}

		//Admin control
		if user.RoleID != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to do this."})
		}

		//Add user to req context
		c.Set("user", user)

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized."})
	}

	c.Next()

}
