package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"flag"
	"log"
	"sync"
	"sync/atomic"
	"time"
	"hlj-comet/core/conn"
)

const (
	version = "0.0.1"
)

var (
	vFlag = flag.Bool("v", false, "Version info")

	nFlag   = flag.Int("n", 20000, "Num of connections")
	dFlag   = flag.Int("d", 100, "Num of dialers")
	pFlag   = flag.Int("p", 10, "Ping interval(seconds)")
	rFlag   = flag.Int("i", 3, "Max ping retry")
	urlFlag = flag.String("u",
		"ws://api.hunliji.com:6050/api/ws?token=1b399b8bca4abf8057a5a1bb23ab3b09",
		"WebSocket connection url")
	oFlag = flag.String("o", "http://localhost", "Origin")
)

var finished sync.WaitGroup
var pingInterval int
var pingRetry int

func main() {
	// 解析参数
	flag.Parse()
	if *vFlag {
		fmt.Printf("%v\n", version)
		return
	}
	url := *urlFlag
	origin := *oFlag
	numConnections := *nFlag
	numDialers := *dFlag
	pingInterval = *pFlag
	pingRetry = *rFlag

	// 运行命令
	var doneDialing sync.WaitGroup
	dialed := int64(0)

	fmt.Printf("Launching %v connections\n", numConnections)
	doneDialing.Add(numDialers)
	for j := 0; j < numDialers; j++ {
		go func() {
			for i := 0; i < numConnections/numDialers; i++ {
				ws, err := websocket.Dial(url, "", origin)
				if err != nil {
					log.Printf("Unable to dial: %v", err)
					continue
				}
				finished.Add(1)
				atomic.AddInt64(&dialed, 1)
				fmt.Printf("Current dailed: %d\n", atomic.LoadInt64(&dialed))
				go talk(ws)
			}
			doneDialing.Done()
		}()
	}

	doneDialing.Wait()
	fmt.Printf("Launched %v connections\n", atomic.LoadInt64(&dialed))
	finished.Wait()
}

func talk(ws *websocket.Conn) {
	defer func() {
		finished.Done()
	}()

	peer := conn.NewPeer(ws, conn.PeerOption{
		BufSize:      10,
		PingTry:      pingRetry,
		PingInterval: time.Duration(pingInterval) * time.Second,
		ReadDeadLine: 10 * time.Second,
	})
	peer.Run()
	log.Printf("Unable to talk")
}
