package middleware

import (
	"ginEssential/Model"
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证toekn格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足A"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足B"})
			ctx.Abort()
			return
		}

		//验证通过获取token中的claim 中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user Model.User
		DB.First(&user, userId)

		//用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足C"})
			ctx.Abort()
			return
		}
		//用户存在 将user信息存入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
