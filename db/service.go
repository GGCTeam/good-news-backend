package db

import (
	"context"
	"fmt"
	"time"

	"github.com/kanzitelli/good-news-backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Service <type>
// is used to describe mongo service
type Service struct {
	client *mongo.Client
}

var s *Service

// InitService <function>
// is used to initialize mongo client
func InitService() {
	envVars := utils.GetEnvVars()

	s = new(Service)

	ctx, cncl := s.CTX(120)
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI(envVars.MongoDBURL))
	if err != nil {
		fmt.Println("MONGO CONNECT ERROR", err)
	}
	cncl()

	ctx, cncl = s.CTX(120)
	err = mClient.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("MONGO PING ERROR", err)
	}
	cncl()

	s.client = mClient
}

// GetClient <function>
// is used to return mongo client
func GetClient() *Service {
	return s
}

// CTX <function>
// is used to easily create mongo context
func (s Service) CTX(secs time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), secs*time.Second)
}

// Collection <function>
// is used to get easy access to bussiness collection in database
func (s Service) Collection(colName string) *mongo.Collection {
	envVars := utils.GetEnvVars()

	c := GetClient()
	return c.client.Database(envVars.MongoDBName).Collection(colName)
}

// InsertManyOptionsOrdered <function>
// is used to simplify access to pre-ready insert many options
func (s Service) InsertManyOptionsOrdered() *options.InsertManyOptions {
	// ordered property set to false, so mongo ignore errors with duplicate _ids of news
	return &options.InsertManyOptions{
		Ordered: func(b bool) *bool { return &b }(false), // kind of strange but it needs to be done like this, because it needs pointer of bool, you can check mongodb documentation
	}
}
