package deepapi

/*
#cgo CFLAGS: -I ./include
#cgo LDFLAGS: -L ./lib -lfuzzy
#include "fuzzy.h"
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"unsafe"
	"fmt"
	"log"
	"runtime/debug"
)


//相似度
func Fuzzy_compare(hash1, hash2 string) (result int) {

	ch1 := C.CString(hash1)
	ch2 := C.CString(hash2)
	defer C.free(unsafe.Pointer(ch1))
	defer C.free(unsafe.Pointer(ch2))

	var hashSimilarity C.int
	hashSimilarity = 0
	hashSimilarity = C.fuzzy_compare(ch1, ch2)

	if hashSimilarity > 0 {
		return int(hashSimilarity)
	} else {
		return int(0)
	}

}

//提取文件模糊hash
func Fuzzy_hash_file(filePath string) (filehash string) {
	result := C.CString("0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

	defer C.free(unsafe.Pointer(result))

	filename := C.CString(filePath)

	cmode := C.CString("rb")

	defer C.free(unsafe.Pointer(cmode))

	fp := C.fopen(filename, cmode)
	defer C.fclose(fp)

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			debug.PrintStack()
		}
	}()

	ret := C.fuzzy_hash_file(fp, result)

	if ret == 0 {

		return C.GoString(result)
	} else {
		return "0000000000000000"
	}

}

//提取字符串的模糊hash
func Fuzzy_hash_buf(str string) (strhash string) {

	result := C.CString("0")
	defer C.free(unsafe.Pointer(result))
	bufdata := str
	chdata := (*C.uchar)(unsafe.Pointer(C.CString(bufdata)))
	clen := C.uint32_t(len(bufdata))
	ret := C.fuzzy_hash_buf(chdata, clen, result)
	if ret == 0 {
		fmt.Println(C.GoString(result))
		return C.GoString(result)
	} else {
		return ""
	}

}
