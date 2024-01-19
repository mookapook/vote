package vote

import (
	ttlmap "9mookapook/vote/ttl"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var itemtest = "65a7887f166933dd2b7a834f"

func TestCreateItemModel(t *testing.T) {
	client2, err := ConnectDB()
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	cache := ttlmap.New()
	model := &ModelImpl{db: client2.Database("itemv"), cache: cache, mutex: new(sync.RWMutex)}

	action := &Action{
		Name:        "Test Action",
		Description: "Test Description",
		Vote:        0,
		Status:      "open",
		UserName:    "user1",
	}

	createdAction, err := model.CreateItem(action)
	if err != nil {
		t.Errorf("CreateBan failed: %v", err)
	}

	if createdAction.ID == primitive.NilObjectID {
		t.Error("Created action did not have an ID assigned")
	}

	if createdAction.Name != action.Name || createdAction.Description != action.Description {
		t.Error("Created action did not match the input action")
	}
}

func TestConvertHex(t *testing.T) {
	_, err := Converthex("xzczxasasdsadasdasdsa")
	if err != nil {
		t.Errorf("IDHEX failed: %v", err)
	}
}

func ConnectDB() (*mongo.Client, error) {
	key := "testKey"
	return ClientV2(dbhost, key)
}

func TestUpdateItemModel(t *testing.T) {
	client2, err := ConnectDB()
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	cache := ttlmap.New()
	model := &ModelImpl{db: client2.Database("itemv"), cache: cache, mutex: new(sync.RWMutex)}

	action := &Action{
		Name:        "Test Action",
		Description: "Test DescriptionUpdate",
	}
	// Error Convert ObjectID

	ID, err := Converthex("65a5ff28f189c5fd5dfaa7a2")
	if err != nil {
		t.Errorf("IDHEX failed: %v", err)
	}

	// Error NO ID
	ID, _ = Converthex(itemtest)
	action.ID = ID
	updateAction, err := model.UpdateItem(action)
	if err != nil {
		t.Errorf("updateAction failed: %v", err)
	}

	if updateAction == false {
		t.Error("Update action did not have an ID assigned")
	}

}

func TestDeleteItemModel(t *testing.T) {
	client2, err := ConnectDB()
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	cache := ttlmap.New()
	model := &ModelImpl{db: client2.Database("itemv"), cache: cache, mutex: new(sync.RWMutex)}

	// Error NO ID
	ID, err := Converthex("65a5ff28f189c5fd5dfaa7a1")
	if err != nil {
		t.Errorf("updateAction failed: %v", err)
	}
	updateAction, err := model.DeleteItem(ID)
	if err != nil {
		t.Errorf("updateAction failed: %v", err)
	}
	if updateAction == false {
		t.Error("Update action did not have an ID assigned")
	}

	// ID, err = Converthex("65a5ff28f189c5fd5dfaa7a2")
	// if err != nil {
	// 	t.Errorf("IDHEX failed: %v", err)
	// }

	// updateAction, err = model.DeleteItem(ID)
	// if err != nil {
	// 	t.Errorf("updateAction failed: %v", err)
	// }

	// if updateAction == false {
	// 	t.Error("Update action did not have an ID assigned")
	// }

}

func TestGetdataByID(t *testing.T) {
	client2, err := ConnectDB()
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	cache := ttlmap.New()
	model := &ModelImpl{db: client2.Database("itemv"), cache: cache, mutex: new(sync.RWMutex)}

	// Error NO ID
	ID, err := Converthex("65a5ff28f189c5fd5dfaa7a2")
	if err != nil {
		t.Errorf("updateAction failed: %v", err)
	}
	data, err := model.GetItemVoteByID(ID)
	if err != nil {
		t.Errorf("Nodata failed: %v", err)
	}

	log.Printf("%#v", data)

	data2, err := model.GetItemVoteByID(ID)
	if err != nil {
		t.Errorf("Nodata failed: %v", err)
	}

	log.Printf("%#v", data2)
}

func TestVoteByItem(t *testing.T) {
	client2, err := ConnectDB()
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	cache := ttlmap.New()
	model := &ModelImpl{db: client2.Database("itemv"), cache: cache, mutex: new(sync.RWMutex)}

	v := VoteUser{}
	v.CreateTime = time.Now()
	v.UpDateTime = time.Now()
	v.UserID = "user1"
	v.Itemid, _ = Converthex("65a5ff28f189c5fd5dfaa7a2")
	r, err := model.VoteItemByUser(&v)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			t.Errorf("duplicate failed: %v", err)
		} else {
			t.Errorf("Nodata failed: %v", err)
		}
	}
	if r == false {
		t.Errorf("Nodata Vote: %v", err)
	}
	vote := model.GetAllItem(0, 10+1, "vote", "user1", "all")
	if len(vote) == 11 {
		//temp := datainfo.Content
		vote = vote[:len(vote)-1]
	}
	log.Println(vote)
	vote = model.GetAllItem(10, 10+1, "vote", "user1", "all")
	if len(vote) == 11 {
		//temp := datainfo.Content
		vote = vote[:len(vote)-1]
	}

	log.Println(vote)
}

