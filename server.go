package main

import (
	"log"
	"net"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
	pb "./public"
	core "./core"
	"google.golang.org/grpc"
)

type LshServiceImpl struct {
	index *core.ForestLSH
}

func (s *LshServiceImpl) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResult, error) {
	out := make(chan int)
	defer close(out)

	id, _ := s.index.Data.Index[in.Cid]
	p := s.index.Data.Points[id]
	s.index.Query(p, 1000, out)
	docs := []*pb.SearchDoc{}
	for {
		key := <- out
		if key == -1 {
			break
		}
		docs = append(docs, &pb.SearchDoc{s.index.Data.Points[key].Id, p.L2(s.index.Data.Points[key]), "", })
	}

	return &pb.SearchResult{docs, int32(len(docs)), }, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8999")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	data, err := core.NewDatasetWithHeader("/Users/dwang/Downloads/cls_rep_vec_1489636166", "cid", "col_")
	if err != nil {
		log.Fatalln("failed loading index")
	}

	params := core.NewParams(data.VecLen, 100, 10, 5.0)
	forestLSH := core.NewForestLSHFromData(params, data)
	serviceImpl := &LshServiceImpl{forestLSH}

	s := grpc.NewServer()
	pb.RegisterLSHServer(s, serviceImpl)
	reflection.Register(s)

	log.Printf("start serving, port%s\n", ":8999")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
