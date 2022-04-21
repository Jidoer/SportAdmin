package tool

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

/********
*Userinfo数据库格式
*UID      -> Auto int
*username -> Base64URL-E-String
*password -> DES
*sex      -> int
*Created  -> time.time
*Vip      -> int
*phone    -> 10086 ?
*email    -> Base64URL-E-String
*loginip  -> 0.0.0.0 ? string
 */

/**********
*Server SQL
*UID      -> Auto int
*name     -> string
*ip       -> string
*type     -> string
*info     -> string
 */

func search() {

}

func interface2String(inter interface{}) string {

	switch inter.(type) {

	case string:
		// rt.Println("string", inter.(string))
		return inter.(string)
	case int:
		fmt.Println("int", inter.(int))
		return ""
	case float64:
		fmt.Println("float64", inter.(float64))
		return ""
	}
	return ""

}

func Isnumber(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

func String2Int(str5 string) int {
	int5, err := strconv.Atoi(str5)
	if err != nil {
		fmt.Println(err)
		return -1 //error
	} else {

		return int5
	}
}

func MapToJson(param map[string]map[string]string/*interface{}*/) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func InterfaceToJson(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}


func JsonToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}



/*
func IfLogin( username, key string) string { //id + key batter than username +key!

	kusername := base64.URLEncoding.EncodeToString([]byte(username))
	fileuser, err := mydes.Encrypt(kusername, []byte("2fa6c1e9"))
	if err != nil {
		//log.Fatal(err)//解密失败退出bug!
		log.Println(err)
	}
if file.Exists("Login/cookies/" + fileuser){
	if file.Reader("Login/cookies/" + fileuser) == key {
		//00000x1
		return "yes"
	}else{
		log.Println("KEY不相等!")
		return "no"
	}
}else{
	log.Println("不存在cookies!")
	return "no"
}

}
*/

func URLCode(yoururl string) string{
	return url.QueryEscape(yoururl)
}
func UnURLCode(yoururl string) string{
	decodeurl,err := url.QueryUnescape(yoururl)
	if err != nil {
		fmt.Println(err)
	}
	return decodeurl
}


func GetFileMd5(filename string) (string, error) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("os Open error")
        return "", err
    }
    md5 := md5.New()
    _, err = io.Copy(md5, file)
    if err != nil {
        fmt.Println("io copy error")
        return "", err
    }
    md5Str := hex.EncodeToString(md5.Sum(nil))
    return md5Str, nil
}
  
func GetStringMd5(s string) string {
    md5 := md5.New()
    md5.Write([]byte(s))
    md5Str := hex.EncodeToString(md5.Sum(nil))
    return md5Str
}

func IsEmail(email string) bool{
	result, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	if result {
		return true
	} else {
		return false
	}
}