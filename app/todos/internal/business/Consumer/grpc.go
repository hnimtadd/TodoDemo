package server

// import (
//
//	"log"
//	"todosService/config"
//	"todosService/internal/business"
//	"todosService/protoc"
//
//	"google.golang.org/grpc"
//
// )
//
//	type GrpcConsumer struct {
//		conn   *grpc.ClientConn
//		repo   business.TodosRepository
//		config *config.GrpcConfig
//		protoc.TodosServiceServer
//	}
//
//	func NewGrpcConsumer(repo business.TodosRepository, config config.GrpcConfig) business.TodoConsumer {
//		grpcConsumer := &GrpcConsumer{
//			repo:   repo,
//			config: &config,
//		}
//		if err := grpcConsumer.Setup(); err != nil {
//			log.Fatalf("Failed to dial...: %v", err)
//		}
//		return grpcConsumer
//	}
//
//	func (sv *GrpcConsumer) Setup() error {
//		conn, err := grpc.Dial(sv.config.Source)
//		if err != nil {
//			return err
//		}
//		sv.conn = conn
//		return nil
//	}
// func (sv *GrpcConsumer) Listen() error {
// 	return nil
// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// defer func() {
// 	cancel()
// 	sv.conn.Close()
// }()
// s := grpc.NewServer()
// protoc.RegisterTodosServiceServer(s, sv)
// fmt.Println("server connecting")
// err = s.Serve(lis)
// return nil
// }
