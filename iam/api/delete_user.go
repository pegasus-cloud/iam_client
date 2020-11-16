package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

func deleteUser(c *gin.Context) {
	if err := iam.DeleteUser(&protos.UserID{
		ID: c.Param(userIDParams),
	}); err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: databaseErrMsg,
		})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
