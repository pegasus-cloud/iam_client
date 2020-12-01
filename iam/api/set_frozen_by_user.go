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
	setFrozenByUserInput struct {
		Frozen *bool `json:"frozen"`
	}
)

func (sfi *setFrozenByUserInput) convertToMap() (output map[string]*any.Any) {
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

func setFrozenByUser(c *gin.Context) {
	if c.Param(groupIDParams) == adminGroupID {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: adminGroupShouldNotBeFrozen,
		})
		return
	}

	if c.Param(userIDParams) == adminUserID || c.Param(userIDParams) == systemUserID {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: adminAndSystemUserShouldNotBeFrozen,
		})
		return
	}

	setFrozenByUserInput := &setFrozenByUserInput{}

	if err := c.ShouldBindWith(setFrozenByUserInput, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: utility.ConvertError(err).Error(),
		})
		return
	}

	setFrozenByUserInputMap := make(map[string]*any.Any)
	setFrozenByUserInputMap = setFrozenByUserInput.convertToMap()

	if err := iam.UpdateMembership(&protos.UpdateMembershipInput{
		UserID:  c.Param(userIDParams),
		GroupID: c.Param(groupIDParams),
		Data:    setFrozenByUserInputMap,
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}

	utility.ResponseWithType(c, http.StatusOK, nil)
}
