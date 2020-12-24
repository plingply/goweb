/*
 * @Author: 彭林
 * @Date: 2020-12-24 14:59:06
 * @LastEditTime: 2020-12-24 19:30:51
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/service/zuowen/zuowen.go
 */
package zuowen

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"goframe-web/app/model"
	"io/ioutil"
	"reflect"
	"strconv"
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

func SyncZuowenById(start, end uint) bool {
	go forEach(start, end)
	return true
}

func forEach(start, end uint) {
	for i := start; i < end; i++ {
		BatchCreateZuowen(i)
	}
}

func BatchCreateZuowen(id uint) (bool, uint) {
	fmt.Println("i:======:", id)
	var zuowen model.Zuowen
	zuowen.ZuowenId = id
	db := model.GetDB()
	db.Where("zuowen_id = ?", id).Find(&zuowen)
	if zuowen.Id == 0 {
		fmt.Println("xxxxxxx: 1")
		var wg sync.WaitGroup
		wg.Add(1) //
		result, err := getArt(id, &wg)
		wg.Wait()
		fmt.Println("xxxxxxx: 3")
		if err == nil {
			var new_zuowen model.Zuowen
			new_zuowen.Comments = result["comment"].(string)
			new_zuowen.Genre = result["genre"].(string)
			new_zuowen.Grade = result["grade"].(string)
			new_zuowen.ZuowenId = id
			new_zuowen.Quality = result["quality"].(string)
			new_zuowen.Theme = result["theme"].(string)
			new_zuowen.Title = result["title"].(string)
			new_zuowen.WordNumber = uint(result["wordNumber"].(float64))
			new_zuowen.ZwContent = result["zwContent"].(string)
			zuowen_id := CreateZuowen(&new_zuowen)
			return true, zuowen_id
		}
		return false, 0
	}
	return false, 0
}

func CreateZuowen(zuowen *model.Zuowen) uint {

	db := model.GetDB()

	var zuowens model.Zuowen
	db.Where("zuowen_id = ?", zuowen.ZuowenId).Find(&zuowens)

	if zuowens.Id != 0 {
		return 0
	}
	db.Create(&zuowen)
	return zuowen.ZuowenId
}

func getArt(id uint, wg *sync.WaitGroup) (map[string]interface{}, error) {

	gtime.SetTimeZone("Asia/Shanghai")
	data := make(map[string]string)
	data["date"] = gtime.Datetime()
	data["zuoWenId"] = strconv.Itoa(int(id))

	data["ciphertext"] = getSign(data)

	fmt.Println("data:=====>", data)

	var postUrl = "https://api.xhzapp.com/api/v5/AppBook/GetZuoWenModel"
	params := "zuoWenId=1&date=" + data["date"] + "&ciphertext=" + data["ciphertext"]
	c := *g.Client()
	c.SetHeader("User-Agent:", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_16_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")
	req, err := c.Post(postUrl, params)
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	body, _ := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	typeOfA := reflect.TypeOf(result["code"]).Name()

	fmt.Println("xxxxxxx: code:", typeOfA, result)

	if typeOfA == "string" {
		if result["code"].(string) == "0" {
			return nil, errors.New(result["msg"].(string))
		}
	}

	if typeOfA == "float64" {
		if result["code"].(float64) == 0 {
			return nil, errors.New(result["msg"].(string))
		}
	}

	fmt.Println("xxxxxxx: 2:", result["data"])
	wg.Done()
	return result["data"].(map[string]interface{}), nil
}

func getSign(data map[string]string) string {
	t := md5V(data["date"])
	r := data["date"] + data["zuoWenId"] + "CVrFcVJu564ArjsShh058xSjb"
	s := t[0:16] + md5V(r) + t[16:32]
	return s
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
