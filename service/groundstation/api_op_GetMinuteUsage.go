// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the number of reserved minutes used by account.
func (c *Client) GetMinuteUsage(ctx context.Context, params *GetMinuteUsageInput, optFns ...func(*Options)) (*GetMinuteUsageOutput, error) {
	if params == nil {
		params = &GetMinuteUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMinuteUsage", params, optFns, c.addOperationGetMinuteUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMinuteUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMinuteUsageInput struct {

	// The month being requested, with a value of 1-12.
	//
	// This member is required.
	Month *int32

	// The year being requested, in the format of YYYY.
	//
	// This member is required.
	Year *int32

	noSmithyDocumentSerde
}

type GetMinuteUsageOutput struct {

	// Estimated number of minutes remaining for an account, specific to the month
	// being requested.
	EstimatedMinutesRemaining *int32

	// Returns whether or not an account has signed up for the reserved minutes
	// pricing plan, specific to the month being requested.
	IsReservedMinutesCustomer *bool

	// Total number of reserved minutes allocated, specific to the month being
	// requested.
	TotalReservedMinuteAllocation *int32

	// Total scheduled minutes for an account, specific to the month being requested.
	TotalScheduledMinutes *int32

	// Upcoming minutes scheduled for an account, specific to the month being
	// requested.
	UpcomingMinutesScheduled *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMinuteUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMinuteUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMinuteUsage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMinuteUsage"); err != nil {
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
	if err = addOpGetMinuteUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMinuteUsage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMinuteUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMinuteUsage",
	}
}
