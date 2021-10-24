package v1

import (
	"github.com/akazwz/go-gin-restful-api/model"
	"github.com/akazwz/go-gin-restful-api/model/request"
	"github.com/akazwz/go-gin-restful-api/model/response"
	"github.com/akazwz/go-gin-restful-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"log"
)

func GetALlSubWords(c *gin.Context) {
	var pageInfo request.PageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.CommonFailed("Bind Query Error", CodeBindError, c)
	}
	if err, list, total := service.GetSubList(pageInfo); err != nil {
		response.CommonFailed("Get Subs Error", CodeDbErr, c)
	} else {
		response.CommonSuccess(0, response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "Get Subs Success", c)
	}
}

func GetUserSub(c *gin.Context) {
	// 获取 user uuid
	claims, _ := c.Get("claims")
	// convent claims to type *request.CustomClaims
	customClaims := claims.(*request.CustomClaims)
	// get user uuid to store who upload this file
	userUuid := customClaims.UUID
	err, sub := service.GetUserSub(userUuid)
	if err != nil {
		log.Println(err)
		response.CommonFailed("Get User Sub Error", CodeDbErr, c)
		return
	}
	log.Println(sub.SubWords)
}

func CreateSub(c *gin.Context) {
	var sub request.Sub
	err := c.ShouldBindJSON(&sub)
	if err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
		return
	}
	// 获取 user uuid
	claims, _ := c.Get("claims")
	// convent claims to type *request.CustomClaims
	customClaims := claims.(*request.CustomClaims)
	// get user uuid to store who upload this file
	userUuid := customClaims.UUID
	s := &model.Sub{
		UserUUID: userUuid,
		SubWords: datatypes.JSON(sub.SubWord),
	}
	if err, subAdded := service.CreateSub(s); err != nil {
		response.CommonFailed("Create Error", CodeDbErr, c)
		return
	} else {
		response.Created(subAdded, "Create Sub Success", c)
	}
}
