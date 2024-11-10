package interfacedemo

type Shapes interface {
	Area()
}
type Square struct {
	side float32
}
type Rectangle struct {
	length float32
	width  float32
}

func (s Square) Area() float32 {
	result := s.side * s.side
	return result
}
func (r Rectangle) Area() float32 {
	result := r.length * r.width
	return result
}
func InterfaceDemo() {

}
