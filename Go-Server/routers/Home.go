package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/db"
	"main/jwtToken"
	"net/http"
)

func BasicHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		str, err := jwtToken.ParseToken(token)
		if err != nil {
			fmt.Println(err)
			context.AbortWithStatusJSON(http.StatusOK, gin.H{"msg": "TokenError"})
		} else {
			fmt.Println(*str)
			context.Next()
		}
	}
}

func GetData(data *gin.Context) {
	res, _ := db.Engine.QueryInterface("select * from userdata")
	data.JSON(http.StatusOK, gin.H{"result": res})
}
