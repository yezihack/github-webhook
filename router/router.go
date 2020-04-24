package router

import (
	"fmt"
	"github.com/yezihack/github-webhook/config"
	"github.com/yezihack/github-webhook/internal"
	"github.com/yezihack/github-webhook/logger"
	"github.com/yezihack/github-webhook/util"
	"github.com/yezihack/gorestful"

	"net/http"
)

var (
	log logger.Logger
)

// new api router server
func New(cfg config.Config) error {
	if cfg.Quiet {
		gorestful.SetMode(gorestful.ModeRelease)
	}
	log = logger.NewLogger(cfg.Quiet, cfg.Verbose)
	router := gorestful.New()
	router.GET("/ping", pong)
	router.POST("/web-hook", webHookLog(cfg))
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}

func pong(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		util.Response(w, 405, "Method Not Allowed")
		return
	}
	util.Response(w, 200, "PONG")
}

// hook
func webHookLog(conf config.Config) http.HandlerFunc {
	return internal.Handler(conf.Secret, log, func(event string, payload *internal.GitHubRepo, req *http.Request) error {
		// Log webhook
		log.Println("Received:", event, " for:", payload.Name)
		// You'll probably want to do some real processing
		log.Println("Can clone repo at:", payload.CloneURL)
		log.Printf("Commit INFO Name:%s\n, Email:%s\n, Branch:%s\n, CommitID:%s\n, At:%s\n",
			payload.CommitName, payload.CommitEmail, payload.BranchName, payload.CommitID, payload.CommitAt)

		// All is good (return an error to fail)
		str, err := util.CallScript(conf.ScriptBash)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(str)
		return nil
	})
}
