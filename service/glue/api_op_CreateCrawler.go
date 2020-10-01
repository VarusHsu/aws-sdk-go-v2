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

// Creates a new crawler with specified targets, role, configuration, and optional
// schedule. At least one crawl target must be specified, in the s3Targets field,
// the jdbcTargets field, or the DynamoDBTargets field.
func (c *Client) CreateCrawler(ctx context.Context, params *CreateCrawlerInput, optFns ...func(*Options)) (*CreateCrawlerOutput, error) {
	stack := middleware.NewStack("CreateCrawler", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateCrawlerMiddlewares(stack)
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
	addOpCreateCrawlerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCrawler(options.Region), middleware.Before)
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
			OperationName: "CreateCrawler",
			Err:           err,
		}
	}
	out := result.(*CreateCrawlerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCrawlerInput struct {

	// The IAM role or Amazon Resource Name (ARN) of an IAM role used by the new
	// crawler to access customer resources.
	//
	// This member is required.
	Role *string

	// The tags to use with this crawler request. You may use tags to limit access to
	// the crawler. For more information about tags in AWS Glue, see AWS Tags in AWS
	// Glue (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html) in the
	// developer guide.
	Tags map[string]*string

	// A list of custom classifiers that the user has registered. By default, all
	// built-in classifiers are included in a crawl, but these custom classifiers
	// always override the default classifiers for a given classification.
	Classifiers []*string

	// The table prefix used for catalog tables that are created.
	TablePrefix *string

	// The name of the SecurityConfiguration structure to be used by this crawler.
	CrawlerSecurityConfiguration *string

	// The AWS Glue database where results are written, such as:
	// arn:aws:daylight:us-east-1::database/sometable/*.
	DatabaseName *string

	// A cron expression used to specify the schedule (see Time-Based Schedules for
	// Jobs and Crawlers
	// (https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html).
	// For example, to run something every day at 12:15 UTC, you would specify: cron(15
	// 12 * * ? *).
	Schedule *string

	// Name of the new crawler.
	//
	// This member is required.
	Name *string

	// A description of the new crawler.
	Description *string

	// Crawler configuration information. This versioned JSON string allows users to
	// specify aspects of a crawler's behavior. For more information, see Configuring a
	// Crawler (https://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html).
	Configuration *string

	// A list of collection of targets to crawl.
	//
	// This member is required.
	Targets *types.CrawlerTargets

	// The policy for the crawler's update and deletion behavior.
	SchemaChangePolicy *types.SchemaChangePolicy
}

type CreateCrawlerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateCrawlerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCrawler{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCrawler{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCrawler(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateCrawler",
	}
}
