package user

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-sso/models"
	"go-sso/modules/app"
	"go-sso/utils/response"
)

// 记录用户的对于某一个电影的评分
func Rating(c *gin.Context) {

}

func GetAllRating(c *gin.Context) {
	model := models.Ratings{}
	var res []models.Ratings
	var err error
	if res, err = model.GetAll(); err != nil {
		fmt.Println(err)
		response.ShowError(c, "get ratings err")
		return
	}
	fmt.Printf("%v", res)
}

func UpdateMovieRating(c *gin.Context) {
	var rating models.Ratings
	var res sql.Result
	var affected int64
	var err error
	if err = c.Bind(&rating); err != nil {
		fmt.Printf("%v\n", err)
		response.ShowError(c, "bind err")
		return
	}
	if rating.UserId, err = c.Cookie(app.COOKIE_TOKEN); err != nil {
		response.ShowError(c, "get cookie err")
	}
	if res, err = rating.Update(&rating); err != nil {
		fmt.Printf("%v\n", err)
		response.ShowError(c, "update err")
		return
	}
	affected, _ = res.RowsAffected()
	response.ShowData(c, affected)
}

func RateAMovie(c *gin.Context) {
	var rating models.Ratings
	var affected int64
	var err error
	if err = c.Bind(&rating); err != nil {
		fmt.Printf("%v\n", err)
		response.ShowError(c, "bind err")
		return
	}
	if rating.UserId, err = c.Cookie(app.COOKIE_TOKEN); err != nil {
		response.ShowError(c, "get cookie err")
	}
	if affected, err = rating.Add(&rating); err != nil {
		fmt.Printf("%v\n", err)
		response.ShowError(c, "add err")
		return
	}
	response.ShowData(c, affected)
}
