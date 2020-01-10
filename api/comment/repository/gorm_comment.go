package repository

import (
"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/entity"
)

// CommentGormRepo implements comment.CommentRepository
type CommentGormRepo struct {
	conn *gorm.DB
}

// NewCommentGormRepo creates object of CommentGormRepo
func NewCommentGormRepo(conn *gorm.DB) *CommentGormRepo {
	return &CommentGormRepo{conn: conn}
}

// Comments returns all health center comments from database
func (cr *CommentGormRepo) Comments(id uint) ([]entity.Comment, []error) {
	comments := []entity.Comment{}
	errs := cr.conn.Where("health_center_id = ?", id).Find(&comments).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return comments, nil
}

// Comment returns single healthcenter comment from database
func (cr *CommentGormRepo) Comment(id uint) (*entity.Comment, []error) {
	comment := entity.Comment{}
	errs := cr.conn.First(&comment, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &comment, nil
}

// UpdateComment updates comment from the database
func (cr *CommentGormRepo) UpdateComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmt := comment

	errs := cr.conn.Save(cmt).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return cmt, nil
}

// StoreComment stores comment to the database
func (cr *CommentGormRepo) StoreComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmt := comment
	errs := cr.conn.Create(cmt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmt, nil
}
// DeleteComment deletes single comment from the database
func (cr *CommentGormRepo) DeleteComment(id uint) (*entity.Comment, []error) {
	comment, errs := cr.Comment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = cr.conn.Delete(&comment, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return comment, nil
}