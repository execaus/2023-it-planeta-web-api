package handler

import (
	"encoding/base64"
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

	c.Next()
}
