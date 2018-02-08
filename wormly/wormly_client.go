package wormly

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coreos/pkg/flagutil"
)

const baseURL = "https://api.wormly.com/"

type HostStatus struct {
	errorcode int `json:"errorcode"`
	// status    string `json:"status"`
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

func (s *HostService) List() ([]HostStatus, *http.Response, error) {
	hosts := new([]Host)
	body := &WormlyRequestParams{
		key:      "",
		response: "json",
		cmd:      "getHostStatus",
	}

	resp, err := s.sling.New().Post().BodyJSON(body).Receive(hosts, githubError)

	return hosts, resp, err
}

func NewWormlyClient() *WormlyClient {
	return &WormlyClient{
		HostService: NewHostService(nil),
	}
}
