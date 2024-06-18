// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new ReplicationConfigurationTemplate.
func (c *Client) CreateReplicationConfigurationTemplate(ctx context.Context, params *CreateReplicationConfigurationTemplateInput, optFns ...func(*Options)) (*CreateReplicationConfigurationTemplateOutput, error) {
	if params == nil {
		params = &CreateReplicationConfigurationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicationConfigurationTemplate", params, optFns, c.addOperationCreateReplicationConfigurationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicationConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReplicationConfigurationTemplateInput struct {

	// Whether to associate the default Elastic Disaster Recovery Security group with
	// the Replication Configuration Template.
	//
	// This member is required.
	AssociateDefaultSecurityGroup *bool

	// Configure bandwidth throttling for the outbound data transfer rate of the
	// Source Server in Mbps.
	//
	// This member is required.
	BandwidthThrottling int64

	// Whether to create a Public IP for the Recovery Instance by default.
	//
	// This member is required.
	CreatePublicIP *bool

	// The data plane routing mechanism that will be used for replication.
	//
	// This member is required.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// The Staging Disk EBS volume type to be used during replication.
	//
	// This member is required.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// The type of EBS encryption to be used during replication.
	//
	// This member is required.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// The Point in time (PIT) policy to manage snapshots taken during replication.
	//
	// This member is required.
	PitPolicy []types.PITPolicyRule

	// The instance type to be used for the replication server.
	//
	// This member is required.
	ReplicationServerInstanceType *string

	// The security group IDs that will be used by the replication server.
	//
	// This member is required.
	ReplicationServersSecurityGroupsIDs []string

	// The subnet to be used by the replication staging area.
	//
	// This member is required.
	StagingAreaSubnetId *string

	// A set of tags to be associated with all resources created in the replication
	// staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
	//
	// This member is required.
	StagingAreaTags map[string]string

	// Whether to use a dedicated Replication Server in the replication staging area.
	//
	// This member is required.
	UseDedicatedReplicationServer *bool

	// Whether to allow the AWS replication agent to automatically replicate newly
	// added disks.
	AutoReplicateNewDisks *bool

	// The ARN of the EBS encryption key to be used during replication.
	EbsEncryptionKeyArn *string

	// A set of tags to be associated with the Replication Configuration Template
	// resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateReplicationConfigurationTemplateOutput struct {

	// The Replication Configuration Template ID.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string

	// The Replication Configuration Template ARN.
	Arn *string

	// Whether to associate the default Elastic Disaster Recovery Security group with
	// the Replication Configuration Template.
	AssociateDefaultSecurityGroup *bool

	// Whether to allow the AWS replication agent to automatically replicate newly
	// added disks.
	AutoReplicateNewDisks *bool

	// Configure bandwidth throttling for the outbound data transfer rate of the
	// Source Server in Mbps.
	BandwidthThrottling int64

	// Whether to create a Public IP for the Recovery Instance by default.
	CreatePublicIP *bool

	// The data plane routing mechanism that will be used for replication.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// The Staging Disk EBS volume type to be used during replication.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// The type of EBS encryption to be used during replication.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// The ARN of the EBS encryption key to be used during replication.
	EbsEncryptionKeyArn *string

	// The Point in time (PIT) policy to manage snapshots taken during replication.
	PitPolicy []types.PITPolicyRule

	// The instance type to be used for the replication server.
	ReplicationServerInstanceType *string

	// The security group IDs that will be used by the replication server.
	ReplicationServersSecurityGroupsIDs []string

	// The subnet to be used by the replication staging area.
	StagingAreaSubnetId *string

	// A set of tags to be associated with all resources created in the replication
	// staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
	StagingAreaTags map[string]string

	// A set of tags to be associated with the Replication Configuration Template
	// resource.
	Tags map[string]string

	// Whether to use a dedicated Replication Server in the replication staging area.
	UseDedicatedReplicationServer *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReplicationConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateReplicationConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateReplicationConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateReplicationConfigurationTemplate"); err != nil {
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
	if err = addOpCreateReplicationConfigurationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationConfigurationTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicationConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateReplicationConfigurationTemplate",
	}
}
