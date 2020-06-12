package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"math"
)

func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getHandler)
	router.HandleFunc("/radius/", radiusHandler)
	router.HandleFunc("/location/", locationHandler)
	log.Fatal(http.ListenAndServe(":9000", router))
}


func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func defaultHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world")
}

func radiusHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello radius")
}

func locationHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello location")
}

type Users struct{
	Users []User `json:"users"`
}

type User struct{
	Id int `json:"id`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Email string `json:"email"`
	Ip_address string `json:"ip_address"`
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

    fmt.Println("GET... users within 50miles")
    resp, err := http.Get("https://bpdts-test-app.herokuapp.com/users")
    if err != nil {
        log.Fatalln(err)
    }

	defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
	}

	var users Users

	json.Unmarshal(body, users)

	// var arr =  data.filter(function(person) {
    //     var distance = GetDistanceBetweenTwoCoordinates(51.509865, -0.118092, person.latitude, person.longitude);
    //     distance < radius ? person['distance'] = (distance * 0.62137).toFixed(2) : null;
    //     return distance <= radius;
	//   }).sort((a,b) => { return a.distance - b.distance; });
	  
	// fmt.Println(users)
	// s := []Users
    // for key := range users User {
    //     fmt.Println(key)
    // }


	
	json.NewEncoder(w).Encode(string(body))
}

// func FilterMagicalCreatures(record string) bool {
// 	magicalCreatures := []string{
// 		"Unicorn",
// 		"Dragon",
// 		"Griffin",
// 		"Minotaur",
// 	}

// 	for _, c := range magicalCreatures {
// 		if record == c {
// 			return false
// 		}
// 	}

// 	return true
// }

func GetDistanceBetweenTwoCoordinates(lat1 float64,lon1 float64,lat2 float64,lon2 float64) float64{
	var radius float64 = 6371;
  
	var dLat = ConvertToRadian(lat2-lat1); 
	var dLon = ConvertToRadian(lon2-lon1); 
  
	var a = 
	  math.Sin(dLat/2) * math.Sin(dLat/2) +
	  math.Cos(ConvertToRadian(lat1)) * math.Cos(ConvertToRadian(lat2)) * 
	  math.Sin(dLon/2) * math.Sin(dLon/2); 
  
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a)); 
	var d = radius * c;
	return d;
  }

  func ConvertToRadian(degrees float64) float64 {
	return degrees * (math.Pi/180)
  }
