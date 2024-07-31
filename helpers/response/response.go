package response

import "github.com/gin-gonic/gin"

func InvalidRequest(c *gin.Context) {
	c.JSON(400, gin.H{
		"message": "Invalid request",
	})
}

func WrongCredentials(c *gin.Context) {
	c.JSON(401, gin.H{
		"message": "Wrong credentials",
	})
}

func InternalServerError(c *gin.Context) {
	c.JSON(500, gin.H{
		"message": "Internal server error",
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(401, gin.H{
		"message": "Unauthorized",
	})
}

func Forbidden(c *gin.Context) {
	c.JSON(403, gin.H{
		"message": "Forbidden",
	})
}

func NotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "Not found",
	})
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"message": "Success",
		"content": data,
	})
}
