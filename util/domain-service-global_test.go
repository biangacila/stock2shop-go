package util

import (
	"testing"
)

func TestDomainServiceGlobal_Delete(t *testing.T) {
	/*//req :=myTestData() // `{"AutoGenerate":false,"AutoGenerateField":"","Data":{"address":"01 Peddie Road Milnerton","appname":"pos","code":"7891","contactemail":"marlenetshilombi@gmail.com","contactnumber":"+27794611664","contactperson":"Marlene Tshilombi","date":"2020-10-10","id":"6675d78b-957f-4848-bf42-f03e6dad05a2","key":1,"name":"Marlene Tshilombi","org":"C100000","orgdatetime":"2020-10-10 09:51:08","profile":null,"status":"active","time":"09:51:08"},"Entity":"insurer","HasUniqueKey":false,"Params":[{"key":"appname","type":"string","val":"pos"},{"key":"org","type":"string","val":"C100000"},{"key":"code","type":"string","val":"7891"}]}`
	hub := DomainServiceGlobal{}
	//	_ = json.Unmarshal([]byte(req), &hub)

	var o interface{}
	var posted PostedRequest
	_ = json.Unmarshal([]byte(myTestData()), &o)
	_ = json.Unmarshal([]byte(myTestData()), &posted)
	str1, _ := json.Marshal(posted.Data)
	dataString := string(str1)

	hub.Entity = "Fields"
	var entity Fields
	if err := json.Unmarshal([]byte(dataString), &entity); err == nil {
		hub.New(entity)
	} else {
		fmt.Println("ERROR: ", err)
	}*/

}
func myTestData() string {
	return `
 {"AutoGenerate":false,"AutoGenerateField":"","Data":{"DataType":"string","DefaultValue":"","DisplaySelection":true,"DisplayTable":true,"Email":false,"FieldName":"Hospital name","LinkEntity":"","Mandatory":true,"Module":"GBV pos","Options":"","Org":"C100002","Phone":false,"Position":1,"Ref":"Hospital","Unique":true,"category":"activity"},"HasUniqueKey":false,"Params":{}}

`
}
