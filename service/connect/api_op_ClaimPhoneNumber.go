// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Claims an available phone number to your Amazon Connect instance or traffic
// distribution group. You can call this API only in the same Amazon Web Services
// Region where the Amazon Connect instance or traffic distribution group was
// created.
//
// For more information about how to use this operation, see [Claim a phone number in your country] and [Claim phone numbers to traffic distribution groups] in the Amazon
// Connect Administrator Guide.
//
// You can call the [SearchAvailablePhoneNumbers] API for available phone numbers that you can claim. Call the [DescribePhoneNumber]
// API to verify the status of a previous [ClaimPhoneNumber]operation.
//
// If you plan to claim and release numbers frequently, contact us for a service
// quota exception. Otherwise, it is possible you will be blocked from claiming and
// releasing any more numbers until up to 180 days past the oldest number released
// has expired.
//
// By default you can claim and release up to 200% of your maximum number of
// active phone numbers. If you claim and release phone numbers using the UI or API
// during a rolling 180 day cycle that exceeds 200% of your phone number service
// level quota, you will be blocked from claiming any more numbers until 180 days
// past the oldest number released has expired.
//
// For example, if you already have 99 claimed numbers and a service level quota
// of 99 phone numbers, and in any 180 day period you release 99, claim 99, and
// then release 99, you will have exceeded the 200% limit. At that point you are
// blocked from claiming any more numbers until you open an Amazon Web Services
// support ticket.
//
// [Claim phone numbers to traffic distribution groups]: https://docs.aws.amazon.com/connect/latest/adminguide/claim-phone-numbers-traffic-distribution-groups.html
// [Claim a phone number in your country]: https://docs.aws.amazon.com/connect/latest/adminguide/claim-phone-number.html
// [SearchAvailablePhoneNumbers]: https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchAvailablePhoneNumbers.html
// [DescribePhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribePhoneNumber.html
// [ClaimPhoneNumber]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ClaimPhoneNumber.html
func (c *Client) ClaimPhoneNumber(ctx context.Context, params *ClaimPhoneNumberInput, optFns ...func(*Options)) (*ClaimPhoneNumberOutput, error) {
	if params == nil {
		params = &ClaimPhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ClaimPhoneNumber", params, optFns, c.addOperationClaimPhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ClaimPhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ClaimPhoneNumberInput struct {

	// The phone number you want to claim. Phone numbers are formatted [+] [country
	// code] [subscriber number including area code] .
	//
	// This member is required.
	PhoneNumber *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see [Making retries safe with idempotent APIs].
	//
	// Pattern: ^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$
	//
	// [Making retries safe with idempotent APIs]: https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/
	ClientToken *string

	// The identifier of the Amazon Connect instance that phone numbers are claimed
	// to. You can [find the instance ID]in the Amazon Resource Name (ARN) of the instance. You must enter
	// InstanceId or TargetArn .
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	InstanceId *string

	// The description of the phone number.
	PhoneNumberDescription *string

	// The tags used to organize, track, or control access for this resource. For
	// example, { "Tags": {"key1":"value1", "key2":"value2"} }.
	Tags map[string]string

	// The Amazon Resource Name (ARN) for Amazon Connect instances or traffic
	// distribution groups that phone number inbound traffic is routed through. You
	// must enter InstanceId or TargetArn .
	TargetArn *string

	noSmithyDocumentSerde
}

type ClaimPhoneNumberOutput struct {

	// The Amazon Resource Name (ARN) of the phone number.
	PhoneNumberArn *string

	// A unique identifier for the phone number.
	PhoneNumberId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationClaimPhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpClaimPhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpClaimPhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ClaimPhoneNumber"); err != nil {
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
	if err = addIdempotencyToken_opClaimPhoneNumberMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpClaimPhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opClaimPhoneNumber(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpClaimPhoneNumber struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpClaimPhoneNumber) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpClaimPhoneNumber) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ClaimPhoneNumberInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ClaimPhoneNumberInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opClaimPhoneNumberMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpClaimPhoneNumber{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opClaimPhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ClaimPhoneNumber",
	}
}
