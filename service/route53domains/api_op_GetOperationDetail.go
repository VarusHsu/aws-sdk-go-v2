// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation returns the current status of an operation that is not completed.
func (c *Client) GetOperationDetail(ctx context.Context, params *GetOperationDetailInput, optFns ...func(*Options)) (*GetOperationDetailOutput, error) {
	if params == nil {
		params = &GetOperationDetailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOperationDetail", params, optFns, c.addOperationGetOperationDetailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOperationDetailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The [GetOperationDetail] request includes the following element.
//
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
type GetOperationDetailInput struct {

	// The identifier for the operation for which you want to get the status. Route 53
	// returned the identifier in the response to the original request.
	//
	// This member is required.
	OperationId *string

	noSmithyDocumentSerde
}

// The GetOperationDetail response includes the following elements.
type GetOperationDetailOutput struct {

	// The name of a domain.
	DomainName *string

	//  The date when the operation was last updated.
	LastUpdatedDate *time.Time

	// Detailed information on the status including possible errors.
	Message *string

	// The identifier for the operation.
	OperationId *string

	// The current status of the requested operation in the system.
	Status types.OperationStatus

	//  Lists any outstanding operations that require customer action. Valid values
	// are:
	//
	//   - PENDING_ACCEPTANCE : The operation is waiting for acceptance from the
	//   account that is receiving the domain.
	//
	//   - PENDING_CUSTOMER_ACTION : The operation is waiting for customer action, for
	//   example, returning an email.
	//
	//   - PENDING_AUTHORIZATION : The operation is waiting for the form of
	//   authorization. For more information, see [ResendOperationAuthorization].
	//
	//   - PENDING_PAYMENT_VERIFICATION : The operation is waiting for the payment
	//   method to validate.
	//
	//   - PENDING_SUPPORT_CASE : The operation includes a support case and is waiting
	//   for its resolution.
	//
	// [ResendOperationAuthorization]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ResendOperationAuthorization.html
	StatusFlag types.StatusFlag

	// The date when the request was submitted.
	SubmittedDate *time.Time

	// The type of operation that was requested.
	Type types.OperationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOperationDetailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOperationDetail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOperationDetail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOperationDetail"); err != nil {
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
	if err = addOpGetOperationDetailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOperationDetail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOperationDetail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOperationDetail",
	}
}
