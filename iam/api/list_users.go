package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pegasus-cloud/iam_client/iam"
	"github.com/pegasus-cloud/iam_client/utility"
)

type (
	listUserOutput struct {
		Users []User `json:"users"`
		Total int    `json:"total"`
	}
	// Pagination ...
	pagination struct {
		Limit  int    `json:"limit" form:"limit,default=100" binding:"max=100"`
		Offset int    `json:"offset" form:"offset,default=0" binding:"min=0"`
		Order  string `json:"order" form:"order"`
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

	users, err := iam.ListUser(pagination.Limit, pagination.Offset)
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	for _, user := range users.Data {
		listUserOutput.Users = append(listUserOutput.Users, User{
			UserID:      user.ID,
			DisplayName: user.DisplayName,
			Description: user.Description,
			Extra:       user.Extra,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		})
	}

	total, err := iam.CountUser()
	if err != nil {
		utility.ResponseWithType(c, http.StatusInternalServerError, &utility.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	listUserOutput.Total = int(total.Data)
	c.JSON(http.StatusOK, listUserOutput)
}
