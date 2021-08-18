package main

import (
	".main.go/factory"
	".main.go/vehicle"
	"fmt"
)

const carsAmount = 100

func main() {

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	StartAssemblingProcess()
}

func StartAssemblingProcess() {
	factory := factory.New()
	ch := make(chan vehicle.Car, carsAmount)
	factory.StartAssemblingProcess(carsAmount, ch)

	for i := 0; i < carsAmount; i++ {
		car := <-ch
		log := fmt.Sprintf(`
Created car with ID: %d

-- Assembly log: %s

-- Test log: %s


`, car.Id, car.AssembleLog, car.TestingLog)
		fmt.Println(log)
	}

}
