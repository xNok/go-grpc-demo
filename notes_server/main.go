package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/xNok/go-grpc-demo/notes"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// Implement the notes service (notes.NotesServer interface)
type notesServer struct {
	notes.UnimplementedNotesServer
}

// Implment the notes.NotesServer interface
func (s *notesServer) Save(ctx context.Context, n *notes.Note) (*notes.NoteSaveReply, error) {
	log.Printf("Recieved a note to save: %v", n.Title)
	err := notes.SaveToDisk(n, "testdata")

	if err != nil {
		return &notes.NoteSaveReply{Saved: false}, err
	}

	return &notes.NoteSaveReply{Saved: true}, nil
}

// Implment the notes.NotesServer interface
func (s *notesServer) Load(ctx context.Context, search *notes.NoteSearch) (*notes.Note, error) {
	log.Printf("Recieved a note to save: %v", search.Keyword)
	n, err := notes.LoadFromDisk(search.Keyword, "testdata")

	if err != nil {
		return &notes.Note{}, err
	}

	return n, nil
}

func main() {
	// parse arguments from the command line
	// this lets use define the port for the server
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// Check for errors
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Instanciate the server
	s := grpc.NewServer()

	// Register server method (actions the server il do)
	notes.RegisterNotesServer(s, &notesServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
