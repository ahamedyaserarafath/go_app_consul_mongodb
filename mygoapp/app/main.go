package main

import (
    "encoding/json"
    "log"
    "net/http"
    "gopkg.in/mgo.v2"
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

// Json stuture in DB
type Greeting struct {
        Id string `bson:"id"`
        Title string `bson:"title"`
}
var allResult []Greeting

type consulService struct {
    Node string
    Address string
    ServicePort int
}

// Function helps to get all the greeting string from database
func getGreetingString(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside getGreetingString function with url %s",r.URL.Path[1:])

	mongodbHostName,mongodbPort := getMongodbIPFromConsul()
        if mongodbHostName == "" {
		fmt.Fprintf(w, "Please wait consul need to sync with other agents,please refresh after 5sec..")
	}else{
	   session, err := mgo.Dial(mongodbHostName+":"+mongodbPort)
	   if err != nil {
	        panic(err)
	   }
	   defer session.Close()
	   log.Printf("Connected with mongodb database %s : %s",mongodbHostName,mongodbPort)

	   collection := session.DB(mongodbDatabaseName).C(mongodbCollectionName)

	   err = collection.Find(nil).All(&allResult)
	   if err != nil {
	        log.Fatal(err)
	   }
	   fmt.Fprintf(w, "%s", allResult[0].Title)
	   log.Printf("Greeging string  from database : %s ",allResult[0].Title)
        }
}

func getMongodbIPFromConsul()(string, string) {
    log.Printf("Inside getMongodbIPFromConsul function")
    serviceUrl := "http://"+ consulIP +":" + consulHTTPPort + 
    		"/v1/catalog/service/" + serviceName
    response, err := http.Get(serviceUrl)
    if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
	if strings.TrimSuffix(string(data),"\n") != "[]"{
           var consulservice []consulService
           json.Unmarshal([]byte(data), &consulservice)
           mongodbHostName :=  consulservice[0].Address
           mongodbPort :=  strconv.Itoa(consulservice[0].ServicePort)
           return mongodbHostName,mongodbPort
	}
    }
    return "",""
}


func main() {
      http.HandleFunc("/", getGreetingString)
      log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}

