package comment

import "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"

// CommentService is
type CommentService interface {
	Comment(id int) (*entity.Comment, []error)
	Comments() ([]entity.Comment, []error)
	UpdateComment(hc entity.Comment) (*entity.Comment, []error)
	StoreComment(hc entity.Comment) (*entity.Comment, []error)
	DeleteComment(id int) (*entity.Comment, []error)
}