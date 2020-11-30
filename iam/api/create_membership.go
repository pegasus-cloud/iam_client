package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

func createMembership(c *gin.Context) {
	membershipInfo := &membership{}
	if err := c.ShouldBindWith(membershipInfo, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: utility.ConvertError(err).Error(),
		})
		return
	}
	if statusCode, err := checkUserExist(c, membershipInfo.UserID, true); err != nil {
		utility.ResponseWithType(c, statusCode, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	if statusCode, err := checkGroupExist(c, membershipInfo.GroupID, true); err != nil {
		utility.ResponseWithType(c, statusCode, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	if statusCode, err := checkMembershipExist(c, membershipInfo.UserID, membershipInfo.GroupID, false); err != nil {
		utility.ResponseWithType(c, statusCode, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	for _, permission := range []string{membershipInfo.GlobalPermissionID, membershipInfo.UserPermissionID} {
		if permission == "" {
			permission = "DEFAULT"
		}
		if statusCode, err := checkPermissionExist(c, permission, membershipInfo.GroupID, true); err != nil {
			utility.ResponseWithType(c, statusCode, &utility.ErrResponse{
				Message: err.Error(),
			})
			return
		}
	}
	input := &protos.MembershipInfo{
		ID:                 utility.GetRandAlphanumericUpper(16),
		GroupID:            membershipInfo.GroupID,
		UserID:             membershipInfo.UserID,
		GlobalPermissionID: membershipInfo.GlobalPermissionID,
		UserPermissionID:   membershipInfo.UserPermissionID,
		Frozen:             membershipInfo.Frozen,
	}
	if membershipInfo.Quota != nil && *membershipInfo.Quota != "{}" && *membershipInfo.Quota != "" {
		quota := &quota{}
		if err := json.Unmarshal([]byte(*membershipInfo.Quota), quota); err != nil {
			utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
				Message: quotaIsInvalid,
			})
			return
		}
		quotaByte, err := json.Marshal(quota)
		if err != nil {
			utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
				Message: err.Error(),
			})
			return
		}
		quotaString := string(quotaByte)
		input.Quota = &quotaString
	}
	createMembershipWithRespOutput, err := iam.CreateMembershipWithResp(input)
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}
	utility.ResponseWithType(c, http.StatusCreated, membership{
		MembershipID: createMembershipWithRespOutput.ID,
		GroupID:      createMembershipWithRespOutput.GroupID,
		UserID:       createMembershipWithRespOutput.UserID,
		GlobalPermission: &permissionInMembership{
			ID:    createMembershipWithRespOutput.GlobalPermissionID,
			Label: createMembershipWithRespOutput.GlobalPermissionLabel,
		},
		UserPermission: &permissionInMembership{
			ID:    createMembershipWithRespOutput.UserPermissionID,
			Label: createMembershipWithRespOutput.UserPermissionLabel,
		},
		Frozen: createMembershipWithRespOutput.Frozen,
		Quota:  input.Quota,
	})
}
