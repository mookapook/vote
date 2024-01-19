package vote

import (
	"sync"
	"time"

	ttlmap "9mookapook/vote/ttl"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	digit  = "qazwsx123"
	client *mongo.Client
	dbname string
	//dbhost = "mongodb:27017"
	dbhost = "localhost:27017"

	_collectionItem = "item"
	_collectionVote = "vote"
	tokenTest       = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMxNjI1MjQsIm5hbWUiOiJOYW1lMSBBbGljZSIsInVzZXJJZCI6InVzZXIxIn0.sHjTfXXEN4OEM7V4hi6aem3R4dJEEOr9FBgk-iRhT2I"
	telegramid      = -390713380
	m               *ModelImpl
)

const _dateLayout = "2006/01/02 15:04"

type ModelImpl struct {
	db    *mongo.Database
	cache *ttlmap.TTLMap
	mutex *sync.RWMutex

	// define any necessary dependencies
}

type Controller struct {
	// define any necessary dependencies
	model *ModelImpl
}

const (
	Open  string = "open"
	Close        = "close"
	All          = "all"
)

type Action struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty" `
	Vote        int                `json:"vote" bson:"vote" `
	Votes       []string           `json:"-" bson:"votedata,omitempty" `
	HasVoted    bool               `bson:"hasVoted,omitempty"`
	UserName    string             `json:"username" bson:"username"`
	CreateTime  time.Time          `json:"-" bson:"createtime"`
	UpDateTime  time.Time          `json:"-" bson:"updatetime"`
}

type OpenClose struct {
	Status string `json:"status" bson:"status" validate:"required"`
}

type VoteUser struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     string             `json:"userid" bson:"userid"`
	Itemid     primitive.ObjectID `json:"-" bson:"itemid"`
	ItemidHex  string             `json:"itemid" bson:"-"`
	CreateTime time.Time          `json:"-" bson:"createtime"`
	UpDateTime time.Time          `json:"-" bson:"updatetime"`
}

type DataUser struct {
	UserID   string `json:"userid" bson:"userid"`
	Password string `json:"password" bson:"itemid" validate:"required"`
	UserName string `json:"username" bson:"-" validate:"required"`
}

type User struct {
	Name  string `json:"name" bson:"name"`
	Token string `json:"token" bson:"token" `
}

type VoteRepository interface {
	CreateItem(item *Action) (*Action, error)
	DeleteItem(id primitive.ObjectID) (bool, error)
	UpdateItem(item *Action) (bool, error)
	OpenCloseItem(id primitive.ObjectID, status string) (bool, error)
	ClearItemAndVoteByID(id primitive.ObjectID) (bool, error)
	ClearItemAndVoteALL() (bool, error)
	GetItemVoteByID(id primitive.ObjectID) (*Action, error)
	VoteItemByUser(vote *VoteUser) (bool, error)
	GetAllItem(skip, limit int, sortby string) []Action
}
