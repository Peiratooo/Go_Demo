package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"main/db"
	"main/jwtToken"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 允许 Origin 字段中的域发送请求
		context.Writer.Header().Add("Access-Control-Allow-Origin", context.GetHeader("origin"))
		// 设置预验请求有效期为 70000 秒
		context.Writer.Header().Set("Access-Control-Max-Age", "70000")
		// 设置允许请求的方法
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		// 设置允许请求的 Header
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length,token")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Headers,token")
		// 配置是否可以带认证信息
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// OPTIONS请求返回200
		if context.Request.Method == "OPTIONS" {
			fmt.Println(context.Request.Header)
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}

func initData(rData []byte, err error) map[string]interface{} {
	var data map[string]interface{}
	if err != nil {
		return data
	}
	_ = json.Unmarshal(rData, &data)
	return data
}

func checkData(data map[string]interface{}, keyWords []string) bool {
	//fmt.Println(data)
	for _, i := range keyWords {
		//fmt.Println(i, data[i])
		if _, ok := data[i]; !ok {
			return false
		}
	}
	return true
}

func Login(data *gin.Context) {
	loginData := initData(data.GetRawData())
	fmt.Println(loginData["uidoremail"])
	if checkData(loginData, []string{"uidoremail", "password"}) {
		res, _ := db.Engine.QueryInterface("select * from userdata where uid = ? or email = ?", loginData["uidoremail"], loginData["uidoremail"])
		if len(res) > 0 {
			//fmt.Printf("%T", res)
			if res[0]["password"] == loginData["password"] {
				token, err := jwtToken.GenToken(res[0])
				if err != nil {
					fmt.Println(err)
				}
				//str, _ := jwtToken.ParseToken(token)
				//fmt.Println("pToken =>", str)
				data.JSON(http.StatusOK, gin.H{"msg": "OK", "token": token})
				return
			}
		}
	}
	data.JSON(http.StatusOK, gin.H{"msg": "账号或密码错误！"})

}

func Regsiter(data *gin.Context) {
	errMsg := "注册失败! 请稍后再试!"
	regData := initData(data.GetRawData())
	fmt.Printf("%T:%v", regData, regData)
	if checkData(regData, []string{"name", "email", "password"}) {
		checkExist, _ := db.Engine.QueryInterface("select * from userdata where email = ?", regData["email"])
		if len(checkExist) == 0 {
			sql := "insert into userdata (name, email, password) VALUES (?, ?, ?)"
			_, err := db.Engine.Exec(sql, regData["name"], regData["email"], regData["password"])
			if err != nil {
				fmt.Println(err)
			}
			data.JSON(http.StatusOK, gin.H{"msg": "OK"})
			return
		} else {
			errMsg = "注册失败！邮箱重复！"
		}
	}
	data.JSON(http.StatusOK, gin.H{"msg": errMsg})
}
