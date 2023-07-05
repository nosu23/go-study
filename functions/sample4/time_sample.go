package main

import (
	"fmt"
	"time"
)

func main() {
	nowT := time.Now()
	//nowTStr := nowT.Format("2006-01-02")
	nowTStr := nowT.Format("2006-01-02-150405")
	fmt.Printf("%s\n", nowTStr)
}
