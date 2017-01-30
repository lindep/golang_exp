
/*
Using this Meduim artical "Go and memory layout"
https://hackernoon.com/go-and-memory-layout-6ef30c730d51#.c72d7kxns
*/

package main

import (
	"fmt"
  "reflect"
)

type MyData struct {
  aByte   byte
  //anotherByte byte
  aanInt32 int32
  aString string
  aShort  int16
  anInt32 int32
  aSlice  []byte
}

func main() {
  // First ask Go to give us some information about the MyData type
  typ := reflect.TypeOf(MyData{})
  fmt.Printf("Struct is %d bytes long\n\n", typ.Size())

  // We can run through the fields in the structure in order
  n := typ.NumField()
  for i := 0; i < n; i++ {
    field := typ.Field(i)
    fmt.Printf("%s at offset %v, size=%d, align=%d\n",
        field.Name, field.Offset, field.Type.Size(),
        field.Type.Align())
   }
}
