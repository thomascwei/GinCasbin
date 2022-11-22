package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var SecretKey = "thomaswei"

// MyClaims Customer jwt.StandardClaims
type MyClaims struct {
	Account string `json:"account"`
	jwt.RegisteredClaims
}

// GenToken Create a new token
func GenToken(account string) (string, error) {
	thisSecretKey := []byte(SecretKey)
	claims := MyClaims{
		account,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "thomaswei",
		},
	}

	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Choose specific Signature
	return token.SignedString(thisSecretKey)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	// Valid token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware Middleware of JWT
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get token from Header.Authorization field.
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "fail",
				"msg":    "Authorization is null in Header",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "fail",
				"msg":    "Format of Authorization is wrong",
			})
			return
		}
		// parts[0] is Bearer, parts is token.
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "fail",
				"msg":    "Invalid Token.",
			})
			return
		}
		// Store Account info into Context
		c.Set("account", mc.Account)
		// After that, we can get Account info from c.Get("account")
		c.Next()
	}
}
