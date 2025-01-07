package handlers

import (
	"net/http"
	"time"
	"url_checker/utils"
)

// URLStatus representa el estado de una URL
type URLStatus struct {
	URL          string  `json:"url"`
	StatusCode   int     `json:"status_code"`
	ResponseTime float64 `json:"response_time"`
}

func CheckURLs(urls []string) []URLStatus {
	results := make([]URLStatus, len(urls))
	ch := make(chan URLStatus)

	for _, url := range urls {
		if !utils.IsValidURL(url) {
			ch <- URLStatus{
				URL:          url,
				StatusCode:   0, // Código de error para URL inválida
				ResponseTime: 0,
			}
			continue
		}

		go func(u string) {
			start := time.Now()
			resp, err := http.Get(u)
			status := 0
			if err == nil {
				status = resp.StatusCode
				resp.Body.Close()
			}
			ch <- URLStatus{
				URL:          u,
				StatusCode:   status,
				ResponseTime: float64(time.Since(start).Milliseconds()),
			}
		}(url)
	}

	for i := range urls {
		results[i] = <-ch
	}
	return results
}
