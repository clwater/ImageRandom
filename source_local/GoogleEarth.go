package google_earth

import (
	"ImageRandom/model"
	"math/rand"
	"time"
	//"ImageRandom/model"
	"ImageRandom/util"
	"encoding/json"
	"log"
)

var googleEarthInfo []model.GoogleEarthModel
var googleEarthInfoSize int

//
//  Init
//  @Description: 进行初始化
//
func Init()  {
	log.Println("Init Google Earth")
	allInfo := util.ReadFile(util.GetFilePath("\\file\\google_earth_images.json"))
	err := json.Unmarshal([]byte(allInfo), &googleEarthInfo)
	if err !=nil {
	}else {
		googleEarthInfoSize = len(googleEarthInfo)
	}
}

//
//  Random
//  @Description: 随机获取当前范围内某张图片
//  @return string
//
func Random() string {
	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(googleEarthInfoSize)
	return googleEarthInfo[randIndex].Image
}

