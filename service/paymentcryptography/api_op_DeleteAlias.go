// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptography

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the alias, but doesn't affect the underlying key.
//
// Each key can have multiple aliases. To get the aliases of all keys, use the [UpdateAlias]
// operation. To change the alias of a key, first use [DeleteAlias]to delete the current alias
// and then use [CreateAlias]to create a new alias. To associate an existing alias with a
// different key, call [UpdateAlias].
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateAlias]
//
// [GetAlias]
//
// [ListAliases]
//
// [UpdateAlias]
//
// [ListAliases]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListAliases.html
// [DeleteAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteAlias.html
// [UpdateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_UpdateAlias.html
// [CreateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateAlias.html
// [GetAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetAlias.html
func (c *Client) DeleteAlias(ctx context.Context, params *DeleteAliasInput, optFns ...func(*Options)) (*DeleteAliasOutput, error) {
	if params == nil {
		params = &DeleteAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAlias", params, optFns, c.addOperationDeleteAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAliasInput struct {

	// A friendly name that you can use to refer Amazon Web Services Payment
	// Cryptography key. This value must begin with alias/ followed by a name, such as
	// alias/ExampleAlias .
	//
	// This member is required.
	AliasName *string

	noSmithyDocumentSerde
}

type DeleteAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteAlias"); err != nil {
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
	if err = addOpDeleteAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteAlias",
	}
}
