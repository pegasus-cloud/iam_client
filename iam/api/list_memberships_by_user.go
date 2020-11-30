package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/utility"
)

type (
	listMembershipsByUserOutput struct {
		Memberships []membership `json:"memberships"`
		Total       int          `json:"total"`
	}
)

func listMembershipsByUser(c *gin.Context) {
	listMembershipsByUserOutput := &listMembershipsByUserOutput{}

	pagination := &pagination{}
	if err := c.ShouldBindWith(pagination, binding.Query); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	memberships, err := iam.ListMembershipsByUser(c.Param(userIDParams), pagination.Limit, pagination.Offset)
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	for _, membershipInfo := range memberships.Data {
		listMembershipsByUserOutput.Memberships = append(listMembershipsByUserOutput.Memberships, membership{
			MembershipID: membershipInfo.ID,
			UserID:       membershipInfo.UserID,
			Group: &group{
				GroupID:     membershipInfo.GroupID,
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

	listMembershipsByUserOutput.Total = int(memberships.Count)
	utility.ResponseWithType(c, http.StatusOK, listMembershipsByUserOutput)
}
