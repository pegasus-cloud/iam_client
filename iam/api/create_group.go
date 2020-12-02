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
			Message: utility.ConvertError(err).Error(),
		})
		return
	}
	if groupInfo.Force {
		if statusCode, err := checkGroupExist(c, groupInfo.GroupID, false); err != nil {
			utility.ResponseWithType(c, statusCode, &utility.ErrResponse{
				Message: err.Error(),
			})
			return
		}
	} else {
		for true {
			groupInfo.GroupID = utility.GetRandNumeric(16)
			if output, err := iam.GetGroup(&protos.GroupID{
				ID: c.Param(groupIDParams),
			}); err != nil {
				utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
					Message: iamServerErrMsg,
				})
				return
			} else if output.ID == "" {
				break
			}
		}
	}
	createGroupWithRespOutput, err := iam.CreateGroupWithResp(&protos.GroupInfo{
		ID:          groupInfo.GroupID,
		DisplayName: groupInfo.DisplayName,
		Description: groupInfo.Description,
		Extra:       groupInfo.Extra,
	})
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	utility.ResponseWithType(c, http.StatusCreated, group{
		GroupID:     createGroupWithRespOutput.ID,
		DisplayName: createGroupWithRespOutput.DisplayName,
		Description: createGroupWithRespOutput.Description,
		Extra:       createGroupWithRespOutput.Extra,
		CreatedAt:   createGroupWithRespOutput.CreatedAt,
		UpdatedAt:   createGroupWithRespOutput.UpdatedAt,
	})
}
