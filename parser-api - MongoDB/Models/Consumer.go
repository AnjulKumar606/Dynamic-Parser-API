package Models

import (
	"appdirs/cns-parser/Config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// enter your collection name in MongoDB
var collection_name = "consumer"

// GetAllConsumer Fetch all consumer data
func GetAllConsumers(consumer *[]Consumer) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Config.DB.Collection(collection_name)

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result Consumer
		if err := cur.Decode(&result); err != nil {
			log.Fatal("Failed to decode data:", err)
		}
		*consumer = append(*consumer, result)
	}

	if err := cur.Err(); err != nil {
		return err
	}

	return nil
}

// CreateConsumer ... Insert New data
func CreateConsumer(consumer *Consumer) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Config.DB.Collection(collection_name)
	//adding timestamp
	consumer.CreatedAt = time.Now()
	consumer.UpdatedAt = time.Now()
	_, err = collection.InsertOne(ctx, consumer)
	if err != nil {
		return err
	}

	return nil
}

// GetConsumerByID ... Fetch only one consumer by Id
func GetConsumerByID(consumer *Consumer, id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Config.DB.Collection(collection_name)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	if err := collection.FindOne(ctx, filter).Decode(consumer); err != nil {
		return err
	}

	return nil
}

// UpdateConsumer ... Update consumer
func UpdateConsumer(consumer *Consumer, id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Config.DB.Collection(collection_name)
	//adding timestamp
	consumer.UpdatedAt = time.Now()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": consumer}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteConsumer ... Delete consumer
func DeleteConsumer(consumer *Consumer, id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Config.DB.Collection(collection_name)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
