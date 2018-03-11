//Understanding Go Maps
//Maps are hash tables with [key, value] pair
package main

import "fmt"

func main() {
//A map can be declared as map[keyType]valueType
//A reference to map can be created as follows
var m map[string]int
fmt.Println("Size of map 'm':", m, "and length is:", len(m))
//The above statement only creates a reference and no memory is allocated, it behaves like a empty map
//attempting to write in it will cause runtime panic

//Instead initialize a map using make
m = make(map[string]int)
fmt.Println("Size of map 'm':", m, "and length is:", len(m))
//Adding a value in map
m["Alice"] = 66
m["Bruce"] = 54
m["Kiera"] = 34

fmt.Println("Size of map:", len(m))
fmt.Println("Values in map 'm':", m)

//Retrieve the value stored in the key "Alice"
i := m["Alice"]
fmt.Println("Retreiving a value, i =", i)

//Trying to retrieve a non existent value returns value types zero value, e.g if int it returns 0
j := m["John"]
fmt.Println("The value of John is:", j)

//Deleting a item from map
delete(m, "Alice")
fmt.Println("Map after deleting is:", m)
//Delete will return nothing if the key doesn't exist
delete(m, "Wonder")
fmt.Println("Map after deleting is:", m)

//Two value assignment
if value, ok := m["Bruce"]; ok {
	fmt.Println("The age of Bruce is:", value)
}
//value -> gives the value of the key , ok -> gives a boolean
// _, ok := m["John"] would check if the key exists in the map
if _, ok := m["Emma"]; !ok {
	fmt.Println("The given key does not exist")
}

m["Stark"] = 72
m["Robbie"] = 23
//Iterating the map using range
for key, val := range m {
	fmt.Println("'Key':", key, "'Value':", val)
}

}
