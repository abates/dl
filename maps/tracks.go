package main

import "fmt"
import "github.com/serjvanilla/go-overpass"

func main() {
	client := overpass.New()

	//Retrive relation with all its members, recursively.
	result, _ := client.Query(`
[out:json];
(
  // query part for: “highway=raceway”
  node["highway"="raceway"];
  way["highway"="raceway"];
  relation["highway"="raceway"];
);
out body;
>;
`)
	//Take a note that you shoud use "[out:json]" in your query for correct work.
	fmt.Printf("Result: %v\n", result)
}
