package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	//fmt.Println(resp)
	bs := make([]byte, 99999) //create byteslice that has 99999 empty bytes,
	// it needs to be large because Reader interface cannot modify slices size,
	// when it is full, it will stop reading when the slice is full instead
	resp.Body.Read(bs)
	fmt.Println("Hasil Read:", string(bs))
}
