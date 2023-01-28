package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	token, err := h.parseAuthHeader(c)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "authMiddleware",
		}).Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	userId, err := h.userService.ParseToken(c.Request.Context(), token)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "authMiddleware",
		}).Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("Id", userId)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}
