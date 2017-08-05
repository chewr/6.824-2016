package viewservice

//
// see directions in pbc.go
//

import "time"
import "os"
import "fmt"

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: viewd port\n")
		os.Exit(1)
	}

	StartServer(os.Args[1])

	for {
		time.Sleep(100 * time.Second)
	}
}
