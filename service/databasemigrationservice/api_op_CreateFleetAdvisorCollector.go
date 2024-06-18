// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Fleet Advisor collector using the specified parameters.
func (c *Client) CreateFleetAdvisorCollector(ctx context.Context, params *CreateFleetAdvisorCollectorInput, optFns ...func(*Options)) (*CreateFleetAdvisorCollectorOutput, error) {
	if params == nil {
		params = &CreateFleetAdvisorCollectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleetAdvisorCollector", params, optFns, c.addOperationCreateFleetAdvisorCollectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetAdvisorCollectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetAdvisorCollectorInput struct {

	// The name of your Fleet Advisor collector (for example, sample-collector ).
	//
	// This member is required.
	CollectorName *string

	// The Amazon S3 bucket that the Fleet Advisor collector uses to store inventory
	// metadata.
	//
	// This member is required.
	S3BucketName *string

	// The IAM role that grants permissions to access the specified Amazon S3 bucket.
	//
	// This member is required.
	ServiceAccessRoleArn *string

	// A summary description of your Fleet Advisor collector.
	Description *string

	noSmithyDocumentSerde
}

type CreateFleetAdvisorCollectorOutput struct {

	// The name of the new Fleet Advisor collector.
	CollectorName *string

	// The unique ID of the new Fleet Advisor collector, for example:
	// 22fda70c-40d5-4acf-b233-a495bd8eb7f5
	CollectorReferencedId *string

	// A summary description of the Fleet Advisor collector.
	Description *string

	// The Amazon S3 bucket that the collector uses to store inventory metadata.
	S3BucketName *string

	// The IAM role that grants permissions to access the specified Amazon S3 bucket.
	ServiceAccessRoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFleetAdvisorCollectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFleetAdvisorCollector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFleetAdvisorCollector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFleetAdvisorCollector"); err != nil {
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
	if err = addOpCreateFleetAdvisorCollectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleetAdvisorCollector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFleetAdvisorCollector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFleetAdvisorCollector",
	}
}
