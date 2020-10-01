// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the model packages that have been created.
func (c *Client) ListModelPackages(ctx context.Context, params *ListModelPackagesInput, optFns ...func(*Options)) (*ListModelPackagesOutput, error) {
	stack := middleware.NewStack("ListModelPackages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListModelPackagesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListModelPackages(options.Region), middleware.Before)
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
			OperationName: "ListModelPackages",
			Err:           err,
		}
	}
	out := result.(*ListModelPackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelPackagesInput struct {

	// The maximum number of model packages to return in the response.
	MaxResults *int32

	// The sort order for the results. The default is Ascending.
	SortOrder types.SortOrder

	// The parameter by which to sort the results. The default is CreationTime.
	SortBy types.ModelPackageSortBy

	// A filter that returns only model packages created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only model packages created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// If the response to a previous ListModelPackages request was truncated, the
	// response includes a NextToken. To retrieve the next set of model packages, use
	// the token in the next request.
	NextToken *string

	// A string in the model package name. This filter returns only model packages
	// whose name contains the specified string.
	NameContains *string
}

type ListModelPackagesOutput struct {

	// An array of ModelPackageSummary objects, each of which lists a model package.
	//
	// This member is required.
	ModelPackageSummaryList []*types.ModelPackageSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of model packages, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListModelPackagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListModelPackages{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModelPackages{}, middleware.After)
}

func newServiceMetadataMiddleware_opListModelPackages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListModelPackages",
	}
}
