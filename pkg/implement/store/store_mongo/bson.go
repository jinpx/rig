package store_mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	// MongoDB 连接 URI
	mongoURI = "mongodb://admin:secret@localhost:27017"
)

func MongoSave() {

	// 设置连接选项
	clientOptions := options.Client().ApplyURI(mongoURI)

	// 创建 MongoDB 客户端
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB 连接失败: %v", err)
	}

	// 检查连接状态
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("无法连接到 MongoDB: %v", err)
	}
	fmt.Println("成功连接到 MongoDB")

	// 选择数据库和集合
	db := client.Database("testdb")
	collection := db.Collection("users")

	// 插入数据
	user := bson.M{"name": "Kim", "age": 15}
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatalf("插入数据失败: %v", err)
	}
	fmt.Println("插入成功，ID:", insertResult.InsertedID)

	// 查询数据
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.M{"name": "Alice"}).Decode(&result)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}
	fmt.Println("查询结果:", result)

	// 关闭连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("关闭 MongoDB 连接失败: %v", err)
	}
	fmt.Println("MongoDB 连接已关闭")
}

func MongoRead() {
	// 设置 MongoDB 连接 URI
	clientOptions := options.Client().ApplyURI(mongoURI)

	// 连接 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = client.Disconnect(context.TODO())
	}()

	// 选择数据库和集合
	collection := client.Database("testdb").Collection("users")

	// 读取数据
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.M{"name": "Kim"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("查询结果:", result)
}
