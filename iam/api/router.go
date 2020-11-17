package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
)

//EnableAdminIAMRouter 啟動預設的IAM Routers
func EnableAdminIAMRouter(rg *gin.RouterGroup) {
	user := rg.Group("user")
	{
		iam.POST(user, "", createUser, "admin:CreateUser", true)
		iam.GET(user, "", listUsers, "admin:ListUser", true)
		spUser := user.Group(":user-id") // TODO: Need to check that the user has been created
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
		spGroup := group.Group(":group-id") // TODO: Need to check that the group has been created
		{
			iam.GET(spGroup, "", getGroup, "admin:GetGroup", true)
			iam.PUT(spGroup, "", updateGroup, "admin:UpdateGroup", true)
			iam.DELETE(spGroup, "", deleteGroup, "admin:DeleteGroup", true)
		}
	}
	member := rg.Group("members")
	{
		iam.POST(member, "", createMembership, "admin:CreateMembership", true)
		user := member.Group("user/:user-id") // TODO: Need to check that the user has been created
		{
			iam.GET(user, "groups", listGroupsByMembership, "admin:ListGroupsByMembership", true)
		}
		indevGroup := member.Group("group/:group-id") // TODO: Need to check that the group has been created
		{
			iam.GET(indevGroup, "users", listMembershipsByGroup, "admin:ListMembershipsByGroup", true)

			indevUser := indevGroup.Group("user/:user-id") // TODO: Need to check that the membership has been created
			{
				iam.GET(indevUser, "", getMembership, "admin:GetMembership", true)
				iam.PUT(indevUser, "", updateMembership, "admin:UpdateMembership", true)
				iam.DELETE(indevUser, "", deleteMembership, "admin:DeleteMembership", true)
			}
		}
	}
}
