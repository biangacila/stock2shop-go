package api


import (
	"encoding/json"
	"github.com/biangacila/luvungula-go/global"
	"github.com/gorilla/mux"
	"net/http"
	"stock2shop-go/domain"
	"stock2shop-go/util"
)

type PostedRequest struct {
	AutoGenerate      bool
	HasUniqueKey      bool
	AutoGenerateField string
	Params            []domain.Params
	Data              interface{}
}
type PostedRequestBulk struct {
	AutoGenerate      bool
	HasUniqueKey      bool
	AutoGenerateField string
	Params            []domain.Params
	Data              []interface{}
}

func WsEntityManagement(w http.ResponseWriter, r *http.Request) {
	_, dataString := global.GetPostedDataMapAndString(r)
	var o interface{}
	var posted PostedRequest
	_ = json.Unmarshal([]byte(dataString), &o)
	_ = json.Unmarshal([]byte(dataString), &posted)
	vars := mux.Vars(r)
	module, _ := vars["module"]
	action, _ := vars["action"]
	hub := util.DomainServiceGlobal{}
	Params := make(map[string]domain.Params)
	_ = json.Unmarshal([]byte(dataString), &Params)
	var requestParam interface{}
	hub.AutoGenerate = posted.AutoGenerate
	hub.HasUniqueKey = posted.HasUniqueKey
	hub.AutoGenerateField = posted.AutoGenerateField
	hub.Params = posted.Params
	str1, _ := json.Marshal(posted.Data)
	dataString = string(str1)

	/**
	=== Global endpoint
	*/
	if module == "company" {
		hub.Entity = "Company"
		var entity domain.Company
		if err := json.Unmarshal([]byte(dataString), &entity); err == nil {
			if action == "insert" {
				requestParam = hub.New(entity)
			}
			if action == "list" {
				requestParam = hub.List()
			}
			if action == "delete" {
				requestParam = hub.Delete()
			}
		}
	}
	if module == "util" {
		hub.Entity = "util"
		var entity domain.Util
		if err := json.Unmarshal([]byte(dataString), &entity); err == nil {
			if action == "insert" {
				requestParam = hub.New(entity)
			}
			if action == "list" {
				requestParam = hub.List()
			}
			if action == "delete" {
				requestParam = hub.Delete()
			}
		}
	}
	if module == "user" {
		hub.Entity = "user"
		var entity domain.User
		if err := json.Unmarshal([]byte(dataString), &entity); err == nil {
			if action == "insert" {
				requestParam = hub.New(entity)
			}
			if action == "list" {
				requestParam = hub.List()
			}
			if action == "delete" {
				requestParam = hub.Delete()
			}
		}
	}
	if module == "role" {
		hub.Entity = "role"
		var entity domain.Roles
		if err := json.Unmarshal([]byte(dataString), &entity); err == nil {
			if action == "insert" {
				requestParam = hub.New(entity)
			}
			if action == "list" {
				requestParam = hub.List()
			}
			if action == "delete" {
				requestParam = hub.Delete()
			}
		}
	}

	my := make(map[string]interface{})
	my["STATUS"] = "OK"
	my["REQUEST"] = o
	my["RESULT"] = requestParam
	global.PublishToReact(w, r, my, 200)
}

