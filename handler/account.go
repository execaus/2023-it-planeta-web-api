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

	isAccountExist, err := h.services.Account.IsExistByEmail(input.Email)
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
	stringID := c.Param("accountId")
	if stringID == stringEmpty || stringID == stringNull {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	id, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	if id <= 0 {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	isExist, err := h.services.Account.IsExistByID(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendNotFound(c)
		return
	}

	account, err := h.services.Account.Get(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := &models.GetAccountOutput{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}
	h.sendOKWithBody(c, output)
}

func (h *Handler) getAccounts(c *gin.Context) {
	var input models.GetAccountsInput
	if err := c.ShouldBindQuery(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	if input.From != nil && *input.From < 0 {
		h.sendBadRequest(c, "invalid parameter from")
		return
	}

	if input.Size != nil && *input.Size <= 0 {
		h.sendBadRequest(c, "invalid parameter size")
		return
	}

	accounts, err := h.services.Account.GetList(&input)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOKWithBody(c, accounts)
}

func (h *Handler) updateAccount(c *gin.Context) {
	stringID := c.Param("accountId")
	if stringID == stringEmpty || stringID == stringNull {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	id, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	if id <= 0 {
		h.sendBadRequest(c, "id is not valid")
		return
	}

	isExist, err := h.services.Account.IsExistByID(id)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendForbidden(c)
		return
	}

	var input models.UpdateAccountInput
	if err = c.ShouldBindJSON(&input); err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	// TODO Обновление не своего аккаунта, sendForbidden

	isExist, err = h.services.Account.IsExistByEmail(input.Email)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isExist {
		h.sendConflict(c)
		return
	}

	account, err := h.services.Account.Update(id, &input)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	output := &models.UpdateAccountOutput{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}
	h.sendOKWithBody(c, output)
}
