package goshopify

const (
	graphqlBasePath = "graphql.json"
)

type GraphqlErrors struct {
	Errors []struct {
		Message string `json:"message,omitempty"`
	} `json:"errors,omitempty"`
}
