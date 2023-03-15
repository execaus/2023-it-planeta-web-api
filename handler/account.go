package handler

import (
	"2023-it-planeta-web-api/models"
	"2023-it-planeta-web-api/utils"
	"github.com/gin-gonic/gin"
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
	id, err := utils.GetNumberParam(c, "accountId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
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
	animalID, err := utils.GetNumberParam(c, "accountId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Account.IsExistByID(animalID)
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

	// 403 Обновление не своего аккаунта
	email, err := h.getAccountContext(c)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if input.Email != email {
		h.sendForbidden(c)
		return
	}

	isExist, err = h.services.Account.IsExistByEmailExcept(input.Email, animalID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isExist {
		h.sendConflict(c)
		return
	}

	account, err := h.services.Account.Update(animalID, &input)
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

func (h *Handler) deleteAccount(c *gin.Context) {
	accountID, err := utils.GetNumberParam(c, "accountId")
	if err != nil {
		h.sendBadRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Account.IsExistByID(accountID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if !isExist {
		h.sendForbidden(c)
		return
	}

	// 400 Аккаунт связан с животным
	isLinkedAnimal, err := h.services.Account.IsLinkedAnimal(accountID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if isLinkedAnimal {
		h.sendBadRequest(c, "the account is associated with an animal")
		return
	}

	// 403 Удаление не своего аккаунта
	email, err := h.getAccountContext(c)
	if err != nil {
		h.sendUnauthorized(c)
		return
	}

	account, err := h.services.Account.Get(accountID)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if account.Email != email {
		h.sendForbidden(c)
		return
	}

	if err = h.services.Account.Remove(accountID); err != nil {
		h.sendInternalServerError(c)
		return
	}

	h.sendOk(c)
}
