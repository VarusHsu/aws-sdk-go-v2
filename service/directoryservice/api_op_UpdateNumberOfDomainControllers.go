// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or removes domain controllers to or from the directory. Based on the
// difference between current value and new value (provided through this API call),
// domain controllers will be added or removed. It may take up to 45 minutes for
// any new domain controllers to become fully active once the requested number of
// domain controllers is updated. During this time, you cannot make another update
// request.
func (c *Client) UpdateNumberOfDomainControllers(ctx context.Context, params *UpdateNumberOfDomainControllersInput, optFns ...func(*Options)) (*UpdateNumberOfDomainControllersOutput, error) {
	if params == nil {
		params = &UpdateNumberOfDomainControllersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNumberOfDomainControllers", params, optFns, c.addOperationUpdateNumberOfDomainControllersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNumberOfDomainControllersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNumberOfDomainControllersInput struct {

	// The number of domain controllers desired in the directory.
	//
	// This member is required.
	DesiredNumber *int32

	// Identifier of the directory to which the domain controllers will be added or
	// removed.
	//
	// This member is required.
	DirectoryId *string

	noSmithyDocumentSerde
}

type UpdateNumberOfDomainControllersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNumberOfDomainControllersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNumberOfDomainControllers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNumberOfDomainControllers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateNumberOfDomainControllers"); err != nil {
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
	if err = addOpUpdateNumberOfDomainControllersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNumberOfDomainControllers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNumberOfDomainControllers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateNumberOfDomainControllers",
	}
}
