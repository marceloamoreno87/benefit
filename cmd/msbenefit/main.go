package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/marceloamoreno87/benefit/internal/benefit"
	"github.com/marceloamoreno87/benefit/internal/token"
	msbenefit "github.com/marceloamoreno87/benefit/proto"
	pb "github.com/marceloamoreno87/benefit/proto"
	"google.golang.org/grpc"
)

type BenefitServer struct {
	pb.UnimplementedBenefitServiceServer
}

func main() {

	godotenv.Load()

	lis, err := net.Listen("tcp", os.Getenv("PORT"))

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterBenefitServiceServer(s, &BenefitServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func (s *BenefitServer) GetBenefit(ctx context.Context, in *pb.NewBenefit) (*pb.User, error) {
	res := scrape(in.GetDoc())

	benefits := &pb.User{
		Cpf:        res.Cpf,
		Beneficios: res.Beneficios,
	}

	return benefits, nil

}

func scrape(Doc string) msbenefit.User {

	url_api := os.Getenv("API_URL")
	credential := token.Credentials{
		Login: os.Getenv("API_LOGIN"),
		Senha: os.Getenv("API_PASSWORD"),
	}

	bearerToken := token.GetBearerToken(url_api, credential)
	return benefit.GetBenefits(url_api, bearerToken, Doc)
}
