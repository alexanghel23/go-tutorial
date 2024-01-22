package main

/*
1. Function naming
- give functions a good name
- give variable a good name

2. Functional cohesion
- function should perform only one operation. Example: Geometry application -> PointDist(), DrawCircle(), TriangleArea()
- merging behaviors makes code complicated. Example: DrawCircle() + TriangleArea()

3. Few parameters
- more diffucult with a large number of parameters
- function maybe have bad cohesion -> large number of parameters

Reducing Parameter number
- may need to group related arguments into structures

Example of bad function:
TriangleArea() - with 3 points needed to define a triangle; in 3D there are 9 arguments (xyz to define a point in space)

Example of better function:
TriangleArea() - with 3 points defined as type Point struct {x,y,z float}; total 3 arguments

Example of good solution
TriangleArea() - one argument using a Triangle type

4. Function Complexity
- function length
Option1:
func a() {100 lines}

Option2
func a() {
	b()
	c()
}

func b() {50 lines}
func c() {50 lines}




*/
