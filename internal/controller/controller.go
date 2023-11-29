package controller

import (
	"archive/zip"
	"fmt"
	"github.com/e4t4g/test_FT/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type Controller struct {
	router *chi.Mux
	cfg    *config.Config
	log    *zap.SugaredLogger
}

func (c *Controller) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c.router.ServeHTTP(writer, request)
}

func NewController(router *chi.Mux, cfg *config.Config, log *zap.SugaredLogger) *Controller {

	ctr := &Controller{router: router, cfg: cfg, log: log}

	ctr.router.Use(middleware.Logger)
	ctr.router.Get("/", ctr.FileList)

	return ctr
}

func (c *Controller) FileList(w http.ResponseWriter, r *http.Request) {

	fl, err := zip.OpenReader(c.cfg.File)
	if err != nil {
		c.log.Errorf("%v", err)
		http.NotFound(w, r)
		return
	}

	defer fl.Close()

	w.Write([]byte(fmt.Sprintf("<body>")))

	w.Write([]byte(fmt.Sprintf("<h3>Ext: %v</h3>", c.cfg.Ext)))

	w.Write([]byte(fmt.Sprintf("<ul>")))

	for _, file := range fl.File {
		if strings.HasSuffix(file.Name, c.cfg.Ext) == true {
			w.Write([]byte(fmt.Sprintf("<li> %s </li>", file.Name)))
			c.log.Infof(file.Name)
		}

	}

	w.Write([]byte(fmt.Sprintf("</ul>")))
	w.Write([]byte(fmt.Sprintf("</body>")))
	w.WriteHeader(http.StatusOK)
}
