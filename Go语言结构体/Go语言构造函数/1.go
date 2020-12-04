package main

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor (color string) *Cat{
   return &Cat {
	Color:color 
   }
}

func main() {

}
