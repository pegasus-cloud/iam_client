package api

import "encoding/xml"

const (
	databaseErrMsg          = "The database errors occurred."
	userExistErrMsg         = "The user is already exist"
	userDoesNotExistErrMsg  = "The user does not exist"
	groupExistErrMsg        = "The group is already exist"
	groupDoesNotExistErrMsg = "The group does not exist"

	userIDParams  = "user-id"
	groupIDParams = "group-id"
)

type (
	user struct {
		XMLName     *xml.Name `json:",omitempty" xml:"user"`
		UserID      string    `json:"userId" binding:"required_with=Force,min=8,max=16" xml:"userId,"`
		DisplayName string    `json:"displayName" binding:"required,max=255" xml:"displayName,"`
		Password    string    `json:"password" binding:"required" xml:"password,"`
		Description string    `json:"description" xml:"description,"`
		Extra       string    `json:"extra" xml:"extra,"`
		CreatedAt   string    `json:"createdAt" xml:"createdAt,"`
		UpdatedAt   string    `json:"updatedAt" xml:"updatedAt,"`
		Force       bool      `json:"force,omitempty" binding:"required_with" xml:"force,"`
		_           struct{}
	}
	group struct {
		XMLName     *xml.Name `json:",omitempty" xml:"group"`
		GroupID     string    `json:"groupId" binding:"required_with=Force,min=8,max=16" xml:"groupID"`
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
	}
)
