/*
 Copyright (C) 2016 Ulbora Labs Inc. (www.ulboralabs.com)
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
	db "UlboraTemplateService/database"
	"fmt"
	"strconv"
	"strings"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

//Response res
type Response struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
}

//Template Template
type Template struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Application string `json:"application"`
	Active      bool   `json:"active"`
	ClientID    int64  `json:"clientId"`
}

//TemplateDB db config
type TemplateDB struct {
	DbConfig db.DbConfig
}

//ConnectDb to database
func (db *TemplateDB) ConnectDb() bool {
	rtn := db.DbConfig.ConnectDb()
	if rtn == true {
		fmt.Println("db connect")
	}
	return rtn
}

//InsertTemplate in database
func (db *TemplateDB) InsertTemplate(template *Template) *Response {
	var rtn Response
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	template.Name = stripSpace(template.Name)
	var a []interface{}
	template.Active = false
	a = append(a, template.Name, template.Application, template.Active, template.ClientID)
	success, insID := db.DbConfig.InsertTemplate(a...)
	if success == true {
		fmt.Println("inserted record")
	}
	rtn.ID = insID
	rtn.Success = success
	return &rtn
}

//UpdateActiveTemplate in database
func (db *TemplateDB) UpdateActiveTemplate(template *Template) *Response {
	var rtn Response
	//var existingTemp *Template
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	template.Active = true
	a = append(a, template.Active, template.ID, template.ClientID)
	success := db.DbConfig.UpdateTemplate(a...)
	if success == true {
		fmt.Println("update record")
	}
	var a2 []interface{}
	template.Active = false
	a2 = append(a2, template.Active, template.ID, template.ClientID)
	success2 := db.DbConfig.UpdateClearTemplate(a2...)
	if success2 == true {
		fmt.Println("update cleared record")
	}

	rtn.ID = template.ID
	rtn.Success = success
	return &rtn
}

//GetActiveTemplate template from database
func (db *TemplateDB) GetActiveTemplate(template *Template) *Template {
	var a []interface{}
	a = append(a, template.Application, template.ClientID)
	var rtn *Template
	rowPtr := db.DbConfig.GetActiveTemplate(a...)
	if rowPtr != nil {
		//print("template row: ")
		//println(rowPtr.Row)
		foundRow := rowPtr.Row
		rtn = parseTemplateRow(&foundRow)
	}
	return rtn
}

//GetTemplateByClient content by Client
func (db *TemplateDB) GetTemplateByClient(template *Template) *[]Template {
	var rtn []Template
	var a []interface{}
	a = append(a, template.Application, template.ClientID)
	rowsPtr := db.DbConfig.GetTemplateByClient(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowTmpt := parseTemplateRow(&foundRow)
			rtn = append(rtn, *rowTmpt)
		}
	}
	return &rtn
}

//DeleteTemplate in database
func (db *TemplateDB) DeleteTemplate(template *Template) *Response {
	var rtn Response
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, template.ID, template.ClientID)
	success := db.DbConfig.DeleteTemplate(a...)
	if success == true {
		fmt.Println("inserted record")
	}
	rtn.Success = success
	rtn.ID = template.ID
	return &rtn
}

//CloseDb connection to database
func (db *TemplateDB) CloseDb() bool {
	rtn := db.DbConfig.CloseDb()
	if rtn == true {
		fmt.Println("db connect closed")
	}
	return rtn
}

func parseTemplateRow(foundRow *[]string) *Template {
	var rtn Template
	if len(*foundRow) > 0 {
		id, errID := strconv.ParseInt((*foundRow)[0], 10, 0)
		if errID != nil {
			fmt.Print(errID)
		} else {
			rtn.ID = id
		}
		rtn.Name = (*foundRow)[1]
		rtn.Application = (*foundRow)[2]

		if (*foundRow)[3] == "1" {
			rtn.Active = true
		} else {
			rtn.Active = false
		}
		clientID, errClient := strconv.ParseInt((*foundRow)[4], 10, 0)
		if errClient != nil {
			fmt.Print(errClient)
		} else {
			rtn.ClientID = clientID
		}
	}
	return &rtn
}

func stripSpace(val string) string {
	var rtn = ""
	rtn = strings.Replace(val, " ", "", -1)
	return rtn
}
