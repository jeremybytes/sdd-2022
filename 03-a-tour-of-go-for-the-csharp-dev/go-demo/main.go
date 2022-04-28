package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	//ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ids, err := getIds()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("IDs: %v\n", ids)

	var waiter sync.WaitGroup
	for _, currentId := range ids {
		waiter.Add(1)
		go func(id int) {
			defer waiter.Done()
			selectedPerson, err := getPerson(id)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%d: %v\n", selectedPerson.ID, selectedPerson)
		}(currentId)
	}

	waiter.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Total time: %v\n", elapsed)
}

// func fetchAndDisplay(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	selectedPerson, err := getPerson(id)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Printf("%d: %v\n", selectedPerson.ID, selectedPerson)
// }

func getIds() ([]int, error) {
	url := "http://localhost:9874/people/ids"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting data: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(
			"error fetching url (%s) - status %d", url, resp.StatusCode)
	}
	var ids []int
	err = json.NewDecoder(resp.Body).Decode(&ids)
	if err != nil {
		return nil, fmt.Errorf("error parsing data: %v", err)
	}
	return ids, nil
}

func getPerson(id int) (person, error) {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return person{}, fmt.Errorf("error getting data: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return person{}, fmt.Errorf(
			"error fetching url (%s) - status %d", url, resp.StatusCode)
	}
	var p person
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return person{}, fmt.Errorf("error parsing data: %v", err)
	}
	return p, nil
}

type person struct {
	ID         int
	GivenName  string
	FamilyName string
	StartDate  time.Time
	Rating     int
}

func (p person) String() string {
	return fmt.Sprintf("%s %s", p.GivenName, p.FamilyName)
}
