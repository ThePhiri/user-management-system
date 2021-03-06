package middleware

import (
	"fmt"
	"net/http"

	helper "github.com/thephiri/user-management-system/go-backend/helpers"

	"github.com/gin-gonic/gin"
)

// Authz validates token and authorizes users
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization")
		c.Header("Content-Type", "application/json")

		// Second, we handle the OPTIONS problem
		if c.Request.Method != "OPTIONS" {

			clientToken := c.Request.Header.Get("token")
			if clientToken == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
				c.Abort()
				return
			}

			claims, err := helper.ValidateToken(clientToken)
			if err != "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				c.Abort()
				return
			}

			c.Set("email", claims.Email)
			c.Set("first_name", claims.First_name)
			c.Set("last_name", claims.Last_name)
			c.Set("uid", claims.Uid)
			c.Set("user_type", claims.User_type)

			c.Next()

		} else {

			c.AbortWithStatus(http.StatusOK)
		}

	}
}
