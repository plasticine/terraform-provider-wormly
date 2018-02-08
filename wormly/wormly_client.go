package wormly

import (
	// "flag"
	// "log"
	"net/http"
	"fmt"
	// "os"

	// "github.com/coreos/pkg/flagutil"
	"github.com/dghubble/sling"
)

const baseURL = "https://api.wormly.com/"

type HostStatus struct {
	errorcode int `json:"errorcode"`
	// status    string `json:"status"`
}

type WormlyError struct {
    ErrorCode          int   `json:"errorcode"`
    ErrorMessage     string   `json:"errormsg"`
}

type WormlyRequestParams struct {
	Key          string   `json:"key"`
	ResponseType string   `json:"response"`
	Command      string   `json:"cmd"`
	HostID       []string `json:"hostids,omitempty"`
}

type WormlyClient struct {
	HostService *HostService
}

type HostService struct {
	sling *sling.Sling
}

func NewHostService(httpClient *http.Client) *HostService {
	return &HostService{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}

func (s *HostService) List() (*[]HostStatus, *http.Response, error) {
	wormlyHosts := new([]HostStatus)
	wormlyError := new(WormlyError)
	body := &WormlyRequestParams{
		Key:          "",
		ResponseType: "json",
		Command:      "getHostStatus",
	}

	resp, err := s.sling.New().Post("/").BodyJSON(body).Receive(wormlyHosts, wormlyError)

	fmt.Println(wormlyHosts, wormlyError, resp, err)

	return wormlyHosts, resp, err
}

func NewWormlyClient() *WormlyClient {
	return &WormlyClient{
		HostService: NewHostService(nil),
	}
}
