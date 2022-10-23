# slice

## new slice one
```go
slice := []int{10, 20, 30, 40, 50, 60}
newSlice := slice[1:3] // [20, 30]
newSlice = append(newSlice, 60) // [20, 30, 60]
// slice: [10, 20, 30, 60, 50, 60]
```

## new slice two
```go
slice := []int{10, 20, 30, 40, 50, 60}
newSlice := slice[1:3:3] // [20, 30]
newSlice = append(newSlice, 60) // [20, 30, 60]
// slice: [10, 20, 30, 40, 50, 60]
```

## slice with append
```go
slice := []int{10, 20, 30, 40, 50, 60}
newSlice := slice[1:4]
// slice: [10 20 30 40 50 60], newSlice: [20 30 40]
newSlice = append(newSlice[:1], newSlice[2:]...)
// slice: [10 20 40 40 50 60], newSlice: [20 40]
slice = append(slice[:2], slice[4:]...)
// slice: [10 20 50 60], newSlice: [20 50]
```

## slice as func param
```go
slice := []int{10, 20, 30, 40, 50, 60}
func(slice []int) {
    slice[1] = 40
}(slice)
// slice: [10 40 30 40 50 60]
```

## slice as func param with append
```go
slice := []int{10, 20, 30, 40, 50, 60}
func(slice []int) {
    slice = append(slice, 40)
}(slice)
// slice: [10 20 30 40 50 60]
```

# map
## map as func param
```go
colors := map[string]string{
    "AliceBlue":   "#f0f8ff",
    "Coral":       "#ff7F50",
    "DarkGray":    "#a9a9a9",
    "ForestGreen": "#228b22",
}
func(colors map[string]string, key string) {
    delete(colors, key)
}(colors, "Coral")
// colors: { "AliceBlue": "#f0f8ff", "DarkGray": "#a9a9a9", "ForestGreen": "#228b22" }
```

# struct
## struct as func param
```go
type treeNode struct {
	value       int
	left, right *treeNode
}
root:=&treeNode{666,nil,nil}

// value passing
func (node treeNode) setValue1(val int) {
	node.value = val
}
root.setValue1(8)
// root: 666

// point passing
func (node *treeNode) setValue2(val int) {
	node.value = val
}
root.setValue1(8)
// root: 8
```