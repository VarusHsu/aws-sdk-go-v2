// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The description of the domain.
func (c *Client) DescribeDomain(ctx context.Context, params *DescribeDomainInput, optFns ...func(*Options)) (*DescribeDomainOutput, error) {
	if params == nil {
		params = &DescribeDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomain", params, optFns, c.addOperationDescribeDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDomainInput struct {

	// The domain ID.
	//
	// This member is required.
	DomainId *string

	noSmithyDocumentSerde
}

type DescribeDomainOutput struct {

	// Specifies the VPC used for non-EFS traffic. The default value is
	// PublicInternetOnly .
	//   - PublicInternetOnly - Non-EFS traffic is through a VPC managed by Amazon
	//   SageMaker, which allows direct internet access
	//   - VpcOnly - All Studio traffic is through the specified VPC and subnets
	AppNetworkAccessType types.AppNetworkAccessType

	// The entity that creates and manages the required security groups for inter-app
	// communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType
	// is VPCOnly and
	// DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.
	AppSecurityGroupManagement types.AppSecurityGroupManagement

	// The domain's authentication mode.
	AuthMode types.AuthMode

	// The creation time.
	CreationTime *time.Time

	// The default settings used to create a space.
	DefaultSpaceSettings *types.DefaultSpaceSettings

	// Settings which are applied to UserProfiles in this domain if settings are not
	// explicitly specified in a given UserProfile.
	DefaultUserSettings *types.UserSettings

	// The domain's Amazon Resource Name (ARN).
	DomainArn *string

	// The domain ID.
	DomainId *string

	// The domain name.
	DomainName *string

	// A collection of Domain settings.
	DomainSettings *types.DomainSettings

	// The failure reason.
	FailureReason *string

	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId *string

	// Use KmsKeyId .
	//
	// Deprecated: This property is deprecated, use KmsKeyId instead.
	HomeEfsFileSystemKmsKeyId *string

	// The Amazon Web Services KMS customer managed key used to encrypt the EFS volume
	// attached to the domain.
	KmsKeyId *string

	// The last modified time.
	LastModifiedTime *time.Time

	// The ID of the security group that authorizes traffic between the RSessionGateway
	// apps and the RStudioServerPro app.
	SecurityGroupIdForDomainBoundary *string

	// The IAM Identity Center managed application instance ID.
	SingleSignOnManagedApplicationInstanceId *string

	// The status.
	Status types.DomainStatus

	// The VPC subnets that Studio uses for communication.
	SubnetIds []string

	// The domain's URL.
	Url *string

	// The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for
	// communication.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDomain"); err != nil {
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
	if err = addOpDescribeDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDomain",
	}
}
