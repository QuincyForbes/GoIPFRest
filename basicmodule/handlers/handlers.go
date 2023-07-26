// Package handlers contains the HTTP request handlers for the application.
package handlers

import (
	"basicmodule/models"
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

// GetAllItems handles the request to fetch all items from the DynamoDB table.
func GetAllItems(c *gin.Context) {
	// Load the AWS SDK configuration.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error loading default config: %v", err)})
		return
	}

	// Create a new DynamoDB service client.
	svc := dynamodb.NewFromConfig(cfg)

	// Prepare the Scan input for the DynamoDB table.
	input := &dynamodb.ScanInput{
		TableName: aws.String("MetadataTable"),
	}

	// Execute the Scan operation.
	resp, err := svc.Scan(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error scanning DynamoDB table: %v", err)})
		return
	}

	// Unmarshal the response items into cids.
	var cids []*models.Cid
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &cids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error unmarshalling DynamoDB items: %v", err)})
		return
	}

	// Send the result as a JSON response.
	c.JSON(http.StatusOK, cids)
}

// GetItemByCID handles the request to fetch a specific item by its CID from the DynamoDB table.
func GetItemByCID(c *gin.Context) {
	// Extract the CID parameter from the request.
	cid := c.Param("cid")

	// Load the AWS SDK configuration.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("unable to load SDK config: %v", err)})
		return
	}

	// Create a new DynamoDB service client.
	svc := dynamodb.NewFromConfig(cfg)

	// Prepare the key for the item to be fetched.
	inputKey := struct {
		Cid string `dynamodbav:"Cid"`
	}{
		Cid: cid,
	}

	key, err := attributevalue.MarshalMap(inputKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to marshal key: %v", err)})
		return
	}

	// Prepare the GetItem input for the DynamoDB table.
	input := &dynamodb.GetItemInput{
		TableName: aws.String("MetadataTable"),
		Key:       key,
	}

	// Execute the GetItem operation.
	resp, err := svc.GetItem(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get item from DynamoDB: %v", err)})
		return
	}

	// Unmarshal the response item into metadata.
	var metadata models.Metadata
	err = attributevalue.UnmarshalMap(resp.Item, &metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to unmarshal item: %v", err)})
		return
	}

	// Send the result as a JSON response.
	c.JSON(http.StatusOK, metadata)
}
