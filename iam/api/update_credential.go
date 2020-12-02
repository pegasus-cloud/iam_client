package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
	"google.golang.org/protobuf/types/known/anypb"
)

type (
	updateCredentialInput struct {
		Access *string `json:"access_key"`
		Secret *string `json:"secret_key"`
	}
)

func (uci *updateCredentialInput) convertToMap() (output map[string]*any.Any) {
	output = make(map[string]*anypb.Any)
	e := reflect.ValueOf(uci).Elem()
	for i := 0; i < e.NumField(); i++ {
		key, value := e.Type().Field(i).Name, e.Field(i).Interface()
		if value.(*string) != nil {
			output[key], _ = ptypes.MarshalAny(&protos.GString{Val: value.(*string)})
		}
	}
	return output
}

func updateCredential(c *gin.Context) {
	accessKey, secretKey := utility.GetRandAlphanumericUpper(24), utility.GetRandAlphanumeric(40)
	updateCredentialInput := &updateCredentialInput{
		Access: &accessKey,
		Secret: &secretKey,
	}
	updateCredentialInputMap := make(map[string]*any.Any)
	updateCredentialInputMap = updateCredentialInput.convertToMap()
	if err := iam.UpdateCredential(&protos.UpdateInput{
		ID:   c.MustGet(membershipIDParams).(string),
		Data: updateCredentialInputMap,
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}
	getCredentialOutput, err := iam.GetCredential(&protos.CredUserGroupInput{
		UserID:  c.Param(userIDParams),
		GroupID: c.Param(groupIDParams),
	})
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
	}
	utility.ResponseWithType(c, http.StatusOK, credential{
		AccessKey:    getCredentialOutput.Access,
		SecretKey:    getCredentialOutput.Secret,
		MembershipID: getCredentialOutput.MembershipID,
		GroupID:      getCredentialOutput.GroupID,
		UserID:       getCredentialOutput.UserID,
	})
}
