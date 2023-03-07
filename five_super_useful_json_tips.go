package main

import (
	"encoding/json"
	"fmt"
)

// Tip0:
// Outer Struct Override Inner Struct

type Inner struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Outer struct {
	Inner
	Prince string `json:"prince"`
}

func tip0() {
	m := new(Outer)
	m.Name = "river"
	m.Prince = "100"
	r, _ := json.Marshal(m)
	fmt.Println(string(r))
	// {"name":"river","price":0,"prince":"100"}
}

// Tip1: ignore field 2 ways:

type Meta struct {
	Username string `json:"username" :"Username"`
	Password string `json:"password" :"Password"`
}

type MetaProxy struct {
	Meta
	// ignore by downCase
	score string `json:"score" :"Score"`
	// ignore by tag
	Password string `json:"-"`
}

func tip1() {
	m := new(MetaProxy)
	m.Username = "river"
	m.score = "100"
	m.Password = "123456"
	r, _ := json.Marshal(m)
	fmt.Println(string(r))
	// {"username":"river","password":""}
}

type Resp struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	ErrorInf string `json:"error"`
}

type RespProxy struct {
	Resp
	ErrorInf string `json:"error,omitempty"`
}

func tip2() {
	r := new(RespProxy)
	r.Code = 200
	r.Msg = "ok"
	r.ErrorInf = ""

	r2, _ := json.Marshal(r)
	fmt.Println(string(r2))
	// {"code":200,"msg":"ok"}
}

// Tip3: Type convert

type Base struct {
	Code int    `json:"code,string"`
	Msg  string `json:"msg"`
}

func tip3() {
	r := new(Base)
	r.Code = 200
	r.Msg = "ok"

	r2, _ := json.Marshal(r)
	fmt.Println(string(r2))
	// {"code":"200","msg":"ok"}
}

// Tip4: Store Raw Data Without Type

type Raw struct {
	Code int             `json:"code"`
	Resp json.RawMessage `json:"resp"`
}

func tip4() {
	rawStr := `{
		"code": 200,
		"resp": {"name": "river", "age": 18}
	}`
	var raw Raw
	json.Unmarshal([]byte(rawStr), &raw)
	fmt.Println(string(raw.Resp))
	// {"name": "river", "age": 18}
}

func main() {
	tip0()
	tip1()
	tip2()
	tip3()
	tip4()
}
