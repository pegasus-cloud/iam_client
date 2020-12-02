package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

func getCredential(c *gin.Context) {
	getCredentialOutput, err := iam.GetCredential(&protos.CredUserGroupInput{
		UserID:  c.Param(userIDParams),
		GroupID: c.Param(groupIDParams),
	})
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}

	utility.ResponseWithType(c, http.StatusOK, credential{
		AccessKey:    getCredentialOutput.Access,
		SecretKey:    getCredentialOutput.Secret,
		MembershipID: getCredentialOutput.MembershipID,
		GroupID:      getCredentialOutput.GroupID,
		UserID:       getCredentialOutput.UserID,
	})
}