////

func TestCreateitemVote(t *testing.T) {
	// Set up mock dependencies
	//mockModel := &ModelImpl{}

	controller := NewController()
	// Create a mock Echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/itemcreate", strings.NewReader(`{"name": "Test Action", "description": "Test Description"}`))
	req.Header.Set("Authorization", "Bearer "+tokenTest)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if err := JWTAuthMiddleware(controller.CreateitemVote)(c); err == nil {
		var resp interface{}
		if err := json.Unmarshal(res.Body.Bytes(), &resp); err == nil {
			log.Println(res.Code)
		} else {
			t.Errorf("test.getMeAffiliateHandler.Unmarshal: %#v  %s\n", err.Error(), res.Body.String())
		}
	} else {
		t.Errorf("test.getMeAffiliateHandler: %s", err.Error())
	}

	if res.Code == 200 {

		var responseBody map[string]Action
		if err := json.Unmarshal(res.Body.Bytes(), &responseBody); err != nil {
			t.Errorf("Failed to unmarshal response body: %v", err)
		}

		data := responseBody["data"]

		if data.Name != "Test Action" || data.Description != "Test Description" {
			t.Errorf("Created action did not match the input data")
		}

		if data.Vote != 0 || data.Status != "open" {
			t.Errorf("Created action fields were not set correctly")
		}
	}

}

func TestModelCheckCanEdit(t *testing.T) {
	id := []string{itemtest, "65a693a572b0cbfc9cbf1eaa", "65a693a572b0cbfc9cbf1eaa1"}
	for _, data := range id {
		controller := NewController()
		// Create a mock Echo context
		e := echo.New()

		urlstring := "/v1/itemcreate/" + data
		log.Println(urlstring)
		req := httptest.NewRequest(http.MethodPut, urlstring, strings.NewReader(`{"name": "Test ActionU", "description": "Test Description"}`))
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMyNTkxMjAsIm5hbWUiOiJBbGljZSIsInVzZXJJZCI6InVzZXIxIn0._reT2S8x9r96oHbjslc7U6-idEheSIuudEH3ot64kgc")
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues(data)
		if err := JWTAuthMiddleware(controller.UpdateitemVote)(c); err == nil {
			var resp interface{}
			if err := json.Unmarshal(res.Body.Bytes(), &resp); err == nil {
				log.Println(res.Code)
			} else {
				t.Errorf("test.getMeAffiliateHandler.Unmarshal: %#v  %s\n", err.Error(), res.Body.String())
			}
		} else {
			t.Errorf("test.getMeAffiliateHandler: %s", err.Error())
		}

		if res.Code == 200 {
			var responseBody map[string]bool
			if err := json.Unmarshal(res.Body.Bytes(), &responseBody); err != nil {
				t.Errorf("Failed to unmarshal response body: %v", err)
			}

			if responseBody["data"] != true {
				t.Errorf("Failed to Update Data")
			}

		} else {
			var responseBody map[string]string
			if err := json.Unmarshal(res.Body.Bytes(), &responseBody); err != nil {
				t.Errorf("Failed to unmarshal response body: %v", err)
			}
			if responseBody["error"] == "" {
				t.Errorf("Failed to Update Data %s", responseBody["error"])
			}
		}
	}
}
func TestVotebyUser(t *testing.T) {
	client2, err := ConnectDB()
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	cache := ttlmap.New()
	model := &ModelImpl{db: client2.Database("itemv"), cache: cache, mutex: new(sync.RWMutex)}
	user := []string{"user1", "user2"}
	for _, data := range user {
		q := primitive.M{}
		q["userid"] = data
		itemid, _ := Converthex("65a5ff28f189c5fd5dfaa7a2")
		q["itemid"] = itemid
		var ck bool
		if val, ok := userVote["65a5ff28f189c5fd5dfaa7a2"]; ok {
			//val = append(val, userid)
			if Contains(val, data) {
				t.Errorf(" Vote  Data MapUser %s", data)
				ck = true
			}

		}
		if ck == false {
			b := model.CheckVote(q)
			if b == true {
				t.Errorf(" Vote  Data User %s", data)
			} else {
				v := VoteUser{}
				v.CreateTime = time.Now()
				v.UpDateTime = time.Now()
				v.Itemid = itemid
				v.UserID = data
				c, err := model.VoteItemByUser(&v)
				if err != nil || c == false {
					t.Errorf(" CannotVote  Data User %s", data)
				}
			}

		}
	}

	if val, ok := userVote["65a5ff28f189c5fd5dfaa7a2"]; ok {
		//val = append(val, userid)
		if Contains(val, "user1") {
			t.Errorf(" Vote  Data MapUser %s", "user1")
		}

	}
}

