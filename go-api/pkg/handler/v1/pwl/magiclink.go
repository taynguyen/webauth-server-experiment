package pwl

import (
	"github.com/dwarvesf/go-api/pkg/handler/v1/view"
	"github.com/dwarvesf/go-api/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateMagicLink godoc
// @Summary create magic link
// @Description create magic link
// @id magiclink
// @Tags magiclink
// @Accept  json
// @Produce  json
// @Param Body body CreateMagicLink true "Body"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /pwl/user/magic-link/create [post]
func (h Handler) CreateMagicLink(c *gin.Context) {
	const spanName = "createMagicLinkHandler"
	ctx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	var req view.CreateMagicLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "failed to bind request")
		util.HandleError(c, err)
		return
	}

	magicLink, err := h.authCtrl.CreateMagicLink(ctx, req.Email)
	if err != nil {
		h.log.Error(err, "failed to create magic link")
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, view.CreateMagicLinkResponse{
		MagicLink: magicLink,
	})
}

func (h Handler) VerifyMagicLink(c *gin.Context) {
	const spanName = "loginHandler"
	ctx, span := h.monitor.Start(c.Request.Context(), spanName)
	defer span.End()

	var req view.VerifyMagicLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "failed to bind request")
		util.HandleError(c, err)
		return
	}

	token, err := h.authCtrl.VerifyMagicLink(ctx, req.Secret)
	if err != nil {
		h.log.Error(err, "failed to verify magic link")
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, view.VerifyMagicLinkResponse{
		AccessToken: token,
	})
}
