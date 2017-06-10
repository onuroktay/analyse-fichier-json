package OnurTPIjsonReader

import (
	"errors"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v5"
	"golang.org/x/net/context"
)

type ELASTICSEARCH struct {
	client *elastic.Client
	_indexName string
	_type string
}

// NewElasticSearch open a connection to the ElasticSearch database
func NewElasticSearch(indexName, typeName string) (es *ELASTICSEARCH, err error) {
	if indexName == "" {
		err = errors.New("Index Name is mandatory")
		return
	}

	if typeName == "" {
		err = errors.New("Type is mandatory")
		return
	}

	es = new(ELASTICSEARCH)

	es._indexName = indexName
	es._type = typeName
	es.client, err = elastic.NewClient()
	if err != nil {
		return
	}

	err = es.createIndexIfNotExist()

	return
}

func (es *ELASTICSEARCH) createIndexIfNotExist() error {
	exists, err := es.client.IndexExists(es._indexName).Do(context.TODO())
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	createIndex, err := es.client.CreateIndex(es._indexName).Do(context.TODO())
	if err != nil {
		return err
	}

	if !createIndex.Acknowledged {
		err = errors.New(fmt.Sprintf("expected IndicesCreateResult.Acknowledged %v; got %v", true, createIndex.Acknowledged))
	}

	return err
}

// SaveItem save an item in ElasticSearch
func (es *ELASTICSEARCH) SaveItem(item *ITEM) (err error) {
	_, err = es.client.Index().
		Index(es._indexName).
		Type(es._type).
		Id(item.ID).
		BodyJson(item).
		Refresh("false").
		Do(context.TODO())

	return
}
