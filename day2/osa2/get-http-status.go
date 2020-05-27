package main

import "fmt"
import "net/http"
import "flag"

func main(){
	url := getFlag()
	checkConnection(url)
}

func getFlag()string{
	var site string
	flag.StringVar(&site, "url", "https://vuorivirta.fi/", "Website url for response status test")
	flag.Parse()
	return site
}

func checkConnection(site string)string{
	resp, err := http.Get(site)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(site,"status:",resp.Status)
	status := resp.Status
	return status
}