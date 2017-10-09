package main

import "github.com/MXi4oyu/gossdeep/api"

func main()  {


	//比较相似度
	api.Fuzzy_compare("3:YD6xL4fYvn:Y2xMwvn","3:YD6xL4fYvn:Y2xMwvk")


	//提取文件模糊hash
	api.Fuzzy_hash_file("/var/www/shell.php")


	//提取字符串的模糊hash
	api.Fuzzy_hash_buf("MXi4oyu")



}


