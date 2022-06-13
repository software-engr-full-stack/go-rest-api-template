package dynamodb

import (
    "context"

    "github.com/pkg/errors"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Client struct {
    Client *dynamodb.Client
}

func NewClient(useActualDynamoDb bool) (*Client, error) {
    if useActualDynamoDb {
        panic("TODO")
    }

    return newClientLocal()
}

func (c *Client) ListTables() (*dynamodb.ListTablesOutput, error) {
    tables, err := c.Client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
    var empty *dynamodb.ListTablesOutput
    if err != nil {
        return empty, errors.WithStack(err)
    }
    return tables, nil
}

func newClientLocal() (*Client, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion("us-west-1"),
        config.WithEndpointResolver(aws.EndpointResolverFunc(
            func(service, region string) (aws.Endpoint, error) {
                return aws.Endpoint{URL: "http://localhost:8000"}, nil
            })),
        config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
            Value: aws.Credentials{
                AccessKeyID: "irrelevant", SecretAccessKey: "irrelevant", SessionToken: "irrelevant",
                Source: "Hard-coded credentials, values are irrelevant for local DynamoDB",
            },
        }),
    )
    var empty *Client
    if err != nil {
        return empty, errors.WithStack(err)
    }

    return &Client{Client: dynamodb.NewFromConfig(cfg)}, nil
}
