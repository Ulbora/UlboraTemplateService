package mysqldb

import (
	"fmt"
	"strconv"
	"testing"
)

var connected bool
var insertID int64
var insertID2 int64

func TestTemplateDb(t *testing.T) {
	connected = ConnectDb("localhost:3306", "admin", "admin", "ulbora_template_service")
	if connected != true {
		fmt.Println("database init failed")
		t.Fail()
	}
}

func TestConnectionTest(t *testing.T) {
	var a []interface{}
	rowPtr := ConnectionTest(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		//fmt.Print("Records found during test ")
		//fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		fmt.Print("Records found during test ")
		fmt.Println(int64Val)
		if err2 != nil {
			fmt.Print(err2)
		}
		if int64Val >= 0 {
			fmt.Println("database connection ok")
		} else {
			fmt.Println("database connection failed")
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestInsertTemplate(t *testing.T) {
	var a []interface{}
	a = append(a, "testinsert1", "cms", false, 125)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := InsertTemplate(a...)
	if success == true && insID != -1 {
		insertID = insID
		fmt.Print("new Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	var a2 []interface{}
	a2 = append(a2, "testinsert2", "cms", false, 125)
	success, insID2 := InsertTemplate(a2...)
	if success == true && insID2 != -1 {
		insertID2 = insID2
		fmt.Print("new Id: ")
		fmt.Println(insID2)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

// func TestUpdateTemplate(t *testing.T) {
// 	var a []interface{}
// 	a = append(a, true, insertID, 125)
// 	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
// 	success := UpdateTemplate(a...)
// 	if success != true {
// 		fmt.Println("database update failed")
// 		t.Fail()
// 	}
// }

func TestUpdateTemplate2(t *testing.T) {
	var a []interface{}
	a = append(a, true, insertID2, 125)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := UpdateTemplate(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestGetActiveTemplate(t *testing.T) {
	a := []interface{}{"cms", 125}
	rowPtr := GetActiveTemplate(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if insertID2 != int64Val {
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
	a := []interface{}{"cms", 125}
	rowsPtr := GetTemplateByClient(a...)
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
							fmt.Print(insertID2)
							fmt.Print(" != ")
							fmt.Println(int64Val)
							t.Fail()
						}
					}
				}
				//fmt.Println(string(foundRow[c]))
			}
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCloseDb(t *testing.T) {
	if connected == true {
		rtn := CloseDb()
		if rtn != true {
			fmt.Println("database close failed")
			t.Fail()
		}
	}
}
