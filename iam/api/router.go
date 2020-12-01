package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/utility"
)

func checkUser(c *gin.Context) {
	if statusCode, err := checkUserExist(c, c.Param(userIDParams), true); err != nil {
		utility.ResponseWithType(c, statusCode, &utility.ErrResponse{
			Message: err.Error(),
		})
	}
	c.Next()
}

func checkGroup(c *gin.Context) {
	if statusCode, err := checkGroupExist(c, c.Param(groupIDParams), true); err != nil {
		utility.ResponseWithType(c, statusCode, &utility.ErrResponse{
			Message: err.Error(),
		})
	}
	c.Next()
}

//EnableAdminIAMRouter 啟動預設的IAM Routers
func EnableAdminIAMRouter(rg *gin.RouterGroup) {
	user := rg.Group("user")
	{
		iam.POST(user, "", createUser, "admin:CreateUser", true)
		iam.GET(user, "", listUsers, "admin:ListUser", true)
		spUser := user.Group(":user-id", checkUser)
		{
			iam.GET(spUser, "", getUser, "admin:GetUser", true)
			iam.PUT(spUser, "", updateUser, "admin:UpdateUser", true)
			iam.PUT(spUser, "password", updatePassword, "admin:UpdatePassword", true)
			iam.DELETE(spUser, "", deleteUser, "admin:DeleteUser", true)
		}
	}
	group := rg.Group("group")
	{
		iam.POST(group, "", createGroup, "admin:CreateGroup", true)
		iam.GET(group, "", listGroups, "admin:ListGroups", true)
		spGroup := group.Group(":group-id", checkGroup)
		{
			iam.GET(spGroup, "", getGroup, "admin:GetGroup", true)
			iam.PUT(spGroup, "", updateGroup, "admin:UpdateGroup", true)
			iam.DELETE(spGroup, "", deleteGroup, "admin:DeleteGroup", true)
		}
	}
	membership := rg.Group("membership")
	{
		iam.POST(membership, "", createMembership, "admin:CreateMembership", true)
		user := membership.Group("user/:user-id", checkUser)
		{
			iam.GET(user, "", listMembershipsByUser, "admin:ListMembershipsByUser", true)
		}
		indevGroup := membership.Group("group/:group-id", checkGroup)
		{
			iam.GET(indevGroup, "", listMembershipsByGroup, "admin:ListMembershipsByGroup", true)

			indevUser := indevGroup.Group("user/:user-id", checkUser)
			{
				iam.GET(indevUser, "", getMembershipAndPermission, "admin:GetMembership", true)
				iam.PUT(indevUser, "", updateMembership, "admin:UpdateMembership", true)
				iam.DELETE(indevUser, "", deleteMembership, "admin:DeleteMembership", true)
			}
		}
	}
	iam.GET(rg, "/actions", listPermissionActions, "admin:ListPermissionActions", true)
}
