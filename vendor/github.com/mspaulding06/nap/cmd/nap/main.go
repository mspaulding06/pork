package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mspaulding06/nap"
)

var api = nap.NewAPI("https://httpbin.org")

func main() {
	list := flag.Bool("list", false, "Get list of all API resources")
	flag.Parse()
	if *list {
		fmt.Println("Available Resources:")
		for _, name := range api.ResourceNames() {
			fmt.Println(name)
		}
		return
	}
	resource := os.Args[1]
	if err := api.Call(resource, nil); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	router := nap.NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(content))
		return nil
	})
	api.AddResource("get", nap.NewResource("/get", "GET", router))
}
