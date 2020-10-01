// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds up to 50 cost allocation tags to the named resource. A cost allocation tag
// is a key-value pair where the key and value are case-sensitive. You can use cost
// allocation tags to categorize and track your AWS costs. When you apply tags to
// your ElastiCache resources, AWS generates a cost allocation report as a
// comma-separated value (CSV) file with your usage and costs aggregated by your
// tags. You can apply tags that represent business categories (such as cost
// centers, application names, or owners) to organize your costs across multiple
// services. For more information, see Using Cost Allocation Tags in Amazon
// ElastiCache
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Tagging.html) in
// the ElastiCache User Guide.
func (c *Client) AddTagsToResource(ctx context.Context, params *AddTagsToResourceInput, optFns ...func(*Options)) (*AddTagsToResourceOutput, error) {
	stack := middleware.NewStack("AddTagsToResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpAddTagsToResourceMiddlewares(stack)
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
	addOpAddTagsToResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddTagsToResource(options.Region), middleware.Before)
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
			OperationName: "AddTagsToResource",
			Err:           err,
		}
	}
	out := result.(*AddTagsToResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an AddTagsToResource operation.
type AddTagsToResourceInput struct {

	// A list of cost allocation tags to be added to this resource. A tag is a
	// key-value pair. A tag key must be accompanied by a tag value.
	//
	// This member is required.
	Tags []*types.Tag

	// The Amazon Resource Name (ARN) of the resource to which the tags are to be
	// added, for example arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster or
	// arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot. ElastiCache
	// resources are cluster and snapshot. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// This member is required.
	ResourceName *string
}

// Represents the output from the AddTagsToResource, ListTagsForResource, and
// RemoveTagsFromResource operations.
type AddTagsToResourceOutput struct {

	// A list of cost allocation tags as key-value pairs.
	TagList []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpAddTagsToResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpAddTagsToResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpAddTagsToResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddTagsToResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "AddTagsToResource",
	}
}
