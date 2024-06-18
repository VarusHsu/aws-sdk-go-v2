// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerusersubscriptions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a product subscription for a user with the specified identity provider.
//
// Your estimated bill for charges on the number of users and related costs will
// take 48 hours to appear for billing periods that haven't closed (marked as
// Pending billing status) in Amazon Web Services Billing. For more information,
// see [Viewing your monthly charges]in the Amazon Web Services Billing User Guide.
//
// [Viewing your monthly charges]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/invoice.html
func (c *Client) StartProductSubscription(ctx context.Context, params *StartProductSubscriptionInput, optFns ...func(*Options)) (*StartProductSubscriptionOutput, error) {
	if params == nil {
		params = &StartProductSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartProductSubscription", params, optFns, c.addOperationStartProductSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartProductSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartProductSubscriptionInput struct {

	// An object that specifies details for the identity provider.
	//
	// This member is required.
	IdentityProvider types.IdentityProvider

	// The name of the user-based subscription product.
	//
	// This member is required.
	Product *string

	// The user name from the identity provider of the user.
	//
	// This member is required.
	Username *string

	// The domain name of the user.
	Domain *string

	noSmithyDocumentSerde
}

type StartProductSubscriptionOutput struct {

	// Metadata that describes the start product subscription operation.
	//
	// This member is required.
	ProductUserSummary *types.ProductUserSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartProductSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartProductSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartProductSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartProductSubscription"); err != nil {
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
	if err = addOpStartProductSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartProductSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartProductSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartProductSubscription",
	}
}
