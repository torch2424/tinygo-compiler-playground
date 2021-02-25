package main

import (
  "fmt"
)

//export fd_write
// func fd_write() (errno uint) {
// 	return 0
// }

//go:wasm-module wasi_snapshot_preview1
//export fd_write
func fd_write(int32, *int32, int32, *int32) int32

type BodyHandle *uintptr
type BodyWriteEnd *uintptr

//go:wasm-module fastly_http_body
//export write
func Fastly_http_body_write(
	body_handle BodyHandle,
	buf *uintptr,
	buf_len *uintptr,
	end BodyWriteEnd,
	nwritten_out *uintptr, // *mut usize
) int32

func main() {
	_ = Fastly_http_body_write(nil, nil, nil, nil, nil)
  fmt.Println("sadasd")
}
