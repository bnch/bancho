package models

import (
	"github.com/jinzhu/gorm"
)

var migrations = map[uint64]func(*gorm.DB){
	1455998503: func(db *gorm.DB) {
		db.CreateTable(&User{})
	},
	1456178412: func(db *gorm.DB) {
		// Users start from ID 3.
		// This is because if an user comes to have ID == 2, then they become peppy and thus can't be contacted
		// (any attempt to message them will result in a browser opening http://osu.ppy.sh/p/doyoureallywanttoaskpeppy)
		db.Exec("ALTER TABLE users AUTO_INCREMENT=3;")
	},
	1456756895: func(db *gorm.DB) {
		db.CreateTable(&UserFriendship{})
	},
	1457036659: func(db *gorm.DB) {
		db.CreateTable(&Channel{})
		db.Create(&Channel{
			Name:        "#osu",
			Description: "Main channel for discussion about anything and everything.",
		})
		db.Create(&Channel{
			Name:        "#announce",
			Description: "The channnel where the announcements should appear. Not really.",
		})
	},
	1457125054: func(db *gorm.DB) {
		db.CreateTable(&UserStats{})
		users := []User{}
		db.Find(&users)
		for _, u := range users {
			db.Create(&UserStats{
				ID: u.ID,
			})
		}
	},
	1457180378: func(db *gorm.DB) {
		db.CreateTable(&Leaderboard{})
		BuildLeaderboard(db)
	},
	1457450104: func(db *gorm.DB) {
		renamesLeaderboard := map[string]string{
			"s_t_d": "std",
			"c_t_b": "ctb",
		}
		renamesUserStats := map[string]string{
			"p_p_s_t_d":          "ppstd",
			"p_p_taiko":          "pp_taiko",
			"p_p_c_t_b":          "ppctb",
			"p_p_mania":          "pp_mania",
			"total_score_s_t_d":  "total_score_std",
			"total_score_c_t_b":  "total_score_ctb",
			"ranked_score_s_t_d": "ranked_score_std",
			"ranked_score_c_t_b": "ranked_score_ctb",
			"accuracy_s_t_d":     "accuracy_std",
			"accuracy_c_t_b":     "accuracy_ctb",
		}
		renameColumns(renamesLeaderboard, "leaderboards", db)
		renameColumns(renamesUserStats, "user_stats", db)
	},
}

func renameColumns(cols map[string]string, tableName string, db *gorm.DB) {
	r, err := db.DB().Query("SHOW FIELDS FROM " + tableName)
	if err != nil {
		panic(err)
	}
	for r.Next() {
		var field string
		var fType string
		var extra string
		var none string
		r.Scan(&field, &fType, &none, &none, &none, &extra)
		if v, ok := cols[field]; ok {
			db.Exec("ALTER TABLE " + tableName + " CHANGE " + field + " " + v + " " + fType + " " + extra)
		}
	}
}
