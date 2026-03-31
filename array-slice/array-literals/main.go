package main

import "fmt"

func main() {
	var arrKeyValue = [5]string{3: "chris", 4: "ron"}
	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("arrKeyValue[%d]=%s\n", i, arrKeyValue[i])
	}
}
