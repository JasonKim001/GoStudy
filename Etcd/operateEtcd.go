package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	// config etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	go func() {
		rch := cli.Watch(context.Background(), "foo")
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Println("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()

	// operate etcd example
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// set key-value
	_, err = cli.Put(ctx, "foo", "caibi")
	if err != nil {
		log.Fatal(err)
	}

	// get key-value
	resp, err := cli.Get(ctx, "foo")
	if err != nil {
		log.Fatal(err)
	}

	for _, ev := range resp.Kvs {
		fmt.Println("%s: %s\n", ev.Key, ev.Value)
	}

}
