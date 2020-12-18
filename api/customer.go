package api

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type customer struct {
	CustomerID      int32  `bson:"customer_id" json:"customer_id"`
	CustomerName    string `bson:"customer_name" json:"customer_name"`
	CustomerAddress string `bson:"customer_address" json:"customer_address"`
	City            string `bson:"city" json:"city"`
	State           string `bson:"state" json:"state"`
	PostalCode      int32  `bson:"postal_code" json:"postal_code"`
}

type orderlines struct {
	OrderID         int32 `bson:"order_id"`
	ProductID       int32 `bson:"product_id"`
	OrderedQuantity int32 `bson:"ordered_quantity"`
}

type product struct {
	ProductID      int32   `bson:"product_id"`
	ProductDescrip string  `bson:"product_description"`
	ProductFinish  string  `bson:"product_finish"`
	StandardPrice  float32 `bson:"standard_price"`
}

type order struct {
	OrderID    int32
	CustomerID int32
}

func insert(c customer, collection *mongo.Collection) {
	insertResult, err := collection.InsertOne(context.TODO(), c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func update(c customer, collection *mongo.Collection) {

	updateResult, err := collection.UpdateOne(context.TODO(), bson.D{{Key: "customer_name", Value: "Peerawat"}}, bson.D{{Key: "$inc", Value: bson.D{{Key: "postal_code", Value: 1}}}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func searchCustomerByID(collection *mongo.Collection, customerID int32) customer {
	filter := bson.D{{Key: "customer_id", Value: customerID}}
	var result customer
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func searchCustomer(collection *mongo.Collection, size int32) []customer {
	var results []customer
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp customer
		_ = cur.Decode(&temp)
		//fmt.Println(temp)
		results = append(results, temp)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}
