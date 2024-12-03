package mongodb

import (
	"context"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	defaultPage  = 1
	defaultLimit = 20
)

type Pagination struct {
	Page  int
	Limit int
}
type PaginationResponse struct {
	CurrentPage     int `json:"current_page"`
	PageCount       int `json:"page_count"`
	TotalPagesCount int `json:"total_pages_count"`
}

func GetPagination(c *gin.Context) Pagination {
	var (
		page  *int
		limit *int
	)
	if c.Query("page") != "" {
		pageInt, err := strconv.Atoi(c.Query("page"))
		if err == nil {
			page = &pageInt
		}
	}
	if c.Query("limit") != "" {
		limitInt, err := strconv.Atoi(c.Query("limit"))
		if err == nil {
			limit = &limitInt
		}
	}

	if page != nil && limit != nil {
		return Pagination{Page: *page, Limit: *limit}
	} else if page == nil && limit != nil {
		return Pagination{Page: defaultPage, Limit: *limit}
	} else if page != nil && limit == nil {
		return Pagination{Page: *page, Limit: defaultLimit}
	} else {
		return Pagination{Page: defaultPage, Limit: defaultLimit}
	}
}

func (db *Database) SelectAllFromDb(order string, receiver interface{}, query map[string]interface{}, result interface{}) error {
	query = AddDefaultGetParams(query)
	order, sortOrder := getSortDetails(order)
	coll, err := db.GetCollectionForModel(receiver)
	if err != nil {
		return err
	}
	cur, err := coll.Find(context.Background(), query, options.Find().SetSort(bson.D{{order, sortOrder}}))
	if err != nil {
		return err
	}
	defer cur.Close(context.Background())
	return cur.All(context.Background(), result)
}
func (db *Database) SelectAllFromDbWithLimit(order string, limit int, receiver interface{}, query map[string]interface{}, result interface{}) error {
	query = AddDefaultGetParams(query)
	order, sortOrder := getSortDetails(order)
	coll, err := db.GetCollectionForModel(receiver)
	if err != nil {
		return err
	}
	cur, err := coll.Find(context.Background(), query, options.Find().SetSort(bson.D{{order, sortOrder}}).SetLimit(int64(limit)))
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	defer cur.Close(context.Background())
	return cur.All(context.Background(), result)
}

func (db *Database) SelectOneFromDb(receiver interface{}, query map[string]interface{}) error {
	query = AddDefaultGetParams(query)
	coll, err := db.GetCollectionForModel(receiver)
	if err != nil {
		return err
	}
	// filter := bson.M{"$and": []bson.M{{"_id": getDocumentID(receiver)}, {query: args}}}
	return coll.FindOne(context.Background(), query).Decode(receiver)
}

func (db *Database) SelectLatestFromDb(receiver interface{}, query map[string]interface{}) error {
	query = AddDefaultGetParams(query)
	coll, err := db.GetCollectionForModel(receiver)
	if err != nil {
		return err
	}
	// filter := bson.M{"$and": []bson.M{{query: args}}}
	opt := options.FindOne().SetSort(bson.M{"_id": -1})
	return coll.FindOne(context.Background(), query, opt).Decode(receiver)
}

func (db *Database) SelectRandomFromDb(receiver interface{}, query map[string]interface{}) error {
	query = AddDefaultGetParams(query)
	coll, err := db.GetCollectionForModel(receiver)
	if err != nil {
		return err
	}
	// filter := bson.M{"$and": []bson.M{{query: args}}}
	opt := options.FindOne().SetSort(bson.M{"rand()": 1})
	return coll.FindOne(context.Background(), query, opt).Decode(receiver)
}

func (db *Database) SelectFirstFromDb(receiver interface{}) error {
	query := AddDefaultGetParams(map[string]interface{}{})
	coll, err := db.GetCollectionForModel(receiver)
	if err != nil {
		return err
	}
	return coll.FindOne(context.Background(), query).Decode(receiver)
}

func (db *Database) SelectPaginatedFromDb(order string, receiver interface{}, query map[string]interface{}, result interface{}, paginator Pagination) (PaginationResponse, error) {
	query = AddDefaultGetParams(query)
	order, sortOrder := getSortDetails(order)
	coll, err := db.GetCollectionForModel(receiver)
	if err != nil {
		return PaginationResponse{}, err
	}

	// Calculate skip count based on the page number and page size
	skip := (paginator.Page - 1) * paginator.Limit

	// Count the total number of documents that match the query
	totalCount, err := coll.CountDocuments(context.Background(), query)
	if err != nil {
		return PaginationResponse{}, err
	}

	// Calculate the total number of pages
	totalPages := int(math.Ceil(float64(totalCount) / float64(paginator.Limit)))

	cur, err := coll.Find(context.Background(), query, options.Find().
		SetSort(bson.D{{order, sortOrder}}).
		SetSkip(int64(skip)).
		SetLimit(int64(paginator.Limit)))
	if err != nil {
		return PaginationResponse{}, err
	}
	defer cur.Close(context.Background())

	if err := cur.All(context.Background(), result); err != nil {
		return PaginationResponse{}, err
	}

	return PaginationResponse{
		CurrentPage:     paginator.Page,
		PageCount:       cur.RemainingBatchLength(),
		TotalPagesCount: totalPages,
	}, nil
}

func (db *Database) CheckExists(receiver interface{}, query map[string]interface{}) bool {
	query = AddDefaultGetParams(query)
	coll, err := db.GetCollectionForModel(receiver)
	if err != nil {
		return false
	}
	// filter := bson.M{"$and": []bson.M{{"_id": getDocumentID(receiver)}, {query: args}}}
	count, err := coll.CountDocuments(context.Background(), query)
	return err == nil && count > 0
}

func (db *Database) CheckExistsInTable1(table string, query map[string]interface{}) bool {
	query = AddDefaultGetParams(query)
	coll := db.GetCollection(table)
	// filter := bson.M{query: args}
	count, err := coll.CountDocuments(context.Background(), query)
	return err == nil && count > 0
}

func (db *Database) CheckExistsInTable(table string, query map[string]interface{}) bool {
	query = AddDefaultGetParams(query)
	coll := db.GetCollection(table)
	// filter := bson.M{query: args}
	count, err := coll.CountDocuments(context.Background(), query)
	return err == nil && count > 0
}

func AddDefaultGetParams(query map[string]interface{}) map[string]interface{} {
	query["deleted"] = false
	return query
}

func getSortDetails(order string) (string, int) {
	if order == "" {
		order = "-_id"
	}

	sortOrder := 1
	if order[0] == '-' {
		sortOrder = -1
		order = order[1:]
	}

	return order, sortOrder
}
