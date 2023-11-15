// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the result of a Lambda validation function. The function validates
// lifecycle hooks during a deployment that uses the Lambda or Amazon ECS compute
// platform. For Lambda deployments, the available lifecycle hooks are
// BeforeAllowTraffic and AfterAllowTraffic . For Amazon ECS deployments, the
// available lifecycle hooks are BeforeInstall , AfterInstall ,
// AfterAllowTestTraffic , BeforeAllowTraffic , and AfterAllowTraffic . Lambda
// validation functions return Succeeded or Failed . For more information, see
// AppSpec 'hooks' Section for an Lambda Deployment  (https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-lambda)
// and AppSpec 'hooks' Section for an Amazon ECS Deployment (https://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file-structure-hooks.html#appspec-hooks-ecs)
// .
func (c *Client) PutLifecycleEventHookExecutionStatus(ctx context.Context, params *PutLifecycleEventHookExecutionStatusInput, optFns ...func(*Options)) (*PutLifecycleEventHookExecutionStatusOutput, error) {
	if params == nil {
		params = &PutLifecycleEventHookExecutionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecycleEventHookExecutionStatus", params, optFns, c.addOperationPutLifecycleEventHookExecutionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecycleEventHookExecutionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecycleEventHookExecutionStatusInput struct {

	// The unique ID of a deployment. Pass this ID to a Lambda function that validates
	// a deployment lifecycle event.
	DeploymentId *string

	// The execution ID of a deployment's lifecycle hook. A deployment lifecycle hook
	// is specified in the hooks section of the AppSpec file.
	LifecycleEventHookExecutionId *string

	// The result of a Lambda function that validates a deployment lifecycle event.
	// The values listed in Valid Values are valid for lifecycle statuses in general;
	// however, only Succeeded and Failed can be passed successfully in your API call.
	Status types.LifecycleEventStatus

	noSmithyDocumentSerde
}

type PutLifecycleEventHookExecutionStatusOutput struct {

	// The execution ID of the lifecycle event hook. A hook is specified in the hooks
	// section of the deployment's AppSpec file.
	LifecycleEventHookExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLifecycleEventHookExecutionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutLifecycleEventHookExecutionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutLifecycleEventHookExecutionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutLifecycleEventHookExecutionStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecycleEventHookExecutionStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLifecycleEventHookExecutionStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutLifecycleEventHookExecutionStatus",
	}
}
