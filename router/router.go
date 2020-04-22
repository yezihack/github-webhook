package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-webhook/config"
	"go-webhook/internal"
	"go-webhook/logger"
	"go-webhook/util"

	"net/http"
)

var (
	conf   config.Config
	engine *gin.Engine
	log    logger.Logger
)

// new api router server
func New(cfg config.Config) error {
	conf = cfg
	log.Print(conf)
	setMode(conf)
	engine = gin.Default()
	addRouter(conf)
	log = logger.NewLogger(conf.Quiet)
	return engine.Run(fmt.Sprintf(":%d", conf.Port))
}

// set server mode
func setMode(cfg config.Config) {
	if cfg.Quiet == false {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// web http
func addRouter(cfg config.Config) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "PONG")
	})
	engine.POST("/web-hook", WebHookLog(cfg))
}

// hook
func WebHookLog(conf config.Config) gin.HandlerFunc {
	return internal.Handler(conf.Secret, log, func(event string, payload *internal.GitHubRepo, req *http.Request) error {
		// Log webhook
		log.Println("Received:", event, "for:", payload.Name)
		// You'll probably want to do some real processing
		log.Println("Can clone repo at:", payload.CloneURL)
		log.Printf("Commit INFO Name:%s, Email:%s, Branch:%s, CommitID:%s, At:%s\n",
			payload.CommitName, payload.CommitEmail, payload.BranchName, payload.CommitID, payload.CommitAt)

		// All is good (return an error to fail)
		str, err := util.CallScript(conf.ScriptBash)
		if err != nil {
			log.Print(err)
			return err
		}
		log.Print(str)
		return nil
	})
}
