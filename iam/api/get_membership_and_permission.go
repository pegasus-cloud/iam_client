package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/utility"
)

func getMembershipAndPermission(c *gin.Context) {
	getMembershipAndPermissionOutput, err := iam.GetMembershipAndPermission(c.Param(userIDParams), c.Param(groupIDParams))
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}
	if getMembershipAndPermissionOutput.ID == "" {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: membershipDoesNotExistErrMsg,
		})
		return
	}
	utility.ResponseWithType(c, http.StatusOK, membership{
		MembershipID: getMembershipAndPermissionOutput.ID,
		GroupID:      getMembershipAndPermissionOutput.GroupID,
		UserID:       getMembershipAndPermissionOutput.UserID,
		GlobalPermission: &permissionInMembership{
			ID:    getMembershipAndPermissionOutput.GlobalPermissionID,
			Label: getMembershipAndPermissionOutput.GlobalPermissionLabel,
		},
		UserPermission: &permissionInMembership{
			ID:    getMembershipAndPermissionOutput.UserPermissionID,
			Label: getMembershipAndPermissionOutput.UserPermissionLabel,
		},
		Frozen: getMembershipAndPermissionOutput.Frozen,
		Quota:  getMembershipAndPermissionOutput.Quota,
	})
}
