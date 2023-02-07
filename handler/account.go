package handler

import (
	"2023-it-planeta-web-api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) registrationAccount(c *gin.Context) {
	var input models.RegistrationAccountInput

	if err := c.BindJSON(&input); err != nil {
		h.sendBadRequest(c)
		return
	}

	isAccountExist, err := h.services.Account.IsExist(input.Email)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}
	if isAccountExist {
		h.sendConflict(c)
		return
	}

	output, err := h.services.Account.Registration(&input)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOKWithBody(c, output)
}
