// fetch.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"syscall/js"
)

func fetchStatus(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("ERROR: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("ERROR: %s", err.Error())
		return
	}

	results <- fmt.Sprintf("%s: %d %s", url, resp.StatusCode, string(body)[:60]) // レスポンスの最初の60文字を取得
}

func fetchAll(urls []string, done js.Func) {
	var wg sync.WaitGroup
	results := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetchStatus(url, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
		done.Invoke()
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func registerCallbacks() {
	js.Global().Set("fetchAll", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		urls := make([]string, len(args)-1)
		for i, arg := range args[:len(args)-1] {
			urls[i] = arg.String()
		}
		doneCallback := args[len(args)-1]

		go fetchAll(urls, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			doneCallback.Invoke()
			return nil
		}))

		return nil
	}))
}

func main() {
	c := make(chan struct{}, 0)

	registerCallbacks()

	<-c
}
