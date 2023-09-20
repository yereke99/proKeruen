package help

type Type interface{}

var taxiGroup = []string{
	"econom",
	"comfort", 
	"minivan",
}

var truckGroup = []string{
	"microbus",
	"small",
	"middle",
	"big",
	"bort",
	"refrizh",
	"truck",
}

func Choose(type_ string) Type {
	for _, v := range taxiGroup{
		if v == type_{
			return "taxiGroup"
		}
	} 

	for _, k := range truckGroup{
		if k == type_{
			return "truckGroup"
		}
	}

	return nil
}