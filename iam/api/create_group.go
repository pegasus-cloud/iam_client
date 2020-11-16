package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

func createGroup(c *gin.Context) {
	groupInfo := &group{}
	if err := c.ShouldBindWith(groupInfo, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	if groupInfo.Force {
		if output, err := iam.GetGroup(groupInfo.GroupID); err != nil {
			utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
				Message: databaseErrMsg,
			})
			return
		} else if output.ID != "" {
			utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
				Message: groupExistErrMsg,
			})
			return
		}
	} else {
		for true {
			groupInfo.GroupID = utility.GetRandNumeric(16)
			if output, err := iam.GetGroup(groupInfo.GroupID); err != nil {
				utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
					Message: databaseErrMsg,
				})
				return
			} else if output.ID == "" {
				break
			}
		}
	}
	if err := iam.CreateGroup(&protos.GroupInfo{
		ID:          groupInfo.GroupID,
		DisplayName: groupInfo.DisplayName,
		Description: groupInfo.Description,
		Extra:       groupInfo.Extra,
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	getGroupOutput, err := iam.GetGroup(groupInfo.GroupID)
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	utility.ResponseWithType(c, http.StatusCreated, group{
		GroupID:     getGroupOutput.ID,
		DisplayName: getGroupOutput.DisplayName,
		Description: getGroupOutput.Description,
		Extra:       getGroupOutput.Extra,
		CreatedAt:   getGroupOutput.CreatedAt,
		UpdatedAt:   getGroupOutput.UpdatedAt,
	})
}
