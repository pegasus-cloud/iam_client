package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

func deleteMembership(c *gin.Context) {
	if err := iam.DeleteMembership(&protos.MemUserGroupInput{
		UserID:  c.Param(userIDParams),
		GroupID: c.Param(groupIDParams),
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: iamServerErrMsg,
		})
		return
	}
	utility.ResponseWithType(c, http.StatusNoContent, nil)
}
