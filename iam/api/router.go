package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam/abac"
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

func checkMembership(c *gin.Context) {
	if statusCode, err := checkMembershipExist(c, c.Param(userIDParams), c.Param(groupIDParams), true); err != nil {
		utility.ResponseWithType(c, statusCode, &utility.ErrResponse{
			Message: err.Error(),
		})
	}
	c.Next()
}

//EnableAdminIAMRouter 啟動預設的IAM Routers
func EnableAdminIAMRouter(rg *gin.RouterGroup) {
	abac.GET(rg, "/actions", listPermissionActions, "admin:ListPermissionActions", true)
	user := rg.Group("user")
	{
		abac.POST(user, "", createUser, "admin:CreateUser", true)
		abac.GET(user, "", listUsers, "admin:ListUser", true)
		spUser := user.Group(":user-id", checkUser)
		{
			abac.GET(spUser, "", getUser, "admin:GetUser", true)
			abac.PUT(spUser, "", updateUser, "admin:UpdateUser", true)
			abac.PUT(spUser, "password", updatePassword, "admin:UpdatePassword", true)
			abac.DELETE(spUser, "", deleteUser, "admin:DeleteUser", true)
		}
	}
	group := rg.Group("group")
	{
		abac.POST(group, "", createGroup, "admin:CreateGroup", true)
		abac.GET(group, "", listGroups, "admin:ListGroups", true)
		spGroup := group.Group(":group-id", checkGroup)
		{
			abac.GET(spGroup, "", getGroup, "admin:GetGroup", true)
			abac.PUT(spGroup, "", updateGroup, "admin:UpdateGroup", true)
			abac.DELETE(spGroup, "", deleteGroup, "admin:DeleteGroup", true)
		}
	}
	membership := rg.Group("membership")
	{
		abac.POST(membership, "", createMembership, "admin:CreateMembership", true)
		user := membership.Group("user/:user-id", checkUser)
		{
			abac.GET(user, "", listMembershipsByUser, "admin:ListMembershipsByUser", true)
		}
		indevGroup := membership.Group("group/:group-id", checkGroup)
		{
			abac.GET(indevGroup, "", listMembershipsByGroup, "admin:ListMembershipsByGroup", true)

			indevUser := indevGroup.Group("user/:user-id", checkUser)
			{
				abac.GET(indevUser, "", getMembershipAndPermission, "admin:GetMembership", true)
				abac.PUT(indevUser, "", updateMembership, "admin:UpdateMembership", true)
				abac.DELETE(indevUser, "", deleteMembership, "admin:DeleteMembership", true)
			}
		}
	}
	frozen := rg.Group("frozen/group/:group-id", checkGroup)
	{
		abac.PUT(frozen, "", setFrozenByGroup, "admin:SetFrozenByGroup", true)
		forzenByUser := frozen.Group("user/:user-id", checkUser, checkMembership)
		{
			abac.PUT(forzenByUser, "", setFrozenByUser, "admin:SetFrozenByUser", true)
		}

	}
	credential := rg.Group("credential/group/:group-id/user/:user-id", checkUser, checkGroup, checkMembership)
	{
		abac.PUT(credential, "", updateCredential, "admin:UpdateCredential", true)
		abac.GET(credential, "", getCredential, "admin:GetCredential", true)
	}
}
