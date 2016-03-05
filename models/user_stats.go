package models

import (
	"github.com/jinzhu/gorm"
)

// UserStats are the statistics of an user that aren't of vital importance.
type UserStats struct {
	ID int `gorm:"primary_key"`

	TotalScoreSTD  uint64
	RankedScoreSTD uint64
	PPSTD          float64
	AccuracySTD    float64

	TotalScoreTaiko  uint64
	RankedScoreTaiko uint64
	PPTaiko          float64
	AccuracyTaiko    float64

	TotalScoreCTB  uint64
	RankedScoreCTB uint64
	PPCTB          float64
	AccuracyCTB    float64

	TotalScoreMania  uint64
	RankedScoreMania uint64
	PPMania          float64
	AccuracyMania    float64

	FavouriteMode  uint8
	ReplaysWatched uint32
}

// RelatedUser finds the User that is related to a certain UserStats.
func (u UserStats) RelatedUser(db gorm.DB) User {
	user := User{}
	db.First(&user, u.ID)
	return user
}
