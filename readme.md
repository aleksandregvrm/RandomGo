# Arrays

An array in Go is a fixed-size sequence of elements of the same type.
Declaration

go

var arr [5]int // An array of 5 integers
arr[0] = 10    // Assign value to the first element
fmt.Println(arr[0]) // Access an element

Key Points:

    Fixed size: Once an array is declared, its size cannot change.
    Zero value: Elements are initialized to the zero value of the element type (e.g., 0 for integers, "" for strings).
    Value type: Arrays are copied by value when passed to functions.

Example:

go

arr := [3]int{1, 2, 3}  // Declare and initialize an array
fmt.Println(arr)        // Output: [1 2 3]


# Slices

A slice is a more flexible, dynamic-sized view into an array. Slices are used more frequently than arrays in Go.
Declaration

go

var slice []int         // Declares a slice of integers
slice = append(slice, 1) // Add an element to the slice
fmt.Println(slice[0])   // Access the first element

Key Points:

    Dynamic size: Slices can grow and shrink using functions like append.
    Underlying array: Slices point to an underlying array and keep track of the length and capacity.
    Reference type: When passed to functions, slices behave like references (i.e., modifications affect the original).

Example:

go

arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // A slice from index 1 to 3
fmt.Println(slice) // Output: [2 3 4]

len() and cap() functions:

    len(slice) returns the number of elements in the slice.
    cap(slice) returns the capacity of the slice (the size of the underlying array from the start of the slice).

# Maps

A map in Go is a key-value data structure that allows fast lookups and inserts.
Declaration

go

var m map[string]int         // Declare a map with string keys and integer values
m = make(map[string]int)     // Initialize the map
m["age"] = 25                // Set key-value pair
fmt.Println(m["age"])        // Access value by key

Key Points:

    Dynamic size: Maps grow as needed when elements are added.
    Zero value: A map that hasn’t been initialized (with make) is nil and cannot be used.
    Check existence: To check if a key exists, use the second return value.

Example:

go

m := map[string]int{"age": 25, "score": 100}
value, exists := m["age"] // Retrieve value and check if key exists
if exists {
    fmt.Println("Age:", value)
} else {
    fmt.Println("Key not found")
}

delete() function:

Removes a key from the map.

go

delete(m, "age")
