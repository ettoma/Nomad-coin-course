package main

import (
	"fmt"

	"github.com/ettoretoma/Nomad-coin-course/person"
)

func main() {
	ettore := person.Person{}
	ettore.SetDetails("Ettore", 32)

	fmt.Println(ettore)

}
