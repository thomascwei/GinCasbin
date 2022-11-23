package middleware

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var ABACEnforcer *casbin.Enforcer

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		Logger.Fatal(err)
	}

	Logger.Info(pwd)
	model := "abac_model.conf"
	csv := "rbac_policy.csv"
	modelPath := filepath.Join(pwd, model)
	csvPath := filepath.Join(pwd, csv)
	_, err = os.Stat(modelPath)
	if err != nil {
		modelPath = filepath.Join(pwd, "middleware", model)
	}
	_, err = os.Stat(csvPath)
	if err != nil {
		csvPath = filepath.Join(pwd, "middleware", csv)
	}
	ABACEnforcer = casbin.NewEnforcer(modelPath, csvPath)
}

type Subject struct {
	Name string
	Hour int
}

func checkABAC(sub Subject, obj, act string) bool {
	ok := ABACEnforcer.Enforce(sub, obj, act)
	if ok {
		Logger.Infof("%s CAN %s %s at %d:00\n", sub.Name, act, obj, sub.Hour)
	} else {
		Logger.Infof("%s CANNOT %s %s at %d:00\n", sub.Name, act, obj, sub.Hour)
	}
	return ok
}

func ABACAuthorizeMiddleware(obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user/subject
		sub, existed := c.Get("account")
		if !existed {
			c.AbortWithStatusJSON(401, gin.H{"msg": "can not find account"})
			return
		}
		subStruct := Subject{Name: sub.(string), Hour: time.Now().Hour()}
		Logger.Infof("%+v", subStruct)
		ok := checkABAC(subStruct, obj, act)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "You are not authorized"})
			return
		}
		c.Next()
	}
}
