package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

func deleteGroup(c *gin.Context) {
	if err := iam.DeleteGroup(&protos.GroupID{
		ID: c.Param(groupIDParams),
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: databaseErrMsg,
		})
		return
	}
	utility.ResponseWithType(c, http.StatusNoContent, nil)
}
