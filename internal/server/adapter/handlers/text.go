package handlers

import (
	"context"
	proto2 "github.com/dlc-01/GophKeeper/internal/general/proto"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TextServer struct {
	proto2.UnimplementedTextServer
	text port.ITextService
}

func NewTextServer(text port.ITextService) *TextServer {
	return &TextServer{
		text: text,
	}
}

func (t *TextServer) CreateText(ctx context.Context, req *proto2.CreateTextRequest) (*proto2.CreateTextResponse, error) {
	var resp proto2.CreateTextResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	text := models.Text{
		Note:     req.Note.Note,
		Metadata: req.Note.Metadata,
	}

	texts, err := t.text.CreateByUserId(ctx, user, text)
	if err != nil {
		return nil, err
	}

	resp.Note = req.Note
	resp.Note.Id = texts.ID

	return &resp, nil
}

func (t *TextServer) GetText(ctx context.Context, req *proto2.GetTextRequest) (*proto2.GetTextResponse, error) {
	var resp proto2.GetTextResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	stored, err := t.text.GetTextsByUserID(ctx, user)
	if err != nil {
		return nil, err
	}

	for _, note := range *stored {
		resp.Notes = append(resp.Notes, &proto2.NoteMsg{
			Id:       note.ID,
			Note:     note.Note,
			Metadata: note.Metadata,
		})
	}

	return &resp, nil
}

func (t *TextServer) UpdateText(ctx context.Context, req *proto2.UpdateTextRequest) (*proto2.UpdateTextResponse, error) {
	var resp proto2.UpdateTextResponse

	_, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	text := models.Text{
		ID:       req.Note.Id,
		Note:     req.Note.Note,
		Metadata: req.Note.Metadata,
	}

	_, err := t.text.Update(ctx, text)
	if err != nil {
		return nil, err
	}

	resp.Note = req.Note

	return &resp, nil
}

func (t *TextServer) DeleteText(ctx context.Context, req *proto2.DeleteTextRequest) (*proto2.DeleteTextResponse, error) {
	var resp proto2.DeleteTextResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	err := t.text.DeleteByUserID(ctx, user)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
