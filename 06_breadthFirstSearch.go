package main

import "fmt"

var graph = make(map[string][]string)
// I use a map of empty struct to keep track of already searched subjects, this allows
// faster lookup instead of ranging over a list of searched
var searched = make(map[string]struct{})

func main() {
	graph["me"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	bfsearch("me")
}

func bfsearch(name string) bool {
	searchQueue := graph[name]

	for len(searchQueue) > 0 {
		person := searchQueue[0]

		if _, ok := searched[person]; !ok {

			if personIsSeller(person) {
				fmt.Printf("%s is a mango seller!\n", person)

				return true
			}
				searched[person] = struct{}{}
		}
		searchQueue = append(searchQueue[1:], graph[person]...)
	}

	return false
}

func personIsSeller(name string) bool {
	return name[len(name)-1] == 'm' // use single-quotes to declare a byte or rune
}
