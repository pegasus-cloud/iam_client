package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

type (
	updatePasswordInput struct {
		Password string `json:"password" binding:"required"`
	}
)

func updatePassword(c *gin.Context) {
	updatePasswordInput := &updatePasswordInput{}

	if err := c.ShouldBindWith(updatePasswordInput, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	updatePasswordInputMap := make(map[string]*any.Any)

	password, err := ptypes.MarshalAny(&protos.GString{Val: &updatePasswordInput.Password})
	if err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	updatePasswordInputMap["Password_Hash"] = password

	if err := iam.UpdateUser(&protos.UpdateInput{
		ID:   c.Param(userIDParams),
		Data: updatePasswordInputMap,
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	utility.ResponseWithType(c, http.StatusOK, nil)
}
