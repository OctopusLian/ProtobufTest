// main project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"pftest/pf"

	"pftest/github.com/golang/protobuf/proto"
)

func write() {
	p1 := &pf.Person{
		Id:   1,
		Name: "小张",
		Phones: []*pf.Phone{
			{pf.PhoneType_HOME, "11111111"},
			{pf.PhoneType_WORK, "22222222"},
		},
	}

	p2 := &pf.Person{
		Id:   2,
		Name: "小王",
		Phones: []*pf.Phone{
			{pf.PhoneType_HOME, "33333333"},
			{pf.PhoneType_WORK, "44444444"},
		},
	}

	p3 := &pf.Person{
		Id:   3,
		Name: "小李",
		Phones: []*pf.Phone{
			{pf.PhoneType_HOME,"55555555"},
			{pf.PhoneType_WORK,"66666666"},
		},
	}

	book := &pf.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)
	book.Persons = append(book.Persons, p3)

	data, _ := proto.Marshal(book)

	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {
	date, _ := ioutil.ReadFile("./test.txt")
	book := &pf.ContactBook{}

	proto.Unmarshal(date, book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}
}

func main() {
	write()
	read()
}
