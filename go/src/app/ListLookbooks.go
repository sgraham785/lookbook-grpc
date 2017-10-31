package main

import (
	"encoding/json"
	"log"
	"fmt"

	pb "app/pb/lookbook"

	"golang.org/x/net/context"
	mgo "gopkg.in/mgo.v2"
)

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
