package main

/*
Make
- Create a slice (and array) using make()
- 2-argument version:type and length(=capacity)
  sli = make([]int, 10)
- 3-argument version: type, length, capacity
  sli = make([]int, 10, 15)

Append
- increase the size of a slice
- adds elements to the end of a slice and inserts into the underlying array
- you can append even beyond the capacity of the array
  sli = make([]int, 0, 3)
  sli = append(sli, 100)

*/
