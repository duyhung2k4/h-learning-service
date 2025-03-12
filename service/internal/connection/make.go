package connection

import (
	constant "app/internal/constants"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/time/rate"
)

func makeVariable() {
	mapSocket = make(map[string]*websocket.Conn)
	mapSocketEvent = make(map[string]map[string]*websocket.Conn)

	// chanel job
	emailChan = make(chan EmailJob_MessPayload)

	// http
	limiter = rate.NewLimiter(rate.Every(time.Second), 500)
}

func makeFolder() {
	folderNames := []string{
		string(constant.MP4),
		string(constant.VIDEO),
	}

	var wg sync.WaitGroup

	for _, f := range folderNames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			permissions := os.ModePerm
			err := os.MkdirAll(f, permissions)
			if err != nil {
				fmt.Printf("Error create folder: %v\n", err)
				return
			}
		}(f)
	}

	wg.Wait()
}
