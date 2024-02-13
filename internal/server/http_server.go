package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"custom-fizzbuzz/internal/handler"
	"custom-fizzbuzz/pkg/model"
)

func NewHttpServer(h *handler.Handler, addr string) *Server {
	mux := http.NewServeMux()
	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	s := &Server{httpServer: httpServer}
	mux.HandleFunc("/numbers/print", HandlerFunc(h))
	return s
}

func HandlerFunc(h *handler.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var req *model.RequestData

		switch r.Method {
		case "GET": // try to retrieve query params
			req, err = parseQueryParams(r.URL)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Printf("error decoding query params: %#v\n", err.Error())
				return
			}
		case "POST": // use request body
			req = &model.RequestData{}
			err = json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Printf("error decoding body: %#v\n", err)
				return
			}
		default:
			http.Error(w, fmt.Sprintf("unknown query method"), http.StatusMethodNotAllowed)
		}

		input := model.FromRequestData(req)
		err = input.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		numbers := h.PrintNumber(input)
		resp, err := json.Marshal(numbers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Printf("error marshalling resp: %#v\n", err)
			return
		}
		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Printf("error writing resp: %#v\n", err)
			return
		}
	}
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) ListenAndServe() {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		fmt.Printf("failed to start http server : %#v", err)
	}
}

// Shutdown stops the HTTP server (so the exporter become unavailable).
func (s *Server) Shutdown(ctx context.Context) {
	_ = s.httpServer.Shutdown(ctx)
}

func parseQueryParams(url *url.URL) (*model.RequestData, error) {
	maxNb, err := parseIntParam(url.Query().Get("maxNumber"))
	if err != nil {
		return nil, err
	}
	firstMult, err := parseIntParam(url.Query().Get("firstMultiple"))
	if err != nil {
		return nil, err
	}
	secondMult, err := parseIntParam(url.Query().Get("secondMultiple"))
	if err != nil {
		return nil, err
	}
	req := &model.RequestData{
		MaxNumber:      maxNb,
		FirstMultiple:  firstMult,
		SecondMultiple: secondMult,
		FirstAlias:     parseStringParam(url.Query().Get("firstAlias")),
		SecondAlias:    parseStringParam(url.Query().Get("secondAlias")),
	}

	return req, nil
}

func parseIntParam(p string) (*int, error) {
	if p == "" {
		return nil, nil
	}
	param, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return nil, err
	}
	res := int(param)
	return &res, nil
}

func parseStringParam(p string) *string {
	if p == "" {
		return nil
	}
	return &p
}
