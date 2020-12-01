package api

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/utility"
)

func listPermissionActions(c *gin.Context) {
	getActions := iam.Actions.GetActions()
	sort.Strings(getActions)
	utility.ResponseWithType(c, http.StatusOK, &actions{
		Actions: getActions,
	})
}
