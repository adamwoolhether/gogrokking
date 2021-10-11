package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var voted map[string]bool
var cache map[string][]byte

func main() {
	// Lookups
	phoneBook := make(map[string]int)

	phoneBook["adam"] = 846257833445
	phoneBook["nikki"] = 59234567895

	fmt.Println(phoneBook["adam"])

	// Preventing duplicates
	voted = make(map[string]bool)
	voted["adam"] = true
	checkVoter("adam")
	checkVoter("nikki")
	checkVoter("nikki")

	// Cache
	cache = make(map[string][]byte)
	get1 := getPage("https://www.google.com")
	get2 := getPage("https://www.wikipedia.com")
	get3 := getPage("https://www.google.com")
	fmt.Println(string(get1))
	fmt.Println(string(get2))
	fmt.Println(string(get3))

}

func checkVoter(name string) bool {
	if voted[name] {
		fmt.Printf("%s already voted!\n", name)
		return false
	} else {
		fmt.Printf("%s hasn't voted yet, let them vote!\n", name)
		voted[name] = true // add them to the list
		return true
	}
}

func getPage(url string) []byte {
	if cache[url] != nil {
		fmt.Printf("%s already in cache, here is the data...\n", url)

		return cache[url]
	} else {
		fmt.Printf("%s not in cache, loading it now...\n", url)

		resp, err := http.Get(url)
		if err != nil {
			return nil
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil
		}

		cache[url] = data

		return data
	}
}
