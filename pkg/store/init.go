package store

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

func createContainer(client *azcosmos.Client, databaseName, containerName, partitionKey string) error {

	//	databaseName = "adventureworks"
	//	containerName = "customer"
	//	partitionKey = "/customerId"

	databaseClient, err := client.NewDatabase(databaseName) // returns a struct that represents a database
	if err != nil {
		log.Fatal("Failed to create a database client:", err)
	}

	// Setting container properties
	containerProperties := azcosmos.ContainerProperties{
		ID: containerName,
		PartitionKeyDefinition: azcosmos.PartitionKeyDefinition{
			Paths: []string{partitionKey},
		},
	}

	// Setting container options
	throughputProperties := azcosmos.NewManualThroughputProperties(400) //defaults to 400 if not set
	options := &azcosmos.CreateContainerOptions{
		ThroughputProperties: &throughputProperties,
	}

	ctx := context.TODO()
	containerResponse, err := databaseClient.CreateContainer(ctx, containerProperties, options)
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("Container [%v] created. ActivityId %s\n", containerName, containerResponse.ActivityID)

	return nil
}
