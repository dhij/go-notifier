package apisvc

import (
	"context"
	"net/http"

	"github.com/dhij/go-notifier"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	grpcClient notifier.NotifierClient
}

func New(client notifier.NotifierClient) *Handler {
	return &Handler{
		grpcClient: client,
	}
}

func (h *Handler) Notify(c *gin.Context) {
	var n NotifyReq
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := &notifier.EnqueueNotificationEventReq{
		UserUuid: n.UserUUID,
		Message:  n.Message,
	}

	_, err := h.grpcClient.EnqueueNotificationEvent(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
