package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var modes = [...]string{
	"std",
	"taiko",
	"ctb",
	"mania",
}

// Leaderboard is separate from users_stats as leaderboard can be rebuilt from scratch, and its data is just cached.
type Leaderboard struct {
	ID    int `gorm:"primary_key"`
	STD   uint32
	Taiko uint32
	CTB   uint32
	Mania uint32
}

// UpdateLeaderboard queries an update to the database for a certain user.
// Partly taken from https://github.com/osuripple/ripple/blob/00287748552fb8b7b59deccfb9347e87e2d07f77/osu.ppy.sh/inc/pages/Leaderboard.php#L146
func (u UserStats) UpdateLeaderboard(mode byte, db *gorm.DB) error {
	if mode > 3 {
		return errors.New("models/UserStats.UpdateLeaderboard: invalid mode number")
	}
	dbmode := modes[mode]

	var isNew bool
	self := Leaderboard{}
	db.First(&self, u.ID)
	if db.NewRecord(self) {
		isNew = true
		self.ID = u.ID
	} else if self.choose(mode) == 0 {
		isNew = true
	}

	var val uint64
	switch mode {
	case 0:
		val = u.RankedScoreSTD
	case 1:
		val = u.RankedScoreTaiko
	case 2:
		val = u.RankedScoreCTB
	case 3:
		val = u.RankedScoreMania
	}

	target := UserStats{}
	db.Where("ranked_score_"+dbmode+" <= ?", val).Order("ranked_score_" + dbmode + " desc").First(&target)

	var targetPos uint32
	if db.NewRecord(target) {
		lbUser := Leaderboard{}
		db.Order(dbmode + " desc").First(&lbUser)
		if db.NewRecord(lbUser) {
			// if no one is on the leaderboard in general
			targetPos = 1
		} else {
			// if the user is the worst player of the game
			targetPos = lbUser.choose(mode)
		}
	} else {
		lbUser := Leaderboard{}
		db.First(&lbUser, target.ID)
		targetPos = lbUser.choose(mode)
	}

	// BEGIN HELL CODE
	tx := db.Begin()
	tx.Delete(&Leaderboard{
		ID: u.ID,
	})
	if isNew {
		tx.Exec("update leaderboards set "+dbmode+" = "+dbmode+" + 1 where "+dbmode+" >= ? order by "+dbmode+" desc", targetPos)
	} else {
		tx.Exec("update leaderboards set "+dbmode+" = "+dbmode+" + 1 where "+dbmode+" >= ? and "+dbmode+" < ? order by "+dbmode+" desc", self.choose(mode), targetPos)
	}

	switch mode {
	case 0:
		self.STD = targetPos
	case 1:
		self.Taiko = targetPos
	case 2:
		self.CTB = targetPos
	case 3:
		self.Mania = targetPos
	}
	tx.Create(&self)
	tx.Commit()
	// END HELL CODE

	return nil
}

func (u Leaderboard) choose(mode byte) uint32 {
	switch mode {
	case 0:
		return u.STD
	case 1:
		return u.Taiko
	case 2:
		return u.CTB
	case 3:
		return u.Mania
	default:
		return 9001
	}
}

// BuildLeaderboard makes the leaderboard.
func BuildLeaderboard(db *gorm.DB) {
	var c int
	db.Model(&User{}).Count(&c)
	leaderboard := make(map[int]*Leaderboard, c)
	for mID, m := range modes {
		var users []UserStats
		db.Order("ranked_score_" + m + " desc").Find(&users)
		for pos, user := range users {
			if leaderboard[user.ID] == nil {
				leaderboard[user.ID] = &Leaderboard{ID: user.ID}
			}
			lb := leaderboard[user.ID]
			switch mID {
			case 0:
				lb.STD = uint32(pos)
			case 1:
				lb.Taiko = uint32(pos)
			case 2:
				lb.CTB = uint32(pos)
			case 3:
				lb.Mania = uint32(pos)
			}
		}
	}

	endData := make([]Leaderboard, len(leaderboard))
	var i = 0
	for _, lb := range leaderboard {
		endData[i] = *lb
		i++
	}

	tx := db.Begin()
	defer func() {
		p := recover()
		if p != nil {
			tx.Rollback()
			panic(p)
		}
	}()
	tx.Exec("TRUNCATE TABLE leaderboards")
	for _, v := range endData {
		tx.Create(&v)
	}
	tx.Commit()
}
