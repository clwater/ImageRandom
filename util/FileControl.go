package util

import (
	"io/ioutil"
	"os"
)

//
//  ReadFile
//  @Description: 读取指定文件内容
//  @param path 文件路径
//  @return string 文件内容
//
func ReadFile(path string) string  {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	//fmt.Println(string(content))
	return string(content)
}

func GetFilePath(path string)  string {
	str, _ := os.Getwd()
	str = str + path
	return str
}

