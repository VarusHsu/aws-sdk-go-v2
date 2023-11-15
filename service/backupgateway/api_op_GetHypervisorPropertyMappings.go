// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backupgateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action retrieves the property mappings for the specified hypervisor. A
// hypervisor property mapping displays the relationship of entity properties
// available from the on-premises hypervisor to the properties available in Amazon
// Web Services.
func (c *Client) GetHypervisorPropertyMappings(ctx context.Context, params *GetHypervisorPropertyMappingsInput, optFns ...func(*Options)) (*GetHypervisorPropertyMappingsOutput, error) {
	if params == nil {
		params = &GetHypervisorPropertyMappingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHypervisorPropertyMappings", params, optFns, c.addOperationGetHypervisorPropertyMappingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHypervisorPropertyMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetHypervisorPropertyMappingsInput struct {

	// The Amazon Resource Name (ARN) of the hypervisor.
	//
	// This member is required.
	HypervisorArn *string

	noSmithyDocumentSerde
}

type GetHypervisorPropertyMappingsOutput struct {

	// The Amazon Resource Name (ARN) of the hypervisor.
	HypervisorArn *string

	// The Amazon Resource Name (ARN) of the IAM role.
	IamRoleArn *string

	// This is a display of the mappings of on-premises VMware tags to the Amazon Web
	// Services tags.
	VmwareToAwsTagMappings []types.VmwareToAwsTagMapping

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetHypervisorPropertyMappingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetHypervisorPropertyMappings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetHypervisorPropertyMappings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetHypervisorPropertyMappings"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpGetHypervisorPropertyMappingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHypervisorPropertyMappings(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetHypervisorPropertyMappings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetHypervisorPropertyMappings",
	}
}
