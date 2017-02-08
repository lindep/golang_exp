
/*
Using this Meduim article "Go and memory layout"
https://hackernoon.com/go-and-memory-layout-6ef30c730d51#.c72d7kxns
*/

package main

import (
	"fmt"
  "reflect"
	"unsafe"
)

type MyData struct {
  aByte   byte
  //anotherByte byte
  //aanInt32 int32
  //aString string
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

	 data := MyData{
		aByte: 0x1,
		//aString: "a",
		aShort: 0x0203,
		anInt32: 0x04050607,
		aSlice: []byte{
			0x08, 0x09, 0x0a,
		},
	}

	 dataBytes := (*[32]byte)(unsafe.Pointer(&data))
	 fmt.Printf("Bytes are \n%#v\n", dataBytes)
}
