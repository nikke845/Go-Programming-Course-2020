package main

import "fmt"
import "net/http"
import "encoding/json"
import "io/ioutil"
import "flag"
import "os"
import "time"
import "github.com/dariubs/percent"
import "strings"

func readApiKey()string{
	bytes, err := ioutil.ReadFile(".env")
	if err != nil {
		panic(err)
	}
	a := (string(bytes))
	key := strings.TrimSuffix(a, "\n")
	return key
}

func makeUrl(apiKey string)(string,string){
	var latitude string
	var longitude string
	var filename string
	flag.StringVar(&latitude, "latitude", "60.192059", "City's latitude")
	flag.StringVar(&longitude, "longitude", "24.945831", "City's longitude")
	flag.StringVar(&filename, "file", "", "Filename for optional data logging")
	flag.Parse()
	completeUrl:= fmt.Sprintf("%v%v%v%v%v%v%v", "https://api.darksky.net/forecast/",apiKey,"/", latitude,",", longitude,"?units=auto&exclude=minutely,hourly,daily,alerts,flags")
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
	f := precipP.(float64)
	converted := percent.PercentOfFloat(f, 1)
	ppc := fmt.Sprintf("%v%%", converted)
	fmt.Println("Chance of rain:", ppc)

	wind := (currently["windSpeed"])
	wc := fmt.Sprintf("%v㎧", wind)
	fmt.Println("Current wind:", wc)

	summary := (currently["summary"])
	sc := fmt.Sprintf("%v", summary)
	fmt.Println("Summary:", sc)

	writeToFile := fmt.Sprintf("%v%v%v%v%v%v%q",temp,",",converted,",",wind,",",summary)
	return writeToFile
}

func dateTime()string{
	const(
		RFC3339 = "2006-01-02T15:04:05Z07:00"
	)
	t := time.Now().Format(RFC3339)
	timeDate := fmt.Sprintf("%q" ,t)
	return timeDate
}

func writeToFile(filename string, a string, t string){
	if filename == ""{
		return
	} else {

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    		if err != nil {
        	fmt.Println(err)
        	return
			}
			firstline := "Time,Temperature,RainProbability,WindSpeed,Summary"
    		_, err = fmt.Fprintln(f, firstline)
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



		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
    	if err != nil {
        	fmt.Println(err)
        	return
		}
		line := fmt.Sprintf("%v%v%v", t , ",", a)
    	_, err = fmt.Fprintln(f, line)
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
	key := readApiKey()
	url, filename := makeUrl(key)
	wheather := getWheather(url)
	data := printWheather(wheather)
	dateTime := dateTime()
	writeToFile(filename, data, dateTime)
	fmt.Println("-------------------")
	fmt.Println("Powered by Dark Sky")
	fmt.Println("https://darksky.net/poweredby/")
}