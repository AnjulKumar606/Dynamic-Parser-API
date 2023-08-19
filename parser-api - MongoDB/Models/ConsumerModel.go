package Models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Consumer struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	ApplicationName string             `json:"application_name" bson:"application_name"`
	Title           string             `json:"title" bson:"title"`
	Url             string             `json:"url" bson:"url"`
	Body            string             `json:"body" bson:"body"`
	Image           string             `json:"image" bson:"image"`
	Extra           string             `json:"extra" bson:"extra"`
	CreatedAt       time.Time        `bson:"createdAt,omitempty"`
	UpdatedAt       time.Time          `bson:"updatedAt,omitempty"`
}

func (b *Consumer) CollectionName() string {
	return "consumer" // The name of the collection in MongoDB
}

//dummy JSON

// "consumer":[
// {
// 	"items": [
// 	  {
// 		"riskScore": "High",
// 		"events": [
// 		  {
// 			"eventSource": "Firewall",
// 			"destination": {
// 			  "device": {
// 				"dnsDomain": "example.com",
// 				"port": "443",
// 				"dnsHostname": "server.example.com",
// 				"macAddress": "00:1A:2B:3C:4D:5E",
// 				"ipAddress": "192.168.1.100"
// 			  },
// 			  "user": {
// 				"adUsername": "johndoe",
// 				"emailAddress": "johndoe@example.com",
// 				"username": "johndoe",
// 				"adDomain": "example.com"
// 			  }
// 			},
// 			"source": {
// 			  "device": {
// 				"dnsDomain": "",
// 				"port": "",
// 				"dnsHostname": "192.168.1.200",
// 				"macAddress": "00:1A:2B:3C:4D:5F",
// 				"ipAddress": "192.168.1.200"
// 			  },
// 			  "user": {
// 				"adUsername": "",
// 				"emailAddress": "",
// 				"username": "",
// 				"adDomain": ""
// 			  }
// 			},
// 			"domain": "example.com",
// 			"eventSourceId": "1234567890"
// 		  }
// 		],
// 		"id": "12345",
// 		"detail": "Unauthorized access attempt",
// 		"title": "Security Alert",
// 		"source": "Security System",
// 		"type": "Intrusion",
// 		"created": "2023-07-25T12:34:56Z"
// 	  }
// 	],
// 	"pageNumber": 1,
// 	"totalPages": 1,
// 	"hasPrevious": false,
// 	"totalItems": 1,
// 	"hasNext": false,
// 	"pageSize": 10
//   }

// ]
