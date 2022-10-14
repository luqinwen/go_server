package models

import "database/sql"

type Ratings struct {
	Rating    string `json:"rating" xorm:"not null default '' comment('评分') VARCHAR(255)"`
	MovieId   string `json:"movieId" xorm:"not null default '' comment('电影id') VARCHAR(255)"`
	UserId    string `json:"userId" xorm:"null default '' comment('用户id') VARCHAR(255)"`
	Timestamp string `json:"timestamp" xorm:"null default '' comment('时间戳') VARCHAR(255)"`
}

func (r *Ratings) GetAll() ([]Ratings, error) {
	var ratings []Ratings
	err := mEngine.Find(&ratings)
	return ratings, err
}

func (r *Ratings) Add(ratings *Ratings) (int64, error) {
	affected, err := mEngine.Insert(ratings)
	return affected, err
}
func (r *Ratings) Update(rating *Ratings) (sql.Result, error) {
	sql := "update `ratings` set rating=? where movie_id=? and user_id = ?"
	res, err := mEngine.Exec(sql, rating.Rating, rating.MovieId, rating.UserId)
	return res, err
}
