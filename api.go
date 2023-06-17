package main

import "net/http"

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {

}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handlecreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
