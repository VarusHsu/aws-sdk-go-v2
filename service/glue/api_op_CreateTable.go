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

// Creates a new table definition in the Data Catalog.
func (c *Client) CreateTable(ctx context.Context, params *CreateTableInput, optFns ...func(*Options)) (*CreateTableOutput, error) {
	stack := middleware.NewStack("CreateTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTableMiddlewares(stack)
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
	addOpCreateTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTable(options.Region), middleware.Before)
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
			OperationName: "CreateTable",
			Err:           err,
		}
	}
	out := result.(*CreateTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTableInput struct {

	// The ID of the Data Catalog in which to create the Table. If none is supplied,
	// the AWS account ID is used by default.
	CatalogId *string

	// The catalog database in which to create the new table. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// The TableInput object that defines the metadata table to create in the catalog.
	//
	// This member is required.
	TableInput *types.TableInput
}

type CreateTableOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateTable",
	}
}
