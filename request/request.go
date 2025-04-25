package request

import (
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

type Result struct {
	StatusCode int
}

func RunLoadTest(url string, totalRequests int, concurrency int) []Result {
	var wg sync.WaitGroup
	results := make([]Result, 0, totalRequests)
	ch := make(chan Result, totalRequests)

	// Cria um worker pool para a concorrência
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < totalRequests/concurrency; j++ {
				result := performRequest(url)
				ch <- result
			}
		}()
	}

	// Espera os workers terminarem
	wg.Wait()
	close(ch)

	// Coleta todos os resultados
	for result := range ch {
		results = append(results, result)
	}

	return results
}

func performRequest(url string) Result {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Error("Erro ao criar requisição: ", err)
		return Result{StatusCode: 500}
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("Erro na requisição: ", err)
		return Result{StatusCode: 500}
	}
	defer resp.Body.Close()

	return Result{StatusCode: resp.StatusCode}
}
