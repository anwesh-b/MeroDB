package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	cli "github.com/anwesh-b/MeroDB/lib/cli"
)

const filePath = "../data/"
const referenceFilePath = filePath + "tables.db"

func doesTableExist(tableName string) (bool, error) {
	data, err := ioutil.ReadFile(referenceFilePath)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("error in reading db. db file might be corrupt")
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, tableName+" ") {
			return true, nil
		}
	}

	return false, nil
}

func getReferenceTableDetails(tableName string) ([]string, error) {
	data, err := ioutil.ReadFile(referenceFilePath)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error in reading db. db file might be corrupt")
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, tableName+" ") {
			columns := strings.Split(line, " ")

			return columns[1:], nil
		}
	}

	return nil, errors.New("table does not exist")
}

func insert(str string) {
	s := strings.Split(str, " ")

	tableName := s[1]

	columns, err := getReferenceTableDetails(tableName)

	if err != nil {
		fmt.Println(err)
		return
	}

	datas := s[2:]

	if len(datas) != len(columns) {
		fmt.Println("Column count does not match. The columns are: " + strings.Join(columns, ", "))
		return
	}

	id, err := strconv.Atoi(datas[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := filePath + tableName + ".db"

	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, strconv.Itoa(id)+" ") {
			cli.CLog("Data with id " + strconv.Itoa(id) + " already exists")
			return
		}
	}

	output := strings.Join(datas, " ") + "\n"

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(output); err != nil {
		log.Println(err)
	}

	cli.CLog("Insert success")
}

func selectData(str string) {
	tableName := strings.Split(str, " ")[1]

	columns, err := getReferenceTableDetails(tableName)

	if err != nil {
		fmt.Println(err)
	}

	fileName := filePath + tableName + ".db"

	fmt.Println("debug")
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintln(w, strings.Join(columns, "\t"))

	for _, line := range lines {
		items := strings.Split(line, " ")

		fmt.Fprintln(w, strings.Join(items, "\t"))
	}

	w.Flush()
}

func update(str string) {
	s := strings.Split(str, " ")

	tableName := s[1]

	columns, err := getReferenceTableDetails(tableName)

	if err != nil {
		fmt.Println(err)
		return
	}

	datas := s[2:]

	if len(datas) != len(columns) {
		fmt.Println("Column count does not match. The columns are: " + strings.Join(columns, ", "))
		return
	}

	id, err := strconv.Atoi(datas[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := filePath + tableName + ".db"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var didUpdate bool = false

	for i, line := range lines {
		if strings.HasPrefix(line, strconv.Itoa(id)+" ") {
			lines[i] = strings.Join(datas, " ")
			output := strings.Join(lines, "\n")
			err = ioutil.WriteFile(fileName, []byte(output), 0644) // can optimize this code
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
	tableName := s[1]

	_, err := getReferenceTableDetails(tableName)

	if err != nil {
		fmt.Println(err)
		return
	}

	deleteId := s[2]

	id, err := strconv.Atoi(deleteId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := filePath + tableName + ".db"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var didDelete bool = false

	for i, line := range lines {
		if strings.HasPrefix(line, strconv.Itoa(id)+" ") {
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

	cli.CLog("Couldnot find the data, delete did not occur")
}

func createTable(str string) {
	s := strings.Split(str, " ")

	tableName := s[1]
	tableColumns := append([]string{tableName, "id"}, s[2:]...)

	isTableExisting, err := doesTableExist(tableName)

	if err != nil {
		fmt.Println(err)
		return
	}

	if isTableExisting {
		fmt.Println("Table already exists")
		return
	}

	fileName := filePath + tableName + ".db"

	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f, err = os.OpenFile(referenceFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(strings.Join(tableColumns, " ") + "\n"); err != nil {
		log.Println(err)
	}
}

func dropTable(str string) {
	s := strings.Split(str, " ")

	tableName := s[1]

	data, err := ioutil.ReadFile(referenceFilePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(data), "\n")

	fileName := filePath + tableName + ".db"

	for i, line := range lines {
		if strings.HasPrefix(line, tableName+" ") {
			lines[i] = ""

			copy(lines[i:], lines[i+1:]) // Shift one index left, until the i.
			lines = lines[:len(lines)-1] // Remove last element (write zero value).

			output := strings.Join(lines, "\n")
			err = ioutil.WriteFile(referenceFilePath, []byte(output), 0644)

			if err != nil {
				fmt.Println(err)
				return
			}

			e := os.Remove(fileName)
			if e != nil {
				log.Fatal(e)
			}

			return
		}
	}

	fmt.Println("Table not found and not deleted")
}

func EvaluateInput(str string) {
	if strings.HasPrefix(str, "insert ") {
		insert(str)
	} else if strings.HasPrefix(str, "select ") {
		selectData(str)
	} else if strings.HasPrefix(str, "update ") {
		update(str)
	} else if strings.HasPrefix(str, "delete ") {
		deleteData(str)
	} else if strings.HasPrefix(str, "create ") {
		createTable(str)
	} else if strings.HasPrefix(str, "drop ") {
		dropTable(str)
	} else {
		cli.CLog("Invalid command")
	}
	cli.CLog("")
}
