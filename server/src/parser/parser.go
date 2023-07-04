package parser

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	cli "github.com/anwesh-b/MeroDB/lib/cli"
)

const fileName = "local.db"

func insert(str string) {
	s := strings.Split(str, " ")

	datas := s[1:]

	id, err := strconv.Atoi(datas[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, strconv.Itoa(id)) {
			cli.CLog("Data with id " + strconv.Itoa(id) + " already exists")
			return
		}
	}

	output := string(file) + strconv.Itoa(id) + "\t" + datas[1] + "\t" + datas[2] + "\n"
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	cli.CLog("Insert success")
}

func selectData(str string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		items := strings.Split(line, "\t")

		for _, item := range items {
			fmt.Print(item + "\t")
		}
		fmt.Println("")
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

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var didUpdate bool = false

	for i, line := range lines {
		if strings.HasPrefix(line, strconv.Itoa(id)) {
			lines[i] = strconv.Itoa(id) + "\t" + datas[1] + "\t" + datas[2]
			output := strings.Join(lines, "\n")
			err = ioutil.WriteFile(fileName, []byte(output), 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			didUpdate = true
		}
	}
	if didUpdate {
		cli.CLog("Updating success")
		return
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

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var didDelete bool = false

	for i, line := range lines {
		if strings.HasPrefix(line, strconv.Itoa(id)) {
			lines[i] = ""

			copy(lines[i:], lines[i+1:]) // Shift one index left, until the i.
			lines = lines[:len(lines)-1] // Remove last element (write zero value).

			output := strings.Join(lines, "\n")
			err = ioutil.WriteFile(fileName, []byte(output), 0644)

			if err != nil {
				fmt.Println(err)
				return
			}
			didDelete = true
		}
	}
	if didDelete {
		cli.CLog("Deleting success")
		return
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
