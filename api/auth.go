package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"searchEngine/models"
	"time"
)

type JWTHandler struct {
}

func CreateJWTHandler() *JWTHandler {
	return &JWTHandler{}
}

func (h *JWTHandler) SignInHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if user.Username != "admin" || user.Password != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid username or password",
			})
		}

		expireTime := time.Now().Add(10 * time.Minute)
		claims := models.Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expireTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
		tokenString, err := token.SignedString(models.JwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		jwtOutput := models.JWTOutput{
			Token:   tokenString,
			Expires: expireTime,
		}
		c.JSON(http.StatusOK, jwtOutput)
	}
}

func (h *JWTHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		claims := models.Claims{}
		tkn, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(models.JwtKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		}
		if tkn == nil || !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		}
		c.Next()
	}

}
