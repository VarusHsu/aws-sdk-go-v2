// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a broker. Note: This API is asynchronous. To create a broker, you must
// either use the AmazonMQFullAccess IAM policy or include the following EC2
// permissions in your IAM policy.
//   - ec2:CreateNetworkInterface This permission is required to allow Amazon MQ
//     to create an elastic network interface (ENI) on behalf of your account.
//   - ec2:CreateNetworkInterfacePermission This permission is required to attach
//     the ENI to the broker instance.
//   - ec2:DeleteNetworkInterface
//   - ec2:DeleteNetworkInterfacePermission
//   - ec2:DetachNetworkInterface
//   - ec2:DescribeInternetGateways
//   - ec2:DescribeNetworkInterfaces
//   - ec2:DescribeNetworkInterfacePermissions
//   - ec2:DescribeRouteTables
//   - ec2:DescribeSecurityGroups
//   - ec2:DescribeSubnets
//   - ec2:DescribeVpcs
//
// For more information, see Create an IAM User and Get Your Amazon Web Services
// Credentials (https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/amazon-mq-setting-up.html#create-iam-user)
// and Never Modify or Delete the Amazon MQ Elastic Network Interface (https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/connecting-to-amazon-mq.html#never-modify-delete-elastic-network-interface)
// in the Amazon MQ Developer Guide.
func (c *Client) CreateBroker(ctx context.Context, params *CreateBrokerInput, optFns ...func(*Options)) (*CreateBrokerOutput, error) {
	if params == nil {
		params = &CreateBrokerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBroker", params, optFns, c.addOperationCreateBrokerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBrokerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a broker using the specified properties.
type CreateBrokerInput struct {

	// Enables automatic upgrades to new minor versions for brokers, as new versions
	// are released and supported by Amazon MQ. Automatic upgrades occur during the
	// scheduled maintenance window of the broker or after a manual broker reboot. Set
	// to true by default, if no value is specified.
	//
	// This member is required.
	AutoMinorVersionUpgrade *bool

	// Required. The broker's name. This value must be unique in your Amazon Web
	// Services account, 1-50 characters long, must contain only letters, numbers,
	// dashes, and underscores, and must not contain white spaces, brackets, wildcard
	// characters, or special characters. Do not add personally identifiable
	// information (PII) or other confidential or sensitive information in broker
	// names. Broker names are accessible to other Amazon Web Services services,
	// including CloudWatch Logs. Broker names are not intended to be used for private
	// or sensitive data.
	//
	// This member is required.
	BrokerName *string

	// Required. The broker's deployment mode.
	//
	// This member is required.
	DeploymentMode types.DeploymentMode

	// Required. The type of broker engine. Currently, Amazon MQ supports ACTIVEMQ and
	// RABBITMQ.
	//
	// This member is required.
	EngineType types.EngineType

	// Required. The broker engine's version. For a list of supported engine versions,
	// see Supported engines (https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker-engine.html)
	// .
	//
	// This member is required.
	EngineVersion *string

	// Required. The broker's instance type.
	//
	// This member is required.
	HostInstanceType *string

	// Enables connections from applications outside of the VPC that hosts the
	// broker's subnets. Set to false by default, if no value is provided.
	//
	// This member is required.
	PubliclyAccessible *bool

	// The list of broker users (persons or applications) who can access queues and
	// topics. For Amazon MQ for RabbitMQ brokers, one and only one administrative user
	// is accepted and created when a broker is first provisioned. All subsequent
	// broker users are created by making RabbitMQ API calls directly to brokers or via
	// the RabbitMQ web console.
	//
	// This member is required.
	Users []types.User

	// Optional. The authentication strategy used to secure the broker. The default is
	// SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// A list of information about the configuration.
	Configuration *types.ConfigurationId

	// The unique ID that the requester receives for the created broker. Amazon MQ
	// passes your ID with the API action. We recommend using a Universally Unique
	// Identifier (UUID) for the creatorRequestId. You may omit the creatorRequestId if
	// your application doesn't require idempotency.
	CreatorRequestId *string

	// Defines whether this broker is a part of a data replication pair.
	DataReplicationMode types.DataReplicationMode

	// The Amazon Resource Name (ARN) of the primary broker that is used to replicate
	// data from in a data replication pair, and is applied to the replica broker. Must
	// be set when dataReplicationMode is set to CRDR.
	DataReplicationPrimaryBrokerArn *string

	// Encryption options for the broker.
	EncryptionOptions *types.EncryptionOptions

	// Optional. The metadata of the LDAP server used to authenticate and authorize
	// connections to the broker. Does not apply to RabbitMQ brokers.
	LdapServerMetadata *types.LdapServerMetadataInput

	// Enables Amazon CloudWatch logging for brokers.
	Logs *types.Logs

	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime *types.WeeklyStartTime

	// The list of rules (1 minimum, 125 maximum) that authorize connections to
	// brokers.
	SecurityGroups []string

	// The broker's storage type.
	StorageType types.BrokerStorageType

	// The list of groups that define which subnets and IP ranges the broker can use
	// from different Availability Zones. If you specify more than one subnet, the
	// subnets must be in different Availability Zones. Amazon MQ will not be able to
	// create VPC endpoints for your broker with multiple subnets in the same
	// Availability Zone. A SINGLE_INSTANCE deployment requires one subnet (for
	// example, the default subnet). An ACTIVE_STANDBY_MULTI_AZ Amazon MQ for ActiveMQ
	// deployment requires two subnets. A CLUSTER_MULTI_AZ Amazon MQ for RabbitMQ
	// deployment has no subnet requirements when deployed with public accessibility.
	// Deployment without public accessibility requires at least one subnet. If you
	// specify subnets in a shared VPC (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-sharing.html)
	// for a RabbitMQ broker, the associated VPC to which the specified subnets belong
	// must be owned by your Amazon Web Services account. Amazon MQ will not be able to
	// create VPC endpoints in VPCs that are not owned by your Amazon Web Services
	// account.
	SubnetIds []string

	// Create tags when creating the broker.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateBrokerOutput struct {

	// The broker's Amazon Resource Name (ARN).
	BrokerArn *string

	// The unique ID that Amazon MQ generates for the broker.
	BrokerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBrokerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBroker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBroker{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBroker"); err != nil {
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
	if err = addIdempotencyToken_opCreateBrokerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateBrokerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBroker(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateBroker struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateBroker) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateBroker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateBrokerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateBrokerInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateBrokerMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateBroker{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateBroker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBroker",
	}
}
