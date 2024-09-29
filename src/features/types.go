package features

import "fmt"

// IceVehicle represents a vehicle with an internal combustion engine (ICE).
// It includes details such as the number of doors, color, engine type, number of wheels,
// fuel efficiency in kilometers per liter (kmpl), current fuel capacity level, and maximum fuel capacity.
type IceVehicle struct {
	doors         int8
	color         string
	engineType    string
	wheels        int8
	kmpl          float32
	capacityLevel float32
	maxCapacity   float32
}

// ElectricVehicle represents an electric vehicle with various attributes.
//
// Fields:
// - doors: The number of doors the vehicle has.
// - color: The color of the vehicle.
// - engineType: The type of engine the vehicle uses.
// - wheels: The number of wheels the vehicle has.
// - kwh: The kilowatt-hour rating of the vehicle's battery.
// - capacityLevel: The current capacity level of the vehicle's battery.
// - maxCapacity: The maximum capacity of the vehicle's battery.
type ElectricVehicle struct {
	doors         int8
	color         string
	engineType    string
	wheels        int8
	kwh           float32
	capacityLevel float32
	maxCapacity   float32
}

// Vehicle represents a generic vehicle interface that defines methods for
// managing and operating a vehicle's fuel or energy capacity.
//
// Methods:
//   - replenish(amt float32) float32: Adds the specified amount of fuel or energy to the vehicle and returns the new capacity.
//   - drive(km float32) float32: Drives the vehicle for the specified distance in kilometers and returns the remaining capacity.
//   - getRemainingCapacity() float32: Returns the current remaining fuel or energy capacity of the vehicle.
//   - getMilage() float32: Returns the mileage of the vehicle, typically measured in distance per unit of fuel or energy.
type Vehicle interface {
	replenish(amt float32) float32
	drive(km float32) float32
	getRemainingCapacity() float32
	getMilage() float32
}

// replenish increases the capacity level of the IceVehicle by the specified amount.
// It returns the new capacity level after replenishment.
//
// Parameters:
//   amt (float32): The amount to increase the capacity level by.
//
// Returns:
//   float32: The new capacity level after replenishment.
func (v *IceVehicle) replenish(amt float32) float32 {
	v.capacityLevel += amt
	return v.capacityLevel
}

// replenish increases the capacity level of the ElectricVehicle by the specified amount.
// It returns the updated capacity level.
//
// Parameters:
//   amt (float32): The amount to increase the capacity level by.
//
// Returns:
//   float32: The updated capacity level of the ElectricVehicle.
func (v *ElectricVehicle) replenish(amt float32) float32 {
	v.capacityLevel += amt
	return v.capacityLevel
}

// drive reduces the capacity level of the IceVehicle based on the kilometers driven.
// It takes the distance driven in kilometers as a parameter and returns the updated capacity level.
//
// Parameters:
// - km: The distance driven in kilometers.
//
// Returns:
// - The updated capacity level after driving the specified distance.
func (v *IceVehicle) drive(km float32) float32 {
	v.capacityLevel -= km / v.kmpl
	return v.capacityLevel
}

// drive reduces the capacity level of the electric vehicle based on the kilometers driven.
// It takes the distance driven in kilometers as a float32 argument and returns the updated
// capacity level as a float32. The capacity level is decreased by the ratio of kilometers
// driven to the vehicle's kWh (kilowatt-hour) efficiency.
func (v *ElectricVehicle) drive(km float32) float32 {
	v.capacityLevel -= km / v.kwh
	return v.capacityLevel
}

// getRemainingCapacity calculates and returns the remaining capacity of the IceVehicle.
// It returns the remaining capacity as a float32 value.
func (v *IceVehicle) getRemainingCapacity() float32 {
	return v.capacityLevel
}

// getRemainingCapacity returns the remaining battery capacity of the electric vehicle.
// It retrieves the current capacity level stored in the ElectricVehicle struct.
//
// Returns:
//   float32: The current capacity level of the electric vehicle.
func (v *ElectricVehicle) getRemainingCapacity() float32 {
	return v.capacityLevel
}

// getMilage returns the mileage of the IceVehicle in kilometers per liter (kmpl).
//
// Returns:
//   float32: The mileage of the vehicle.
func (v *IceVehicle) getMilage() float32 {
	return v.kmpl
}

// getMilage returns the mileage of the ElectricVehicle.
// It calculates the mileage based on the kilowatt-hours (kwh) consumed by the vehicle.
//
// Returns:
//   float32: The mileage of the ElectricVehicle.
func (v *ElectricVehicle) getMilage() float32 {
	return v.kwh
}

