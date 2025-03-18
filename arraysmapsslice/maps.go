package arraysmapsslice

import "fmt"

func Maps() {
	fmt.Println("Maps")

	//websites := []string{"https://google.com", "https://aws.com"}
	websites := map[string]string{
		"google":             "https://google.com",
		"amzon web services": "https://aws.com",
	} /// [string] is the value type

	fmt.Println(websites["google"])
}
