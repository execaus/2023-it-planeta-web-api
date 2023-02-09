package handler

import (
	"2023-it-planeta-web-api/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) registrationAccount(c *gin.Context) {
	var input models.RegistrationAccountInput

	if err := c.BindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
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

	h.sendCreatedWithBody(c, output)
}

func (h *Handler) getAccount(c *gin.Context) {
	stringId := c.Param("id")
	if stringId == "" || stringId == "null" {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	id, err := strconv.Atoi(stringId)
	if err != nil {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	if id <= 0 {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	account, err := h.services.Account.Get(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if account == nil {
		h.sendNotFound(c)
		return
	}

	output := &models.GetAccountOutput{
		Id:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}
	h.sendOKWithBody(c, output)
}
