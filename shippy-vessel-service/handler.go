package main

import (
	"context"
	pb "github.com/CcccFz/shippy/shippy-vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

// 实现微服务的服务端
type handler struct {
	session *mgo.Session
}

func (h *handler) GetRepo() Repository {
	return &VesselRepository{h.session.Clone()}
}

func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	defer h.GetRepo().Close()
	v, err := h.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}
	resp.Vessel = v
	return nil
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	defer h.GetRepo().Close()
	if err := h.GetRepo().Create(req); err != nil {
		return err
	}
	resp.Vessel = req
	resp.Created = true
	return nil
}
