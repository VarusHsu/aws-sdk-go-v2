// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new cluster with the specified parameters. To create a cluster in
// Virtual Private Cloud (VPC), you must provide a cluster subnet group name. The
// cluster subnet group identifies the subnets of your VPC that Amazon Redshift
// uses when creating the cluster. For more information about managing clusters, go
// to Amazon Redshift Clusters
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in
// the Amazon Redshift Cluster Management Guide.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	stack := middleware.NewStack("CreateCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateClusterMiddlewares(stack)
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
	addOpCreateClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before)
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
			OperationName: "CreateCluster",
			Err:           err,
		}
	}
	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateClusterInput struct {

	// The AWS Key Management Service (KMS) key ID of the encryption key that you want
	// to use to encrypt data in the cluster.
	KmsKeyId *string

	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision
	// the cluster. For example, if you have several EC2 instances running in a
	// specific Availability Zone, then you might want the cluster to be provisioned in
	// the same zone in order to decrease network latency. Default: A random,
	// system-chosen Availability Zone in the region that is specified by the endpoint.
	// Example: us-east-2d Constraint: The specified Availability Zone must be in the
	// same region as the current endpoint.
	AvailabilityZone *string

	// The weekly time range (in UTC) during which automated cluster maintenance can
	// occur. Format: ddd:hh24:mi-ddd:hh24:mi Default: A 30-minute window selected at
	// random from an 8-hour block of time per region, occurring on a random day of the
	// week. For more information about the time blocks for each region, see
	// Maintenance Windows
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-maintenance-windows)
	// in Amazon Redshift Cluster Management Guide. Valid Days: Mon | Tue | Wed | Thu |
	// Fri | Sat | Sun Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// Reserved.
	AdditionalInfo *string

	// The name of the first database to be created when the cluster is created. To
	// create additional databases after the cluster is created, connect to the cluster
	// with a SQL client and use SQL commands to create a database. For more
	// information, go to Create a Database
	// (https://docs.aws.amazon.com/redshift/latest/dg/t_creating_database.html) in the
	// Amazon Redshift Database Developer Guide. Default: dev Constraints:
	//
	//     * Must
	// contain 1 to 64 alphanumeric characters.
	//
	//     * Must contain only lowercase
	// letters.
	//
	//     * Cannot be a word that is reserved by the service. A list of
	// reserved words can be found in Reserved Words
	// (https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html) in the
	// Amazon Redshift Database Developer Guide.
	DBName *string

	// If true, the cluster can be accessed from a public network.
	PubliclyAccessible *bool

	// A list of tag instances.
	Tags []*types.Tag

	// The password associated with the master user account for the cluster that is
	// being created. Constraints:
	//
	//     * Must be between 8 and 64 characters in
	// length.
	//
	//     * Must contain at least one uppercase letter.
	//
	//     * Must contain
	// at least one lowercase letter.
	//
	//     * Must contain one number.
	//
	//     * Can be any
	// printable ASCII character (ASCII code 33 to 126) except ' (single quote), "
	// (double quote), \, /, @, or space.
	//
	// This member is required.
	MasterUserPassword *string

	// A list of Virtual Private Cloud (VPC) security groups to be associated with the
	// cluster. Default: The default VPC security group is associated with the cluster.
	VpcSecurityGroupIds []*string

	// An optional parameter for the name of the maintenance track for the cluster. If
	// you don't provide a maintenance track name, the cluster is assigned to the
	// current track.
	MaintenanceTrackName *string

	// The node type to be provisioned for the cluster. For information about node
	// types, go to  Working with Clusters
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#how-many-nodes)
	// in the Amazon Redshift Cluster Management Guide. Valid Values: ds2.xlarge |
	// ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.4xlarge |
	// ra3.16xlarge
	//
	// This member is required.
	NodeType *string

	// A unique identifier for the snapshot schedule.
	SnapshotScheduleIdentifier *string

	// Specifies the name of the HSM configuration that contains the information the
	// Amazon Redshift cluster can use to retrieve and store keys in an HSM.
	HsmConfigurationIdentifier *string

	// The version of the Amazon Redshift engine software that you want to deploy on
	// the cluster. The version selected runs on all the nodes in the cluster.
	// Constraints: Only version 1.0 is currently available. Example: 1.0
	ClusterVersion *string

	// The type of the cluster. When cluster type is specified as
	//
	//     * single-node,
	// the NumberOfNodes parameter is not required.
	//
	//     * multi-node, the
	// NumberOfNodes parameter is required.
	//
	// Valid Values: multi-node | single-node
	// Default: multi-node
	ClusterType *string

	// The port number on which the cluster accepts incoming connections. The cluster
	// is accessible only via the JDBC and ODBC connection strings. Part of the
	// connection string requires the port on which the cluster will listen for
	// incoming connections. Default: 5439 Valid Values: 1150-65535
	Port *int32

	// The number of days that automated snapshots are retained. If the value is 0,
	// automated snapshots are disabled. Even if automated snapshots are disabled, you
	// can still create manual snapshots when you want with CreateClusterSnapshot ().
	// Default: 1 Constraints: Must be a value from 0 to 35.
	AutomatedSnapshotRetentionPeriod *int32

	// An option that specifies whether to create the cluster with enhanced VPC routing
	// enabled. To create a cluster that uses enhanced VPC routing, the cluster must be
	// in a VPC. For more information, see Enhanced VPC Routing
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-routing.html) in
	// the Amazon Redshift Cluster Management Guide. If this option is true, enhanced
	// VPC routing is enabled. Default: false
	EnhancedVpcRouting *bool

	// The user name associated with the master user account for the cluster that is
	// being created. Constraints:
	//
	//     * Must be 1 - 128 alphanumeric characters. The
	// user name can't be PUBLIC.
	//
	//     * First character must be a letter.
	//
	//     *
	// <p>Cannot be a reserved word. A list of reserved words can be found in <a
	// href="https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html">Reserved
	// Words</a> in the Amazon Redshift Database Developer Guide. </p> </li> </ul>
	//
	// This member is required.
	MasterUsername *string

	// The number of compute nodes in the cluster. This parameter is required when the
	// ClusterType parameter is specified as multi-node. For information about
	// determining how many nodes you need, go to  Working with Clusters
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#how-many-nodes)
	// in the Amazon Redshift Cluster Management Guide. If you don't specify this
	// parameter, you get a single-node cluster. When requesting a multi-node cluster,
	// you must specify the number of nodes that you want in the cluster. Default: 1
	// Constraints: Value must be at least 1 and no more than 100.
	NumberOfNodes *int32

	// The Elastic IP (EIP) address for the cluster. Constraints: The cluster must be
	// provisioned in EC2-VPC and publicly-accessible through an Internet gateway. For
	// more information about provisioning clusters in EC2-VPC, go to Supported
	// Platforms to Launch Your Cluster
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#cluster-platforms)
	// in the Amazon Redshift Cluster Management Guide.
	ElasticIp *string

	// If true, the data in the cluster is encrypted at rest. Default: false
	Encrypted *bool

	// Specifies the name of the HSM client certificate the Amazon Redshift cluster
	// uses to retrieve the data encryption keys stored in an HSM.
	HsmClientCertificateIdentifier *string

	// The default number of days to retain a manual snapshot. If the value is -1, the
	// snapshot is retained indefinitely. This setting doesn't change the retention
	// period of existing snapshots. The value must be either -1 or an integer between
	// 1 and 3,653.
	ManualSnapshotRetentionPeriod *int32

	// If true, major version upgrades can be applied during the maintenance window to
	// the Amazon Redshift engine that is running on the cluster. When a new major
	// version of the Amazon Redshift engine is released, you can request that the
	// service automatically apply upgrades during the maintenance window to the Amazon
	// Redshift engine that is running on your cluster. Default: true
	AllowVersionUpgrade *bool

	// The name of a cluster subnet group to be associated with this cluster. If this
	// parameter is not provided the resulting cluster will be deployed outside virtual
	// private cloud (VPC).
	ClusterSubnetGroupName *string

	// The name of the parameter group to be associated with this cluster. Default: The
	// default Amazon Redshift cluster parameter group. For information about the
	// default parameter group, go to Working with Amazon Redshift Parameter Groups
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
	// Constraints:
	//
	//     * Must be 1 to 255 alphanumeric characters or hyphens.
	//
	//     *
	// First character must be a letter.
	//
	//     * Cannot end with a hyphen or contain two
	// consecutive hyphens.
	ClusterParameterGroupName *string

	// A list of security groups to be associated with this cluster. Default: The
	// default cluster security group for Amazon Redshift.
	ClusterSecurityGroups []*string

	// A list of AWS Identity and Access Management (IAM) roles that can be used by the
	// cluster to access other AWS services. You must supply the IAM roles in their
	// Amazon Resource Name (ARN) format. You can supply up to 10 IAM roles in a single
	// request. A cluster can have up to 10 IAM roles associated with it at any time.
	IamRoles []*string

	// A unique identifier for the cluster. You use this identifier to refer to the
	// cluster for any subsequent cluster operations such as deleting or modifying. The
	// identifier also appears in the Amazon Redshift console. Constraints:
	//
	//     * Must
	// contain from 1 to 63 alphanumeric characters or hyphens.
	//
	//     * Alphabetic
	// characters must be lowercase.
	//
	//     * First character must be a letter.
	//
	//     *
	// Cannot end with a hyphen or contain two consecutive hyphens.
	//
	//     * Must be
	// unique for all clusters within an AWS account.
	//
	//     <p>Example:
	// <code>myexamplecluster</code> </p>
	//
	// This member is required.
	ClusterIdentifier *string
}

type CreateClusterOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateCluster",
	}
}
