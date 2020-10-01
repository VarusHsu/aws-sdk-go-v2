// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of connection definitions from the Data Catalog.
func (c *Client) GetConnections(ctx context.Context, params *GetConnectionsInput, optFns ...func(*Options)) (*GetConnectionsOutput, error) {
	stack := middleware.NewStack("GetConnections", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetConnectionsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetConnections(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetConnections",
			Err:           err,
		}
	}
	out := result.(*GetConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConnectionsInput struct {

	// Allows you to retrieve the connection metadata without returning the password.
	// For instance, the AWS Glue console uses this flag to retrieve the connection,
	// and does not display the password. Set this parameter when the caller might not
	// have permission to use the AWS KMS key to decrypt the password, but it does have
	// permission to access the rest of the connection properties.
	HidePassword *bool

	// A continuation token, if this is a continuation call.
	NextToken *string

	// The maximum number of connections to return in one response.
	MaxResults *int32

	// A filter that controls which connections are returned.
	Filter *types.GetConnectionsFilter

	// The ID of the Data Catalog in which the connections reside. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string
}

type GetConnectionsOutput struct {

	// A continuation token, if the list of connections returned does not include the
	// last of the filtered connections.
	NextToken *string

	// A list of requested connection definitions.
	ConnectionList []*types.Connection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetConnectionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetConnections{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetConnections{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetConnections",
	}
}
