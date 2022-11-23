package internal

import (
	"github.com/gin-gonic/gin"
	selfLogger "github.com/thomascwei/golang_logger"
	"net/http"
)

var (
	// Logger create main.log
	Logger = selfLogger.InitLogger("internal")
)

type accountPwd struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// FinishReturnResult request完成返回結果
func FinishReturnResult(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   result,
	})
	return
}

func HelloWorldHandler(c *gin.Context) {
	FinishReturnResult(c, "Hello World!")
}

func LoginHandler(c *gin.Context) {
	accountPwd := accountPwd{}
	err := c.Bind(&accountPwd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"msg":    "request format error",
		})
		return
	}
	Logger.Infof("%+v", accountPwd)
	ok := auth(accountPwd.Account, accountPwd.Password)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"msg":    "account or password error",
		})
		return
	}
	token, err := GenToken(accountPwd.Account)
	if err != nil {
		Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "generate token error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"token":  token,
	})
}

func GetDataHandler(c *gin.Context) {
	id, _ := c.Params.Get("id")
	account, ok := c.Get("account")
	if !ok {
		Logger.Error("account not found")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "get account error",
		})
		return
	}
	token, err := GenToken(account.(string))
	if err != nil {
		Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "generate token error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"id":     id,
		"token":  token,
	})
}

func POSTDataHandler(c *gin.Context) {
	account, ok := c.Get("account")
	if !ok {
		Logger.Error("account not found")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "get account error",
		})
		return
	}
	token, err := GenToken(account.(string))
	if err != nil {
		Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"msg":    "generate token error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"token":  token,
	})
}
