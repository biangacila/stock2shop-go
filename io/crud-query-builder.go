package io

import (
	"encoding/json"
	"fmt"
	"github.com/biangacila/luvungula-go/global"
	"github.com/pborman/uuid"
	"log"
)

type Where struct {
	Key  string
	Val  interface{}
	Type string
}

func RunQueryCass2(qry string) (string, error) {
	qResult := "[]"
	iter := Session.Query(qry).Iter()
	myrow, err := iter.SliceMap()
	if err != nil {
		log.Println("RunQueryCass Cassandra  session.Query Error --->>> ", err, " > ", qry)
		return qResult, err
	}
	str, _ := json.Marshal(myrow)
	qResult = string(str)
	return qResult, nil
}

func RunQueryCass(qry string, fack []string) (string, error) {
	qResult := "[]"
	iter := Session.Query(qry).Iter()
	myrow, err := iter.SliceMap()
	if err != nil {
		log.Println("RunQueryCass Cassandra  session.Query Error --->>> ", err, " > ", qry)
		return qResult, err
	}
	str, _ := json.Marshal(myrow)
	qResult = string(str)
	return qResult, nil
}

func UpdateQuery(dbName, table string, params []Where, setParams []Where) string {
	qry := fmt.Sprintf("update %v.%v ", dbName, table)
	if len(setParams) > 0 {
		var x = 0
		for _, row := range setParams {
			if x == 0 {
				innerQry := fmt.Sprintf(" set %v='%v' ", row.Key, row.Val)
				if row.Type != "string" {
					innerQry = fmt.Sprintf(" set %v=%v ", row.Key, row.Val)
				}
				qry = qry + innerQry
			} else {
				innerQry := fmt.Sprintf(" , %v='%v' ", row.Key, row.Val)
				if row.Type != "string" {
					innerQry = fmt.Sprintf(" , %v=%v ", row.Key, row.Val)
				}
				qry = qry + innerQry
			}
			x++
		}

	}
	if len(params) > 0 {
		var x = 0
		for _, row := range params {
			if x == 0 {
				innerQry := fmt.Sprintf("where %v='%v' ", row.Key, row.Val)
				if row.Type != "string" {
					innerQry = fmt.Sprintf("where %v=%v ", row.Key, row.Val)
				}
				qry = qry + innerQry
			} else {
				innerQry := fmt.Sprintf(" and %v='%v' ", row.Key, row.Val)
				if row.Type != "string" {
					innerQry = fmt.Sprintf(" and %v=%v ", row.Key, row.Val)
				}
				qry = qry + innerQry
			}
			x++
		}
	}
	RunQueryCass2(qry)
	return "OK"
}
func DeleteQuery(dbName, table string, params []Where) string {
	qry := fmt.Sprintf("delete  from %v.%v ", dbName, table)
	if len(params) > 0 {
		var x = 0
		for _, row := range params {
			if x == 0 {
				innerQry := fmt.Sprintf("where %v='%v' ", row.Key, row.Val)
				if row.Type != "string" {
					innerQry = fmt.Sprintf("where %v=%v ", row.Key, row.Val)
				}
				qry = qry + innerQry
			} else {
				innerQry := fmt.Sprintf(" and %v='%v' ", row.Key, row.Val)
				if row.Type != "string" {
					innerQry = fmt.Sprintf(" and %v=%v ", row.Key, row.Val)
				}
				qry = qry + innerQry
			}
			x++
		}
	}
	RunQueryCass2(qry)
	return "OK"
}
func SelectQuery(dbName, table string, params []Where) []byte {
	qry := fmt.Sprintf("select * from %v.%v ", dbName, table)
	if len(params) > 0 {
		var x = 0
		for _, row := range params {
			if x == 0 {
				innerQry := fmt.Sprintf("where %v='%v' ", row.Key, row.Val)
				if row.Type != "string" {
					innerQry = fmt.Sprintf("where %v=%v ", row.Key, row.Val)
				}
				qry = qry + innerQry
			} else {
				innerQry := fmt.Sprintf(" and %v='%v' ", row.Key, row.Val)
				if row.Type != "string" {
					innerQry = fmt.Sprintf(" and %v=%v ", row.Key, row.Val)
				}
				qry = qry + innerQry
			}
			x++
		}
	}
	var ls []interface{}
	//fmt.Println("):(--> ", qry)
	res, _ := RunQueryCass2(qry + " ALLOW FILTERING")
	err := json.Unmarshal([]byte(res), &ls)
	str, err := json.Marshal(ls)
	if err != nil {
		fmt.Println("Error Unmarshal fetch data: ", err, qry)
		/*do nothing*/
	}

	return str
}
func InsertTable(dbName, table string, objIn interface{}) string {
	in := make(map[string]interface{})
	str1, err := json.Marshal(objIn)
	if err != nil {
	}
	err = json.Unmarshal(str1, &in)
	if err != nil {
	}

	dt, hr := global.GetDateAndTimeString()
	if profile, ok := in["Profile"]; ok {
		if profile == nil {
			in["Profile"] = make(map[string]interface{})
		}
	}
	if orgDateTime, ok := in["OrgDateTime"]; ok {
		if fmt.Sprintf("%v", orgDateTime) == "" {
			in["OrgDateTime"] = fmt.Sprintf("%v %v", dt, hr)
		}
	}
	if Status, ok := in["Status"]; ok {
		if fmt.Sprintf("%v", Status) == "" {
			in["Status"] = "active"
		}
	}

	if _, ok := in["Date"]; ok {
		if isValueEmpty(in["Date"]) {
			in["Date"] = dt
		}
	}
	if _, ok := in["Time"]; ok {
		if isValueEmpty(in["Time"]) {
			in["Time"] = hr
		}
	}
	if _, ok := in["Id"]; ok {
		if isValueEmpty(in["Id"]) {
			in["Id"] = uuid.New()
		}

	}

	str, _ := json.Marshal(in)
	qry := fmt.Sprintf("insert into %v.%v  JSON '%v' ", dbName, table, string(str))
	_, err = RunQueryCass2(qry)

	//fmt.Println("InsertTable (:)--> ", err, " > ", qry)
	return "OK"
}

func isValueEmpty(in interface{}) bool {
	str := fmt.Sprintf("%v", in)
	if str == "" {
		return true
	}
	return false
}
