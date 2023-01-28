package rest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my-crud-app/internal/domain"
	"net/http"
)

// @Summary SignUp
// @Description  register user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input body domain.SignUpUser true "createUser"
// @Success      200
// @Failure      404
// @Router       /auth [post]
func (h *Handler) SignUp(c *gin.Context) {
	var inputData domain.SignUpUser
	if err := c.BindJSON(&inputData); err != nil {
		return
	}
	if err := inputData.Validate(); err != nil {
		log.WithFields(log.Fields{
			"handler": "SignUp",
			"problem": err,
		}).Error(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}
	err := h.userService.SignUp(context.TODO(), inputData)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "SignUp",
			"problem": "reading request body",
		}).Error(err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		return
	}
	c.JSON(http.StatusOK, inputData)
}

// @Summary SignIn
// @Description  signIn user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input body domain.SignInUser true "signIn User"
// @Success      200
// @Failure      404
// @Router       /auth [get]
func (h *Handler) SignIn(c *gin.Context) {
	var input domain.SignInUser
	if err := c.BindJSON(&input); err != nil {
		return
	}
	if err := input.Validate(); err != nil {
		log.WithFields(log.Fields{
			"handler": "SignIn",
			"problem": err,
		}).Error(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	token, err := h.userService.SignIn(c.Request.Context(), input)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "SignIn",
			"problem": err,
		}).Error(err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
