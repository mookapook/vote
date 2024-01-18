package vote

import (
	"context"
	"fmt"
	"log"
	"time"

	ttlmap "9mookapook/vote/ttl"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMongo(host, key string) (*mongo.Client, error) {
	return ClientV2(host, key)
}

func model() *ModelImpl {

	if m == nil {
		//dbhost := "" //os.Getenv("MONGOHOST")
		dbname = "itemv"
		cache := ttlmap.New()
		client, err := newMongo(dbhost, "item")
		if err == nil {
			db := client.Database(dbname)
			m = &ModelImpl{db: db, cache: cache}
		}
	}

	return m
}

func (b *ModelImpl) CreateItem(action *Action) (*Action, error) {

	result, err := b.db.Collection(_collectionItem).InsertOne(context.Background(), action)
	if err != nil {
		return nil, err
	}
	action.ID = result.InsertedID.(primitive.ObjectID)
	return action, nil
}

func (b *ModelImpl) UpdateItem(action *Action) (bool, error) {
	success := false
	query := primitive.M{}
	query["_id"] = action.ID
	update := map[string]primitive.M{}
	update["$set"] = primitive.M{
		"name":        action.Name,
		"description": action.Description,
		"updatetime":  time.Now(),
	}
	r, err := b.db.Collection(_collectionItem).UpdateOne(context.Background(), query, update)
	if err != nil {
		return success, err
	}
	if r.ModifiedCount > 0 {
		keyid := "item" + action.ID.Hex()
		b.cache.Del(keyid)
		go b.GetItemVoteByID(action.ID)
		success = true
	}

	return success, nil
}

func (b *ModelImpl) OpenCloseItem(id primitive.ObjectID, status string) (bool, error) {
	success := false
	query := primitive.M{}
	query["_id"] = id

	update := map[string]primitive.M{}
	update["$set"] = primitive.M{
		"status":     status,
		"updatetime": time.Now(),
	}
	r, err := b.db.Collection(_collectionItem).UpdateOne(context.Background(), query, update)
	if err != nil {
		return success, err
	}
	if r.ModifiedCount > 0 {
		keyid := "item" + id.Hex()
		b.cache.Del(keyid)
		go b.GetItemVoteByID(id)
		success = true
	}

	return success, nil
}

func (b *ModelImpl) DeleteItem(id primitive.ObjectID) (bool, error) {
	success := false
	query := primitive.M{}
	query["_id"] = id
	r, err := b.db.Collection(_collectionItem).DeleteOne(context.Background(), query)
	if err != nil {
		return success, err
	}
	// Remove Vote User Filter item
	if r.DeletedCount > 0 {
		keyid := "item" + id.Hex()
		b.cache.Del(keyid)
		//go b.GetItemVoteByID(action.ID)
		filter := primitive.M{}
		filter["itemid"] = id
		success = true
		go b.db.Collection(_collectionVote).DeleteMany(context.Background(), filter)
	}
	return success, nil
}

func (b *ModelImpl) ClearItemAndVoteByID(id primitive.ObjectID) (bool, error) {
	success := false
	query := primitive.M{}
	query["_id"] = id
	update := map[string]primitive.M{}
	update["$set"] = primitive.M{
		"vote":       0,
		"updatetime": time.Now(),
	}
	r, err := b.db.Collection(_collectionItem).UpdateOne(context.Background(), query, update)
	if err != nil {
		return success, err
	}
	// Remove Vote User Filter item
	if r.ModifiedCount > 0 {
		filter := primitive.M{}
		filter["itemid"] = id
		success = true
		keyid := "item" + id.Hex()
		b.cache.Del(keyid)
		delete(userVote, id.Hex())
		go b.db.Collection(_collectionVote).DeleteMany(context.Background(), filter)
	}
	return success, nil
}

func (b *ModelImpl) ClearItemAndVoteALL() (bool, error) {
	success := false
	query := primitive.M{}

	update := map[string]primitive.M{}
	update["$set"] = primitive.M{
		"vote":       0,
		"updatetime": time.Now(),
	}
	r, err := b.db.Collection(_collectionItem).UpdateMany(context.Background(), query, update)
	//r, err := b.db.Collection(_collectionItem).DeleteMany(context.Background(), primitive.M{})
	if err != nil {
		return success, err
	}
	// Remove Vote User Filter item
	if r.ModifiedCount > 0 {
		filter := primitive.M{}
		success = true
		userVote = make(map[string][]string)
		go b.cache.DelALL()
		go b.db.Collection(_collectionVote).DeleteMany(context.Background(), filter)
	}
	return success, nil
}

func (b *ModelImpl) CheckVoteMoreZero(id primitive.ObjectID) (bool, error) {
	data, err := b.GetItemVoteByID(id)
	if err != nil {
		return false, err
	}
	if data.Vote > 0 {
		return false, err
	}
	return true, err
}

func (b *ModelImpl) createIndex() {
	// เลือกคอลเลกชัน
	collection := b.db.Collection(_collectionVote)

	// สร้าง index
	// var isTrue bool = true
	// var isName string = ""
	otp := options.IndexOptions{}
	otp.SetBackground(true)
	otp.SetUnique(true)
	otp.SetName("itemvoteUser")

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "itemid", Value: -1},
				{Key: "userid", Value: -1},
			},
			Options: &otp,
		},
	}
	_, err := collection.Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		log.Panicln(err)
		panic(err)
	}
	// isTrue = false
	// isName = "itemid_vote"
	otp.SetBackground(true)
	otp.SetUnique(false)
	otp.SetName("itemid_vote")
	indexesOne := mongo.IndexModel{
		Keys: bson.M{
			"itemid": 1,
		},
		Options: &otp,
	}
	_, err = collection.Indexes().CreateOne(context.Background(), indexesOne)
	if err != nil {
		panic(err)
	}
}

