package goshopify

import (
	"errors"
	"fmt"
)

type OrderGraphqlService interface {
	GetLocalizationExtensions(int64) (*LocalizationExtensions, error)
}

type LocalizationExtensions struct {
	GraphqlErrors
	Data struct {
		Order struct {
			Id                     string `json:"id"`
			LocalizationExtensions struct {
				Edges []struct {
					Node struct {
						CountryCode string `json:"countryCode"`
						Purpose     string `json:"purpose"`
						Title       string `json:"title"`
						Value       string `json:"value"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"localizationExtensions"`
		} `json:"order"`
	} `json:"data"`
}

type OrderGraphqlServiceOp struct {
	client *Client
}

func (g *OrderGraphqlServiceOp) GetLocalizationExtensions(orderID int64) (*LocalizationExtensions, error) {
	query := fmt.Sprintf(`query {
      order(id: "gid://shopify/Order/%d") {
        id
        localizationExtensions(first: 100) {
          edges {
            node {
              countryCode
              purpose
              title
              value
            }
          }
        }
      }
    }`, orderID)
	resource := new(LocalizationExtensions)
	if err := g.client.PostGraphql(&GraphqlRequest{
		Query: query,
	}, resource); err != nil {
		return nil, err
	}
	if len(resource.Errors) > 0 {
		return nil, errors.New(resource.Errors[0].Message)
	}

	return resource, nil
}
