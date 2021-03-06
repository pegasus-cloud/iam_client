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
	updateGroupInput struct {
		DisplayName *string `json:"displayName" binding:"omitempty,min=1,max=255"`
		Description *string `json:"description"`
		Extra       *string `json:"extra"`
	}
)

func (uui *updateGroupInput) convertToMap() (output map[string]*any.Any) {
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

func updateGroup(c *gin.Context) {
	updateGroupInput := &updateGroupInput{}
	if err := c.ShouldBindWith(updateGroupInput, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: utility.ConvertError(err).Error(),
		})
		return
	}
	updateGroupInputMap := make(map[string]*any.Any)
	updateGroupInputMap = updateGroupInput.convertToMap()
	updateGroupWithRespOutput, err := iam.UpdateGroupWithResp(&protos.UpdateInput{
		ID:   c.Param(groupIDParams),
		Data: updateGroupInputMap,
	})
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}
	utility.ResponseWithType(c, http.StatusOK, group{
		GroupID:     updateGroupWithRespOutput.ID,
		DisplayName: updateGroupWithRespOutput.DisplayName,
		Description: updateGroupWithRespOutput.Description,
		Extra:       updateGroupWithRespOutput.Extra,
		CreatedAt:   updateGroupWithRespOutput.CreatedAt,
		UpdatedAt:   updateGroupWithRespOutput.UpdatedAt,
	})
}
