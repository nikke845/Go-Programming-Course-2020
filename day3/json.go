package main

import "fmt"
import "net/http"
import "flag"
import "io/ioutil"
import "bytes"
import "encoding/json"

func main(){
	url := "https://work-hours-management-app.herokuapp.com/paycycles/user1/"
	paycycle := getFlag()
	getSite(url+paycycle)
}

func getFlag()string{
	var paycycle string
	flag.StringVar(&paycycle, "paycycle", "Joulukuu", "Paycycle name")
	flag.Parse()
	return paycycle
}

func getSite(site string){
	resp, err := http.Get(site)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	trimmedBody := bytes.Trim(body, "[]")
	
	byt := []byte(trimmedBody)
	
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	name := dat["name"].(string)
	fmt.Println("Palkkakauden nimi:", name)	
	
	workdays := dat["workdays"].([]interface{})

	numberOfWorkdays := (len(dat["workdays"].([]interface{})))
	fmt.Println("Työpäiviä tällä kaudella:",numberOfWorkdays)
	
	for i := 0; i < numberOfWorkdays; i++{
		fmt.Println(workdays[i])
	}
}