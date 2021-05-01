package main

import (
	"rogchap.com/v8go"
	"fmt"
	"net/http"
	"sync"
	// "encoding/json"
	"io/ioutil"
	// "io/fs"
	"io"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	
	// resp, err := http.Get("https://deckofcardsapi.com/api/deck/new/draw/?count=10")
	// resp, err := http.Get("https://ll.thespacedevs.com/2.0.0/event/")
	resp, err := http.Get("https://lldev.thespacedevs.com/2.0.0/event/") //dev endpoint has not ratelimits but isn't up to date
	check(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	file, err := os.Create("/tmp/dat2") //temp data file
    check(err)
	defer file.Close()
	
	length := 1 //intended to specifiy the number of go routines to execute
	fmt.Printf("Running %d isolates...\n",length)
	var wg sync.WaitGroup
	
	for i:=0; i<length; i++ {
		if err != nil {
			panic(err.Error())
		}

		wg.Add(1)
		go execute(string(body), &wg, file)
	}

	wg.Wait()
}

func execute(response string, wg *sync.WaitGroup, file *os.File) {
	defer wg.Done()

	iso, _ := v8go.NewIsolate() 

	//call back to write a byteslice to the passed flie
	filewrite, _ := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		data := []byte(info.Args()[0].String())
		numb, err := file.Write(data)
		check(err)
		fmt.Printf("wrote %d bytes\n", numb)

		file.Sync()

		return nil
	})

	global, _ := v8go.NewObjectTemplate(iso)
	global.Set("print", filewrite)

	util, _ := ioutil.ReadFile("util.js")
	yaml, _ := ioutil.ReadFile("node_modules/yaml")
	ctx, _ := v8go.NewContext(iso, global) 
	
	ctx.RunScript(string(util), "util.js") 
	var scr string ="const result = parse("+response+")"
	ctx.RunScript(scr, "main.js") 
	ctx.RunScript("result", "value.js") 

}
