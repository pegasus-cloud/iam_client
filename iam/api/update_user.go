package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
	"google.golang.org/protobuf/types/known/anypb"
)

type (
	updateUserInput struct {
		DisplayName *string `json:"displayName" binding:"omitempty,min=1,max=255"`
		Description *string `json:"description"`
		Extra       *string `json:"extra"`
	}
)

func (uui *updateUserInput) convertToMap() (output map[string]*any.Any) {
	output = make(map[string]*anypb.Any)
	e := reflect.ValueOf(uui).Elem()
	for i := 0; i < e.NumField(); i++ {
		key, value := e.Type().Field(i).Name, e.Field(i).Interface()
		if value.(*string) != nil {
			output[key], _ = ptypes.MarshalAny(&protos.GString{Val: value.(*string)})
		}
	}
	return output
}

func updateUser(c *gin.Context) {
	updateUserInput := &updateUserInput{}
	if err := c.ShouldBindWith(updateUserInput, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: utility.ConvertError(err).Error(),
		})
		return
	}
	updateUserInputMap := make(map[string]*any.Any)
	updateUserInputMap = updateUserInput.convertToMap()
	if err := iam.UpdateUser(&protos.UpdateInput{
		ID:   c.Param(userIDParams),
		Data: updateUserInputMap,
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	getUserOutput, err := iam.GetUser(&protos.UserID{
		ID: c.Param(userIDParams),
	})
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
	}
	utility.ResponseWithType(c, http.StatusOK, user{
		UserID:      getUserOutput.ID,
		DisplayName: getUserOutput.DisplayName,
		Description: getUserOutput.Description,
		Extra:       getUserOutput.Extra,
		CreatedAt:   getUserOutput.CreatedAt,
		UpdatedAt:   getUserOutput.UpdatedAt,
	})
}
