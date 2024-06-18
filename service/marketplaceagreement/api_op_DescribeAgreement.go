// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplaceagreement

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceagreement/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about an agreement, such as the proposer, acceptor, start
// date, and end date.
func (c *Client) DescribeAgreement(ctx context.Context, params *DescribeAgreementInput, optFns ...func(*Options)) (*DescribeAgreementOutput, error) {
	if params == nil {
		params = &DescribeAgreementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAgreement", params, optFns, c.addOperationDescribeAgreementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAgreementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAgreementInput struct {

	// The unique identifier of the agreement.
	//
	// This member is required.
	AgreementId *string

	noSmithyDocumentSerde
}

type DescribeAgreementOutput struct {

	// The date and time the offer was accepted or the agreement was created.
	//
	// AcceptanceTime and StartTime can differ for future dated agreements (FDAs).
	AcceptanceTime *time.Time

	// The details of the party accepting the agreement terms. This is commonly the
	// buyer for PurchaseAgreement .
	Acceptor *types.Acceptor

	// The unique identifier of the agreement.
	AgreementId *string

	// The type of agreement. Values are PurchaseAgreement or VendorInsightsAgreement .
	AgreementType *string

	// The date and time when the agreement ends. The field is null for pay-as-you-go
	// agreements, which don’t have end dates.
	EndTime *time.Time

	// The estimated cost of the agreement.
	EstimatedCharges *types.EstimatedCharges

	// A summary of the proposal received from the proposer.
	ProposalSummary *types.ProposalSummary

	// The details of the party proposing the agreement terms. This is commonly the
	// seller for PurchaseAgreement .
	Proposer *types.Proposer

	// The date and time when the agreement starts.
	StartTime *time.Time

	// The current status of the agreement.
	//
	// Statuses include:
	//
	//   - ACTIVE – The terms of the agreement are active.
	//
	//   - ARCHIVED – The agreement ended without a specified reason.
	//
	//   - CANCELLED – The acceptor ended the agreement before the defined end date.
	//
	//   - EXPIRED – The agreement ended on the defined end date.
	//
	//   - RENEWED – The agreement was renewed into a new agreement (for example, an
	//   auto-renewal).
	//
	//   - REPLACED – The agreement was replaced using an agreement replacement offer.
	//
	//   - ROLLED_BACK (Only applicable to inactive agreement revisions) – The
	//   agreement revision has been rolled back because of an error. An earlier revision
	//   is now active.
	//
	//   - SUPERCEDED (Only applicable to inactive agreement revisions) – The agreement
	//   revision is no longer active and another agreement revision is now active.
	//
	//   - TERMINATED – The agreement ended before the defined end date because of an
	//   AWS termination (for example, a payment failure).
	Status types.AgreementStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAgreementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeAgreement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeAgreement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAgreement"); err != nil {
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
	if err = addOpDescribeAgreementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAgreement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAgreement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAgreement",
	}
}
