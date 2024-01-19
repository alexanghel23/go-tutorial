/*
Type conversions
- most binary operations need operands of the same type

This will fail

var x int32 = 1
var y int16 = 2

x = y

In order to work we have to convert the type with T() operation

x = int32(y)

*/

