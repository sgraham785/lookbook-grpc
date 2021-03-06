package main

import (
	"fmt"
	"log"
	"net"

	pb "app/pb/lookbook"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

type Lookbooks struct {
	Id			string 		`json:'_id'`
	Name		string 		`json:'name'`
	UserId 	string 		`json:'user_id'`
	Items		[]string 	`json:'items_id'`
}

func (s *mgo.Session) ListLookbooks(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("workplacex").C("lookbooks")

	var books []Lookbooks

	err := c.Find(bson.M{}).All(&books)
	if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all books: ", err)
			return
	}

	respBody, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
			log.Fatal(err)
	}
	return &pb.Response{data:[respBody]}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLookbookServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		fmt.Printf("Server running on port: %v", port)
	}
}
