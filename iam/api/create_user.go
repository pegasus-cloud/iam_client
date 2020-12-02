package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

func createUser(c *gin.Context) {
	userInfo := &user{}
	if err := c.ShouldBindWith(userInfo, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: utility.ConvertError(err).Error(),
		})
		return
	}
	if userInfo.Force {
		if statusCode, err := checkUserExist(c, userInfo.UserID, false); err != nil {
			utility.ResponseWithType(c, statusCode, &utility.ErrResponse{
				Message: err.Error(),
			})
			return
		}
	} else {
		for true {
			userInfo.UserID = utility.GetRandNumeric(16)
			if output, err := iam.GetUser(&protos.UserID{
				ID: userInfo.UserID,
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
	createUserWithRespOutput, err := iam.CreateUserWithResp(&protos.UserInfo{
		ID:           userInfo.UserID,
		DisplayName:  userInfo.DisplayName,
		PasswordHash: userInfo.Password,
		Description:  userInfo.Description,
		Extra:        userInfo.Extra,
	})
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	utility.ResponseWithType(c, http.StatusCreated, user{
		UserID:      createUserWithRespOutput.ID,
		DisplayName: createUserWithRespOutput.DisplayName,
		Description: createUserWithRespOutput.Description,
		Extra:       createUserWithRespOutput.Extra,
		CreatedAt:   createUserWithRespOutput.CreatedAt,
		UpdatedAt:   createUserWithRespOutput.UpdatedAt,
	})
}
