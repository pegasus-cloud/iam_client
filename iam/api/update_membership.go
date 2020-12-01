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
	updateMembershipInput struct {
		GlobalPermissionID *string `json:"globalPermissionId"`
		UserPermissionID   *string `json:"userPermissionId"`
		Quota              *string `json:"quota"`
	}
)

func (umi *updateMembershipInput) convertToMap() (output map[string]*any.Any) {
	output = make(map[string]*anypb.Any)
	e := reflect.ValueOf(umi).Elem()

	for i := 0; i < e.NumField(); i++ {
		key, value := e.Type().Field(i).Name, e.Field(i).Interface()
		if value.(*string) != nil {
			output[key], _ = ptypes.MarshalAny(&protos.GString{Val: value.(*string)})
		}
	}
	return output
}

func updateMembership(c *gin.Context) {
	updateMembershipInput := &updateMembershipInput{}

	if err := c.ShouldBindWith(updateMembershipInput, binding.JSON); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: utility.ConvertError(err).Error(),
		})
		return
	}

	updateMembershipInputMap := make(map[string]*any.Any)
	updateMembershipInputMap = updateMembershipInput.convertToMap()

	if err := iam.UpdateMembership(&protos.UpdateMembershipInput{
		UserID:  c.Param(userIDParams),
		GroupID: c.Param(groupIDParams),
		Data:    updateMembershipInputMap,
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}

	getMembershipAndPermissionOutput, err := iam.GetMembershipAndPermission(c.Param(userIDParams), c.Param(groupIDParams))
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}

	utility.ResponseWithType(c, http.StatusOK, membership{
		MembershipID: getMembershipAndPermissionOutput.ID,
		GroupID:      getMembershipAndPermissionOutput.GroupID,
		UserID:       getMembershipAndPermissionOutput.UserID,
		GlobalPermission: &permissionInMembership{
			ID:    getMembershipAndPermissionOutput.GlobalPermissionID,
			Label: getMembershipAndPermissionOutput.GlobalPermissionLabel,
		},
		UserPermission: &permissionInMembership{
			ID:    getMembershipAndPermissionOutput.UserPermissionID,
			Label: getMembershipAndPermissionOutput.UserPermissionLabel,
		},
		Frozen: getMembershipAndPermissionOutput.Frozen,
		Quota:  getMembershipAndPermissionOutput.Quota,
	})
}
