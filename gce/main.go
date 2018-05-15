package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if err := handleRoot2(w, r); err != nil {
		http.Error(w, err.Error(), 500)
		log.Printf("%v", err)
	}
}

func handleRoot2(w http.ResponseWriter, r *http.Request) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/disks/", nil)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	req.Header.Add("Metadata-Flavor", "Google")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	w.Write(body)
	return nil
}
