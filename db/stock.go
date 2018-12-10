package db

import (
	"context"
	"fmt"
	"stock/schema"
	"stock/util"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

const STOCKCOLLECTION = "stock"

type StockRepository struct {
	Conn *mongo.Database
}

func NewMongo() (*StockRepository, error) {
	client, err := util.GetDBConnection()
	if err != nil {
		return nil, err
	}
	client.Connect(context.Background())
	db := client.Database("stock-database")
	return &StockRepository{
		db,
	}, nil
}

func (r *StockRepository) CreateStock(ctx context.Context, stock *schema.Stock) error {
	_, err := r.Conn.Collection(STOCKCOLLECTION).InsertOne(
		ctx,
		stock)
	if err != nil {
		return err
	}
	return nil
}

func (r *StockRepository) ModifyStock(ctx context.Context, id string, quantity int) error {

	idDoc := bson.D{{"product_id", id}}
	var stock *schema.Stock
	err := r.Conn.Collection(STOCKCOLLECTION).FindOne(ctx, idDoc).Decode(&stock)
	if err != nil {
		return err
	}

	stock.Quantity += quantity
	if stock.Quantity < 0 {
		return fmt.Errorf("Not enough stock for this item")
	}
	_, err = r.Conn.Collection(STOCKCOLLECTION).UpdateOne(
		ctx,
		idDoc,
		bson.D{{"$set", bson.D{{"quantity", stock.Quantity}}}},
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *StockRepository) GetStock(ctx context.Context, id string) (*schema.Stock, error) {
	idDoc := bson.D{{"product_id", id}}
	var stock *schema.Stock
	err := r.Conn.Collection(STOCKCOLLECTION).FindOne(ctx, idDoc).Decode(&stock)
	if err != nil {
		return nil, err
	}

	return stock, nil
}
