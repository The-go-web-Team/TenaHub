package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"

	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/comment"
	"github.com/julienschmidt/httprouter"
)

// CommentHandler handles comment related http requests
type CommentHandler struct {
	cmtService comment.CommentService
}

// NewCommentHandler creates an object of CommentHandler
func NewCommentHandler(cs comment.CommentService) *CommentHandler {
	return &CommentHandler{cmtService: cs}
}

// GetComments hanldes GET /v1/comments/:id
func (ch *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comments, errs := ch.cmtService.Comments(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(&comments, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

// GetComment hanldes GET /v1/comment/:id
func (ch *CommentHandler) GetComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ch.cmtService.Comment(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(&comment, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return
}

// PutComment handles PUT /v1/comments/:id
func (ch *CommentHandler) PutComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ch.cmtService.Comment(uint(id))

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)

	err = json.Unmarshal(body, comment)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs = ch.cmtService.UpdateComment(comment)
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comment, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(output)
	return

}

// DeleteComment handles DELETE /v1/comments/:id
func (ch *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ch.cmtService.DeleteComment(uint(id))
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = json.MarshalIndent(comment, "", "\n")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return

}

// PostComment handles POST /v1/comments
func (ch *CommentHandler) PostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-type", "application/json")

	comment := &entity.Comment{}

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)

	err := json.Unmarshal(body, comment)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ch.cmtService.StoreComment(comment)

	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/comments/%d", comment.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return

}
