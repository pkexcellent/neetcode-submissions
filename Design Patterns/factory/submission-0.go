type Vehicle interface {
    GetType() string
}

type Car struct{}

func (c *Car) GetType() string {
    return "Car"
}

type Bike struct{}

func (b *Bike) GetType() string {
    return "Bike"
}

type Truck struct{}

func (t *Truck) GetType() string {
    return "Truck"
}

type VehicleFactory interface {
    CreateVehicle() Vehicle
}

type CarFactory struct{}

func (cf *CarFactory) CreateVehicle() Vehicle {
    return &Car{}

}

type BikeFactory struct{}

func (bf *BikeFactory) CreateVehicle() Vehicle {
    return &Bike{}
}

type TruckFactory struct{}

func (tf *TruckFactory) CreateVehicle() Vehicle {
    return &Truck{}
}
