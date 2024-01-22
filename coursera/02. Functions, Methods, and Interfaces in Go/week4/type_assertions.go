package main

/*
Concealing type differences
- interfaces hide the differences between types
- sometimes you need to treat different types in different ways


Type Assertions for Disambiguation
- Type assertions can be used to determine and extract the underlying concrete type

func DrawShape(s Shape2D) bool {
	rect, ok := s.(Rectangle)
	if ok {
		DrawRect(rect)
	}
	tri, ok := s.(Triangle)
	if ok {
		DrawRect(tri)
	}
}
-type assertions extracts Rectangle from Shape2D (Concrete type in paranthesis)

Type Switch

*/
