// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Transfers a domain from the current Amazon Web Services account to another
// Amazon Web Services account. Note the following:
//
//   - The Amazon Web Services account that you're transferring the domain to must
//     accept the transfer. If the other account doesn't accept the transfer within 3
//     days, we cancel the transfer. See [AcceptDomainTransferFromAnotherAwsAccount].
//
//   - You can cancel the transfer before the other account accepts it. See [CancelDomainTransferToAnotherAwsAccount].
//
//   - The other account can reject the transfer. See [RejectDomainTransferFromAnotherAwsAccount].
//
// When you transfer a domain from one Amazon Web Services account to another,
// Route 53 doesn't transfer the hosted zone that is associated with the domain.
// DNS resolution isn't affected if the domain and the hosted zone are owned by
// separate accounts, so transferring the hosted zone is optional. For information
// about transferring the hosted zone to another Amazon Web Services account, see [Migrating a Hosted Zone to a Different Amazon Web Services Account]
// in the Amazon Route 53 Developer Guide.
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [RejectDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_RejectDomainTransferFromAnotherAwsAccount.html
// [AcceptDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html
// [Migrating a Hosted Zone to a Different Amazon Web Services Account]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/hosted-zones-migrating.html
// [CancelDomainTransferToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_CancelDomainTransferToAnotherAwsAccount.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
func (c *Client) TransferDomainToAnotherAwsAccount(ctx context.Context, params *TransferDomainToAnotherAwsAccountInput, optFns ...func(*Options)) (*TransferDomainToAnotherAwsAccountOutput, error) {
	if params == nil {
		params = &TransferDomainToAnotherAwsAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TransferDomainToAnotherAwsAccount", params, optFns, c.addOperationTransferDomainToAnotherAwsAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TransferDomainToAnotherAwsAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The TransferDomainToAnotherAwsAccount request includes the following elements.
type TransferDomainToAnotherAwsAccountInput struct {

	// The account ID of the Amazon Web Services account that you want to transfer the
	// domain to, for example, 111122223333 .
	//
	// This member is required.
	AccountId *string

	// The name of the domain that you want to transfer from the current Amazon Web
	// Services account to another account.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// The TransferDomainToAnotherAwsAccount response includes the following elements.
type TransferDomainToAnotherAwsAccountOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use [GetOperationDetail].
	//
	// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
	OperationId *string

	// To finish transferring a domain to another Amazon Web Services account, the
	// account that the domain is being transferred to must submit an [AcceptDomainTransferFromAnotherAwsAccount]request. The
	// request must include the value of the Password element that was returned in the
	// TransferDomainToAnotherAwsAccount response.
	//
	// [AcceptDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html
	Password *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTransferDomainToAnotherAwsAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTransferDomainToAnotherAwsAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTransferDomainToAnotherAwsAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "TransferDomainToAnotherAwsAccount"); err != nil {
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
	if err = addOpTransferDomainToAnotherAwsAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTransferDomainToAnotherAwsAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTransferDomainToAnotherAwsAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TransferDomainToAnotherAwsAccount",
	}
}
