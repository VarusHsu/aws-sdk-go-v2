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

// Modifies the replication instance to apply new settings. You can change one or
// more parameters by specifying these parameters and the new values in the
// request. Some settings are applied during the maintenance window.
func (c *Client) ModifyReplicationInstance(ctx context.Context, params *ModifyReplicationInstanceInput, optFns ...func(*Options)) (*ModifyReplicationInstanceOutput, error) {
	if params == nil {
		params = &ModifyReplicationInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyReplicationInstance", params, optFns, c.addOperationModifyReplicationInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyReplicationInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyReplicationInstanceInput struct {

	// The Amazon Resource Name (ARN) of the replication instance.
	//
	// This member is required.
	ReplicationInstanceArn *string

	// The amount of storage (in gigabytes) to be allocated for the replication
	// instance.
	AllocatedStorage *int32

	// Indicates that major version upgrades are allowed. Changing this parameter does
	// not result in an outage, and the change is asynchronously applied as soon as
	// possible. This parameter must be set to true when specifying a value for the
	// EngineVersion parameter that is a different major version than the replication
	// instance's current version.
	AllowMajorVersionUpgrade bool

	// Indicates whether the changes should be applied immediately or during the next
	// maintenance window.
	ApplyImmediately bool

	// A value that indicates that minor version upgrades are applied automatically to
	// the replication instance during the maintenance window. Changing this parameter
	// doesn't result in an outage, except in the case described following. The change
	// is asynchronously applied as soon as possible. An outage does result if these
	// factors apply:
	//   - This parameter is set to true during the maintenance window.
	//   - A newer minor version is available.
	//   - DMS has enabled automatic patching for the given engine version.
	AutoMinorVersionUpgrade *bool

	// The engine version number of the replication instance. When modifying a major
	// engine version of an instance, also set AllowMajorVersionUpgrade to true .
	EngineVersion *string

	// Specifies whether the replication instance is a Multi-AZ deployment. You can't
	// set the AvailabilityZone parameter if the Multi-AZ parameter is set to true .
	MultiAZ *bool

	// The type of IP address protocol used by a replication instance, such as IPv4
	// only or Dual-stack that supports both IPv4 and IPv6 addressing. IPv6 only is not
	// yet supported.
	NetworkType *string

	// The weekly time range (in UTC) during which system maintenance can occur, which
	// might result in an outage. Changing this parameter does not result in an outage,
	// except in the following situation, and the change is asynchronously applied as
	// soon as possible. If moving this window to the current time, there must be at
	// least 30 minutes between the current time and end of the window to ensure
	// pending changes are applied. Default: Uses existing setting Format:
	// ddd:hh24:mi-ddd:hh24:mi Valid Days: Mon | Tue | Wed | Thu | Fri | Sat | Sun
	// Constraints: Must be at least 30 minutes
	PreferredMaintenanceWindow *string

	// The compute and memory capacity of the replication instance as defined for the
	// specified replication instance class. For example to specify the instance class
	// dms.c4.large, set this parameter to "dms.c4.large" . For more information on the
	// settings and capacities for the available replication instance classes, see
	// Selecting the right DMS replication instance for your migration (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.html#CHAP_ReplicationInstance.InDepth)
	// .
	ReplicationInstanceClass *string

	// The replication instance identifier. This parameter is stored as a lowercase
	// string.
	ReplicationInstanceIdentifier *string

	// Specifies the VPC security group to be used with the replication instance. The
	// VPC security group must work with the VPC containing the replication instance.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type ModifyReplicationInstanceOutput struct {

	// The modified replication instance.
	ReplicationInstance *types.ReplicationInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyReplicationInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyReplicationInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyReplicationInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyReplicationInstance"); err != nil {
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
	if err = addOpModifyReplicationInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReplicationInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyReplicationInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyReplicationInstance",
	}
}
