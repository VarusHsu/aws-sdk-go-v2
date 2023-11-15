// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified instance profile using the provided parameters. All
// migration projects associated with the instance profile must be deleted or
// modified before you can modify the instance profile.
func (c *Client) ModifyInstanceProfile(ctx context.Context, params *ModifyInstanceProfileInput, optFns ...func(*Options)) (*ModifyInstanceProfileOutput, error) {
	if params == nil {
		params = &ModifyInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceProfile", params, optFns, c.addOperationModifyInstanceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceProfileInput struct {

	// The identifier of the instance profile. Identifiers must begin with a letter
	// and must contain only ASCII letters, digits, and hyphens. They can't end with a
	// hyphen, or contain two consecutive hyphens.
	//
	// This member is required.
	InstanceProfileIdentifier *string

	// The Availability Zone where the instance profile runs.
	AvailabilityZone *string

	// A user-friendly description for the instance profile.
	Description *string

	// A user-friendly name for the instance profile.
	InstanceProfileName *string

	// The Amazon Resource Name (ARN) of the KMS key that is used to encrypt the
	// connection parameters for the instance profile. If you don't specify a value for
	// the KmsKeyArn parameter, then DMS uses your default encryption key. KMS creates
	// the default encryption key for your Amazon Web Services account. Your Amazon Web
	// Services account has a different default encryption key for each Amazon Web
	// Services Region.
	KmsKeyArn *string

	// Specifies the network type for the instance profile. A value of IPV4 represents
	// an instance profile with IPv4 network type and only supports IPv4 addressing. A
	// value of IPV6 represents an instance profile with IPv6 network type and only
	// supports IPv6 addressing. A value of DUAL represents an instance profile with
	// dual network type that supports IPv4 and IPv6 addressing.
	NetworkType *string

	// Specifies the accessibility options for the instance profile. A value of true
	// represents an instance profile with a public IP address. A value of false
	// represents an instance profile with a private IP address. The default value is
	// true .
	PubliclyAccessible *bool

	// A subnet group to associate with the instance profile.
	SubnetGroupIdentifier *string

	// Specifies the VPC security groups to be used with the instance profile. The VPC
	// security group must work with the VPC containing the instance profile.
	VpcSecurityGroups []string

	noSmithyDocumentSerde
}

type ModifyInstanceProfileOutput struct {

	// The instance profile that was modified.
	InstanceProfile *types.InstanceProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyInstanceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyInstanceProfile"); err != nil {
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
	if err = addOpModifyInstanceProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyInstanceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyInstanceProfile",
	}
}
