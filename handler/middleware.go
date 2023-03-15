package handler

import (
	"2023-it-planeta-web-api/constants"
	"encoding/base64"
	"errors"
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *Handler) requiredAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		h.sendUnauthorized(c)
		return
	}

	if !strings.HasPrefix(authHeader, "Basic ") {
		h.sendUnauthorized(c)
		return
	}

	encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		h.sendUnauthorized(c)
		return
	}

	credentials := string(decodedCredentials)
	splitCredentials := strings.SplitN(credentials, ":", 2)
	login := splitCredentials[0]
	password := splitCredentials[1]

	account, err := h.services.Account.Auth(login, password)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if account == nil {
		h.sendUnauthorized(c)
		return
	}

	h.setAccountContext(c, account.Email)

	c.Next()
}

func (h *Handler) notAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.Next()
		return
	}

	if !strings.HasPrefix(authHeader, "Basic ") {
		h.sendUnauthorized(c)
		return
	}

	encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		h.sendUnauthorized(c)
		return
	}

	credentials := string(decodedCredentials)
	splitCredentials := strings.SplitN(credentials, ":", 2)
	login := splitCredentials[0]
	password := splitCredentials[1]

	account, err := h.services.Account.Auth(login, password)
	if err != nil {
		h.sendInternalServerError(c)
		return
	}

	if account == nil {
		h.sendUnauthorized(c)
		return
	}

	h.setAccountContext(c, account.Email)

	c.Next()
}

func (h *Handler) setAccountContext(c *gin.Context, email string) {
	c.Set(constants.AccountContextKey, email)
}

func (h *Handler) getAccountContext(c *gin.Context) (string, error) {
	email := c.GetString(constants.AccountContextKey)
	if email == "" {
		exloggo.Error("account context not found")
		return "", errors.New("account context not found")
	}
	return email, nil
}
