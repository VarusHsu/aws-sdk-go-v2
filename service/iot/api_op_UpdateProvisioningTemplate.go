// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a fleet provisioning template.
func (c *Client) UpdateProvisioningTemplate(ctx context.Context, params *UpdateProvisioningTemplateInput, optFns ...func(*Options)) (*UpdateProvisioningTemplateOutput, error) {
	stack := middleware.NewStack("UpdateProvisioningTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateProvisioningTemplateMiddlewares(stack)
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
	addOpUpdateProvisioningTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProvisioningTemplate(options.Region), middleware.Before)
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
			OperationName: "UpdateProvisioningTemplate",
			Err:           err,
		}
	}
	out := result.(*UpdateProvisioningTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProvisioningTemplateInput struct {

	// The name of the fleet provisioning template.
	//
	// This member is required.
	TemplateName *string

	// True to enable the fleet provisioning template, otherwise false.
	Enabled *bool

	// Removes pre-provisioning hook template.
	RemovePreProvisioningHook *bool

	// Updates the pre-provisioning hook template.
	PreProvisioningHook *types.ProvisioningHook

	// The ID of the default provisioning template version.
	DefaultVersionId *int32

	// The description of the fleet provisioning template.
	Description *string

	// The ARN of the role associated with the provisioning template. This IoT role
	// grants permission to provision a device.
	ProvisioningRoleArn *string
}

type UpdateProvisioningTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateProvisioningTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateProvisioningTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateProvisioningTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateProvisioningTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateProvisioningTemplate",
	}
}
