package main

import (
	"context"
	"flag"
	"fmt"
	"io"
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

// Implment the notes.NotesServer interface
func (s *notesServer) SaveLargeNote(stream notes.Notes_SaveLargeNoteServer) error {
	var finalBody []byte
	var finalTitle string
	for {
		// Get a packet
		note, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Recieved a note to save: %v", finalTitle)
			err := notes.SaveToDisk(&notes.Note{
				Title: finalTitle,
				Body:  finalBody,
			}, "testdata")

			if err != nil {
				stream.SendAndClose(&notes.NoteSaveReply{Saved: false})
				return err
			}

			stream.SendAndClose(&notes.NoteSaveReply{Saved: true})
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Recieved a chunk of the note to save: %v", note.Body)
		// Concat packet to create final note
		finalBody = append(finalBody, note.Body...)
		finalTitle = note.Title
	}
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
