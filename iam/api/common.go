package api

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
)

const (
	iamServerErrMsg              = "The iam server errors occurred"
	inputFormatErrMsg            = "The body only JSON supported"
	userExistErrMsg              = "The user is already exist"
	userDoesNotExistErrMsg       = "The user does not exist"
	groupExistErrMsg             = "The group is already exist"
	groupDoesNotExistErrMsg      = "The group does not exist"
	membershipExistErrMsg        = "The membership is already exist"
	membershipDoesNotExistErrMsg = "The membership does not exist"
	permissionExistErrMsg        = "The permission is already exist"
	permissionDoesNotExistErrMsg = "The permission does not exist"
	quotaIsInvalid               = "The quota is invalid"

	userIDParams  = "user-id"
	groupIDParams = "group-id"
)

type (
	user struct {
		XMLName     *xml.Name `json:",omitempty" xml:"user"`
		UserID      string    `json:"userId" binding:"required_with=Force,omitempty,min=8,max=16" xml:"userId"`
		DisplayName string    `json:"displayName" binding:"required,max=255" xml:"displayName"`
		Password    string    `json:"password,omitempty" binding:"required" xml:"password,omitempty"`
		Description string    `json:"description" xml:"description"`
		Extra       string    `json:"extra" xml:"extra"`
		CreatedAt   string    `json:"createdAt" xml:"createdAt"`
		UpdatedAt   string    `json:"updatedAt" xml:"updatedAt"`
		Force       bool      `json:"force" binding:"required_with" xml:"force"`
		_           struct{}
	}
	group struct {
		XMLName     *xml.Name `json:",omitempty" xml:"group"`
		GroupID     string    `json:"groupId" binding:"required_with=Force,omitempty,min=8,max=16" xml:"groupID"`
		DisplayName string    `json:"displayName" binding:"required,max=255" xml:"displayName"`
		Description string    `json:"description" xml:"description"`
		Extra       string    `json:"extra" xml:"extra"`
		CreatedAt   string    `json:"createAt" xml:"createdAt"`
		UpdatedAt   string    `json:"updatedAt" xml:"updatedAt"`
		Force       bool      `json:"force" binding:"required_with=GroupID" xml:"force"`
		_           struct{}
	}
	pagination struct {
		Limit  int    `json:"limit" form:"limit,default=100" binding:"max=100"`
		Offset int    `json:"offset" form:"offset,default=0" binding:"min=0"`
		Order  string `json:"order" form:"order"`
		_      struct{}
	}
	membership struct {
		MembershipID       string                  `json:"membershipId"`
		GroupID            string                  `json:"groupId,omitempty" binding:"required"`
		Group              *group                  `json:"group,omitempty"`
		UserID             string                  `json:"userId,omitempty" binding:"required"`
		User               *user                   `json:"user,omitempty"`
		GlobalPermissionID string                  `json:"globalPermissionId,omitempty"`
		GlobalPermission   *permissionInMembership `json:"globalPermission,omitempty"`
		UserPermissionID   string                  `json:"userPermissionId,omitempty"`
		UserPermission     *permissionInMembership `json:"userPermission,omitempty"`
		Frozen             bool                    `json:"frozen"`
		Quota              *string                 `json:"quota,omitempty"`
		_                  struct{}
	}
	permissionInMembership struct {
		ID    string `json:"id"`
		Label string `json:"label"`
		_     struct{}
	}
	quota struct {
		Triton tritonQuota `json:"triton,omitempty"`
		_      struct{}
	}
	tritonQuota struct {
		MaxTopicCount        int `json:"max_topic_count,omitempty"`
		MaxSubscriptionCount int `json:"max_subscription_count,omitempty"`
		MaxQueueCount        int `json:"max_queue_count,omitempty"`
		_                    struct{}
	}
	actions struct {
		Actions []string `json:"actions" xml:"actions"`
		_       struct{}
	}
)

func checkUserExist(c *gin.Context, userID string, expectExist bool) (statusCode int, err error) {
	getUserOutput, err := iam.GetUser(userID)
	if err != nil {
		return http.StatusInternalServerError, errors.New(iamServerErrMsg)
	}
	if getUserOutput.ID == "" && expectExist {
		return http.StatusBadRequest, errors.New(userDoesNotExistErrMsg)
	} else if getUserOutput.ID != "" && !expectExist {
		return http.StatusBadRequest, errors.New(userExistErrMsg)
	}
	return http.StatusOK, nil
}

func checkGroupExist(c *gin.Context, groupID string, expectExist bool) (statusCode int, err error) {
	getGroupOutput, err := iam.GetGroup(groupID)
	if err != nil {
		return http.StatusInternalServerError, errors.New(iamServerErrMsg)
	}
	if getGroupOutput.ID == "" && expectExist {
		return http.StatusBadRequest, errors.New(groupDoesNotExistErrMsg)
	} else if getGroupOutput.ID != "" && !expectExist {
		return http.StatusBadRequest, errors.New(groupExistErrMsg)
	}
	return http.StatusOK, nil
}

func checkMembershipExist(c *gin.Context, userID, groupID string, expectExist bool) (statusCode int, err error) {
	getMembershipOutput, err := iam.GetMembership(userID, groupID)
	if err != nil {
		return http.StatusInternalServerError, errors.New(iamServerErrMsg)
	}
	if getMembershipOutput.ID == "" && expectExist {
		return http.StatusBadRequest, errors.New(membershipDoesNotExistErrMsg)
	} else if getMembershipOutput.ID != "" && !expectExist {
		return http.StatusBadRequest, errors.New(membershipExistErrMsg)
	}
	return http.StatusOK, nil
}

func checkPermissionExist(c *gin.Context, permissionID, groupID string, expectExist bool) (statusCode int, err error) {
	permissionExist, err := iam.CheckPermissionByGroup(&protos.PermissionGroupInput{
		PermissionID: permissionID,
		GroupID:      groupID,
	})
	fmt.Println(permissionID, groupID)
	fmt.Println(permissionExist.GetVal())
	if err != nil {
		return http.StatusInternalServerError, errors.New(iamServerErrMsg)
	}
	if !permissionExist.GetVal() && expectExist {
		return http.StatusBadRequest, errors.New(permissionDoesNotExistErrMsg)
	} else if permissionExist.GetVal() && !expectExist {
		return http.StatusBadRequest, errors.New(permissionExistErrMsg)
	}
	return http.StatusOK, nil
}
