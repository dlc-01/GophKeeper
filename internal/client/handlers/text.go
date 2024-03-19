package handlers

import (
	"context"
	"github.com/dlc-01/GophKeeper/internal/general/proto"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type TextClient struct {
	conn *grpc.ClientConn
}

func NewTextClient(conn *grpc.ClientConn) *TextClient {
	return &TextClient{
		conn: conn,
	}
}

func (c *TextClient) CreateText(ctx context.Context, token string, note models.Text) (*models.Text, error) {
	client := proto.NewTextClient(c.conn)
	req := &proto.CreateTextRequest{
		Token: token,
		Note:  &proto.NoteMsg{Note: note.Note, Metadata: note.Metadata},
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	resp, err := client.CreateText(ctx, req)
	if err != nil {
		return nil, err
	}
	note.ID = resp.Note.Id
	return &note, nil
}

func (c *TextClient) GetText(ctx context.Context, token string) ([]models.Text, error) {
	client := proto.NewTextClient(c.conn)
	req := &proto.GetTextRequest{
		Token: token,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	resp, err := client.GetText(ctx, req)
	if err != nil {
		return nil, err
	}

	out := make([]models.Text, len(resp.Notes))
	for i, note := range resp.GetNotes() {
		out[i] = models.Text{
			ID:       note.GetId(),
			Note:     note.GetNote(),
			Metadata: note.GetMetadata(),
		}
	}

	return out, nil
}

func (c *TextClient) UpdateText(ctx context.Context, note models.Text) (*models.Text, error) {
	client := proto.NewTextClient(c.conn)
	req := &proto.UpdateTextRequest{
		Note: &proto.NoteMsg{Id: note.ID, Note: note.Note, Metadata: note.Metadata},
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()

	resp, err := client.UpdateText(ctx, req)
	if err != nil {
		return nil, err
	}
	note = models.Text{ID: resp.Note.GetId(), Note: req.Note.GetNote(), Metadata: req.Note.Metadata}
	return &note, nil
}

func (c *TextClient) DeleteText(ctx context.Context, token string) error {
	client := proto.NewTextClient(c.conn)
	req := &proto.DeleteTextRequest{
		Token: token,
	}

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	_, err := client.DeleteText(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
