# gossdeep

	//比较相似度
	Fuzzy_compare("3:YD6xL4fYvn:Y2xMwvn","3:YD6xL4fYvn:Y2xMwvk")


	//提取文件模糊hash
	Fuzzy_hash_file("/var/www/shell.php")


	//提取字符串的模糊hash
	Fuzzy_hash_buf("MXi4oyu")

## gcc fuzzy.c  -fPIC -shared -o libfuzzy.so

### error while loading shared libraries: libfuzzy.so.2

#### 方法一：

* 将/usr/local/lib/ 加入到动态链接库的路径
	echo "/usr/local/lib/" >> vi /etc/ld.so.conf
* 更新一下新的库文件
	sudo ldconfig
	
#### 方法二：

* export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH
* sudo ldconfig

#### 方法三:
* apt-get install libfuzzy-dev

#### https://ssdeep-project.github.io/ssdeep/doc/api/html/fuzzy_8h.html