func (b *ModelImpl) GetItemVoteByID(id primitive.ObjectID) (action Action, err error) {

	keyid := "item" + id.String()
	action = Action{}
	item, _b := b.cache.Get(keyid)
	if _b == false {
		filter := primitive.M{}
		filter["_id"] = id

		err := b.db.Collection(_collectionItem).FindOne(context.Background(), filter, nil).Decode(&action)
		if err != nil {
			return action, err
		}
		b.cache.Add(keyid, action, 60*time.Minute)
	} else {
		//log.Println("cache")
		action = item.(Action)
	}

	return action, err
}

func (b *ModelImpl) VoteItemByUser(vote *VoteUser) (r bool, err error) {
	r = false
	//  Check Vote Exits with index Uniq userid , itemid
	_, err = b.db.Collection(_collectionVote).InsertOne(context.Background(), vote)
	if err != nil {
		return r, err
	}
	VoteUserMap(vote.Itemid.Hex(), vote.UserID)
	query := primitive.M{}
	query["_id"] = vote.Itemid
	update := map[string]primitive.M{}
	update["$set"] = primitive.M{
		"updatetime": time.Now(),
	}
	update["$inc"] = primitive.M{
		"vote": 1,
	}
	//otp := options.Update().SetUpsert(true)
	_, errz := b.db.Collection(_collectionItem).UpdateOne(context.Background(), query, update)
	if errz != nil {
		return r, errz
	}
	keyid := "item" + vote.Itemid.Hex()
	b.cache.Del(keyid)
	go b.GetItemVoteByID(vote.Itemid)
	return true, err
}

