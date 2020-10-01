// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new version of a slot type based on the $LATEST version of the
// specified slot type. If the $LATEST version of this resource has not changed
// since the last version that you created, Amazon Lex doesn't create a new
// version. It returns the last version that you created. You can update only the
// $LATEST version of a slot type. You can't update the numbered versions that you
// create with the CreateSlotTypeVersion operation.  <p>When you create a version
// of a slot type, Amazon Lex sets the version to 1. Subsequent versions increment
// by 1. For more information, see <a>versioning-intro</a>. </p> <p>This operation
// requires permissions for the <code>lex:CreateSlotTypeVersion</code> action.</p>
func (c *Client) CreateSlotTypeVersion(ctx context.Context, params *CreateSlotTypeVersionInput, optFns ...func(*Options)) (*CreateSlotTypeVersionOutput, error) {
	stack := middleware.NewStack("CreateSlotTypeVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateSlotTypeVersionMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateSlotTypeVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSlotTypeVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateSlotTypeVersion",
			Err:           err,
		}
	}
	out := result.(*CreateSlotTypeVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSlotTypeVersionInput struct {

	// Checksum for the $LATEST version of the slot type that you want to publish. If
	// you specify a checksum and the $LATEST version of the slot type has a different
	// checksum, Amazon Lex returns a PreconditionFailedException exception and doesn't
	// publish the new version. If you don't specify a checksum, Amazon Lex publishes
	// the $LATEST version.
	Checksum *string

	// The name of the slot type that you want to create a new version for. The name is
	// case sensitive.
	//
	// This member is required.
	Name *string
}

type CreateSlotTypeVersionOutput struct {

	// The strategy that Amazon Lex uses to determine the value of the slot. For more
	// information, see PutSlotType ().
	ValueSelectionStrategy types.SlotValueSelectionStrategy

	// Configuration information that extends the parent built-in slot type.
	SlotTypeConfigurations []*types.SlotTypeConfiguration

	// The date that the slot type was created.
	CreatedDate *time.Time

	// The built-in slot type used a the parent of the slot type.
	ParentSlotTypeSignature *string

	// The date that the slot type was updated. When you create a resource, the
	// creation date and last update date are the same.
	LastUpdatedDate *time.Time

	// The name of the slot type.
	Name *string

	// Checksum of the $LATEST version of the slot type.
	Checksum *string

	// A description of the slot type.
	Description *string

	// A list of EnumerationValue objects that defines the values that the slot type
	// can take.
	EnumerationValues []*types.EnumerationValue

	// The version assigned to the new slot type version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateSlotTypeVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateSlotTypeVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSlotTypeVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateSlotTypeVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "CreateSlotTypeVersion",
	}
}
