package api

import "encoding/xml"

const (
	databaseErrMsg                = "The database errors occurred."
	userExistErrMsg               = "The UserID is already exist."
	unKnownUserNameOrBadPwdErrMsg = "Unknown username or bad password"
	recordNotFoundErrMsg          = "record not found"

	userIDParams = "user-id"
)

type (
	// User ...
	User struct {
		XMLName     *xml.Name `json:",omitempty" xml:"user"`
		UserID      string    `json:"userId,omitempty" binding:"required_with=Force,omitempty,min=8,max=16" xml:"userId,omitempty"`
		DisplayName string    `json:"displayName,omitempty" binding:"required,max=255" xml:"displayName,omitempty"`
		Password    string    `json:"password,omitempty" binding:"required" xml:"password,omitempty"`
		Description string    `json:"description,omitempty" xml:"description,omitempty"`
		Extra       string    `json:"extra,omitempty" xml:"extra,omitempty"`
		CreatedAt   string    `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
		UpdatedAt   string    `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
		Force       bool      `json:"force,omitempty" binding:"required_with=UserID" xml:"force,omitempty"`
		_           struct{}
	}
)
