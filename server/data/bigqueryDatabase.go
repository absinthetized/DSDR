package data

import (
	"context"
	"dsdr/models"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const bqLocation = "us-central1"

// the roles repository mimiking an actual data layer (eg. a DB)
type BqDB struct {
	client *bigquery.Client
}

// implement the DB interface for the FileSystemDB struct
// you must 'defer b.client.Close()'
// folder can be empty.. I'm in the middle of a refactor and the interface sucks
func (b *BqDB) Connect(folder string) error {
	ctx := context.Background()
	var err error
	b.client, err = bigquery.NewClient(ctx, bigquery.DetectProjectID)
	if err != nil {
		log.Println("bigquery.NewClient:", err)
		return err
	}

	log.Println("BQ Client created!")
	return nil
}

func (b *BqDB) Close() {
	b.client.Close()
}

// this must be removed
func (b *BqDB) Roles() []models.BasicIAMRole {
	var patch_for_refactor []models.BasicIAMRole
	return patch_for_refactor
}

func (b *BqDB) Client() *bigquery.Client {
	return b.client
}

// Query returns an array of bigquery values and an error
func Query[T any](b *BqDB, queryString string) ([]T, error) {
	q := b.client.Query(queryString)
	q.Location = bqLocation

	ctx := context.Background()

	// Run the query and print results when the query job is completed.
	job, err := q.Run(ctx)
	if err != nil {
		log.Println("DB.Query error:", err)
		return nil, err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		log.Println("DB.Query error:", err)
		return nil, err
	}
	if err := status.Err(); err != nil {
		log.Println("DB.Query error:", err)
		return nil, err
	}

	it, err := job.Read(ctx)

	if err != nil {
		log.Println("DB.Query error:", err)
		return nil, err
	}

	if err := status.Err(); err != nil {
		log.Println("DB.Query error:", err)
		return nil, err
	}

	var rows []T
	for {
		var row T

		err := it.Next(&row)

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println("DB.Query error:", err)
			return rows, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}
