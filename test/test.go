package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	i := 0
	for {
		log.Println(Create("User"+strconv.Itoa(i), "password", "15265858452","10"))
		i++
		time.Sleep(500)
		if(i >=1000){
			break
		}
	}
}

func Create(name, pw, phone, money string) string {
	resp, err := http.Get("http://127.0.0.1:84/testadduser?username=" + name + "&password=" + pw + "&phone=" + phone + "&money=" + money + "&vip=0&sex=0")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	input, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	/*
		result := gojson.Json(string(input)).Get("translateResult").Getindex(1).Getindex(1).Get("tgt").Tostring()
		println(string(input))
	*/
	return string(input)
}

//http://127.0.0.1:84/adduser?username=&password=&phone=&money=&vip=0&sex=0