func (b *ModelImpl) GetAllItem(skip, limit int, sortby, user, status string) []Action {

	sort := primitive.M{}

	sort["vote"] = -1

	and := []primitive.M{
		{"$eq": []interface{}{"$itemid", "$$contentid"}},
		{"$eq": []interface{}{user, "$userid"}},
	}
	pipelineStatement := []primitive.M{}
	pipelineops := primitive.M{}

	if status != "all" {
		pipelineops["$match"] = primitive.M{"status": status}
		pipelineStatement = append(pipelineStatement, pipelineops)
	}

	pipelineops = primitive.M{}
	pipelineops["$lookup"] = primitive.M{
		"from": "vote",
		"let": primitive.M{
			"contentid": "$_id",
		},
		"pipeline": []primitive.M{
			{
				"$match": primitive.M{
					"$expr": primitive.M{
						"$and": and,
					},
				},
			},
		},
		"as": "votedata",
	}
	pipelineStatement = append(pipelineStatement, pipelineops)
	pipelineops = primitive.M{}
	pipelineops["$project"] = primitive.M{
		"_id":         1,
		"name":        1,
		"description": 1,
		"vote":        1,
		"votedata":    "$votedata.userid",
	}
	pipelineStatement = append(pipelineStatement, pipelineops)
	pipelineops = primitive.M{}
	pipelineops["$skip"] = skip
	pipelineStatement = append(pipelineStatement, pipelineops)
	pipelineops = primitive.M{}
	pipelineops["$limit"] = limit
	pipelineStatement = append(pipelineStatement, pipelineops)
	pipelineops = primitive.M{}
	pipelineops["$sort"] = sort
	pipelineStatement = append(pipelineStatement, pipelineops)
	//pipelineStatement = append(pipelineStatement, pipelineops)
	// pipeline := []primitive.M{
	// 	{
	// 		"$match": primitive.M{
	// 			"status": status,
	// 		},
	// 	},
	// 	{
	// 		"$lookup": primitive.M{
	// 			"from": "vote",
	// 			"let": primitive.M{
	// 				"contentid": "$_id",
	// 			},
	// 			"pipeline": []primitive.M{
	// 				{
	// 					"$match": primitive.M{
	// 						"$expr": primitive.M{
	// 							"$and": and,
	// 						},
	// 					},
	// 				},
	// 			},
	// 			"as": "votedata",
	// 		},
	// 	}, {
	// 		"$project": primitive.M{
	// 			"_id":         1,
	// 			"name":        1,
	// 			"description": 1,
	// 			"vote":        1,
	// 			"votedata":    "$votedata.userid",
	// 		},
	// 	},
	// 	{
	// 		"$skip": skip,
	// 	},
	// 	{
	// 		"$limit": limit,
	// 	},
	// 	{
	// 		"$sort": sort,
	// 	},
	// }

	// pipeline = []primitive.M{
	// 	{
	// 		"$sort": primitive.M{
	// 			"vote": -1, // เรียงลำดับจากมากไปน้อย
	// 		},
	// 	},
	// }
	//pipeline = append(pipeline, P)

	// ทำ aggregation และดึงข้อมูล
	cursor, err := b.db.Collection(_collectionItem).Aggregate(context.Background(), pipelineStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var items []Action
	if err := cursor.All(context.Background(), &items); err != nil {
		log.Fatal(err)
	}

	for i := range items {
		log.Println(len(items[i].Votes))
		if len(items[i].Votes) > 0 {
			items[i].HasVoted = true
		} else {
			items[i].HasVoted = false
		}

	}
	return items
}

func (b *ModelImpl) CheckVote(q primitive.M) bool {
	var data VoteUser
	if err := b.db.Collection(_collectionVote).FindOne(context.Background(), q).Decode(&data); err != nil {
		//log.Println(err.Error())
		return false
	}
	if data.UserID != "" {
		//log.Println("Set")
		VoteUserMap(data.Itemid.Hex(), data.UserID)
		go UserUniq(data.Itemid.Hex())
	}
	return true
}

func (b *ModelImpl) ReportItem(status, st, end string) []Action {
	filter := primitive.M{}
	if status != "" {
		filter["status"] = status
	}
	loc, _ := time.LoadLocation("Asia/Bangkok")
	if !IsEmpty(st) {
		if len(st) == 10 {
			if tt, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59+07:00", st)); err == nil {
				tt = tt.In(loc)
				filter["cd"] = primitive.M{"$gte": tt}
			}
		} else {
			if tt, err := time.Parse(time.RFC3339, st); err == nil {
				tt = tt.In(loc)
				filter["cd"] = primitive.M{"$gte": tt}
			}
		}
	}
	if !IsEmpty(end) {
		if len(st) == 10 {
			if tt, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59+07:00", end)); err == nil {
				tt = tt.In(loc)
				filter["cd"] = primitive.M{"$lte": tt}
			}
		} else {
			if tt, err := time.Parse(time.RFC3339, end); err == nil {
				tt = tt.In(loc)
				filter["cd"] = primitive.M{"$lte": tt}
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := b.db.Collection(_collectionItem).Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var items []Action
	if err := cursor.All(context.Background(), &items); err != nil {
		log.Fatal(err)
	}

	return items
}

func (b *ModelImpl) ReportVoteItemById(id primitive.ObjectID, st, end string) []VoteUser {
	filter := primitive.M{}

	filter["itemid"] = id

	loc, _ := time.LoadLocation("Asia/Bangkok")
	if !IsEmpty(st) {
		if len(st) == 10 {
			if tt, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59+07:00", st)); err == nil {
				tt = tt.In(loc)
				filter["cd"] = primitive.M{"$gte": tt}
			}
		} else {
			if tt, err := time.Parse(time.RFC3339, st); err == nil {
				tt = tt.In(loc)
				filter["cd"] = primitive.M{"$gte": tt}
			}
		}
	}
	if !IsEmpty(end) {
		if len(st) == 10 {
			if tt, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59+07:00", end)); err == nil {
				tt = tt.In(loc)
				filter["cd"] = primitive.M{"$lte": tt}
			}
		} else {
			if tt, err := time.Parse(time.RFC3339, end); err == nil {
				tt = tt.In(loc)
				filter["cd"] = primitive.M{"$lte": tt}
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := b.db.Collection(_collectionVote).Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var items []VoteUser
	if err := cursor.All(context.Background(), &items); err != nil {
		log.Fatal(err)
	}

	return items
}
