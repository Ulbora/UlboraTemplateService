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

package main

import (
	templateManager "UlboraTemplateService/managers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	uoauth "github.com/Ulbora/go-ulbora-oauth2"
	"github.com/gorilla/mux"
)

func handleTemplateChange(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "admin"
	me.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			me.URI = "/ulbora/rs/template/add"
			valid := auth.Authorize(me)
			if valid != true {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				tmpl := new(templateManager.Template)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&tmpl)
				tmpl.ClientID = auth.ClientID
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if tmpl.Name == "" || tmpl.Application == "" || tmpl.ClientID == 0 {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					resOut := templateDB.InsertTemplate(tmpl)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					resJSON, err := json.Marshal(resOut)
					if err != nil {
						log.Println(error.Error())
						http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, string(resJSON))
				}
			}
		case "PUT":
			me.URI = "/ulbora/rs/template/update"
			valid := auth.Authorize(me)
			if valid != true {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				tmpl := new(templateManager.Template)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&tmpl)
				tmpl.ClientID = auth.ClientID
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if tmpl.Application == "" || tmpl.ID == 0 || tmpl.ClientID == 0 {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					resOut := templateDB.UpdateActiveTemplate(tmpl)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					resJSON, err := json.Marshal(resOut)
					if err != nil {
						log.Println(error.Error())
						http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, string(resJSON))
				}
			}
		}
	}
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	app := vars["app"]
	switch r.Method {
	case "GET":
		clientID, errClient := strconv.ParseInt(vars["clientId"], 10, 0)
		if errClient != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
		}
		tmpl := new(templateManager.Template)
		tmpl.Application = app
		tmpl.ClientID = clientID
		resOut := templateDB.GetActiveTemplate(tmpl)
		//fmt.Print("response: ")
		//fmt.Println(resOut)
		resJSON, err := json.Marshal(resOut)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "json output failed", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resJSON))

	}
}

func handleTemplateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	app := vars["app"]
	clientID, errClient := strconv.ParseInt(vars["clientId"], 10, 0)
	if errClient != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	switch r.Method {
	case "GET":
		tmpl := new(templateManager.Template)
		tmpl.ClientID = clientID
		tmpl.Application = app
		resOut := templateDB.GetTemplateByClient(tmpl)
		//fmt.Print("response: ")
		//fmt.Println(resOut)

		resJSON, err := json.Marshal(resOut)
		//fmt.Print("response json: ")
		//fmt.Println(string(resJSON))
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "json output failed", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		if string(resJSON) == "null" {
			fmt.Fprint(w, "[]")
		} else {
			fmt.Fprint(w, string(resJSON))
		}

	}
}

func handleTemplateDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, errID := strconv.ParseInt(vars["id"], 10, 0)
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	clientID, errClient := strconv.ParseInt(vars["clientId"], 10, 0)
	if errClient != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	switch r.Method {
	case "DELETE":
		auth := getAuth(r)
		me := new(uoauth.Claim)
		me.Role = "admin"
		me.Scope = "write"
		me.URI = "/ulbora/rs/template/delete"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			tmpl := new(templateManager.Template)
			tmpl.ClientID = clientID
			tmpl.ID = id
			resOut := templateDB.DeleteTemplate(tmpl)
			//fmt.Print("response: ")
			//fmt.Println(resOut)

			resJSON, err := json.Marshal(resOut)
			//fmt.Print("response json: ")
			//fmt.Println(string(resJSON))
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "json output failed", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			if string(resJSON) == "null" {
				fmt.Fprint(w, "[]")
			} else {
				fmt.Fprint(w, string(resJSON))
			}
		}
	}
}
