package handlers

import (
	av "av-service/src/proto"
	"context"
)

type AvServer struct {}

func (this *AvServer) GetAVList(ctx context.Context, req *av.GetAVListReq) (*av.GetAVListRep, error) {
	return &av.GetAVListRep{
		Total: 10,
		Datas: []*av.AV{&av.AV{Id:1, Title: "Hello"},},
	}, nil
}