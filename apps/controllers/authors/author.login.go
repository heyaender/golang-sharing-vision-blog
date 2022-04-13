package authors

import (
	"net/http"
	"sv-article/schemas"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func LoginHandler(c *gin.Context) {
	var user schemas.Author
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Can't bind JSON",
		})
	}
	if user.Username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
	} else {
		if user.Password == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		}
	}
	
	login := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := login.SignedString([]byte("secret"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	
}