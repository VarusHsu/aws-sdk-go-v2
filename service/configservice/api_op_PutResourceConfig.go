// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Records the configuration state for the resource provided in the request.  The
// configuration state of a resource is represented in AWS Config as Configuration
// Items. Once this API records the configuration item, you can retrieve the list
// of configuration items for the custom resource type using existing AWS Config
// APIs. </p> <note> <p>The custom resource type must be registered with AWS
// CloudFormation. This API accepts the configuration item registered with AWS
// CloudFormation.</p> <p>When you call this API, AWS Config only stores
// configuration state of the resource provided in the request. This API does not
// change or remediate the configuration of the resource. </p> <p>Write-only schema
// properites are not recorded as part of the published configuration item.</p>
// </note>
func (c *Client) PutResourceConfig(ctx context.Context, params *PutResourceConfigInput, optFns ...func(*Options)) (*PutResourceConfigOutput, error) {
	stack := middleware.NewStack("PutResourceConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutResourceConfigMiddlewares(stack)
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
	addOpPutResourceConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourceConfig(options.Region), middleware.Before)
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
			OperationName: "PutResourceConfig",
			Err:           err,
		}
	}
	out := result.(*PutResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourceConfigInput struct {

	// Unique identifier of the resource.
	//
	// This member is required.
	ResourceId *string

	// Version of the schema registered for the ResourceType in AWS CloudFormation.
	//
	// This member is required.
	SchemaVersionId *string

	// Tags associated with the resource.
	Tags map[string]*string

	// The configuration object of the resource in valid JSON format. It must match the
	// schema registered with AWS CloudFormation. The configuration JSON must not
	// exceed 64 KB.
	//
	// This member is required.
	Configuration *string

	// Name of the resource.
	ResourceName *string

	// The type of the resource. The custom resource type must be registered with AWS
	// CloudFormation. You cannot use the organization names “aws”, “amzn”, “amazon”,
	// “alexa”, “custom” with custom resource types. It is the first part of the
	// ResourceType up to the first ::.
	//
	// This member is required.
	ResourceType *string
}

type PutResourceConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutResourceConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourceConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourceConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutResourceConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutResourceConfig",
	}
}