func TestUnVotebyUser(t *testing.T) {
	client2, err := ConnectDB()
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	cache := ttlmap.New()
	model := &ModelImpl{db: client2.Database("itemv"), cache: cache, mutex: new(sync.RWMutex)}

	user := []string{"user1"}
	for _, data := range user {
		q := primitive.M{}
		q["userid"] = data
		itemid, _ := Converthex("65a7887f166933dd2b7a834f")
		q["itemid"] = itemid
		var ck bool
		userVote["65a7887f166933dd2b7a834f"] = append(userVote["65a7887f166933dd2b7a834f"], data)
		if val, ok := userVote["65a7887f166933dd2b7a834f"]; ok {
			//val = append(val, userid)
			if !Contains(val, data) {
				t.Errorf(" Vote  Data MapUser %s", data)
				ck = true
			}

		}
		if ck == false {
			b := model.CheckVote(q)
			if b == false {
				t.Errorf(" Vote  Data User %s", data)
			} else {
				v := VoteUser{}
				v.CreateTime = time.Now()
				v.UpDateTime = time.Now()
				v.Itemid = itemid
				v.UserID = data
				c, err := model.UnvoteItem(&v)
				if err != nil || c == false {
					t.Errorf(" CannotVote  Data User %s", data)
				}
			}

		}
	}
	data := "user1"
	if val, ok := userVote["65a7887f166933dd2b7a834f"]; ok {
		//val = append(val, userid)
		if Contains(val, "user1") {
			t.Errorf(" Vote  Data MapUser 2 %s", "user1")
		} else {
			q := primitive.M{}
			q["userid"] = data
			itemid, _ := Converthex("65a7887f166933dd2b7a834f")
			b := model.CheckVote(q)
			if b == false {
				t.Errorf(" Vote  Data User %s", data)
			} else {
				v := VoteUser{}
				v.CreateTime = time.Now()
				v.UpDateTime = time.Now()
				v.Itemid = itemid
				v.UserID = data
				c, err := model.VoteItemByUser(&v)
				if err != nil || c == false {
					t.Errorf(" CannotVote  Data User %s", data)
				}
			}
		}

	}
	if val, ok := userVote["65a7887f166933dd2b7a834f"]; ok {
		if Contains(val, "user1") {
			t.Errorf(" Vote  Data MapUser 2 %s", "user1")
		}

	}
}

