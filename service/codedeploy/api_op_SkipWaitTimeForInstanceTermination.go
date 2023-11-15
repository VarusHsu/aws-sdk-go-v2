// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// In a blue/green deployment, overrides any specified wait time and starts
// terminating instances immediately after the traffic routing is complete.
//
// Deprecated: This operation is deprecated, use ContinueDeployment with
// DeploymentWaitType instead.
func (c *Client) SkipWaitTimeForInstanceTermination(ctx context.Context, params *SkipWaitTimeForInstanceTerminationInput, optFns ...func(*Options)) (*SkipWaitTimeForInstanceTerminationOutput, error) {
	if params == nil {
		params = &SkipWaitTimeForInstanceTerminationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SkipWaitTimeForInstanceTermination", params, optFns, c.addOperationSkipWaitTimeForInstanceTerminationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SkipWaitTimeForInstanceTerminationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SkipWaitTimeForInstanceTerminationInput struct {

	// The unique ID of a blue/green deployment for which you want to skip the
	// instance termination wait time.
	DeploymentId *string

	noSmithyDocumentSerde
}

type SkipWaitTimeForInstanceTerminationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSkipWaitTimeForInstanceTerminationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSkipWaitTimeForInstanceTermination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSkipWaitTimeForInstanceTermination{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SkipWaitTimeForInstanceTermination"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSkipWaitTimeForInstanceTermination(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opSkipWaitTimeForInstanceTermination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SkipWaitTimeForInstanceTermination",
	}
}
