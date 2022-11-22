package middleware

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	selfLogger "github.com/thomascwei/golang_logger"
	"net/http"
	"os"
	"path/filepath"
)

var (
	// Logger create main.log
	Logger   = selfLogger.InitLogger("middleware")
	Enforcer *casbin.Enforcer
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		Logger.Fatal(err)
	}

	Logger.Info(pwd)
	model := "rbac_model.conf"
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
	Enforcer = casbin.NewEnforcer(modelPath, csvPath)
}

func RBAC(sub, obj, act string) bool {
	Logger.Info("RBAC start")
	//if Enforcer.Enforce("admin", "data", "GET") {
	allow := Enforcer.Enforce(sub, obj, act)
	if allow {
		Logger.Infof("%s can use this API", sub)
	} else {
		Logger.Error("ERROR: admin can not read project")
	}
	return allow
}

// AuthorizeMiddleware determines if current user has been authorized to take an action on an object.
func AuthorizeMiddleware(obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user/subject
		sub, existed := c.Get("account")
		if !existed {
			c.AbortWithStatusJSON(401, gin.H{"msg": "can not find account"})
			return
		}

		ok := RBAC(sub.(string), obj, act)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "You are not authorized"})
			return
		}
		c.Next()
	}
}
