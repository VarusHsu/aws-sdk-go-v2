// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the overall health of the specified environment. The
// DescribeEnvironmentHealth operation is only available with AWS Elastic Beanstalk
// Enhanced Health.
func (c *Client) DescribeEnvironmentHealth(ctx context.Context, params *DescribeEnvironmentHealthInput, optFns ...func(*Options)) (*DescribeEnvironmentHealthOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentHealth", params, optFns, c.addOperationDescribeEnvironmentHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// See the example below to learn how to create a request body.
type DescribeEnvironmentHealthInput struct {

	// Specify the response elements to return. To retrieve all attributes, set to All
	// . If no attribute names are specified, returns the name of the environment.
	AttributeNames []types.EnvironmentHealthAttribute

	// Specify the environment by ID.
	//
	// You must specify either this or an EnvironmentName, or both.
	EnvironmentId *string

	// Specify the environment by name.
	//
	// You must specify either this or an EnvironmentName, or both.
	EnvironmentName *string

	noSmithyDocumentSerde
}

// Health details for an AWS Elastic Beanstalk environment.
type DescribeEnvironmentHealthOutput struct {

	// Application request metrics for the environment.
	ApplicationMetrics *types.ApplicationMetrics

	// Descriptions of the data that contributed to the environment's current health
	// status.
	Causes []string

	// The [health color] of the environment.
	//
	// [health color]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced-status.html
	Color *string

	// The environment's name.
	EnvironmentName *string

	// The [health status] of the environment. For example, Ok .
	//
	// [health status]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/health-enhanced-status.html
	HealthStatus *string

	// Summary health information for the instances in the environment.
	InstancesHealth *types.InstanceHealthSummary

	// The date and time that the health information was retrieved.
	RefreshedAt *time.Time

	// The environment's operational status. Ready , Launching , Updating , Terminating
	// , or Terminated .
	Status types.EnvironmentHealth

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEnvironmentHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironmentHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironmentHealth{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEnvironmentHealth"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentHealth(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeEnvironmentHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEnvironmentHealth",
	}
}
