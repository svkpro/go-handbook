# Maps
One of the most useful data structures in computer science is the hash table. 
Many hash table implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes. 
Go provides a built-in map type that implements a hash table.

A Go map type looks like this:

```
map[KeyType]ValueType
```
where KeyType may be any type that is comparable (more on this later), and ValueType may be any type at all, including another map!

This variable m is a map of string keys to int values:

```
var m map[string]int
```
Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn't point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. To initialize a map, use the built in make function:

```
m = make(map[string]int)
```
The make function allocates and initializes a hash map data structure and returns a map value that points to it. The specifics of that data structure are an implementation detail of the runtime and are not specified by the language itself. In this article we will focus on the use of maps, not their implementation.