func TestVoteByItemHttp(t *testing.T) {
	id := []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMyNTkxMjAsIm5hbWUiOiJBbGljZSIsInVzZXJJZCI6InVzZXIxIn0._reT2S8x9r96oHbjslc7U6-idEheSIuudEH3ot64kgc"}
	for _, data := range id {
		controller := NewController()
		// Create a mock Echo context
		e := echo.New()

		urlstring := "/v1/itemvote/" + itemtest
		log.Println(urlstring)
		req := httptest.NewRequest(http.MethodPut, urlstring, nil)
		req.Header.Set("Authorization", "Bearer "+data)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues(itemtest)
		if err := JWTAuthMiddleware(controller.itemVoteByID)(c); err == nil {
			var resp interface{}
			if err := json.Unmarshal(res.Body.Bytes(), &resp); err == nil {
				log.Println(res.Code)
			} else {
				t.Errorf("test.getMeAffiliateHandler.Unmarshal: %#v  %s\n", err.Error(), res.Body.String())
			}
		} else {
			t.Errorf("test.getMeAffiliateHandler: %s", err.Error())
		}

		if res.Code == 200 {
			var responseBody map[string]interface{}
			if err := json.Unmarshal(res.Body.Bytes(), &responseBody); err != nil {
				t.Errorf("Failed to unmarshal response body: %v", err)
			}

			if responseBody["data"].(bool) != true {
				t.Errorf("Failed to Update Data %s", responseBody["error"].(string))
			}

		} else {
			var responseBody map[string]string

			if err := json.Unmarshal(res.Body.Bytes(), &responseBody); err != nil {
				t.Errorf("Failed to unmarshal response body: %v", err)
			}
			log.Println(responseBody)
			if responseBody["error"] == "" {
				t.Errorf("Failed to Update Data %s", responseBody["error"])
			}
		}
	}
}

func TestHttpRemoveByID(t *testing.T) {
	id := []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMyNTkxMjAsIm5hbWUiOiJBbGljZSIsInVzZXJJZCI6InVzZXIxIn0._reT2S8x9r96oHbjslc7U6-idEheSIuudEH3ot64kgc"}
	for _, data := range id {
		controller := NewController()
		// Create a mock Echo context
		e := echo.New()

		urlstring := "/v1/itemcreate/" + itemtest
		log.Println(urlstring)
		req := httptest.NewRequest(http.MethodDelete, urlstring, nil)
		req.Header.Set("Authorization", "Bearer "+data)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetParamNames("id")
		c.SetParamValues(itemtest)
		if err := JWTAuthMiddleware(controller.itemVoteByID)(c); err == nil {
			var resp interface{}
			if err := json.Unmarshal(res.Body.Bytes(), &resp); err == nil {
				log.Println(res.Code)
			} else {
				t.Errorf("test.getMeAffiliateHandler.Unmarshal: %#v  %s\n", err.Error(), res.Body.String())
			}
		} else {
			t.Errorf("test.getMeAffiliateHandler: %s", err.Error())
		}

		if res.Code == 200 {
			var responseBody map[string]interface{}
			if err := json.Unmarshal(res.Body.Bytes(), &responseBody); err != nil {
				t.Errorf("Failed to unmarshal response body: %v", err)
			}

			if responseBody["data"].(bool) != true {
				t.Errorf("Failed to Update Data %s", responseBody["error"].(string))
			}

		} else {
			var responseBody map[string]string

			if err := json.Unmarshal(res.Body.Bytes(), &responseBody); err != nil {
				t.Errorf("Failed to unmarshal response body: %v", err)
			}
			log.Println(responseBody)
			if responseBody["error"] == "" {
				t.Errorf("Failed to Update Data %s", responseBody["error"])
			}
		}
	}
}

func TestGetALLItem(t *testing.T) {
	client2, err := ConnectDB()
	if err != nil {
		t.Fatalf("ClientV2 failed to connect with key: %v", err)
	}
	cache := ttlmap.New()
	model := &ModelImpl{db: client2.Database("itemv"), cache: cache, mutex: new(sync.RWMutex)}
	vote := model.GetAllItem(0, 10+1, "vote", "user1", "open")
	if len(vote) == 11 {
		//temp := datainfo.Content
		vote = vote[:len(vote)-1]
	}
	log.Println(vote)
	vote = model.GetAllItem(10, 10+1, "vote", "user1", "open")
	if len(vote) == 11 {
		//temp := datainfo.Content
		vote = vote[:len(vote)-1]
	}

	log.Println(vote)
}

func TestLoginSystem(t *testing.T) {
	// Set up mock dependencies
	//mockModel := &ModelImpl{}

	controller := NewController()
	// Create a mock Echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(`{"username": "Alice", "password": "123456"}`))
	req.Header.Set("Authorization", "Bearer "+tokenTest)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if err := controller.LoginUser(c); err != nil {

	}
	if res.Code == 200 {

		var responseBody map[string]string
		if err := json.Unmarshal(res.Body.Bytes(), &responseBody); err != nil {
			t.Errorf("Failed to unmarshal response body: %v", err)
		}
		data := responseBody["token"]

		if responseBody["error"] != "" {
			t.Errorf("Created action did not match the input data")
		}
		log.Println(data)
	}

}
