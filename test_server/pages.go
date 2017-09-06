package main

import (
	"html/template"
	"net/http"
)

func MakeMux(cfg *Config) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, cfg.StaticDir+"/img/yd32.ico")
	})
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(cfg.StaticDir+"/img"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(cfg.StaticDir+"/css"))))
	s := MakeServer(cfg)
	mux.HandleFunc("/", s.TestPage)
	return mux
}

type Server struct {
	TemplateDir string
}

func MakeServer(cfg *Config) *Server {
	return &Server{
		TemplateDir: cfg.TemplateDir,
	}
}

func (s *Server) GetTemplate(files ...string) (*template.Template, error) {
	for i, f := range files {
		files[i] = s.TemplateDir + f + ".html"
	}
	return template.ParseFiles(files...)
}

func (s *Server) TestPage(w http.ResponseWriter, r *http.Request) {
	tp, err := s.GetTemplate("frame")
	if err != nil {
		LOG.ServerErr("Failed to read to template: %v", err)
		http.Error(w, "TEMPLATE READ ERROR", 500)
		return
	}
	tp.ExecuteTemplate(w, "frame", nil)
}
