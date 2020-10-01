// Code generated by smithy-go-codegen DO NOT EDIT.

package macie

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the S3 resources associated with Amazon Macie Classic. If
// memberAccountId isn't specified, the action lists the S3 resources associated
// with Amazon Macie Classic for the current master account. If memberAccountId is
// specified, the action lists the S3 resources associated with Amazon Macie
// Classic for the specified member account.
func (c *Client) ListS3Resources(ctx context.Context, params *ListS3ResourcesInput, optFns ...func(*Options)) (*ListS3ResourcesOutput, error) {
	stack := middleware.NewStack("ListS3Resources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListS3ResourcesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListS3Resources(options.Region), middleware.Before)
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
			OperationName: "ListS3Resources",
			Err:           err,
		}
	}
	out := result.(*ListS3ResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListS3ResourcesInput struct {

	// The Amazon Macie Classic member account ID whose associated S3 resources you
	// want to list.
	MemberAccountId *string

	// Use this parameter when paginating results. Set its value to null on your first
	// call to the ListS3Resources action. Subsequent calls to the action fill
	// nextToken in the request with the value of nextToken from the previous response
	// to continue listing data.
	NextToken *string

	// Use this parameter to indicate the maximum number of items that you want in the
	// response. The default value is 250.
	MaxResults *int32
}

type ListS3ResourcesOutput struct {

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to be
	// listed, this parameter is set to null.
	NextToken *string

	// A list of the associated S3 resources returned by the action.
	S3Resources []*types.S3ResourceClassification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListS3ResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListS3Resources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListS3Resources{}, middleware.After)
}

func newServiceMetadataMiddleware_opListS3Resources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie",
		OperationName: "ListS3Resources",
	}
}
