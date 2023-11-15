// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a security profile.
func (c *Client) CreateSecurityProfile(ctx context.Context, params *CreateSecurityProfileInput, optFns ...func(*Options)) (*CreateSecurityProfileOutput, error) {
	if params == nil {
		params = &CreateSecurityProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecurityProfile", params, optFns, c.addOperationCreateSecurityProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityProfileInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The name of the security profile.
	//
	// This member is required.
	SecurityProfileName *string

	// The list of tags that a security profile uses to restrict access to resources
	// in Amazon Connect.
	AllowedAccessControlTags map[string]string

	// This API is in preview release for Amazon Connect and is subject to change. A
	// list of third-party applications that the security profile will give access to.
	Applications []types.Application

	// The description of the security profile.
	Description *string

	// Permissions assigned to the security profile. For a list of valid permissions,
	// see List of security profile permissions (https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html)
	// .
	Permissions []string

	// The list of resources that a security profile applies tag restrictions to in
	// Amazon Connect. Following are acceptable ResourceNames: User | SecurityProfile
	// | Queue | RoutingProfile
	TagRestrictedResources []string

	// The tags used to organize, track, or control access for this resource. For
	// example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSecurityProfileOutput struct {

	// The Amazon Resource Name (ARN) for the security profile.
	SecurityProfileArn *string

	// The identifier for the security profle.
	SecurityProfileId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSecurityProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSecurityProfile"); err != nil {
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
	if err = addOpCreateSecurityProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSecurityProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSecurityProfile",
	}
}
