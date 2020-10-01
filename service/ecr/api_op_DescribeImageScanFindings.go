// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the scan findings for the specified image.
func (c *Client) DescribeImageScanFindings(ctx context.Context, params *DescribeImageScanFindingsInput, optFns ...func(*Options)) (*DescribeImageScanFindingsOutput, error) {
	stack := middleware.NewStack("DescribeImageScanFindings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeImageScanFindingsMiddlewares(stack)
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
	addOpDescribeImageScanFindingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageScanFindings(options.Region), middleware.Before)
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
			OperationName: "DescribeImageScanFindings",
			Err:           err,
		}
	}
	out := result.(*DescribeImageScanFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImageScanFindingsInput struct {

	// The maximum number of image scan results returned by DescribeImageScanFindings
	// in paginated output. When this parameter is used, DescribeImageScanFindings only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another DescribeImageScanFindings request with the returned nextToken value.
	// This value can be between 1 and 1000. If this parameter is not used, then
	// DescribeImageScanFindings returns up to 100 results and a nextToken value, if
	// applicable.
	MaxResults *int32

	// The repository for the image for which to describe the scan findings.
	//
	// This member is required.
	RepositoryName *string

	// An object with identifying information for an Amazon ECR image.
	//
	// This member is required.
	ImageId *types.ImageIdentifier

	// The AWS account ID associated with the registry that contains the repository in
	// which to describe the image scan findings for. If you do not specify a registry,
	// the default registry is assumed.
	RegistryId *string

	// The nextToken value returned from a previous paginated DescribeImageScanFindings
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string
}

type DescribeImageScanFindingsOutput struct {

	// The registry ID associated with the request.
	RegistryId *string

	// The current state of the scan.
	ImageScanStatus *types.ImageScanStatus

	// The repository name associated with the request.
	RepositoryName *string

	// An object with identifying information for an Amazon ECR image.
	ImageId *types.ImageIdentifier

	// The information contained in the image scan findings.
	ImageScanFindings *types.ImageScanFindings

	// The nextToken value to include in a future DescribeImageScanFindings request.
	// When the results of a DescribeImageScanFindings request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeImageScanFindingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImageScanFindings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImageScanFindings{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeImageScanFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "DescribeImageScanFindings",
	}
}
