package deepapi
/*
#cgo CFLAGS: -I ./include
#cgo LDFLAGS: -L ./lib -lfuzzy
#include "fuzzy.h"
#include <stdio.h>
#include <stdint.h>
 */
import "C"

import (
	"fmt"
	"unsafe"
)

//相似度
func Fuzzy_compare(hash1,hash2 string) (result int)  {

	ch1:=C.CString(hash1)
	ch2:=C.CString(hash2)

	hashSimilarity:=0
	hashSimilarity=C.fuzzy_compare(ch1,ch2)

	if(hashSimilarity>0){
		return hashSimilarity
	}else{
		return 0;
	}



}

//提取文件模糊hash
func Fuzzy_hash_file(filepath string) (filehash string) {
	result:=C.CString("000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

	filename:=C.CString(filepath)

	cmode:=C.CString("rb")

	fp:=C.fopen(filename,cmode)
	defer C.fclose(fp)

	ret:=C.fuzzy_hash_file(fp,result)

	if (ret==0) {

		return C.GoString(result)
	}else {
		return ""
	}


}

//提取字符串的模糊hash
func Fuzzy_hash_buf(str string) (strhash string) {

	result:=C.CString("0")
	bufdata:=str
	chdata:=(*C.uchar)(unsafe.Pointer(C.CString(bufdata)))
	clen:=C.uint32_t(len(bufdata))
	ret:=C.fuzzy_hash_buf(chdata,clen,result)
	if (ret==0){
		fmt.Println(C.GoString(result))
		return C.GoString(result)
	}else{
		return ""
	}

}

