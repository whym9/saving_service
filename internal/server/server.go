package server

import (
	"encoding/json"
	"io"
	"os"

	"saving_service/internal/metrics"
	sendpb "saving_service/internal/proto"
	"saving_service/internal/saver"
	"saving_service/internal/storage"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	sendpb.UnimplementedSendServiceServer
	Handle  saver.DB_Handle
	storage storage.Manager
	dir     string
}

func NewServer(storage storage.Manager, dir string) Server {
	return Server{storage: storage, dir: dir}
}

func (s Server) Send(stream sendpb.SendService_SendServer) error {
	metrics.RecordMetrics()
	dirName := "files"
	err := os.MkdirAll(dirName, os.ModePerm)

	if err != nil {

		return err
	}
	counter := saver.Protocols{}

	name := dirName + "/" + time.Now().Format("02-01-1975-547896") + ".pcapng"

	f := storage.NewFile(name)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	req, err := stream.Recv()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	err = json.Unmarshal(req.GetChunk(), &counter)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	for {

		req, err := stream.Recv()

		if err == io.EOF {

			if err = s.storage.Store(f); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
			err = s.Handle.SaveToDB(counter, name)

			if err != nil {

				return status.Error(codes.Internal, err.Error())
			}

			return stream.SendAndClose(&sendpb.SendResponse{Name: "ok"})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err = f.Write(req.GetChunk()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

	}

}
