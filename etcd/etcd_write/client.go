package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/hslam/stats"
	"go.etcd.io/etcd/clientv3"
	"log"
	"math/rand"
	"strings"
	"time"
)

var endpoints string
var conns int
var clients int
var total int
var bar bool

func init() {
	flag.StringVar(&endpoints, "endpoints", "127.0.0.1:2379", "host:port,host:port")
	flag.IntVar(&total, "total", 100000, "-total=100000")
	flag.IntVar(&clients, "clients", 256, "-clients=1")
	flag.IntVar(&conns, "conns", 1, "-conns=1")
	flag.BoolVar(&bar, "bar", true, "-bar=true")
	flag.Parse()
	stats.SetBar(bar)
	fmt.Printf("./client -endpoints=%s -conns=%d -clients=%d -total=%d\n", endpoints, conns, clients, total)
}

func main() {
	if conns < 1 || clients < 1 || total < 1 {
		return
	}
	var wrkClients []stats.Client
	config := clientv3.Config{
		Endpoints:   strings.Split(endpoints, ","),
		DialTimeout: 5 * time.Second,
	}
	for i := 0; i < conns; i++ {
		if client, err := clientv3.New(config); err != nil {
			log.Fatalln("dailing error: ", err)
		} else {
			wrkClients = append(wrkClients, &WrkClient{client})
		}
	}
	stats.StartPrint(clients, total, wrkClients)
}

type WrkClient struct {
	*clientv3.Client
}

func (c *WrkClient) Call() (int64, int64, bool) {
	A := RandString(4)
	B := RandString(32)
	_, err := c.Client.Put(context.TODO(), A, B)

	if err == nil {
		return int64(len(A) + len(B)), 0, true
	}
	return int64(len(A) + len(B)), 0, false
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
