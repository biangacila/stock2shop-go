package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/biangacila/luvungula-go/global"
	"os"
	"stock2shop-go/domain"
	"reflect"
)

type Repository struct {
	Host     string
	Keyspace string
	HasFile  bool

	queryTable []string
	queryDB    string
}

func (obj *Repository) New() {
	obj.CreateKeyspace()
	obj.CreateTable("User", "Username", domain.User{}, getStructType(reflect.ValueOf(&domain.User{}).Elem()))

	if obj.HasFile {
		obj.createFile()
	}
}
func (obj *Repository) createFile() {
	fname := fmt.Sprintf("./keyspace-%v.sql", obj.Keyspace)
	os.Remove(fname)
	//global.CreateFolderUploadIfNotExist(fname)
	global.WriteNewLineToLogFile(obj.queryDB, fname)
	global.WriteNewLineToLogFile("\n", fname)
	for _, line := range obj.queryTable {
		global.WriteNewLineToLogFile(line, fname)
		global.WriteNewLineToLogFile("\n", fname)
	}
}
func (obj *Repository) CreateKeyspace() {
	qry := fmt.Sprintf("CREATE KEYSPACE %v WITH replication = {'class': 'NetworkTopologyStrategy', 'dc1': '3'}  AND durable_writes = true;", obj.Keyspace)
	obj.queryDB = qry
}
func (obj *Repository) CreateTable(table, primaryKey string, in interface{}, sTypeList []StructType) {
	qry := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v.%v (\n", obj.Keyspace, table)
	o := make(map[string]interface{})
	str, _ := json.Marshal(in)
	json.Unmarshal(str, &o)
	var x = 0
	for key, _ := range o {
		mp := obj.fnWitchDataType2(key, sTypeList)
		qry = qry + fmt.Sprintf("\t %v %v, \n", key, mp)
		x++
		//todo delete after
	}
	qry = qry + fmt.Sprintf("\t PRIMARY KEY(%v) \n", primaryKey)
	qry = qry + fmt.Sprintf(");")
	obj.queryTable = append(obj.queryTable, qry)
}

type StructType struct {
	Field int
	Key   string
	Value string
	Type  string
}

func getStructType(val reflect.Value) []StructType {
	var arr []StructType
	typeOfTstObj := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldType := val.Field(i)
		o := StructType{}
		o.Field = i
		o.Key = typeOfTstObj.Field(i).Name
		o.Value = fmt.Sprintf("%v", fieldType.Interface())
		o.Type = fmt.Sprintf("%v", fieldType.Type())
		arr = append(arr, o)
	}
	return arr
}
func (obj *Repository) fnWitchDataType2(key string, list []StructType) string {
	o := StructType{}
	for _, row := range list {
		if row.Key == key {
			o = row
		}
	}

	switch o.Type {
	case "int":
		return "text"
	case "int64":
		return "int"
	case "int32":
		return "int"
	case "float64":
		return "float"
	case "float32":
		return "float"
	case "string":
		return "text"
	case "bool":
		return "boolean"
	case "map[string]float64":
		return "map<text,float>"
	case "map[string]string":
		return "map<text,text>"
	case "map[string]interface {}":
		return "map<text,text>"
	default:
		return fmt.Sprintf("frozen<%v >", o.Key)
	}
}
func (obj *Repository) fnWitchDataType(myInterface interface{}) string {
	switch myInterface.(type) {

	case int:
		return "text"
	case int64:
		return "int"
	case int32:
		return "int"
	case float64:
		return "float"
	case float32:
		return "float"
	case string:
		return "text"
	case bool:
		return "boolean"
	case map[string]float64:
		return "map<text,float>"
	case map[string]string:
		return "map<text,text>"
	case map[string]interface{}:
		return "map<text,text>"
	default:
		return "map<text,text>"
	}
}
