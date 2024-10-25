package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Response struct {
		ErrCode int32  `json:"err_code"`
		ErrMsg  string `json:"err_msg"`
	}

	User struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	UserInfo struct {
		Response
		Users []*User `json:"users"`
	}
)

var userCache = make(map[string]*User, 64)

func (s *Server) RegisterUser(c *gin.Context) {
	user := &User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			ErrCode: 100,
			ErrMsg:  "invalid request",
		})
		return
	}
	if user.UserName == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, Response{
			ErrCode: 1,
			ErrMsg:  "username and password can't be empyt",
		})
		return
	}
	userCache[user.UserName] = user
	c.JSON(http.StatusOK, user)
}

func (s *Server) GetUserInfo(c *gin.Context) {
	users := make([]*User, 0, 16)
	for _, v := range userCache {
		users = append(users, v)
	}
	var rsp UserInfo
	rsp.Users = users

	c.JSON(http.StatusOK, rsp)
}
