// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates one or more scheduled scaling actions for an Auto Scaling
// group.
func (c *Client) BatchPutScheduledUpdateGroupAction(ctx context.Context, params *BatchPutScheduledUpdateGroupActionInput, optFns ...func(*Options)) (*BatchPutScheduledUpdateGroupActionOutput, error) {
	if params == nil {
		params = &BatchPutScheduledUpdateGroupActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchPutScheduledUpdateGroupAction", params, optFns, c.addOperationBatchPutScheduledUpdateGroupActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchPutScheduledUpdateGroupActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchPutScheduledUpdateGroupActionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// One or more scheduled actions. The maximum number allowed is 50.
	//
	// This member is required.
	ScheduledUpdateGroupActions []types.ScheduledUpdateGroupActionRequest

	noSmithyDocumentSerde
}

type BatchPutScheduledUpdateGroupActionOutput struct {

	// The names of the scheduled actions that could not be created or updated,
	// including an error message.
	FailedScheduledUpdateGroupActions []types.FailedScheduledUpdateGroupActionRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchPutScheduledUpdateGroupActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpBatchPutScheduledUpdateGroupAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpBatchPutScheduledUpdateGroupAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchPutScheduledUpdateGroupAction"); err != nil {
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
	if err = addOpBatchPutScheduledUpdateGroupActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchPutScheduledUpdateGroupAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchPutScheduledUpdateGroupAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchPutScheduledUpdateGroupAction",
	}
}
