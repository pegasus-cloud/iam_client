package api

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/utility"
)

func listPermissionActions(c *gin.Context) {
	actions := iam.Actions.GetActions()
	sort.Strings(actions)
	utility.ResponseWithType(c, http.StatusOK, actions)
}
