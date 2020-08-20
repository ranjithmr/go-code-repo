package product

type Item struct {
	Name        string
	Productcode string
	Price       float64
}

//var Items []Item

func Menu() []Item {
	Items := []Item{
		{"Chai", "CH1", 3.11},
		{"Apples", "AP1", 6},
		{"Coffee", "CF1", 11.23},
		{"Milk", "MK1", 4.75},
		{"Oatmeal", "OM1", 3.69},
	}
	return Items
}
