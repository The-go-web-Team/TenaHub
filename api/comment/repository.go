package comment

import (
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
)

// CommentRepository is
type CommentRepository interface {
	Comment(id int)(*entity.Comment, []error)
	Comments()([]entity.Comment, []error)
	UpdateComment(hc entity.Comment)(*entity.Comment, []error)
	StoreComment(hc entity.Comment)(*entity.Comment, []error)
	DeleteComment(id int)(*entity.HealthCenter, []error)
}