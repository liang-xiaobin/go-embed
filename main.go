package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
var f embed.FS

//go:embed hello.txt hello1.txt
var f1 embed.FS

//go:embed hello_world/*
var f2 embed.FS

func main() {
	print(s)
	print(string(b))
	data, _ := f.ReadFile("hello.txt")
	print(string(data))
	data1, _ := f1.ReadFile("hello.txt")
	fmt.Println(string(data1))
	data2, _ := f1.ReadFile("hello1.txt")
	fmt.Println(string(data2))
	data3, _ := f2.ReadFile("hello_world/hello2.txt")
	fmt.Println(string(data3))
	data4, _ := f2.ReadFile("hello_world/hello3.txt")
	fmt.Println(string(data4))
}
