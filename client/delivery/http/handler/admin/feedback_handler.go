package admin

import "html/template"

type FeedBackHandler struct {
	temp *template.Template
}
func NewFeedBackHandlerHandler(T *template.Template) *FeedBackHandler {
	return &FeedBackHandler{temp: T}
}
