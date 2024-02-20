package internal

import (
	"fmt"
	"github.com/andre2ar/stress-test/pkg"
	"sync"
	"time"
)

type StressTest struct {
	Report sync.Map
}

func (st *StressTest) Stress(url string, requests int, concurrency int) {
	responses := make(chan int, concurrency)

	start := time.Now()

	responsesWaitGroup := sync.WaitGroup{}
	responsesWaitGroup.Add(1)
	go st.genReport(responses, &responsesWaitGroup)

	requestWaitGroup := sync.WaitGroup{}
	canRequestBegin := make(chan bool, concurrency)

	processedRequests := 0
	for ; processedRequests < requests; processedRequests++ {
		canRequestBegin <- true
		requestWaitGroup.Add(1)

		go pkg.GetRequest(url, responses, canRequestBegin, &requestWaitGroup)
	}

	requestWaitGroup.Wait()
	close(responses)

	responsesWaitGroup.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %.2f seconds\n", elapsed.Seconds())
	fmt.Println("Processed request: ", processedRequests)

	st.PrintReport()
}

func (st *StressTest) genReport(responses chan int, wg *sync.WaitGroup) {
	for responseCode := range responses {
		processedRequests, ok := st.Report.Load(responseCode)

		if !ok {
			processedRequests = 0
		}

		st.Report.Store(responseCode, processedRequests.(int)+1)
	}

	wg.Done()
}

func (st *StressTest) PrintReport() {
	st.Report.Range(func(key, value any) bool {
		fmt.Printf("Status Code %d:  processed %d requests\n", key, value)
		return true
	})
}
