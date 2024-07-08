package service

import (
	"context"
	"encoding/json"
	pb "github.com/BitofferHub/proto_center/api/xtimer/v1"
	"github.com/BitofferHub/xtimer/internal/biz"
	"github.com/BitofferHub/xtimer/internal/constant"
)

func (s *XTimerService) CreateTimer(ctx context.Context, req *pb.CreateTimerRequest) (*pb.CreateTimerReply, error) {
	param, err := json.Marshal(req.NotifyHTTPParam)
	if err != nil {
		return nil, err
	}
	timer, err := s.timerUC.CreateTimer(ctx, &biz.Timer{
		App:             req.App,
		Name:            req.Name,
		Status:          constant.Unabled.ToInt(),
		Cron:            req.Cron,
		NotifyHTTPParam: string(param),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateTimerReply{Code: 0, Message: "ok", Data: &pb.CreateTimerReplyData{
		TimerId: timer.TimerId,
	}}, nil
}
func (s *XTimerService) EnableTimer(ctx context.Context, req *pb.EnableTimerRequest) (*pb.EnableTimerReply, error) {

	err := s.timerUC.EnableTimer(ctx, req.GetApp(), req.GetTimerId())
	if err != nil {
		return nil, err
	}
	return &pb.EnableTimerReply{Code: 0, Message: "ok"}, nil
}
