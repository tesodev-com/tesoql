package tesoql

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoRepository struct {
	mongo     *mongo.Collection
	fieldsMap *FieldsMap
}

func newMongoRepository(cfg *Config) *mongoRepository {
	var collection *mongo.Collection
	var client *mongo.Client

	if cfg.ConnectionConfig.ConnectionString != "" {
		var err error
		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.ConnectionConfig.ConnectionString))
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to mongo: %v", err))
		}
	} else if cfg.ConnectionConfig.Client != nil {
		var ok bool
		client, ok = cfg.ConnectionConfig.Client.(*mongo.Client)
		if !ok || client == nil {
			panic("Provided client is not a valid mongo client")
		}
	} else {
		panic("Mongo config is not defined correctly")
	}

	if cfg.ConnectionConfig.DBName == "" || cfg.ConnectionConfig.TableName == "" {
		panic("Database or collection name is not provided!")
	}

	collection = client.Database(cfg.ConnectionConfig.DBName).Collection(cfg.ConnectionConfig.TableName)
	return &mongoRepository{
		mongo:     collection,
		fieldsMap: cfg.FieldsMap,
	}
}

//func (r *mongoRepository) repository(jsonMap *JsonMap) ([]map[string]interface{}, int, int, *ErrorResponseDTO) {
//	var err error
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
//	defer cancel()
//
//	var opts *options.FindOptions
//	var filter = bson.D{{}}
//
//	query := jsonMap.NewMongoQuery(r.fieldsMap)
//
//	opts = options.Find().SetLimit(query.Limit).SetSkip(query.Offset)
//
//	//if query.Sort == nil && query.Filter == nil && query.Projection == nil {
//	//	return nil, nil, nil, newResponse(TESOQL_MONGO_ERROR, "There is no querying criteria given!", MONGO_EMPTY_QUERY_ERR_CODE)
//	//}
//
//	if query.Projection != nil {
//		opts = opts.SetProjection(query.Projection)
//	}
//
//	if query.Filter != nil {
//		filter = *query.Filter
//	}
//
//	if query.Sort != nil {
//		opts = opts.SetSort(query.Sort)
//	}
//
//	var results []map[string]interface{}
//	var wg sync.WaitGroup
//	var size int
//	var totalCount int
//	if jsonMap.TotalCount {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			countOpts := &options.CountOptions{}
//			totalCount64, _ := r.mongo.CountDocuments(ctx, filter, countOpts)
//			totalCount = int(totalCount64)
//		}()
//	}
//
//	if !jsonMap.SuppressDataResponse {
//		cur, err := r.mongo.Find(ctx, filter, opts)
//		if err != nil {
//			return nil, 0, 0, newResponse(TESOQL_MONGO_ERROR, err.Error(), MONGO_FIND_ERR_CODE)
//		}
//		defer cur.Close(ctx)
//		err = cur.All(ctx, &results)
//		size = cur.RemainingBatchLength()
//	}
//
//	wg.Wait()
//	if err != nil {
//		return nil, 0, 0, newResponse(TESOQL_MONGO_ERROR, err.Error(), MONGO_CURSOR_ERR_CODE)
//	}
//
//	return results, totalCount, size, nil
//
//}

func (r *mongoRepository) repository(jsonMap *JsonMap) ([]map[string]interface{}, int, int, *ErrorResponseDTO) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var opts *options.FindOptions
	var filter = bson.D{{}}

	query := jsonMap.NewMongoQuery(r.fieldsMap)

	opts = options.Find().SetLimit(query.Limit).SetSkip(query.Offset)

	if query.Projection != nil {
		opts = opts.SetProjection(query.Projection)
	}

	if query.Filter != nil {
		filter = *query.Filter
	}

	if query.Sort != nil {
		opts = opts.SetSort(query.Sort)
	}

	var results []map[string]interface{}
	var size int
	var totalCount int

	if jsonMap.TotalCount {
		countOpts := &options.CountOptions{}
		totalCount64, countErr := r.mongo.CountDocuments(ctx, filter, countOpts)
		if countErr != nil {
			return nil, 0, 0, newResponse(TESOQL_MONGO_ERROR, countErr.Error(), MONGO_FIND_ERR_CODE)
		}
		totalCount = int(totalCount64)
	}

	if !jsonMap.SuppressDataResponse {
		cur, err := r.mongo.Find(ctx, filter, opts)
		if err != nil {
			return nil, 0, 0, newResponse(TESOQL_MONGO_ERROR, err.Error(), MONGO_FIND_ERR_CODE)
		}
		defer cur.Close(ctx)
		err = cur.All(ctx, &results)
		if err != nil {
			return nil, 0, 0, newResponse(TESOQL_MONGO_ERROR, err.Error(), MONGO_CURSOR_ERR_CODE)
		}
		size = len(results)
	}

	return results, totalCount, size, nil
}
