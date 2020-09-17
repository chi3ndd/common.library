package elastic

import (
	es1 "gopkg.in/olivere/elastic.v2"
)

func (con *ConnectorV1) Get(indexName string, docType string, docId string) (*es1.GetResult, error) {
	service := con.client.Get().
		Index(indexName).
		Type(docType).
		Id(docId)
	// Success
	return service.Do()
}

func (con *ConnectorV1) Insert(indexName string, docType string, docId string, body interface{}) (*es1.IndexResult, error) {
	service := con.client.Index().Index(indexName).Type(docType).BodyJson(body)
	if docId != "" {
		service = service.Id(docId)
	}
	// Success
	return service.Do()
}

func (con *ConnectorV1) Update(indexName string, docType string, docId string, update interface{}) (*es1.UpdateResult, error) {
	service := con.client.Update().Index(indexName).Type(docType).Id(docId).Doc(update)
	// Success
	return service.Do()
}

func (con *ConnectorV1) Search(indexName string, docType string, query es1.Query, offset int, size int, sort *es1.SortInfo) (*es1.SearchResult, error) {
	service := con.client.Search().
		Index(indexName).
		Type(docType).
		Query(query).
		From(offset).
		Size(size)
	if sort != nil {
		service.SortBy(sort)
	}
	// Success
	return service.Do()
}

func (con *ConnectorV1) SearchOne(indexName string, docType string, query es1.Query, offset int, sort *es1.SortInfo) (*es1.SearchResult, error) {
	// Success
	return con.Search(indexName, docType, query, offset, 1, sort)
}

func (con *ConnectorV1) SearchScroll(indexName string, docType string, query es1.Query, size int) *es1.ScrollService {
	scroll := con.client.Scroll().
		Index(indexName).
		Type(docType).
		Query(query).
		Size(size)
	// Success
	return scroll
}
