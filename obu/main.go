package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/AyanokojiKiyotaka8/Microservices/types"
	"github.com/gorilla/websocket"
)

const wsEndpoints = "ws://127.0.0.1:30000/ws"

var (
	sendInterval = time.Second
	r            *rand.Rand
)

func main() {
	obuIDs := generateOBUIDs(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoints, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obuIDs); i++ {
			lat, long := genLatLong()
			data := types.OBUData{
				OBUID: obuIDs[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
			fmt.Println(data)
		}
		time.Sleep(sendInterval)
	}
}

func generateOBUIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = r.Intn(math.MaxInt)
	}
	return ids
}

func genCoord() float64 {
	n := float64(r.Intn(100) + 1)
	f := r.Float64()
	return n + f
}

func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}
