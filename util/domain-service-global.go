package util

import (
	"stock2shop-go/domain"
	"stock2shop-go/io"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)


type DomainServiceGlobal struct {
	Org               string
	Params            []domain.Params
	Entity            string
	AutoGenerate      bool
	AutoGenerateField string
	HasUniqueKey      bool
	SetParams         []io.Where
}

func (obj *DomainServiceGlobal) Update() string {
	var conditions = obj.condition()
	str := io.UpdateQuery(DB_NAME, obj.Entity, conditions, obj.SetParams)
	return str
}
func (obj *DomainServiceGlobal) Delete() string {
	var conditions = obj.condition()
	str := io.DeleteQuery(DB_NAME, obj.Entity, conditions)
	return str
}
func (obj *DomainServiceGlobal) List() []map[string]interface{} {
	var ls []map[string]interface{}
	var conditions = obj.condition()
	str := io.SelectQuery(DB_NAME, obj.Entity, conditions)
	_ = json.Unmarshal(str, &ls)
	return ls
}
func (obj *DomainServiceGlobal) New(in interface{}) string {
	str, _ := json.Marshal(in)
	my := make(map[string]interface{})
	_ = json.Unmarshal(str, &my)
	my["AppName"] = io.APP_NAME
	if obj.HasUniqueKey {
		col, _ := UniqueKeys[obj.Entity]
		colVal, _ := my[col]
		org, _ := my["Org"]
		if checkForUniqueKey(io.APP_NAME, obj.Entity, io.DB_NAME, col,
			fmt.Sprintf("%v", colVal), fmt.Sprintf("%v", org)) {
			return errors.New("error Key exist").Error()
		}
	}
	if obj.AutoGenerate {
		if my[obj.AutoGenerateField] == "" {
			hub := ServiceGenerate{}
			hub.AppName = "pos"
			hub.Org = obj.Entity
			hub.Run()
			t, _ := GeneralInfo[obj.Entity]
			my[obj.AutoGenerateField] = fmt.Sprintf("%v%v", t, hub.Ref)
		}
	}
	return io.InsertTable(io.DB_NAME, obj.Entity, my)
}

func (obj *DomainServiceGlobal) condition() []io.Where {
	var conditions []io.Where
	conditions = append(conditions, io.Where{Key: "AppName", Val: io.APP_NAME, Type: "string"})
	for _, row := range obj.Params {
		conditions = append(conditions, io.Where{Key: row.Key, Val: row.Val, Type: row.Type})
	}
	return conditions
}
func checkForUniqueKey(appName, table, dbName, fieldName, fieldVal, org string) bool {
	var ls []interface{}
	qry := fmt.Sprintf("select appname,org,%v from %v.%v where appname='%v' and org='%v'",
		fieldName, dbName, table, appName, org)
	res, _ := io.RunQueryCass2(qry)
	_ = json.Unmarshal([]byte(res), &ls)
	for _, row := range ls {
		var my = make(map[string]string)
		str, _ := json.Marshal(row)
		_ = json.Unmarshal(str, &my)
		for key, val := range my {
			if key == strings.ToLower(fieldName) {
				if val == fieldVal {
					return true
				}
			}
		}
	}
	return false
}