// canMakeTrip determines if a vehicle can make a trip of a specified distance.
// It takes a Vehicle object and a distance in kilometers as parameters.
// The function returns true if the vehicle's remaining capacity is sufficient
// to cover the distance based on its mileage, otherwise it returns false.
func canMakeTrip(v Vehicle, km float32) bool {
	return v.getRemainingCapacity() >= km/v.getMilage()
}

// CreateIceVehicle initializes and returns an IceVehicle with the specified parameters.
//
// Parameters:
//   - doors: The number of doors on the vehicle.
//   - color: The color of the vehicle.
//   - engineType: The type of engine in the vehicle.
//   - wheels: The number of wheels on the vehicle.
//   - kmpl: The kilometers per liter fuel efficiency of the vehicle.
//   - litres: The current fuel level in liters.
//   - capacity: The maximum fuel capacity in liters.
//
// Returns:
//   An IceVehicle struct populated with the provided parameters.
func CreateIceVehicle(doors int8, color string, engineType string, wheels int8, kmpl float32, litres float32, capacity float32) IceVehicle {
	return IceVehicle{doors: doors, color: color, engineType: engineType, wheels: wheels, kmpl: kmpl, capacityLevel: litres, maxCapacity: capacity}
}

// CreateElectricVehicle initializes and returns an ElectricVehicle instance with the specified parameters.
//
// Parameters:
//   - doors: The number of doors on the electric vehicle.
//   - color: The color of the electric vehicle.
//   - engineType: The type of engine in the electric vehicle.
//   - wheels: The number of wheels on the electric vehicle.
//   - kwh: The kilowatt-hour capacity of the electric vehicle's battery.
//   - charge: The current charge level of the electric vehicle's battery.
//   - capacity: The maximum capacity of the electric vehicle's battery.
//
// Returns:
//   An ElectricVehicle instance populated with the provided parameters.
func CreateElectricVehicle(doors int8, color string, engineType string, wheels int8, kwh float32, charge float32, capacity float32) ElectricVehicle {
	return ElectricVehicle{doors: doors, color: color, engineType: engineType, wheels: wheels, kwh: kwh, capacityLevel: charge, maxCapacity: capacity}
}

// DemonstrateTypes showcases the usage of structs in Go, including:
// - Creating a struct with specific values
// - Accessing and printing struct fields
// - Creating and using an anonymous struct
// - Creating a struct with default values
// - Creating a struct using a constructor function
// - Accessing and modifying struct fields using methods
// - Checking if a struct can perform a certain action using a function
func DemonstrateTypes() {
	// Creating a struct
	var car = IceVehicle{doors: 4, color: "red", engineType: "V8", wheels: 4, kmpl: 10, capacityLevel: 50, maxCapacity: 60}

	// Accessing struct fields
	fmt.Printf("Car has %d doors and is %s in color with a %s engine and %d wheels with a mileage of %f kmpl and %f litres of fuel\n", car.doors, car.color, car.engineType, car.wheels, car.kmpl, car.capacityLevel)

	var anonymousCar = struct {
		doors      int8
		color      string
		engineType string
		wheels     int8
	}{doors: 2, color: "blue", engineType: "V6", wheels: 4}

	// Accessing anonymous struct fields
	fmt.Printf("Anonymous car has %d doors and is %s in color with a %s engine and %d wheels with a mileage of %f kmpl and %f litres of fuel\n", anonymousCar.doors, anonymousCar.color, anonymousCar.engineType, anonymousCar.wheels, car.kmpl, car.capacityLevel)

	// Creating a struct with default values
	var defaultCar = IceVehicle{}
	fmt.Printf("Default car has %d doors and is %s in color with a %s engine and %d wheels with a mileage of %f kmpl and %f litres of fuel\n", defaultCar.doors, defaultCar.color, defaultCar.engineType, defaultCar.wheels, defaultCar.kmpl, defaultCar.capacityLevel)

	// Creating a struct with a constructor
	var motorbike = CreateIceVehicle(0, "black", "V2", 2, 40, 10, 20)
	fmt.Printf("Motorbike has %d doors and is %s in color with a %s engine and %d wheels with a mileage of %f kmpl and %f litres of fuel\n", motorbike.doors, motorbike.color, motorbike.engineType, motorbike.wheels, motorbike.kmpl, motorbike.capacityLevel)

	// Accessing struct methods
	fmt.Printf("Car has %f litres of fuel\n", car.capacityLevel)
	car.replenish(10)
	fmt.Printf("Car has %f litres of fuel after refueling\n", car.capacityLevel)
	car.drive(100)
	fmt.Printf("Car has %f litres of fuel after driving\n", car.capacityLevel)

	// Accessing struct methods with return values
	fmt.Printf("Car can make a trip of 100 km: %t\n", canMakeTrip(&car, 100))

}
