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

package main

import (
	templateManager "UlboraTemplateService/managers"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	uoauth "github.com/Ulbora/go-ulbora-oauth2"
	"github.com/gorilla/mux"
)

var templateDB templateManager.TemplateDB

type authHeader struct {
	token    string
	clientID int64
	userID   string
	hashed   bool
}

func main() {
	if os.Getenv("MYSQL_PORT_3306_TCP_ADDR") != "" {
		templateDB.DbConfig.Host = os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	} else if os.Getenv("DATABASE_HOST") != "" {
		templateDB.DbConfig.Host = os.Getenv("DATABASE_HOST")
	} else {
		templateDB.DbConfig.Host = "localhost:3306"
	}

	if os.Getenv("DATABASE_USER_NAME") != "" {
		templateDB.DbConfig.DbUser = os.Getenv("DATABASE_USER_NAME")
	} else {
		templateDB.DbConfig.DbUser = "admin"
	}

	if os.Getenv("DATABASE_USER_PASSWORD") != "" {
		templateDB.DbConfig.DbPw = os.Getenv("DATABASE_USER_PASSWORD")
	} else {
		templateDB.DbConfig.DbPw = "admin"
	}

	if os.Getenv("DATABASE_NAME") != "" {
		templateDB.DbConfig.DatabaseName = os.Getenv("DATABASE_NAME")
	} else {
		templateDB.DbConfig.DatabaseName = "ulbora_template_service"
	}
	templateDB.ConnectDb()
	defer templateDB.CloseDb()

	fmt.Println("Server running!")
	router := mux.NewRouter()
	router.HandleFunc("/rs/template/add", handleTemplateChange).Methods("POST")
	router.HandleFunc("/rs/template/updateActive", handleTemplateChange).Methods("PUT")
	router.HandleFunc("/rs/template/get/{app}/{clientId}", handleTemplate).Methods("GET")
	router.HandleFunc("/rs/template/list/{app}/{clientId}", handleTemplateList).Methods("GET")
	router.HandleFunc("/rs/template/delete/{id}/{clientId}", handleTemplateDelete).Methods("DELETE")
	http.ListenAndServe(":3009", router)
}

func getHeaders(req *http.Request) *authHeader {
	var rtn = new(authHeader)
	authHeader := req.Header.Get("Authorization")
	tokenArray := strings.Split(authHeader, " ")
	if len(tokenArray) == 2 {
		rtn.token = tokenArray[1]
		//fmt.Println(rtn.token)
	}
	userIDHeader := req.Header.Get("userId")
	rtn.userID = userIDHeader

	clientIDHeader := req.Header.Get("clientId")
	clientID, err := strconv.ParseInt(clientIDHeader, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	rtn.clientID = clientID
	if req.Header.Get("hashed") == "true" {
		rtn.hashed = true
	} else {
		rtn.hashed = false
	}
	//fmt.Println(clientIDHeader)
	//fmt.Println(userIDHeader)
	return rtn
}

func getAuth(req *http.Request) *uoauth.Oauth {
	changeHeader := getHeaders(req)
	auth := new(uoauth.Oauth)
	auth.Token = changeHeader.token
	auth.ClientID = changeHeader.clientID
	auth.UserID = changeHeader.userID
	auth.Hashed = changeHeader.hashed
	if os.Getenv("OAUTH2_VALIDATION_URI") != "" {
		auth.ValidationURL = os.Getenv("OAUTH2_VALIDATION_URI")
	} else {
		auth.ValidationURL = "http://localhost:3000/rs/token/validate"
	}
	return auth
}
