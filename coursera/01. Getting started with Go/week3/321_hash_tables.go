package main

/*

Hash tables
- contains key-value pairs
- each values is associated with unique a key
- hash function is used to compute the slot for a key; you never call this explcitiely

Advantages
- faster lookup than lists: constant time vs linear time !!!
- arbitrary keys: not ints like slices or arrays

Disadvantages
- may have collisions: two keys hash to same slot; may slow things down but it's handled internally by go
*/
