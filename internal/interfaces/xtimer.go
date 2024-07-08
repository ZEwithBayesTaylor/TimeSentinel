package interfaces

import (
	"context"
	"fmt"
	"github.com/BitofferHub/pkg/constant"
	"github.com/BitofferHub/pkg/middlewares/log"
	pb "github.com/BitofferHub/proto_center/api/xtimer/v1"
	"github.com/BitofferHub/xtimer/internal/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) CreateTimer(c *gin.Context) {
	traceID := c.Request.Header.Get(constant.TraceID)

	var req pb.CreateTimerRequest

	if err := c.ShouldBind(&req); err != nil {
		log.Errorf("CreateTimer err: %v", err)
		response.Fail(c, response.ParamError, nil)
		return
	}

	ctx := context.WithValue(context.Background(), constant.TraceID, traceID)
	resp, err := h.xTimerService.CreateTimer(ctx, &req)
	if err != nil {
		fmt.Println("Create timer err: %v", err)
		response.Fail(c, response.ParamError, nil)
		return
	}

	response.Success(c, resp)
}

func (h *Handler) EnableTimer(c *gin.Context) {
	traceID := c.Request.Header.Get(constant.TraceID)
	timerId, err := strconv.ParseInt(c.Query("timerId"), 10, 64)
	if err != nil {
		log.Errorf("EnableTimer err: %v", err)
		response.Fail(c, response.ParamError, nil)
		return
	}
	req := pb.EnableTimerRequest{
		TimerId: timerId,
		App:     c.Query("app"),
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		log.Errorf("EnableTimer err: %v", err)
		response.Fail(c, response.ParamError, nil)
		return
	}

	ctx := context.WithValue(context.Background(), constant.TraceID, traceID)
	_, err = h.xTimerService.EnableTimer(ctx, &req)
	if err != nil {
		log.Errorf("EnableTimer err: %v", err)
		response.Fail(c, response.EnableTimerError, nil)
		return

	}
	response.Success(c, nil)
	return
}

func (h *Handler) TestCallback(c *gin.Context) {
	log.Info("callback test: %v", c.Request.Body)

	response.Success(c, "ok: callback receives")
}
