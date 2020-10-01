// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the versions of the themes in the current AWS account.
func (c *Client) ListThemeVersions(ctx context.Context, params *ListThemeVersionsInput, optFns ...func(*Options)) (*ListThemeVersionsOutput, error) {
	stack := middleware.NewStack("ListThemeVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListThemeVersionsMiddlewares(stack)
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
	addOpListThemeVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThemeVersions(options.Region), middleware.Before)
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
			OperationName: "ListThemeVersions",
			Err:           err,
		}
	}
	out := result.(*ListThemeVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThemeVersionsInput struct {

	// The ID for the theme.
	//
	// This member is required.
	ThemeId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The ID of the AWS account that contains the themes that you're listing.
	//
	// This member is required.
	AwsAccountId *string

	// The token for the next set of results, or null if there are no more results.
	NextToken *string
}

type ListThemeVersionsOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// A structure containing a list of all the versions of the specified theme.
	ThemeVersionSummaryList []*types.ThemeVersionSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListThemeVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListThemeVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListThemeVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListThemeVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListThemeVersions",
	}
}
