package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
)

type (
	listUserOutput struct {
		Users []user `json:"users"`
		Total int    `json:"total"`
	}
)

func listUsers(c *gin.Context) {
	listUserOutput := &listUserOutput{}

	pagination := &pagination{}
	if err := c.ShouldBindWith(pagination, binding.Query); err != nil {
		utility.ResponseWithType(c, http.StatusBadRequest, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	users, err := iam.ListUsers(&protos.LimitOffset{
		Limit:  int32(pagination.Limit),
		Offset: int32(pagination.Offset),
	})
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	for _, userInfo := range users.Data {
		listUserOutput.Users = append(listUserOutput.Users, user{
			UserID:      userInfo.ID,
			DisplayName: userInfo.DisplayName,
			Description: userInfo.Description,
			Extra:       userInfo.Extra,
			CreatedAt:   userInfo.CreatedAt,
			UpdatedAt:   userInfo.UpdatedAt,
		})
	}

	listUserOutput.Total = int(users.Count)
	utility.ResponseWithType(c, http.StatusOK, listUserOutput)
}
