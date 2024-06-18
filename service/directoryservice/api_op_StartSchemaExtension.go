// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies a schema extension to a Microsoft AD directory.
func (c *Client) StartSchemaExtension(ctx context.Context, params *StartSchemaExtensionInput, optFns ...func(*Options)) (*StartSchemaExtensionOutput, error) {
	if params == nil {
		params = &StartSchemaExtensionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSchemaExtension", params, optFns, c.addOperationStartSchemaExtensionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSchemaExtensionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSchemaExtensionInput struct {

	// If true, creates a snapshot of the directory before applying the schema
	// extension.
	//
	// This member is required.
	CreateSnapshotBeforeSchemaExtension bool

	// A description of the schema extension.
	//
	// This member is required.
	Description *string

	// The identifier of the directory for which the schema extension will be applied
	// to.
	//
	// This member is required.
	DirectoryId *string

	// The LDIF file represented as a string. To construct the LdifContent string,
	// precede each line as it would be formatted in an ldif file with \n. See the
	// example request below for more details. The file size can be no larger than 1MB.
	//
	// This member is required.
	LdifContent *string

	noSmithyDocumentSerde
}

type StartSchemaExtensionOutput struct {

	// The identifier of the schema extension that will be applied.
	SchemaExtensionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSchemaExtensionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartSchemaExtension{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSchemaExtension{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartSchemaExtension"); err != nil {
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
	if err = addOpStartSchemaExtensionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSchemaExtension(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartSchemaExtension(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartSchemaExtension",
	}
}
