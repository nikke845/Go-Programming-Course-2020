package main

import "fmt"
import "net/http"
import "encoding/json"
import "io/ioutil"
import "flag"
import "os"


func makeUrl()(string,string){
	var latitude string
	var longitude string
	var filename string
	flag.StringVar(&latitude, "latitude", "60.192059", "City's latitude")
	flag.StringVar(&longitude, "longitude", "24.945831", "City's longitude")
	flag.StringVar(&filename, "file", "", "Filename for optional data logging")
	flag.Parse()
	completeUrl:= fmt.Sprintf("%v%v%v%v%v", "https://api.darksky.net/forecast/785d4ef0e760c496fe6fbd387bdbbc2c/", latitude,",", longitude,"?units=auto&exclude=minutely,hourly,daily,alerts,flags")
	return completeUrl, filename
}

func getWheather(url string)[]byte{
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	byt := []byte(body)
	return byt
}

func printWheather(b []byte)string{
	var dat map[string]interface{}
	if err := json.Unmarshal(b, &dat); err != nil {
		panic(err)
	}

	currently := dat["currently"].(map[string]interface{})

	temp := (currently["temperature"])
	ct := fmt.Sprintf("%v℃", temp)
	fmt.Println("Current temperature:", ct)

	precipP := (currently["precipProbability"])
	ppc := fmt.Sprintf("%v%%", precipP)
	fmt.Println("Chance of rain:", ppc)

	wind := (currently["windSpeed"])
	wc := fmt.Sprintf("%v㎧", wind)
	fmt.Println("Current wind:", wc)

	summary := (currently["summary"])
	sc := fmt.Sprintf("%v", summary)
	fmt.Println("Summary:", sc)

	writeToFile := fmt.Sprintf("%v%v%v%v%v%v%v%v","Temperature: ", ct, " Chance of rain: ", ppc, " Wind Speed: ", wc, " Summary: ", sc)
	return writeToFile
}

func writeToFile(filename string, a string){
	if filename == ""{
		return
	} else {
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    	if err != nil {
        	fmt.Println(err)
        	return
    	}
    	_, err = fmt.Fprintln(f, a)
    	if err != nil {
        	fmt.Println(err)
            f.Close()
        	return
    	}
    	err = f.Close()
    	if err != nil {
        	fmt.Println(err)
        	return
		}
	}
}
func main(){
	url, filename := makeUrl()
	wheather := getWheather(url)
	data := printWheather(wheather)
	writeToFile(filename,data)
	fmt.Println("-------------------")
	fmt.Println("Powered by Dark Sky")
	fmt.Println("https://darksky.net/poweredby/")
}