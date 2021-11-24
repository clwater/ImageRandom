package control

import (
	"ImageRandom/source_local"
	"log"
)

var isInit = false

//
//  InitGoogleEarthImage
//  @Description: 初始化静态数据信息（Google Earth）
//
func InitGoogleEarthImage()  {
	google_earth.Init()
}

//
//  Init
//  @Description: 初始化
//
func Init()  {
	log.Println("init")
	InitGoogleEarthImage()
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

