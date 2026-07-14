package esc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"

	"chatskn/app/message/model"
)

const (
	_indexName = "chatskn-message"
)

type MessageSearcher interface {
	Index(ctx context.Context, doc interface{}) error
	Search(ctx context.Context, msg model.Message) error
}

type EsClient struct {
	c *elasticsearch.TypedClient
}

func NewEsClient(es *elasticsearch.TypedClient) *EsClient {
	c := &EsClient{es}
	if err := c.CreateIndexIfNotExist(context.Background()); err != nil {
		panic(err)
	}
	return c
}

func (es *EsClient) CreateIndexIfNotExist(ctx context.Context) error {
	ok, err := es.c.Indices.Exists(_indexName).Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", ok)

	res, err := es.c.Indices.Create(_indexName).
		Request(&create.Request{
			Mappings: &types.TypeMapping{
				Properties: map[string]types.Property{
					"Id":         types.NewKeywordProperty(),
					"SenderId":   types.NewKeywordProperty(),
					"ReceiverId": types.NewKeywordProperty(),
					"ChannelId":  types.NewKeywordProperty(),
					"Type":       types.NewKeywordProperty(),
					"Content":    types.NewTextProperty(),
					"CreateAt":   types.NewDateProperty(),
				},
			},
		}).
		Do(ctx)
	fmt.Printf("%v\n", res)
	return err
}

func (es *EsClient) Index(ctx context.Context, doc interface{}) error {
	res, err := es.c.Index(_indexName).
		Request(doc).Do(ctx)
	fmt.Printf("%v\n", res)
	return err
}

func (es *EsClient) Search(ctx context.Context, msg model.Message) error {
	match := make(map[string]types.MatchQuery)
	if msg.Id != 0 {
		match["Id"] = types.MatchQuery{Query: strconv.FormatInt(msg.Id, 10)}
	}
	if msg.SenderId != 0 {
		match["SenderId"] = types.MatchQuery{Query: strconv.FormatInt(msg.SenderId, 10)}
	}
	if msg.ReceiverId != 0 {
		match["ReceiverId"] = types.MatchQuery{Query: strconv.FormatInt(msg.ReceiverId, 10)}
	}
	if msg.ChannelId != 0 {
		match["ChannelId"] = types.MatchQuery{Query: strconv.FormatInt(msg.ChannelId, 10)}
	}
	if msg.Content != "" {
		match["Content"] = types.MatchQuery{Query: msg.Content}
	}

	res, err := es.c.Search().
		Index(_indexName).
		Request(&search.Request{
			Query: &types.Query{
				Match: match,
			},
		}).Do(ctx)
	fmt.Printf("%v\n", res)
	return err
}
