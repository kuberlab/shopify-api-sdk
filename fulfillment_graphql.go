package goshopify

import (
	"errors"
	"github.com/dsers/shopify-api-sdk/v3/graphql"
)

type FulfillmentGraphqlService interface {
	FulfillmentCreateV2(request *graphql.FulfillmentCreateV2Request) (*graphql.FulfillmentResource, error)
}

type FulfillmentGraphqlServiceOp struct {
	client *Client
}

type FulfillmentCreateV2Reply struct {
	GraphqlErrors
	Data struct {
		FulfillmentCreateV2 struct {
			Fulfillment graphql.FulfillmentResource `json:"fulfillment"`
			UserErrors  []any                       `json:"userErrors"`
		} `json:"fulfillmentCreateV2"`
	} `json:"data"`
}

func (g *FulfillmentGraphqlServiceOp) FulfillmentCreateV2(data *graphql.FulfillmentCreateV2Request) (*graphql.FulfillmentResource, error) {
	query := `
mutation fulfillmentCreateV2($fulfillment: FulfillmentV2Input!) {
  fulfillmentCreateV2(fulfillment: $fulfillment) {
    fulfillment {
      id
      status
    }
    userErrors {
      field
      message
    }
  }
}`
	resource := new(FulfillmentCreateV2Reply)
	if err := g.client.PostGraphql(&GraphqlRequest{
		Query:     query,
		Variables: data,
	}, resource); err != nil {
		return nil, err
	}
	if len(resource.Errors) > 0 {
		return nil, errors.New(resource.Errors[0].Message)
	}

	return &resource.Data.FulfillmentCreateV2.Fulfillment, nil
}
