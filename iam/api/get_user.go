package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/utility"
)

func getUser(c *gin.Context) {
	getUserOutput, err := iam.GetUser(c.Param(userIDParams))
	if err != nil {
		if err.Error() == recordNotFoundErrMsg {
			utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
				Message: unKnownUserNameOrBadPwdErrMsg,
			})
			return
		}
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: databaseErrMsg,
		})
		return
	}
	c.JSON(http.StatusOK, User{
		UserID:      getUserOutput.ID,
		DisplayName: getUserOutput.DisplayName,
		Description: getUserOutput.Description,
		Extra:       getUserOutput.Extra,
		CreatedAt:   getUserOutput.CreatedAt,
		UpdatedAt:   getUserOutput.UpdatedAt,
	})
}
