// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Runs an on-demand evaluation for the specified resource to determine whether
// the resource details will comply with configured Config rules. You can also use
// it for evaluation purposes. Config recommends using an evaluation context. It
// runs an execution against the resource details with all of the Config rules in
// your account that match with the specified proactive mode and resource type.
//
// Ensure you have the cloudformation:DescribeType role setup to validate the
// resource type schema.
//
// You can find the [Resource type schema] in "Amazon Web Services public extensions" within the
// CloudFormation registry or with the following CLI commmand: aws cloudformation
// describe-type --type-name "AWS::S3::Bucket" --type RESOURCE .
//
// For more information, see [Managing extensions through the CloudFormation registry] and [Amazon Web Services resource and property types reference] in the CloudFormation User Guide.
//
// [Resource type schema]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html
// [Amazon Web Services resource and property types reference]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html
// [Managing extensions through the CloudFormation registry]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html#registry-view
func (c *Client) StartResourceEvaluation(ctx context.Context, params *StartResourceEvaluationInput, optFns ...func(*Options)) (*StartResourceEvaluationOutput, error) {
	if params == nil {
		params = &StartResourceEvaluationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartResourceEvaluation", params, optFns, c.addOperationStartResourceEvaluationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartResourceEvaluationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartResourceEvaluationInput struct {

	// The mode of an evaluation. The valid values for this API are DETECTIVE and
	// PROACTIVE .
	//
	// This member is required.
	EvaluationMode types.EvaluationMode

	// Returns a ResourceDetails object.
	//
	// This member is required.
	ResourceDetails *types.ResourceDetails

	// A client token is a unique, case-sensitive string of up to 64 ASCII characters.
	// To make an idempotent API request using one of these actions, specify a client
	// token in the request.
	//
	// Avoid reusing the same client token for other API requests. If you retry a
	// request that completed successfully using the same client token and the same
	// parameters, the retry succeeds without performing any further actions. If you
	// retry a successful request using the same client token, but one or more of the
	// parameters are different, other than the Region or Availability Zone, the retry
	// fails with an IdempotentParameterMismatch error.
	ClientToken *string

	// Returns an EvaluationContext object.
	EvaluationContext *types.EvaluationContext

	// The timeout for an evaluation. The default is 900 seconds. You cannot specify a
	// number greater than 3600. If you specify 0, Config uses the default.
	EvaluationTimeout int32

	noSmithyDocumentSerde
}

type StartResourceEvaluationOutput struct {

	// A unique ResourceEvaluationId that is associated with a single execution.
	ResourceEvaluationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartResourceEvaluationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartResourceEvaluation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartResourceEvaluation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartResourceEvaluation"); err != nil {
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
	if err = addOpStartResourceEvaluationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartResourceEvaluation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartResourceEvaluation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartResourceEvaluation",
	}
}
