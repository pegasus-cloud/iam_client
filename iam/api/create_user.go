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
	user := &User{}
	if err := c.ShouldBindWith(user, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	if user.Force {
		if output, err := iam.GetUser(user.UserID); err != nil {
			utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
				Message: databaseError,
			})
			return
		} else if output.ID != "" {
			utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
				Message: userExistError,
			})
			return
		}
	} else {
		for true {
			user.UserID = utility.GetRandNumeric(16)
			if output, err := iam.GetUser(user.UserID); err != nil {
				utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
					Message: databaseError,
				})
				return
			} else if output.ID == "" {
				break
			}
		}
	}
	if err := iam.CreateUser(&protos.UserInfo{
		ID:           user.UserID,
		DisplayName:  user.DisplayName,
		PasswordHash: user.Password,
		Description:  user.Description,
		Extra:        user.Extra,
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	getUserOutput, err := iam.GetUser(user.UserID)
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	utility.ResponseWithType(c, http.StatusCreated, User{
		UserID:      getUserOutput.ID,
		DisplayName: getUserOutput.DisplayName,
		Description: getUserOutput.Description,
		Extra:       getUserOutput.Extra,
		CreatedAt:   getUserOutput.CreatedAt,
		UpdatedAt:   getUserOutput.UpdatedAt,
	})
}
