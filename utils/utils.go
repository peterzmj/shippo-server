package utils

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ReadConfigFromFile(path string, config interface{}) error {
	dir, _ := os.Getwd()
	fmt.Printf("ReadConfigFromFile:%v\n", dir)

	file, _ := os.Open(path)
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	return json.Unmarshal(bytes, &config)
}

func PhoneMasking(s string) string {
	if len(s) < 11 {
		return s
	}
	return s[:3] + "******" + s[9:]
}

func QQMasking(s string) string {
	if len(s) < 5 {
		return s
	}
	return s[:1] + "******" + s[len(s)-2:]
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ParseTime(s string) (time.Time, error) {
	return time.Parse(s, "2006-01-02 15:04:05")
}

func GenerateCaptcha() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(899999) + 100000)
}

func GenerateToken() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

//判断文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func In(val interface{}, arr interface{}) bool {
	arrValue := reflect.ValueOf(arr)
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arrValue.Len(); i++ {
			if arrValue.Index(i).Interface() == val {
				return true
			}
		}
	case reflect.Map:
		if arrValue.MapIndex(reflect.ValueOf(val)).IsValid() {
			return true
		}
	}
	return false
}
