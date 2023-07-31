package store

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

const (
	cosmosDbEndpoint = "https://dec92f4d-0ee0-4-231-b9ee.documents.azure.com:443/"
	cosmosDbKey      = ""
	databaseName     = "ToDoList"
	partitionKey     = "test"
	containerName    = "tests"
)

type SignalConfig struct {
	shortAv    int
	longAv     int
	daysback   int
	tradeCount int
	Gain       float64
}

func ini() {

	item := struct {
		ID           string `json:"id"`
		CustomerId   string `json:"customerId"`
		Title        string
		FirstName    string
		LastName     string
		EmailAddress string
		PhoneNumber  string
		CreationDate string
		Testspkey    string `json:"testspkey"`
		Test         SignalConfig
	}{
		ID:           "eee2e2",
		CustomerId:   "1",
		Title:        "Mr",
		FirstName:    "Luke",
		LastName:     "Hayes",
		EmailAddress: "luke12@adventure-works.com",
		PhoneNumber:  "879-555-0197",
		Testspkey:    "test",
		Test:         SignalConfig{Gain: 1.2},
	}

	ctx := context.TODO()
	/* 	item := map[string]string{
		"id":     "1",
		"value":  "2",
		"user":   "test",
		"pw":     "test",
		"drink":  "snapp",
		"editor": "vi",
	} */
	/* marshalled, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	} */
	cred, err := azcosmos.NewKeyCredential(cosmosDbKey)
	if err != nil {
		log.Fatal("Failed to create a credential: ", err)
	}

	// Create a CosmosDB client
	client, err := azcosmos.NewClientWithKey(cosmosDbEndpoint, cred, nil)
	if err != nil {
		log.Fatal("Failed to create Azure Cosmos DB client: ", err)
	}

	// Create database client
	_, err = client.NewDatabase(databaseName)
	if err != nil {
		log.Fatal("Failed to create database client:", err)
	}

	containerClient, err := client.NewContainer(databaseName, containerName)
	if err != nil {
		log.Fatal("Failed to create a container client:", err)
	}

	// Setting container properties
	/* containerProperties := azcosmos.ContainerProperties{
		ID: containerName,
		PartitionKeyDefinition: azcosmos.PartitionKeyDefinition{
			Paths: []string{partitionKey},
		},
	} */
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	itemResponse, err := containerClient.ReadItem(ctx, pk, "replace_with_new_document_id", nil)
	fmt.Println(itemResponse)
	fmt.Println("oooooooooooooo")
	fmt.Println(err)
	fmt.Println(itemResponse)
	fmt.Println("itemResponse")
	err = CreateItem(client, databaseName, containerName, partitionKey, item)
	log.Println(err)
	err = replaceItem(client, databaseName, containerName, partitionKey, "replace_with_new_document_id", item)
	log.Println(err)
	err = deleteItem(client, databaseName, containerName, partitionKey, "replace_with_new_document_id")
	log.Println(err)
}

func CreateItem(client *azcosmos.Client, databaseName, containerName, partitionKey string, item any) error {
	fmt.Println("--!--")
	containerClient, err := client.NewContainer(databaseName, containerName)
	if err != nil {
		return fmt.Errorf("failed to create a container client: %s", err)
	}

	// Specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	// setting item options upon creating ie. consistency level
	itemOptions := azcosmos.ItemOptions{
		ConsistencyLevel: azcosmos.ConsistencyLevelSession.ToPtr(),
	}
	ctx := context.TODO()
	itemResponse, err := containerClient.CreateItem(ctx, pk, b, &itemOptions)
	fmt.Println("---")
	if err != nil {
		return err
	}
	log.Printf("Status %d. Item %v created. ActivityId %s. Consuming %v Request Units.\n", itemResponse.RawResponse.StatusCode, pk, itemResponse.ActivityID, itemResponse.RequestCharge)
	fmt.Println("#####")
	return nil
}

func replaceItem(client *azcosmos.Client, databaseName, containerName, partitionKey, oldItem string, item any) error {
	fmt.Println("---")
	containerClient, err := client.NewContainer(databaseName, containerName)
	if err != nil {
		return fmt.Errorf("failed to create a container client: %s", err)
	}

	// Specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	// setting item options upon creating ie. consistency level
	itemOptions := azcosmos.ItemOptions{
		ConsistencyLevel: azcosmos.ConsistencyLevelSession.ToPtr(),
	}
	ctx := context.TODO()
	itemResponse, err := containerClient.ReplaceItem(ctx, pk, oldItem, b, &itemOptions)
	fmt.Println("---")
	if err != nil {
		return err
	}
	log.Printf("Status %d. Item %v created. ActivityId %s. Consuming %v Request Units.\n", itemResponse.RawResponse.StatusCode, pk, itemResponse.ActivityID, itemResponse.RequestCharge)
	fmt.Println("#####")
	return nil
}

func readItem(client *azcosmos.Client, databaseName, containerName, partitionKey, itemId string) error {
	//	databaseName = "adventureworks"
	//	containerName = "customer"
	//	partitionKey = "1"
	//	itemId = "1"

	// Create container client
	containerClient, err := client.NewContainer(databaseName, containerName)
	if err != nil {
		return fmt.Errorf("Failed to create a container client: %s", err)
	}

	// Specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	// Read an item
	ctx := context.TODO()
	itemResponse, err := containerClient.ReadItem(ctx, pk, itemId, nil)
	if err != nil {
		return err
	}

	itemResponseBody := struct {
		ID           string `json:"id"`
		CustomerId   string `json:"customerId"`
		Title        string
		FirstName    string
		LastName     string
		EmailAddress string
		PhoneNumber  string
		CreationDate string
	}{}

	err = json.Unmarshal(itemResponse.Value, &itemResponseBody)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(itemResponseBody, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("Read item with customerId %s\n", itemResponseBody.CustomerId)
	fmt.Printf("%s\n", b)

	log.Printf("Status %d. Item %v read. ActivityId %s. Consuming %v Request Units.\n", itemResponse.RawResponse.StatusCode, pk, itemResponse.ActivityID, itemResponse.RequestCharge)

	return nil
}

func deleteItem(client *azcosmos.Client, databaseName, containerName, partitionKey, itemId string) error {
	//	databaseName = "adventureworks"
	//	containerName = "customer"
	//	partitionKey = "1"
	//	itemId = "1"

	// Create container client
	containerClient, err := client.NewContainer(databaseName, containerName)
	if err != nil {
		return fmt.Errorf("Failed to create a container client: %s", err)
	}
	// Specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(partitionKey)

	// Delete an item
	ctx := context.TODO()
	res, err := containerClient.DeleteItem(ctx, pk, itemId, nil)
	if err != nil {
		return err
	}

	log.Printf("Status %d. Item %v deleted. ActivityId %s. Consuming %v Request Units.\n", res.RawResponse.StatusCode, pk, res.ActivityID, res.RequestCharge)

	return nil
}
