package spot

import (
	"fmt"
	"io"

	"github.com/liujunren93/bian/client/http"
)

type Spot struct {
	apiBaseURL string
	wsBaseURL  string
	ApiKey     string
	SecretKey  string
	httpClient *http.Client
}

var DefaultSpot = Spot{
	apiBaseURL: "https://api-gcp.binance.com",
	wsBaseURL:  "wss://stream.binance.com:443",
}

func (s *Spot) apiClient() *http.Client {
	if s.httpClient == nil {
		s.httpClient = http.NewClient(http.Config{
			BaseURL:   s.apiBaseURL,
			ApiKey:    s.ApiKey,
			SecretKey: s.SecretKey,
		})
	}
	return s.httpClient
}

func NewSpot(apiKey, secrat string) *Spot {
	s := &Spot{
		apiBaseURL: "https://api.binance.com",
		wsBaseURL:  "wss://stream.binance.com:443",
		ApiKey:     apiKey,
		SecretKey:  secrat,
	}
	// 	httpClient: http.NewClient(http.Config{
	// 		BaseURL:   "https://api.binance.com",
	// 		ApiKey:    apiKey,
	// 		SecretKey: secrat,
	// 	}),
	// }
	return s
}

func (s *Spot) SystenStatus() {
	res, err := s.apiClient().Get("/sapi/v1/system/status", nil, nil, nil)
	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data), err)
}

func (s *Spot) Systentime() {
	res, err := s.apiClient().Get("/api/v3/time", nil, nil, nil)
	fmt.Print(err)
	if err == nil {
		data, _ := io.ReadAll(res.Body)
		fmt.Println(string(data), err)
	}
}

func (s *Spot) Account() {
	res, err := s.apiClient().Get("/api/v3/account", nil, nil, nil)
	fmt.Print(err)
	if err == nil {
		data, _ := io.ReadAll(res.Body)
		fmt.Println(string(data), err)
	}

}
