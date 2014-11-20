package main

import (
	"fmt"
	groupcache "github.com/golang/groupcache"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Usage: ./test_groupcache port
	me := ":" + os.Args[1]
	peers := groupcache.NewHTTPPool("http://localhost" + me)
	peers.Set("http://localhost:8081", "http://localhost:8082", "http://localhost:8083")

	helloworld := groupcache.NewGroup("helloworld", 1024*1024*1024*16, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			log.Println(me)
			dest.SetString(me)
			return nil
		}))

	fmt.Println("GroupName: ", helloworld.Name())
	http.HandleFunc("/xbox/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.SplitN(r.URL.Path[len("/xbox/"):], "/", 1)
		if len(parts) != 1 {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		var data []byte
		helloworld.Get(nil, parts[0], groupcache.AllocatingByteSliceSink(&data))
		w.Write(data)
		log.Println("Gets: ", helloworld.Stats.Gets.String())
		log.Println("Load: ", helloworld.Stats.Loads.String())
		log.Println("LocalLoad: ", helloworld.Stats.LocalLoads.String())
		log.Println("PeerError: ", helloworld.Stats.PeerErrors.String())
		log.Println("PeerLoad: ", helloworld.Stats.PeerLoads.String())
	})

	http.ListenAndServe(me, nil)
}
