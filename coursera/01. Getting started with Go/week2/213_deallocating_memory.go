// when a variable is no longer needed, it should be deallocated (memory space is made available)
// Traditionally, there are two types of memory: Stack vs Heap
// Stack is dedicated to function calls. Local variables are stored here and is deallocated after the function completes
// Heap on the other hand is persistent. Data on the heap must be deallocated when it is done being used. In most compiled languages this is done manually: x = malloc(32); free(x). It is fast but error prone

package main
