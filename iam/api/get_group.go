package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/utility"
)

func getGroup(c *gin.Context) {
	getGroupOutput, err := iam.GetGroup(c.Param(groupIDParams))
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}
	if getGroupOutput.ID == "" {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: groupDoesNotExistErrMsg,
		})
		return
	}
	utility.ResponseWithType(c, http.StatusOK, group{
		GroupID:     getGroupOutput.ID,
		DisplayName: getGroupOutput.DisplayName,
		Description: getGroupOutput.Description,
		Extra:       getGroupOutput.Extra,
		CreatedAt:   getGroupOutput.CreatedAt,
		UpdatedAt:   getGroupOutput.UpdatedAt,
	})
}
