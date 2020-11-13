package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
)

//EnableAdminIAMRouter 啟動預設的IAM Routers
func EnableAdminIAMRouter(rg *gin.RouterGroup) {
	user := rg.Group("user")
	{
		iam.GET(user, "", listUsers, "admin:ListUser", true)
		iam.POST(user, "", createUser, "admin:CreateUser", true)
		spUser := user.Group(":user-id") //TODO: Need check user exist
		{
			iam.GET(spUser, "", getUser, "admin:GetUser", true)
			iam.PUT(spUser, "", updateUser, "admin:UpdateUser", true)
			iam.PUT(spUser, "password", updatePassword, "admin:UpdatePassword", true)
			iam.DELETE(spUser, "", deleteUser, "admin:DeleteUser", true)
		}
	}
}
