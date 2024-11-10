package wallmart

import "fmt"

type Reactangle struct {
	Length float32
	Width float32
}
type RectangleInterface interface{
	Area()
	Perimeter()
}
func (r Reactangle) Area() float32{
	result := r.Length * r.Width
	return result
}
func (r Reactangle) Perimeter() float32{
	res := 2*r.Length * 2*r.Width
	return res
}
func ReactangleCal(){
	res := Reactangle{Length: 10, Width: 30}
	fmt.Printf("Area of rectangle is :%v\n", res.Area())
	fmt.Printf("Peremeter of rectangle is :%v\n", res.Perimeter())
}