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
package database

import (
	tmpDb "UlboraTemplateService/database/mysqldb"
	"fmt"
	"strconv"
)

//DbConfig db config
type DbConfig struct {
	Host         string
	DbUser       string
	DbPw         string
	DatabaseName string
}

//TemplateRow database row
type TemplateRow struct {
	Columns []string
	Row     []string
}

//TemplateRows array of database rows
type TemplateRows struct {
	Columns []string
	Rows    [][]string
}

//ConnectDb to database
func (db *DbConfig) ConnectDb() bool {
	rtn := tmpDb.ConnectDb(db.Host, db.DbUser, db.DbPw, db.DatabaseName)
	if rtn == true {
		fmt.Println("db connect")
	}
	return rtn
}

//ConnectionTest of database
func (db *DbConfig) ConnectionTest(args ...interface{}) bool {
	var rtn = false
	rowPtr := tmpDb.ConnectionTest(args...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		fmt.Print("Records found during test ")
		fmt.Println(int64Val)
		if err2 != nil {
			fmt.Print(err2)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}

//InsertTemplate in database
func (db *DbConfig) InsertTemplate(args ...interface{}) (bool, int64) {
	success, insID := tmpDb.InsertTemplate(args...)
	if success == true {
		fmt.Println("inserted record")
	}
	return success, insID
}

//UpdateTemplate in database
func (db *DbConfig) UpdateTemplate(args ...interface{}) bool {
	success := tmpDb.UpdateTemplate(args...)
	if success == true {
		fmt.Println("updated record")
	}
	return success
}

//GetActiveTemplate get a row. Passing in tx allows for transactions
func (db *DbConfig) GetActiveTemplate(args ...interface{}) *TemplateRow {
	var templateRow TemplateRow
	rowPtr := tmpDb.GetActiveTemplate(args...)
	if rowPtr != nil {
		templateRow.Columns = rowPtr.Columns
		templateRow.Row = rowPtr.Row
	}
	return &templateRow
}

//GetTemplateByClient get a row. Passing in tx allows for transactions
func (db *DbConfig) GetTemplateByClient(args ...interface{}) *TemplateRows {
	var templateRows TemplateRows
	rowsPtr := tmpDb.GetTemplateByClient(args...)
	if rowsPtr != nil {
		templateRows.Columns = rowsPtr.Columns
		templateRows.Rows = rowsPtr.Rows
	}
	return &templateRows
}

//CloseDb database connection
func (db *DbConfig) CloseDb() bool {
	rtn := tmpDb.CloseDb()
	if rtn == true {
		fmt.Println("db connection closed")
	}
	return rtn
}
