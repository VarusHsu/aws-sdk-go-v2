// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates multiple self-service actions with provisioning artifacts.
func (c *Client) BatchAssociateServiceActionWithProvisioningArtifact(ctx context.Context, params *BatchAssociateServiceActionWithProvisioningArtifactInput, optFns ...func(*Options)) (*BatchAssociateServiceActionWithProvisioningArtifactOutput, error) {
	stack := middleware.NewStack("BatchAssociateServiceActionWithProvisioningArtifact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchAssociateServiceActionWithProvisioningArtifactMiddlewares(stack)
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
	addOpBatchAssociateServiceActionWithProvisioningArtifactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateServiceActionWithProvisioningArtifact(options.Region), middleware.Before)
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
			OperationName: "BatchAssociateServiceActionWithProvisioningArtifact",
			Err:           err,
		}
	}
	out := result.(*BatchAssociateServiceActionWithProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateServiceActionWithProvisioningArtifactInput struct {

	// One or more associations, each consisting of the Action ID, the Product ID, and
	// the Provisioning Artifact ID.
	//
	// This member is required.
	ServiceActionAssociations []*types.ServiceActionAssociation

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
}

type BatchAssociateServiceActionWithProvisioningArtifactOutput struct {

	// An object that contains a list of errors, along with information to help you
	// identify the self-service action.
	FailedServiceActionAssociations []*types.FailedServiceActionAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchAssociateServiceActionWithProvisioningArtifactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchAssociateServiceActionWithProvisioningArtifact{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchAssociateServiceActionWithProvisioningArtifact{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchAssociateServiceActionWithProvisioningArtifact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "BatchAssociateServiceActionWithProvisioningArtifact",
	}
}
