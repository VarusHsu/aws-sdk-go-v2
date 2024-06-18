// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new version of a slot type based on the $LATEST version of the
// specified slot type. If the $LATEST version of this resource has not changed
// since the last version that you created, Amazon Lex doesn't create a new
// version. It returns the last version that you created.
//
// You can update only the $LATEST version of a slot type. You can't update the
// numbered versions that you create with the CreateSlotTypeVersion operation.
//
// When you create a version of a slot type, Amazon Lex sets the version to 1.
// Subsequent versions increment by 1. For more information, see versioning-intro.
//
// This operation requires permissions for the lex:CreateSlotTypeVersion action.
func (c *Client) CreateSlotTypeVersion(ctx context.Context, params *CreateSlotTypeVersionInput, optFns ...func(*Options)) (*CreateSlotTypeVersionOutput, error) {
	if params == nil {
		params = &CreateSlotTypeVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSlotTypeVersion", params, optFns, c.addOperationCreateSlotTypeVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSlotTypeVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSlotTypeVersionInput struct {

	// The name of the slot type that you want to create a new version for. The name
	// is case sensitive.
	//
	// This member is required.
	Name *string

	// Checksum for the $LATEST version of the slot type that you want to publish. If
	// you specify a checksum and the $LATEST version of the slot type has a different
	// checksum, Amazon Lex returns a PreconditionFailedException exception and
	// doesn't publish the new version. If you don't specify a checksum, Amazon Lex
	// publishes the $LATEST version.
	Checksum *string

	noSmithyDocumentSerde
}

type CreateSlotTypeVersionOutput struct {

	// Checksum of the $LATEST version of the slot type.
	Checksum *string

	// The date that the slot type was created.
	CreatedDate *time.Time

	// A description of the slot type.
	Description *string

	// A list of EnumerationValue objects that defines the values that the slot type
	// can take.
	EnumerationValues []types.EnumerationValue

	// The date that the slot type was updated. When you create a resource, the
	// creation date and last update date are the same.
	LastUpdatedDate *time.Time

	// The name of the slot type.
	Name *string

	// The built-in slot type used a the parent of the slot type.
	ParentSlotTypeSignature *string

	// Configuration information that extends the parent built-in slot type.
	SlotTypeConfigurations []types.SlotTypeConfiguration

	// The strategy that Amazon Lex uses to determine the value of the slot. For more
	// information, see PutSlotType.
	ValueSelectionStrategy types.SlotValueSelectionStrategy

	// The version assigned to the new slot type version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSlotTypeVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSlotTypeVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSlotTypeVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSlotTypeVersion"); err != nil {
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
	if err = addOpCreateSlotTypeVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSlotTypeVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSlotTypeVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSlotTypeVersion",
	}
}
