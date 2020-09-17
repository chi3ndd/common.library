package elastic

import (
	"context"

	es7 "github.com/olivere/elastic/v7"
)

func (con *ConnectorV7) Get(indexName string, docId string) (*es7.GetResult, error) {
	service := con.client.Get().
		Index(indexName).
		Id(docId)
	// Success
	return service.Do(context.TODO())
}

func (con *ConnectorV7) Insert(indexName string, docId string, body interface{}) (*es7.IndexResponse, error) {
	service := con.client.Index().Index(indexName).BodyJson(body)
	if docId != "" {
		service = service.Id(docId)
	}
	// Success
	return service.Do(context.TODO())
}

func (con *ConnectorV7) Update(indexName string, docId string, update interface{}) (*es7.UpdateResponse, error) {
	service := con.client.Update().Index(indexName).Id(docId).Doc(update)
	// Success
	return service.Do(context.TODO())
}

func (con *ConnectorV7) Search(indexName string, query es7.Query, offset int, size int, sort *es7.SortInfo) (*es7.SearchResult, error) {
	service := con.client.Search().
		Index(indexName).
		Query(query).
		From(offset).
		Size(size)
	if sort != nil {
		service.SortBy(sort)
	}
	// Success
	return service.Do(context.TODO())
}

func (con *ConnectorV7) SearchOne(indexName string, query es7.Query, offset int, sort *es7.SortInfo) (*es7.SearchResult, error) {
	// Success
	return con.Search(indexName, query, offset, 1, sort)
}

func (con *ConnectorV7) SearchScroll(indexName string, query es7.Query, size int) *es7.ScrollService {
	scroll := con.client.Scroll().
		Index(indexName).
		Query(query).
		Size(size)
	// Success
	return scroll
}
