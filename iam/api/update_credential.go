package api

import "github.com/gin-gonic/gin"

type (
	credential struct {
		AccessKey    string `json:"access_key"`
		SecretKey    string `json:"secret_key"`
		MembershipID string `json:"membership_id"`
		_            struct{}
	}
)

func updateCredential(c *gin.Context) {

}
