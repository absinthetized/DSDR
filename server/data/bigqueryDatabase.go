package data

import (
	"context"
	"dsdr/models"
	"log"

	"cloud.google.com/go/bigquery"
)

const bqLocation = "us-central1"

// the roles repository mimiking an actual data layer (eg. a DB)
type BqDB struct {
	roles  []models.BasicIAMRole
	client *bigquery.Client
}

// implemente the DB interface for the FileSystemDB struct
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

func (b *BqDB) Roles() []models.BasicIAMRole {
	return b.roles
}

func (b *BqDB) Client() *bigquery.Client {
	return b.client
}

// Query returns either an iterable bigquery.Job pointer or nil + the-error-cause
func (b *BqDB) Query(queryString string) (*bigquery.Job, error) {
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

	return job, nil
}
