package util

import (
	"stock2shop-go/domain"
	"stock2shop-go/io"
	"encoding/json"
	"fmt"
)

type ServiceGenerate struct {
	Org     string
	AppName string
	Ref     string
	Last    domain.Generate
}

func (obj *ServiceGenerate) Run() {
	//todo get last
	obj.getLast()
	obj.generateNew()
	obj.save()
}
func (obj *ServiceGenerate) save() {
	table := "Generate"
	obj.Last.Id = obj.Ref
	io.InsertTable(DB_NAME, table, obj.Last)
}
func (obj *ServiceGenerate) generateNew() {
	obj.Last.Ref++
	obj.Ref = fmt.Sprintf("%v", obj.Last.Ref)
}
func (obj *ServiceGenerate) getLast() {
	var ls []domain.Generate
	table := "Generate"
	qry := fmt.Sprintf("select * from %v.%v where org='%v' and appname='%v'",
		DB_NAME, table, obj.Org, obj.AppName)
	res, _ := io.RunQueryCass(qry, io.ArrayHost)
	_ = json.Unmarshal([]byte(res), &ls)
	if len(ls) > 0 {
		obj.Last = ls[0]
	} else {
		obj.Last.AppName = obj.AppName
		obj.Last.Org = obj.Org
		if obj.Org == "pos" && obj.Last.AppName == "policy" {
			obj.Last.Ref = 100001
		} else if obj.Org == "pos" && obj.Last.AppName == "product" {
			obj.Last.Ref = 10000001
		} else {
			obj.Last.Ref = 100001
		}

	}
}
