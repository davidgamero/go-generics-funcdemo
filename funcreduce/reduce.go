package funcreduce

import "fmt"

type DayOfSales struct {
	soldUnits      int
	revenueDollars float32
}

type CarModel struct {
	name        string
	maxSpeedMPH int
}

// RunReduceDemo executes the `Reduce Sales Unit Total Demo`
func RunReduceDemo() {
	fmt.Println("------Reduce Sales Unit Total Demo------")
	daysOfSales := []DayOfSales{
		{
			soldUnits:      5,
			revenueDollars: 100.00,
		},
		{
			soldUnits:      0,
			revenueDollars: 0.00,
		},
		{
			soldUnits:      5,
			revenueDollars: 400.00,
		},
		{
			soldUnits:      10,
			revenueDollars: 500.00,
		},
	}

	soldUnitsSum := func(day DayOfSales, totalSales int) int {
		return totalSales + day.soldUnits
	}
	totalUnitsSold := Reduce[DayOfSales, int](daysOfSales, soldUnitsSum, 0)
	fmt.Println("All Days of Sales", daysOfSales)
	fmt.Printf("Total Units Sold: %d\n", totalUnitsSold)
	fmt.Println("")
}

// RunReduceDemo2 executes the `Reduce Car Max Speed Demo`
func RunReduceDemo2() {
	fmt.Println("------Reduce Car Max Speed Demo------")

	allCarModels := []CarModel{
		{
			name:        "car1",
			maxSpeedMPH: 100,
		},
		{
			name:        "car2",
			maxSpeedMPH: 70,
		},
		{
			name:        "car3",
			maxSpeedMPH: 150,
		},
	}
	compareCarToPreviousMaxSpeed := func(car CarModel, prevMaxSpeed int) int {
		if car.maxSpeedMPH > prevMaxSpeed {
			return car.maxSpeedMPH
		}
		return prevMaxSpeed
	}
	initialMaxSpeed := 0
	maxCarSpeed := Reduce[CarModel, int](allCarModels, compareCarToPreviousMaxSpeed, initialMaxSpeed)
	fmt.Println("All car models: ", allCarModels)
	fmt.Println("Max Car Speed: ", maxCarSpeed)
	fmt.Println("")
}

// Reduce executes a "reducer" function on each element of the slice of ItemType, in order, passing in the return value from the reducer on the preceding element.
// The final result of running the reducer across all elements of the array is a single value of type AggregatorType.
// Adapted from https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reduce
func Reduce[ItemType any, AggregatorType any](items []ItemType,
	reducer func(item ItemType, aggregator AggregatorType) AggregatorType,
	initialAggregator AggregatorType) AggregatorType {

	aggregator := initialAggregator

	for _, thisItem := range items {
		aggregator = reducer(thisItem, aggregator)
	}
	return aggregator
}
