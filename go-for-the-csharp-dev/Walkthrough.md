# Walkthrough

* package
* func main
* braces

```go
package main

func main() {

}
```

* import
* fmt.Println
* no semi-colons (removed on save)

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

* go build
* go mod init

```
PS C:\GoDemo\async> go build
go: go.mod file not found in current directory or any parent directory; see 'go help modules'
```

```
PS C:\GoDemo\async> go mod init github.com/jeremybytes/tour-of-go
```

* go.mod file
```
module github.com/jeremybytes/tour-of-go

go 1.18
```

* compiled output
* first run

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> ls

    Directory: C:\GoDemo\async

   Length Name
   ------ ----
       50 go.mod
       81 main.go
  1891840 tour-of-go.exe
```

```
PS C:\GoDemo\async> .\tour-of-go.exe
Hello, World!
```

* variable declaration/assignment
* unused variables

```go
// build failure (unused variable)
func main() {
	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Hello, World!")
}
```

```
PS C:\GoDemo\async> go build
# github.com/jeremybytes/tour-of-go
.\main.go:6:2: ids declared but not used
```

* fmt.Printf

```go
func main() {
	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
```

* struct

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)
}

type person struct {
	ID         int
	GivenName  string
	FamilyName string
	StartDate  time.Time
	Rating     int
}
```

* func parameters / return
* capitalization

```go
// This step does not build (getPerson has no return)
package main

import (
	"fmt"
	"time"
)

func main() {
	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)
}

func getPerson(id int) person {
	
}

type person struct {
	ID         int
	GivenName  string
	FamilyName string
	StartDate  time.Time
	Rating     int
}
```

* Sprintf / Fprintf / Errorf

```go
func getPerson(id int) person {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
}
```

* multiple return
* multiple assignment

```go
func getPerson(id int) person {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, err := http.Get(url)
}
```

* blank identifier / discard

```go
func getPerson(id int) person {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, _ := http.Get(url)
}
```

* interface (io.Reader)
* generics
* variable declaration
* pointer

```go
func getPerson(id int) person {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, _ := http.Get(url)
	var p person
	json.NewDecoder(resp.Body).Decode(&p)
	return p
}
```

* io.Closer

```go
func getPerson(id int) person {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, _ := http.Get(url)
	var p person
	json.NewDecoder(resp.Body).Decode(&p)
	resp.Body.Close()
	return p
}
```

* call getPerson

```go
func main() {
	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)

	p := getPerson(2)
	fmt.Printf("%d: %v\n", p.ID, p)
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
2: {2 Dylan Hunt 2000-10-02 00:00:00 -0700 PDT 8}
```

* interface (fmt.Stringer)
* method receiver

```go
func (p person) String() string {
	return fmt.Sprintf("%s %s", p.GivenName, p.FamilyName)
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
2: Dylan Hunt
```

* for loop (index)
* len()
* no parens

```go
func main() {
	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)

	for i := 0; i < len(ids); i++ {
		p := getPerson(ids[i])
		fmt.Printf("%d: %v\n", p.ID, p)
	}
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
1: John Koenig
2: Dylan Hunt
3: Leela Turanga
4: John Crichton
5: Dave Lister
6: Laura Roslin
7: John Sheridan
8: Dante Montana
9: Isaac Gampu
0:
```

* for loop (range)

```go
	for i := range ids {
		p := getPerson(ids[i])
		fmt.Printf("%d: %v\n", p.ID, p)
	}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
1: John Koenig
2: Dylan Hunt
3: Leela Turanga
4: John Crichton
5: Dave Lister
6: Laura Roslin
7: John Sheridan
8: Dante Montana
9: Isaac Gampu
0:
```

* for (foreach)

```go
	for _, id := range ids {
		p := getPerson(id)
		fmt.Printf("%d: %v\n", p.ID, p)
	}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
1: John Koenig
2: Dylan Hunt
3: Leela Turanga
4: John Crichton
5: Dave Lister
6: Laura Roslin
7: John Sheridan
8: Dante Montana
9: Isaac Gampu
0:
```

* time.Now
* time.Since

```go
func main() {
	start := time.Now()

	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)

	for _, id := range ids {
		p := getPerson(id)
		fmt.Printf("%d: %v\n", p.ID, p)
	}

	elapsed := time.Since(start)
	fmt.Printf("Total time: %v\n", elapsed)
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
1: John Koenig
2: Dylan Hunt
3: Leela Turanga
4: John Crichton
5: Dave Lister
6: Laura Roslin
7: John Sheridan
8: Dante Montana
9: Isaac Gampu
0:
Total time: 10.1133401s
```

* extract function

```go
func main() {
	start := time.Now()

	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)

	for _, id := range ids {
		fetchAndDisplay(id)
	}

	elapsed := time.Since(start)
	fmt.Printf("Total time: %v", elapsed)
}

func fetchAndDisplay(id int) {
	p := getPerson(id)
	fmt.Printf("%d: %v\n", p.ID, p)
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
1: John Koenig
2: Dylan Hunt
3: Leela Turanga
4: John Crichton
5: Dave Lister
6: Laura Roslin
7: John Sheridan
8: Dante Montana
9: Isaac Gampu
0:
Total time: 10.1112296s
```

* concurrent operations
* goroutine

```go
func main() {
	start := time.Now()

	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)

	for _, id := range ids {
		go fetchAndDisplay(id)
	}

	elapsed := time.Since(start)
	fmt.Printf("Total time: %v", elapsed)
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
Total time: 0s
```

* sync.WaitGroup
* Add / Done / Wait

```go
// does not work as expected
func main() {
	start := time.Now()

	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)

	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go fetchAndDisplay(id, wg)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Total time: %v", elapsed)
}

func fetchAndDisplay(id int, wg sync.WaitGroup) {
	p := getPerson(id)
	fmt.Printf("%d: %v\n", p.ID, p)
	wg.Done()
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
0:
3: Leela Turanga
1: John Koenig
7: John Sheridan
9: Isaac Gampu
2: Dylan Hunt
4: John Crichton
5: Dave Lister
6: Laura Roslin
8: Dante Montana
```

* pointers

```go
func main() {
	start := time.Now()

	ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("IDs: %v\n", ids)

	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go fetchAndDisplay(id, &wg)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Total time: %v", elapsed)
}

func fetchAndDisplay(id int, wg *sync.WaitGroup) {
	p := getPerson(id)
	fmt.Printf("%d: %v\n", p.ID, p)
	wg.Done()
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
0:
6: Laura Roslin
5: Dave Lister
2: Dylan Hunt
3: Leela Turanga
9: Isaac Gampu
7: John Sheridan
1: John Koenig
4: John Crichton
8: Dante Montana
Total time: 1.0126087s
```

* panic (stop service & run)

```
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x48 pc=0xb9fd20]

goroutine 12 [running]: ...
```

* error
* conditionals (no parens)
* nil

```go
func getPerson(id int) person {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		
	}
	var p person
	json.NewDecoder(resp.Body).Decode(&p)
	resp.Body.Close()
	return p
}
```

* multiple return

```go
func getPerson(id int) (person, error) {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return person{}, err
	}
	var p person
	json.NewDecoder(resp.Body).Decode(&p)
	resp.Body.Close()
	return p, nil
}

func fetchAndDisplay(id int, wg *sync.WaitGroup) {
	p, _ := getPerson(id)
	fmt.Printf("%d: %v\n", p.ID, p)
	wg.Done()
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
0:
0:
0:
0:
0:
0:
0:
0:
0:
0:
Total time: 2.3505089s
```

* print error

```go
func fetchAndDisplay(id int, wg *sync.WaitGroup) {
	p, err := getPerson(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d: %v\n", p.ID, p)
	wg.Done()
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
Get "http://localhost:9874/people/2": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
Get "http://localhost:9874/people/8": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
...
```

* defer (wg.Done)

```go
func fetchAndDisplay(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	p, err := getPerson(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d: %v\n", p.ID, p)
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
Get "http://localhost:9874/people/10": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
Get "http://localhost:9874/people/1": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
...
Total time: 2.3402189s
```

* fmt.Errorf

```go
func getPerson(id int) (person, error) {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return person{}, fmt.Errorf("error getting data: %v", err)
	}
	var p person
	json.NewDecoder(resp.Body).Decode(&p)
	resp.Body.Close()
	return p, nil
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
error getting data: Get "http://localhost:9874/people/3": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
error getting data: Get "http://localhost:9874/people/2": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
...
Total time: 2.3219417s
```

* log error

```go
func fetchAndDisplay(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	p, err := getPerson(id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%d: %v\n", p.ID, p)
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
2022/03/16 14:01:50 error getting data: Get "http://localhost:9874/people/5": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
2022/03/16 14:01:50 error getting data: Get "http://localhost:9874/people/9": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
...
Total time: 2.3305185s
```

* restart service

```
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
0:
7: John Sheridan
9: Isaac Gampu
5: Dave Lister
8: Dante Montana
6: Laura Roslin
4: John Crichton
2: Dylan Hunt
3: Leela Turanga
1: John Koenig
Total time: 1.0876394s
```

* parse error (id 10)
* fmt.Errorf
* defer (resp.Body.Close())

```go
func getPerson(id int) (person, error) {
	url := fmt.Sprintf("http://localhost:9874/people/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return person{}, fmt.Errorf("error getting data: %v", err)
	}
	defer resp.Body.Close()
	var p person
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return person{}, fmt.Errorf("error parsing data: %v", err)
	}
	return p, nil
}
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
2022/03/16 14:04:06 error parsing data: EOF
7: John Sheridan
5: Dave Lister
6: Laura Roslin
2: Dylan Hunt
1: John Koenig
9: Isaac Gampu
3: Leela Turanga
8: Dante Montana
4: John Crichton
Total time: 1.0108326s
```

* HTTP 204

```go
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
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
2022/03/16 14:05:14 error fetching url (http://localhost:9874/people/10) - status 204
1: John Koenig
5: Dave Lister
3: Leela Turanga
6: Laura Roslin
7: John Sheridan
2: Dylan Hunt
4: John Crichton
9: Isaac Gampu
8: Dante Montana
Total time: 1.0240546s
```

---
## Additional Features

* anonymous method (inline fetchAndDisplay)

```go
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func (id int, wg *sync.WaitGroup) {
			defer wg.Done()
			p, err := getPerson(id)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%d: %v\n", p.ID, p)
		}(id, &wg)
	}

	wg.Wait()
```

* closure (wg)

```go
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func (id int) {
			defer wg.Done()
			p, err := getPerson(id)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%d: %v\n", p.ID, p)
		}(id)
	}

	wg.Wait()
```

* closure on index (id)

```go
	// does not work as expected
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func () {
			defer wg.Done()
			p, err := getPerson(id)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%d: %v\n", p.ID, p)
		}()
	}

	wg.Wait()
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
2022/03/16 14:09:37 error fetching url (http://localhost:9874/people/10) - status 204
Total time: 1.0076644s
```

* put indexer back (id)

```go
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			p, err := getPerson(id)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%d: %v\n", p.ID, p)
		}(id)
	}

	wg.Wait()
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9 10]
4: John Crichton
2022/03/16 14:10:39 error fetching url (http://localhost:9874/people/10) - status 204
8: Dante Montana
1: John Koenig
6: Laura Roslin
2: Dylan Hunt
3: Leela Turanga
5: Dave Lister
7: John Sheridan
9: Isaac Gampu
Total time: 1.0148639s
```

* add getIDs (copy getPerson and change types)

```go
func getIDs() ([]int, error) {
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
```

* call and check error

```go
	//ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ids, err := getIDs()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("IDs: %v\n", ids)
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9]
6: Laura Roslin
4: John Crichton
2: Dylan Hunt
7: John Sheridan
9: Isaac Gampu
5: Dave Lister
1: John Koenig
3: Leela Turanga
8: Dante Montana
Total time: 1.0141731s
```

* stop service

```
PS C:\GoDemo\async> .\tour-of-go.exe
2022/03/16 14:15:25 error getting data: Get "http://localhost:9874/people/ids": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
IDs: []
Total time: 2.3365592s
```

* os.Exit

```go
	//ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ids, err := getIDs()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("IDs: %v\n", ids)
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
2022/03/16 14:16:17 error getting data: Get "http://localhost:9874/people/ids": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
```

* fmt.Fatalln

```go
	//ids := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ids, err := getIDs()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("IDs: %v\n", ids)
```

```
PS C:\GoDemo\async> go build
PS C:\GoDemo\async> .\tour-of-go.exe
2022/03/16 14:18:25 error getting data: Get "http://localhost:9874/people/ids": dial tcp [::1]:9874: connectex: No connection could be made because the target machine actively refused it.
```

* start service

```
PS C:\GoDemo\async> .\tour-of-go.exe
IDs: [1 2 3 4 5 6 7 8 9]
6: Laura Roslin
1: John Koenig
5: Dave Lister
9: Isaac Gampu
2: Dylan Hunt
8: Dante Montana
3: Leela Turanga
4: John Crichton
7: John Sheridan
Total time: 1.1198255s
```

*** **END** ***
