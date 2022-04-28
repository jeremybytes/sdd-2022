package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/people/ids", idHandler)
	rtr.HandleFunc("/people", peopleHandler)
	rtr.HandleFunc("/people/{id:[0-9]+}", personHandler)
	http.Handle("/", rtr)
	//	log.Fatal(http.ListenAndServe("localhost:9874", nil))

	go func() {
		if err := http.ListenAndServe("localhost:9874", nil); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("listening at http://localhost:9874/")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("shutting down")
	os.Exit(0)
}

func idHandler(w http.ResponseWriter, r *http.Request) {
	p := getPeople()
	ids := make([]int, 0, len(p))
	for _, n := range p {
		ids = append(ids, n.ID)
	}
	output, err := json.Marshal(ids)
	if err != nil {
		log.Printf("error converting to json: %v", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(output)
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	p := getPeople()
	output, err := json.Marshal(p)
	if err != nil {
		log.Printf("error converting to json: %v", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(output)
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p, err := getPerson(id)
	if err != nil {
		log.Printf("call to getPerson failed (%s%s): %v", r.Host, r.URL, err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	output, err := json.Marshal(p)
	if err != nil {
		log.Printf("error converting to json: %v", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write(output)
}

type person struct {
	ID           int       `json:"id"`
	GivenName    string    `json:"givenName"`
	FamilyName   string    `json:"familyName"`
	StartDate    time.Time `json:"startDate"`
	Rating       int       `json:"rating"`
	FormatString string    `json:"formatString"`
}

func parseDate(input string) time.Time {
	const dateForm = "2006-01-02"
	date, err := time.Parse(dateForm, input)
	if err != nil {
		log.Printf("error parsing %s to  date: %v", input, err)
	}
	return date
}

func getPerson(id int) (person, error) {
	for _, n := range getPeople() {
		if n.ID == id {
			return n, nil
		}
	}
	return person{}, fmt.Errorf("Cannot find Person ID %d", id)
}

func getPeople() (people []person) {
	people = append(people, person{1, "John", "Koenig", parseDate("1975-10-17"), 6, ""})
	people = append(people, person{2, "Dylan", "Hunt", parseDate("2000-10-02"), 8, ""})
	people = append(people, person{3, "Leela", "Turanga", parseDate("1999-03-28"), 8, "{1} {0}"})
	people = append(people, person{4, "John", "Crichton", parseDate("1999-03-19"), 7, ""})
	people = append(people, person{5, "Dave", "Lister", parseDate("1988-02-15"), 9, ""})
	people = append(people, person{6, "Laura", "Roslin", parseDate("2003-12-08"), 6, ""})
	people = append(people, person{7, "John", "Sheridan", parseDate("1994-01-26"), 6, ""})
	people = append(people, person{8, "Dante", "Montana", parseDate("2000-11-01"), 5, ""})
	people = append(people, person{9, "Isaac", "Gampu", parseDate("1977-09-10"), 4, ""})
	return
}
