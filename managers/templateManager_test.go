/*
 Copyright (C) 2016 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2016 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package managers

import (
	"fmt"
	"testing"
)

var templateDB TemplateDB
var connected bool
var insertID int64
var insertID2 int64

func TestConnectDb(t *testing.T) {
	templateDB.DbConfig.Host = "localhost:3306"
	templateDB.DbConfig.DbUser = "admin"
	templateDB.DbConfig.DbPw = "admin"
	templateDB.DbConfig.DatabaseName = "ulbora_template_service"
	connected = templateDB.ConnectDb()
	if connected != true {
		t.Fail()
	}
}

func TestInsertTemplate(t *testing.T) {
	var tm Template
	tm.Name = "test insert"
	tm.Application = "cms"
	tm.Active = false
	tm.ClientID = 127

	res := templateDB.InsertTemplate(&tm)
	if res.Success == true && res.ID != -1 {
		insertID = res.ID
		fmt.Print("new Id: ")
		fmt.Println(res.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
	var tm2 Template
	tm2.Name = "test insert2"
	tm2.Application = "cms"
	tm2.Active = true
	tm2.ClientID = 127

	res2 := templateDB.InsertTemplate(&tm2)
	if res2.Success == true && res2.ID != -1 {
		insertID2 = res2.ID
		fmt.Print("new Id: ")
		fmt.Println(res2.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestUpdateTemplate(t *testing.T) {
	var tm Template
	tm.ID = insertID
	tm.Active = true
	tm.Application = "cms"
	tm.ClientID = 127
	res := templateDB.UpdateActiveTemplate(&tm)
	if res.Success != true {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestActiveTemplate(t *testing.T) {
	var tm Template
	tm.Application = "cms"
	tm.ClientID = 127
	res := templateDB.GetActiveTemplate(&tm)
	fmt.Println("")
	fmt.Print("found template: ")
	fmt.Println(res)
	if res.Active != true {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestGetTemplateByClient(t *testing.T) {
	var tm Template
	tm.Application = "cms"
	tm.ClientID = 127
	res := templateDB.GetTemplateByClient(&tm)
	fmt.Println("")
	fmt.Print("found list of templates: ")
	fmt.Println(res)
	if len(*res) == 0 {
		fmt.Println("database read failed")
		t.Fail()
	} else {
		row1 := (*res)[0]
		if row1.Active != true {
			t.Fail()
		}
	}
}

func TestDeleteTemplate(t *testing.T) {
	var tm Template
	tm.ID = insertID
	tm.ClientID = 127

	res := templateDB.DeleteTemplate(&tm)
	if res.Success == true {
		fmt.Print("deleted Id: ")
		fmt.Println(insertID)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}

	var tm2 Template
	tm2.ID = insertID2
	tm2.ClientID = 127

	res2 := templateDB.DeleteTemplate(&tm2)
	if res2.Success == true {
		fmt.Print("deleted Id: ")
		fmt.Println(insertID2)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestCloseDb(t *testing.T) {
	success := templateDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
