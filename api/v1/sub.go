package v1

import (
	"encoding/json"
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
	var s []string
	err = json.Unmarshal(sub.SubWords, &s)
	if err != nil {
		log.Println("unmarshal error")
		return
	}
	response.CommonSuccess(0, s, "get sub words success", c)
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

	err, subAlready := service.GetUserSub(userUuid)
	if err != nil {
		log.Println(err)
		response.CommonFailed("Get User Sub Error", CodeDbErr, c)
		return
	}

	var subsAdd []string
	err = json.Unmarshal(subAlready.SubWords, &subsAdd)
	if err != nil {
		log.Println("unmarshal error")
		return
	}

	subsAdd = append(subsAdd, sub.SubWord)
	marshal, err := json.Marshal(subsAdd)
	if err != nil {
		log.Println("json marshall error")
		return
	}

	s := &model.Sub{
		UserUUID: userUuid,
		SubWords: datatypes.JSON(marshal),
	}
	if err, subAdded := service.CreateSub(s); err != nil {
		response.CommonFailed("Create Error", CodeDbErr, c)
		return
	} else {
		response.Created(subAdded, "Create Sub Success", c)
	}
}
