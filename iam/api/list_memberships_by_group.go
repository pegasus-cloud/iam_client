package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

type (
	listMembershipsByGroupOutput struct {
		Memberships []membership `json:"memberships"`
		Total       int          `json:"total"`
	}
)

func listMembershipsByGroup(c *gin.Context) {
	listMembershipsByGroupOutput := &listMembershipsByGroupOutput{}

	pagination := &pagination{}
	if err := c.ShouldBindWith(pagination, binding.Query); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	memberships, err := iam.ListMembershipsByGroup(&protos.ListMembershipByGroupInput{
		GroupID: c.Param(groupIDParams),
		Data: &protos.LimitOffset{
			Limit:  int32(pagination.Limit),
			Offset: int32(pagination.Offset),
		},
	})
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	for _, membershipInfo := range memberships.Data {
		listMembershipsByGroupOutput.Memberships = append(listMembershipsByGroupOutput.Memberships, membership{
			MembershipID: membershipInfo.ID,
			GroupID:      membershipInfo.GroupID,
			User: &user{
				UserID:      membershipInfo.UserID,
				DisplayName: membershipInfo.DisplayName,
				Description: membershipInfo.Description,
				Extra:       membershipInfo.Extra,
				UpdatedAt:   membershipInfo.UpdatedAt,
				CreatedAt:   membershipInfo.CreatedAt,
			},
			GlobalPermissionID: membershipInfo.GlobalPermissionID,
			UserPermissionID:   membershipInfo.UserPermissionID,
			Frozen:             membershipInfo.Frozen,
		})
	}

	listMembershipsByGroupOutput.Total = int(memberships.Count)
	utility.ResponseWithType(c, http.StatusOK, listMembershipsByGroupOutput)
}
