package pkg

import (
	"net/http"
	"sync"
)

func GetRequest(url string, responses chan<- int, canRequestBegin chan bool, wg *sync.WaitGroup) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		responses <- 500
		wg.Done()
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		responses <- res.StatusCode
		wg.Done()
		return
	}
	defer res.Body.Close()

	responses <- res.StatusCode
	<-canRequestBegin
	wg.Done()
}
