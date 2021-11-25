package control

import (
	"ImageRandom/source"
	"log"
)

var isInit = false


//
//  Init
//  @Description: 初始化
//
func Init()  {
	log.Println("init")
	//初始化静态数据信息（Google Earth）
}

//
//  Random
//  @Description: 随机获取一张图片
//  @return string
//
func Random() string {
	if !isInit {
		Init()
		isInit = true
	}
	return google_earth.Random()
}

