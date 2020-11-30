package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/utility"
)

type (
	listGroupsOutput struct {
		Groups []group `json:"groups"`
		Total  int     `json:"total"`
	}
)

func listGroups(c *gin.Context) {
	listGroupsOutput := &listGroupsOutput{}

	pagination := &pagination{}
	if err := c.ShouldBindWith(pagination, binding.Query); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	groups, err := iam.ListGroups(pagination.Limit, pagination.Offset)
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	for _, groupInfo := range groups.Data {
		listGroupsOutput.Groups = append(listGroupsOutput.Groups, group{
			GroupID:     groupInfo.ID,
			DisplayName: groupInfo.DisplayName,
			Description: groupInfo.Description,
			Extra:       groupInfo.Extra,
			CreatedAt:   groupInfo.CreatedAt,
			UpdatedAt:   groupInfo.UpdatedAt,
		})
	}

	listGroupsOutput.Total = int(groups.Count)
	utility.ResponseWithType(c, http.StatusOK, listGroupsOutput)
}
