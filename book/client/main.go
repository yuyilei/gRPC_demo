package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/gRPC_demo/book/book-proto"
	"time"
	"io"
)

const (
	address     = "localhost:30001"
)

func runGetBookInfo(client book.BookClient, bid *book.BookId) {
	log.Printf("Get BookInfo of Id: %v\n",bid.Id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bookinfo, err := client.GetBookInfo(ctx,bid)
	if err != nil {
		log.Fatalf("%v.GetBookInfo() = %v\n\n",client,err)
	}
	log.Printf("%v\n\n",bookinfo)
}

func runPostBookList(client book.BookClient) {
	log.Printf("Post Book List of Info:\n")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.PostBookList(ctx)
	if err != nil {
		log.Fatalf("%v.PostBookList() = %v\n",client,err)
	}
	book1 := &book.BookInfo{
		Author: "Luciano Ramalho",
		Name: "Fluent Python",
		Page: 2333,
	}
	book2 := &book.BookInfo{
		Author: "ahahhahahhahhah",
		Name: "1234",
		Page: 1234,
	}
	stream.Send(book1)
	stream.Send(book2)
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.PostBookList() got error %v, want %v\n", stream, err, nil)
	}
	log.Printf("post summary: %v\n\n", reply)
}

func runGetBookList(client book.BookClient, page *book.BookPage) {
	log.Printf("Get Book List of Info:\n")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetBookList(ctx,page)
	if err != nil {
		log.Fatalf("%v.GetBookList() = %v\n",client,err)
	}
	for {
		bookinfo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetBookList() got error %v, want %v\n", stream, err, nil)
		}
		log.Printf("Book Info: %v\n",bookinfo)
	}
	log.Println()
}

func main() {
	// Set up a connection to the server.
	var err error
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := book.NewBookClient(conn)

	// Contact the server and print out its response.
	runGetBookInfo(c,&book.BookId{Id:"1"})
	runGetBookList(c,&book.BookPage{Page:123})
	runPostBookList(c)
	runGetBookList(c,&book.BookPage{Page:123})
}