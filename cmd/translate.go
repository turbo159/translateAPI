package main

/*
translateAPI

Returns name value pairs in CSV, JSON, or XML based on Content-Type.
Data source is checked in the data folder as a CSV file where the name is the API service.

Return specific value specifying the key in as the last item in the URI.
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Version of the API that will appear in the URI
const version = "v1"

// HealthCheck required by CCP health check function
type HealthCheck struct {
	Code int    `json:"healthCode"`
	Msg  string `json:"healthMsg"`
}

func main() {
	println("Starting translateAPI...")
	handleRequests()
	println("Shutting down translateAPI...")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Router
func handleRequests() {
	http.HandleFunc("/api/util/"+version+"/translation/health", health)
	http.HandleFunc("/api/util/"+version+"/translation/valueCreation", allValueCreation)
	http.HandleFunc("/api/util/"+version+"/translation/categoryBucket", allCategoryBucket)
	log.Fatal(http.ListenAndServe(":9005", nil))
}

// Health check.
func health(w http.ResponseWriter, r *http.Request) {
	exTime := time.Now()
	uri := r.RequestURI

	h := HealthCheck{http.StatusOK, "translateAPI - OK"}
	out, err := json.Marshal(h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(out))

	fmt.Print(exTime.Format("2006-01-02 15:04:05.00000"), " - ", uri)
	fmt.Println(" OK")
}

// Return all shipToSoldTo values.
func allValueCreation(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	t := r.Header.Get("Content-Type")
	var path string
	if t == "text/csv" {
		t = "csv"
		path = "/usr/local/translateAPI/data/db-value-creation.csv"
		w.Header().Add("Content-Type", "text/csv")
	} else {
		t = "json"
		path = "/usr/local/translateAPI/data/db-value-creation.json"
		w.Header().Add("Content-Type", "application/json")

	}
	dat, err := ioutil.ReadFile(path)
	check(err)

	//fmt.Fprintln(w, "All shipToSoldTo "+t)
	exTime := time.Now()

	fmt.Fprintf(w, string(dat))
	fmt.Print(exTime.Format("2006-01-02 15:04:05.00000"), " - ", uri)
	fmt.Println("  Type:", t)
}

func allCategoryBucket(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	t := r.Header.Get("Content-Type")
	var path string
	if t == "text/csv" {
		t = "csv"
		path = "/usr/local/translateAPI/data/db-category-bucket.csv"
		w.Header().Add("Content-Type", "text/csv")
	} else {
		t = "json"
		path = "/usr/local/translateAPI/data/db-category-bucket.json"
		w.Header().Add("Content-Type", "application/json")

	}
	dat, err := ioutil.ReadFile(path)
	check(err)

	//fmt.Fprintln(w, "All shipToSoldTo "+t)
	exTime := time.Now()

	fmt.Fprintf(w, string(dat))
	fmt.Print(exTime.Format("2006-01-02 15:04:05.00000"), " - ", uri)
	fmt.Println("  Type:", t)
}
