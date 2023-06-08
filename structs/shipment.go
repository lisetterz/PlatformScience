package structs

type Shipment struct {
	Address string
	Driver  string
}

type File struct {
	Shipment *Shipment
	Ss       float32
}
