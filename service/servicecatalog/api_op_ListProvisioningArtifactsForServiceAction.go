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

// Lists all provisioning artifacts (also known as versions) for the specified
// self-service action.
func (c *Client) ListProvisioningArtifactsForServiceAction(ctx context.Context, params *ListProvisioningArtifactsForServiceActionInput, optFns ...func(*Options)) (*ListProvisioningArtifactsForServiceActionOutput, error) {
	stack := middleware.NewStack("ListProvisioningArtifactsForServiceAction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListProvisioningArtifactsForServiceActionMiddlewares(stack)
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
	addOpListProvisioningArtifactsForServiceActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProvisioningArtifactsForServiceAction(options.Region), middleware.Before)
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
			OperationName: "ListProvisioningArtifactsForServiceAction",
			Err:           err,
		}
	}
	out := result.(*ListProvisioningArtifactsForServiceActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProvisioningArtifactsForServiceActionInput struct {

	// The self-service action identifier. For example, act-fs7abcd89wxyz.
	//
	// This member is required.
	ServiceActionId *string

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The maximum number of items to return with this call.
	PageSize *int32
}

type ListProvisioningArtifactsForServiceActionOutput struct {

	// An array of objects with information about product views and provisioning
	// artifacts.
	ProvisioningArtifactViews []*types.ProvisioningArtifactView

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListProvisioningArtifactsForServiceActionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListProvisioningArtifactsForServiceAction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProvisioningArtifactsForServiceAction{}, middleware.After)
}

func newServiceMetadataMiddleware_opListProvisioningArtifactsForServiceAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListProvisioningArtifactsForServiceAction",
	}
}
