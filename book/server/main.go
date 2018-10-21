package main

import (
	"net"
	"fmt"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/gRPC_demo/book/book-proto"
	"io"
	"sync"
)


const (
	port = "30001"
)

type server struct{
	mu         sync.Mutex // protects
	books      []*book.BookInfo
}

func (s *server) GetBookInfo(ctx context.Context, in *book.BookId) (*book.BookInfo, error) {
	return &book.BookInfo{Author: "Luciano Ramalho",Name: "Fluent Python", Page: 233}, nil
}

func (s *server) GetBookList(in *book.BookPage,stream book.Book_GetBookListServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, each := range s.books {
		stream.Send(each)
	}
	return nil
}

func (s *server) PostBookList(stream book.Book_PostBookListServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for {
		bookinfo, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&book.Postresponse{Msg:"success!"})
			break
		}
		if err != nil {
			return err
		}
		log.Printf("Book Info: %v\n",bookinfo)
		s.books = append(s.books,bookinfo)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	book.RegisterBookServer(s, &server{})
	s.Serve(lis)
}
