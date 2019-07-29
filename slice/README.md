# Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. 
In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```
a[low : high]
```
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

```
a[1:4]
```

# Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

# Slice literals
A slice literal is like an array literal without the length.

This is an array literal:

```
[3]bool{true, true, false}
```
And this creates the same array as above, then builds a slice that references it:

```
[]bool{true, true, false}
```


Slice defaults
When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

```
var a [10]int
```
these slice expressions are equivalent:

```
a[0:10]
a[:10]
a[0:]
a[:]
```

# Slice length and capacity
A slice has both a length and a capacity.
  
The length of a slice is the number of elements it contains.
  
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
  
The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
  
You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

# Nil slices
The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.

# Slices of slices
Slices can contain any type, including other slices.

# Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in append function. 
The documentation of the built-in package describes append.

```
func append(s []T, vs ...T) []T
```
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. 
The returned slice will point to the newly allocated array.

# Range
The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. 
The first is the index, and the second is a copy of the element at that index.