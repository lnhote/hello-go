package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}


func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

    var m1 = map[string]Vertex{
    	"Bell Labs": Vertex{
    		40.68433, -74.39967,
    	},
    	"Google": Vertex{
    		37.42202, -122.08408,
    	},
    }
    var m2 = map[string]Vertex{
    	"Bell Labs": {40.68433, -74.39967},
    	"Google":    {37.42202, -122.08408},
    }

    fmt.Println(m1)
    fmt.Println(m2)

    delete(m2, "Google")
    fmt.Println(m2)

    elem, exist := m["Google"]
    fmt.Println("The value:", elem, "Present?", exist)
}
