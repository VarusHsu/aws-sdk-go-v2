// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an application instance and deploys it to a device.
func (c *Client) CreateApplicationInstance(ctx context.Context, params *CreateApplicationInstanceInput, optFns ...func(*Options)) (*CreateApplicationInstanceOutput, error) {
	if params == nil {
		params = &CreateApplicationInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApplicationInstance", params, optFns, c.addOperationCreateApplicationInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApplicationInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApplicationInstanceInput struct {

	// A device's ID.
	//
	// This member is required.
	DefaultRuntimeContextDevice *string

	// The application's manifest document.
	//
	// This member is required.
	ManifestPayload types.ManifestPayload

	// The ID of an application instance to replace with the new instance.
	ApplicationInstanceIdToReplace *string

	// A description for the application instance.
	Description *string

	// Setting overrides for the application manifest.
	ManifestOverridesPayload types.ManifestOverridesPayload

	// A name for the application instance.
	Name *string

	// The ARN of a runtime role for the application instance.
	RuntimeRoleArn *string

	// Tags for the application instance.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateApplicationInstanceOutput struct {

	// The application instance's ID.
	//
	// This member is required.
	ApplicationInstanceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateApplicationInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApplicationInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApplicationInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateApplicationInstance"); err != nil {
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
	if err = addOpCreateApplicationInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplicationInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApplicationInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateApplicationInstance",
	}
}
