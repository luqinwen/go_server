package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-sso/models"
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

func UpdateAMovie(c *gin.Context) {

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
	if affected, err = rating.Add(&rating); err != nil {
		fmt.Printf("%v\n", err)
		response.ShowError(c, "add err")
		return
	}
	response.ShowData(c, affected)
}
