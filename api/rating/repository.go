package rating

import (
	"github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"
)

// RatingRepository is
type RatingRepository interface{
	Rating(id uint) (float64, []error)
	StoreRating(rating *entity.Comment) (*entity.Comment, []error)
}
