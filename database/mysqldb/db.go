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
package mysqldb

import (
	crud "github.com/Ulbora/go-crud-mysql"
)

//ConnectDb connect to db
func ConnectDb(host, user, pw, dbName string) bool {
	res := crud.InitializeMysql(host, user, pw, dbName)
	return res
}

//ConnectionTest get a row. Passing in tx allows for transactions
func ConnectionTest(args ...interface{}) *crud.DbRow {
	rowPtr := crud.Get(ConnectionTestQuery, args...)
	return rowPtr
}

//InsertTemplate insert
func InsertTemplate(args ...interface{}) (bool, int64) {
	success, insID := crud.Insert(nil, InsertTemplateQuery, args...)
	return success, insID
}

//UpdateTemplate update a template
func UpdateTemplate(args ...interface{}) bool {
	success := crud.Update(nil, UpdateTemplateQuery, args...)
	return success
}

//UpdateClearTemplate update a template
func UpdateClearTemplate(args ...interface{}) bool {
	success := crud.Update(nil, UpdateClearNonActiveTemplateQuery, args...)
	return success
}

//GetActiveTemplate get the active template
func GetActiveTemplate(args ...interface{}) *crud.DbRow {
	rowPtr := crud.Get(TemplateGetActiveQuery, args...)
	return rowPtr
}

//GetTemplateByClient templates for a client and app
func GetTemplateByClient(args ...interface{}) *crud.DbRows {
	rowsPtr := crud.GetList(TemplateGetByClientQuery, args...)
	return rowsPtr
}

//DeleteTemplate templates for a client and app
func DeleteTemplate(args ...interface{}) bool {
	rowsPtr := crud.Delete(nil, TemplateDeleteQuery, args...)
	return rowsPtr
}

//CloseDb close connection to db
func CloseDb() bool {
	res := crud.Close()
	return res
}
