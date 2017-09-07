package main

import (
	"html/template"
	"net/http"

	cp "github.com/sore0159/star_system/captains"
)

func MakeMux(cfg *Config, r *Resources) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, cfg.StaticDir+"/img/yd32.ico")
	})
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(cfg.StaticDir+"/img"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(cfg.StaticDir+"/css"))))
	s := MakeServer(cfg, r)
	mux.HandleFunc("/", s.CaptainRouter())
	return mux, nil
}

type Server struct {
	*Resources
	TemplateDir string
	DataFile    string
}

func MakeServer(cfg *Config, r *Resources) *Server {
	return &Server{
		Resources:   r,
		TemplateDir: cfg.TemplateDir,
		DataFile:    cfg.DataFile,
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
		s.Log.ServerErr("Failed to read to template: %v", err)
		http.Error(w, "TEMPLATE READ ERROR", 500)
		return
	}
	tp.ExecuteTemplate(w, "frame", nil)
}

func (s *Server) CaptainRouter() http.HandlerFunc {
	return cp.CaptainRouter(
		s.Log,
		s.Academy,
		s.Key,
		s.FoundCaptain,
		s.MadeCaptain,
		s.ServerError,
	)
}

func (s *Server) MadeCaptain(w http.ResponseWriter, r *http.Request, c *cp.Captain) {
	if m, ok := s.Academy.(*cp.MockDB); ok {
		m.Save(s.DataFile)
	}
	tp, err := s.GetTemplate("frame", "made")
	if err != nil {
		s.Log.ServerErr("Failed to read to template: %v", err)
		http.Error(w, "TEMPLATE READ ERROR", 500)
		return
	}
	tp.ExecuteTemplate(w, "frame", c)
}
func (s *Server) FoundCaptain(w http.ResponseWriter, r *http.Request, c *cp.Captain) {
	tp, err := s.GetTemplate("frame", "found")
	if err != nil {
		s.Log.ServerErr("Failed to read to template: %v", err)
		http.Error(w, "TEMPLATE READ ERROR", 500)
		return
	}
	tp.ExecuteTemplate(w, "frame", c)
}
func (s *Server) ServerError(w http.ResponseWriter, r *http.Request, msg string) {
	http.Error(w, msg, 500)
}
