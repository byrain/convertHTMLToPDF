package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/byrain/convertHTMLToPDF/crawler"
)

const (
	// TIMEOUT timeout Second
	TIMEOUT = 90
)

func main() {
	urlToTitleMap := crawler.ExampleScrape()

	var wg sync.WaitGroup
	wg.Add(len(urlToTitleMap))

	for url, title := range urlToTitleMap {
		ctx, _ := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
		go func(url, title string) {
			defer func() {
				wg.Done()
			}()
			cmd := exec.CommandContext(ctx, "wkhtmltopdf", url, title+".pdf")
			err := cmd.Start()
			if err != nil {
				log.Printf(err.Error())
			}
			log.Printf("Waiting for command to finish...")
			fmt.Println(cmd.Args)
			err = cmd.Wait()
			if err != nil {
				log.Printf("Command finished with error: %v", err)
			}
		}(url, title)
	}
	wg.Wait()
	fmt.Println("Done")
}
