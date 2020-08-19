package main

import (
	"go-listofstructs/specs"
	"fmt"	
)

func main ()  {
	devices := make([]specs.Computer,5)

	for i:= range devices {
		devices[i].Mfg = "Apple"
		devices[i].Model = "Mackbook Pro"
		devices[i].Price = 2000 + 100 + i
		devices[i].Year = 2020 - i
		devices[i].Configurations.Hdd = "500GB" // uses sub structs
		devices[i].Configurations.Memory = "16GB"
		devices[i].Configurations.Processor = "i7"		
		}

	for j := range devices {
		fmt.Println (" Value of devices stored :: " , devices[j])
	}


}
