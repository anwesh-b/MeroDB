package parser

import (
	"fmt"
	"strconv"
	"strings"

	cli "github.com/anwesh-b/MeroDB/lib/cli"
)

type Data struct {
	id    int
	name  string
	email string
}

var myData []Data

func insert(str string) {
	s := strings.Split(str, " ")

	datas := s[1:]

	id, err := strconv.Atoi(datas[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	myData = append(myData, Data{
		id:    id,
		name:  datas[1],
		email: datas[2],
	})

	cli.CLog("Insert success")
}

func selectData(str string) {
	for _, data := range myData {
		cli.CLog(strconv.Itoa(data.id) + "\t" + data.name + "\t" + data.email)
	}
}

func update(str string) {
	s := strings.Split(str, " ")

	datas := s[1:]

	id, err := strconv.Atoi(datas[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, data := range myData {
		if data.id == id {
			myData[i].name = datas[1]
			myData[i].email = datas[2]
			cli.CLog("Updating success")
			return
		}
	}

	cli.CLog("Couldnot find the data, update failed")
}

func deleteData(str string) {
	s := strings.Split(str, " ")

	datas := s[1:]

	id, err := strconv.Atoi(datas[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, data := range myData {
		if data.id == id {
			myData = append(myData[:i], myData[i+1:]...)
			cli.CLog("Deleting success")
			return
		}
	}

	cli.CLog("Couldnot find the data, delete failed")
}

func EvaluateInput(str string) {
	if strings.HasPrefix(str, "insert") {
		insert(str)
	} else if strings.HasPrefix(str, "select") {
		selectData(str)
	} else if strings.HasPrefix(str, "update") {
		update(str)
	} else if strings.HasPrefix(str, "delete") {
		deleteData(str)
	} else {
		cli.CLog("Invalid command")
	}
	cli.CLog("")
}
