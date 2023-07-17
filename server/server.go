package server

import (
	"finance/config"
	"finance/plaid"
	"fmt"
	"net/http"
	"encoding/json"
)

type Server struct {
	config       *config.Configuration
	plaidService *plaid.PlaidService
}

func (s *Server) Health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, 0)
}

func (s *Server) FetchToken(w http.ResponseWriter, req *http.Request) {
	if s.plaidService == nil {
		fmt.Fprint(w, nil)
	}

	res := s.plaidService.CreateLinkToken(*s.config, req)

	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)
	response["link"] = res

	jsonResp, _ := json.Marshal(response)
	w.Write(jsonResp)
}

func Serve(config *config.Configuration, plaidService *plaid.PlaidService) {
	server := Server{
		config:       config,
		plaidService: plaidService,
	}

	http.HandleFunc("/health", server.Health)
	http.HandleFunc("/token/get", server.FetchToken)
	http.ListenAndServe(":8090", nil)
}
