// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Designates the Amazon Security Lake delegated administrator account for the
// organization. This API can only be called by the organization management
// account. The organization management account cannot be the delegated
// administrator account.
func (c *Client) RegisterDataLakeDelegatedAdministrator(ctx context.Context, params *RegisterDataLakeDelegatedAdministratorInput, optFns ...func(*Options)) (*RegisterDataLakeDelegatedAdministratorOutput, error) {
	if params == nil {
		params = &RegisterDataLakeDelegatedAdministratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterDataLakeDelegatedAdministrator", params, optFns, c.addOperationRegisterDataLakeDelegatedAdministratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterDataLakeDelegatedAdministratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterDataLakeDelegatedAdministratorInput struct {

	// The Amazon Web Services account ID of the Security Lake delegated administrator.
	//
	// This member is required.
	AccountId *string

	noSmithyDocumentSerde
}

type RegisterDataLakeDelegatedAdministratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterDataLakeDelegatedAdministratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterDataLakeDelegatedAdministrator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterDataLakeDelegatedAdministrator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterDataLakeDelegatedAdministrator"); err != nil {
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
	if err = addOpRegisterDataLakeDelegatedAdministratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDataLakeDelegatedAdministrator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterDataLakeDelegatedAdministrator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterDataLakeDelegatedAdministrator",
	}
}
