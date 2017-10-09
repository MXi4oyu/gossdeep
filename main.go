package main

import (
	"github.com/MXi4oyu/gossdeep/deepapi"
	"fmt"
)

func main()  {


	//比较相似度
	similary:=deepapi.Fuzzy_compare("3:YD6xL4fYvn:Y2xMwvn","3:YD6xL4fYvn:Y2xMwvk")

	fmt.Println(similary)

	//提取文件模糊hash
	deepapi.Fuzzy_hash_file("/var/www/shell.php")


	//提取字符串的模糊hash
	deepapi.Fuzzy_hash_buf("MXi4oyu")



}


