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

func GetAllItems(c *gin.Context) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error loading default config: %v", err)})
		return
	}

	svc := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.ScanInput{
		TableName: aws.String("MetadataTable"),
	}

	resp, err := svc.Scan(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error scanning DynamoDB table: %v", err)})
		return
	}

	var tokens []*models.Cid
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &tokens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error unmarshalling DynamoDB items: %v", err)})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func GetItemByCID(c *gin.Context) {
	cid := c.Param("cid")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("unable to load SDK config: %v", err)})
		return
	}

	svc := dynamodb.NewFromConfig(cfg)

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

	input := &dynamodb.GetItemInput{
		TableName: aws.String("MetadataTable"),
		Key:       key,
	}

	resp, err := svc.GetItem(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get item from DynamoDB: %v", err)})
		return
	}

	var metadata models.Metadata
	err = attributevalue.UnmarshalMap(resp.Item, &metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to unmarshal item: %v", err)})
		return
	}

	c.JSON(http.StatusOK, metadata)
}
