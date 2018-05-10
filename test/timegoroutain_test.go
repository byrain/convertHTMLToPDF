package test

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"sync"
	"testing"
	"time"
)

const (
	COUNT   = 30
	TIMEOUT = 60
)

func Benchmark_With_Routine_test(b *testing.B) {
	urlToTitle_temp := map[int]string{}
	convertedURL := []string{}

	for i := 0; i < COUNT; i++ {
		urlToTitle_temp[i] = "www.baidu.com"
	}

	var wg sync.WaitGroup
	wg.Add(len(urlToTitle_temp))
	for title, url := range urlToTitle_temp {
		ctx, _ := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
		go func(url string, title int) {
			defer func() {
				wg.Done()
			}()
			cmd := exec.CommandContext(ctx, "wkhtmltopdf", url, strconv.Itoa(title)+".pdf")
			err := cmd.Start()
			if err != nil {
				log.Printf(err.Error())
			}
			err = cmd.Wait()
			if err != nil {
				log.Printf("Timeout HTML is %s", strconv.Itoa(title)+".pdf")
				return
			}
			convertedURL = append(convertedURL, strconv.Itoa(title)+".pdf")
		}(url, title)
	}
	wg.Wait()

	fmt.Printf("Done, success ratio is %d\n", len(urlToTitle_temp)/len(convertedURL))
}

func Benchmark_Time_Without_Routine_test(b *testing.B) {
	urlToTitle_temp := map[int]string{}
	convertedURL := []string{}

	for i := 0; i < COUNT; i++ {
		urlToTitle_temp[i] = "www.baidu.com"
	}

	for title, url := range urlToTitle_temp {
		ctx, _ := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
		cmd := exec.CommandContext(ctx, "wkhtmltopdf", url, strconv.Itoa(title)+".pdf")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = cmd.Wait()
		if err != nil {
			log.Printf("Command finished with error: %v", err)
		}
		convertedURL = append(convertedURL, strconv.Itoa(title)+".pdf")
	}
	// fmt.Printf("Done, success ratio is %d\n", len(urlToTitle_temp)/len(convertedURL))
}
