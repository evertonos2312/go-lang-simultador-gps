package main

import (
	"fmt"
	route2 "github.com/codeedu/imersaofsfc2-simulator/application/route"
)

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	strinJson, _ := route.ExportJsonPosition()
	fmt.Printf(strinJson[0])
}
