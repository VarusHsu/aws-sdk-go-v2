// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Stops logging information, such as queries and connection attempts, for the
// specified Amazon Redshift cluster.
func (c *Client) DisableLogging(ctx context.Context, params *DisableLoggingInput, optFns ...func(*Options)) (*DisableLoggingOutput, error) {
	if params == nil {
		params = &DisableLoggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableLogging", params, optFns, c.addOperationDisableLoggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableLoggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableLoggingInput struct {

	// The identifier of the cluster on which logging is to be stopped.
	//
	// Example: examplecluster
	//
	// This member is required.
	ClusterIdentifier *string

	noSmithyDocumentSerde
}

// Describes the status of logging for a cluster.
type DisableLoggingOutput struct {

	// The name of the S3 bucket where the log files are stored.
	BucketName *string

	// The message indicating that logs failed to be delivered.
	LastFailureMessage *string

	// The last time when logs failed to be delivered.
	LastFailureTime *time.Time

	// The last time that logs were delivered.
	LastSuccessfulDeliveryTime *time.Time

	// The log destination type. An enum with possible values of s3 and cloudwatch .
	LogDestinationType types.LogDestinationType

	// The collection of exported log types. Possible values are connectionlog ,
	// useractivitylog , and userlog .
	LogExports []string

	// true if logging is on, false if logging is off.
	LoggingEnabled *bool

	// The prefix applied to the log file names.
	S3KeyPrefix *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableLoggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisableLogging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableLogging{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisableLogging"); err != nil {
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
	if err = addOpDisableLoggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableLogging(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableLogging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisableLogging",
	}
}
