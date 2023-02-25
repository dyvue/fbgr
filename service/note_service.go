package service

import (
	"fbgr/data/request"
	"fbgr/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(noteId int)
	FindById(noteId int) response.NoteResponse
	FindByAll() []response.NoteResponse
}
