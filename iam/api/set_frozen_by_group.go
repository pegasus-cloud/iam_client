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
	setFrozenByGroupInput struct {
		Frozen *bool `json:"frozen"`
	}
)

func (sfi *setFrozenByGroupInput) convertToMap() (output map[string]*any.Any) {
	output = make(map[string]*anypb.Any)
	e := reflect.ValueOf(sfi).Elem()

	for i := 0; i < e.NumField(); i++ {
		key, value := e.Type().Field(i).Name, e.Field(i).Interface()
		if value.(*bool) != nil {
			output[key], _ = ptypes.MarshalAny(&protos.GBoolean{Val: value.(*bool)})
		}
	}
	return output
}

func setFrozenByGroup(c *gin.Context) {
	if c.Param(groupIDParams) == adminGroupID {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: adminGroupShouldNotBeFrozen,
		})
		return
	}

	setFrozenByGroupInput := &setFrozenByGroupInput{}

	if err := c.ShouldBindWith(setFrozenByGroupInput, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: utility.ConvertError(err).Error(),
		})
		return
	}

	setFrozenByGroupInputMap := make(map[string]*any.Any)
	setFrozenByGroupInputMap = setFrozenByGroupInput.convertToMap()

	if err := iam.UpdateMembershipByGroup(&protos.UpdateInput{
		ID:   c.Param(groupIDParams),
		Data: setFrozenByGroupInputMap,
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}

	utility.ResponseWithType(c, http.StatusOK, nil)
}
