package database

import (
	"fmt"
	"strconv"
	"testing"
)

var connected bool
var dbConfig DbConfig
var insertID int64
var insertID2 int64

func TestConnectDb(t *testing.T) {
	//var dbConfig DbConfig
	dbConfig.Host = "localhost:3306"
	dbConfig.DbUser = "admin"
	dbConfig.DbPw = "admin"
	dbConfig.DatabaseName = "ulbora_template_service"
	connected = dbConfig.ConnectDb()
	if connected != true {
		t.Fail()
	}
}

func TestConnectionTest(t *testing.T) {
	var a []interface{}
	success := dbConfig.ConnectionTest(a...)
	if success == true {
		fmt.Print("Connection test: ")
		fmt.Println(success)
	} else {
		fmt.Println("database connection test failed")
		t.Fail()
	}
}

func TestInsertTemplate(t *testing.T) {
	var a []interface{}
	a = append(a, "testinsert1", "cms", false, 126)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfig.InsertTemplate(a...)
	if success == true && insID != -1 {
		insertID = insID
		fmt.Print("new Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	success, insID2 := dbConfig.InsertTemplate(a...)
	if success == true && insID2 != -1 {
		insertID2 = insID2
		fmt.Print("new Id: ")
		fmt.Println(insID2)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestUpdateTemplate(t *testing.T) {
	var a []interface{}
	a = append(a, true, insertID, 126)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfig.UpdateTemplate(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestGetActiveTemplate(t *testing.T) {
	a := []interface{}{"cms", 126}
	rowPtr := dbConfig.GetActiveTemplate(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if insertID != int64Val {
			fmt.Print(insertID)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		} else {
			fmt.Print("found id")
			fmt.Print(" = ")
			fmt.Println(int64Val)
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestGetTemplateByClient(t *testing.T) {
	a := []interface{}{"cms", 126}
	rowsPtr := dbConfig.GetTemplateByClient(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get by client ")
		fmt.Println(foundRows)
		//fmt.Println("GetList results: --------------------------")
		for r := range foundRows {
			foundRow := foundRows[r]
			for c := range foundRow {
				if c == 0 {
					int64Val, err2 := strconv.ParseInt(foundRow[c], 10, 0)
					if err2 != nil {
						fmt.Print(err2)
					}
					if r == 0 {
						if insertID != int64Val {
							fmt.Print(insertID)
							fmt.Print(" != ")
							fmt.Println(int64Val)
							t.Fail()
						}
					} else if r == 1 {
						if insertID2 != int64Val {
							fmt.Print(insertID)
							fmt.Print(" != ")
							fmt.Println(int64Val)
							t.Fail()
						}
					}
				}
			}
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCloseDb(t *testing.T) {
	if connected == true {
		rtn := dbConfig.CloseDb()
		if rtn != true {
			fmt.Println("database close failed")
			t.Fail()
		}
	}
}
