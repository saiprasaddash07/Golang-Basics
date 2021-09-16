package main

import "fmt"

func main() {
	websites := map[string]string{
		"google": "https://google.com",
		"aws":    "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])
	websites["x"] = "y"
	websites["z"] = "w"

	delete(websites,"x");
	delete(websites,"z");
}
