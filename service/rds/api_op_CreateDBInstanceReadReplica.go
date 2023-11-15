// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new DB instance that acts as a read replica for an existing source DB
// instance or Multi-AZ DB cluster. You can create a read replica for a DB instance
// running MySQL, MariaDB, Oracle, PostgreSQL, or SQL Server. You can create a read
// replica for a Multi-AZ DB cluster running MySQL or PostgreSQL. For more
// information, see Working with read replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.html)
// and Migrating from a Multi-AZ DB cluster to a DB instance using a read replica (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html#multi-az-db-clusters-migrating-to-instance-with-read-replica)
// in the Amazon RDS User Guide. Amazon Aurora doesn't support this operation. To
// create a DB instance for an Aurora DB cluster, use the CreateDBInstance
// operation. All read replica DB instances are created with backups disabled. All
// other attributes (including DB security groups and DB parameter groups) are
// inherited from the source DB instance or cluster, except as specified. Your
// source DB instance or cluster must have backup retention enabled.
func (c *Client) CreateDBInstanceReadReplica(ctx context.Context, params *CreateDBInstanceReadReplicaInput, optFns ...func(*Options)) (*CreateDBInstanceReadReplicaOutput, error) {
	if params == nil {
		params = &CreateDBInstanceReadReplicaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBInstanceReadReplica", params, optFns, c.addOperationCreateDBInstanceReadReplicaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBInstanceReadReplicaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBInstanceReadReplicaInput struct {

	// The DB instance identifier of the read replica. This identifier is the unique
	// key that identifies a DB instance. This parameter is stored as a lowercase
	// string.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The amount of storage (in gibibytes) to allocate initially for the read
	// replica. Follow the allocation rules specified in CreateDBInstance . Be sure to
	// allocate enough storage for your read replica so that the create operation can
	// succeed. You can also allocate additional storage for future growth.
	AllocatedStorage *int32

	// Specifies whether to automatically apply minor engine upgrades to the read
	// replica during the maintenance window. This setting doesn't apply to RDS Custom
	// DB instances. Default: Inherits the value from the source DB instance.
	AutoMinorVersionUpgrade *bool

	// The Availability Zone (AZ) where the read replica will be created. Default: A
	// random, system-chosen Availability Zone in the endpoint's Amazon Web Services
	// Region. Example: us-east-1d
	AvailabilityZone *string

	// Specifies whether to copy all tags from the read replica to snapshots of the
	// read replica. By default, tags aren't copied.
	CopyTagsToSnapshot *bool

	// The instance profile associated with the underlying Amazon EC2 instance of an
	// RDS Custom DB instance. The instance profile must meet the following
	// requirements:
	//   - The profile must exist in your account.
	//   - The profile must have an IAM role that Amazon EC2 has permissions to
	//   assume.
	//   - The instance profile name and the associated IAM role name must start with
	//   the prefix AWSRDSCustom .
	// For the list of permissions required for the IAM role, see  Configure IAM and
	// your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc)
	// in the Amazon RDS User Guide. This setting is required for RDS Custom DB
	// instances.
	CustomIamInstanceProfile *string

	// The compute and memory capacity of the read replica, for example db.m4.large.
	// Not all DB instance classes are available in all Amazon Web Services Regions, or
	// for all database engines. For the full list of DB instance classes, and
	// availability for your engine, see DB Instance Class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. Default: Inherits the value from the source DB
	// instance.
	DBInstanceClass *string

	// The name of the DB parameter group to associate with this DB instance. If you
	// don't specify a value for DBParameterGroupName , then Amazon RDS uses the
	// DBParameterGroup of the source DB instance for a same Region read replica, or
	// the default DBParameterGroup for the specified DB engine for a cross-Region
	// read replica. Specifying a parameter group for this operation is only supported
	// for MySQL DB instances for cross-Region read replicas and for Oracle DB
	// instances. It isn't supported for MySQL DB instances for same Region read
	// replicas or for RDS Custom. Constraints:
	//   - Must be 1 to 255 letters, numbers, or hyphens.
	//   - First character must be a letter.
	//   - Can't end with a hyphen or contain two consecutive hyphens.
	DBParameterGroupName *string

	// A DB subnet group for the DB instance. The new DB instance is created in the
	// VPC associated with the DB subnet group. If no DB subnet group is specified,
	// then the new DB instance isn't created in a VPC. Constraints:
	//   - If supplied, must match the name of an existing DB subnet group.
	//   - The specified DB subnet group must be in the same Amazon Web Services
	//   Region in which the operation is running.
	//   - All read replicas in one Amazon Web Services Region that are created from
	//   the same source DB instance must either:
	//   - Specify DB subnet groups from the same VPC. All these read replicas are
	//   created in the same VPC.
	//   - Not specify a DB subnet group. All these read replicas are created outside
	//   of any VPC.
	// Example: mydbsubnetgroup
	DBSubnetGroupName *string

	// Indicates whether the DB instance has a dedicated log volume (DLV) enabled.
	DedicatedLogVolume *bool

	// Specifies whether to enable deletion protection for the DB instance. The
	// database can't be deleted when deletion protection is enabled. By default,
	// deletion protection isn't enabled. For more information, see Deleting a DB
	// Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html)
	// .
	DeletionProtection *bool

	// The Active Directory directory ID to create the DB instance in. Currently, only
	// MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created
	// in an Active Directory Domain. For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html)
	// in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom DB
	// instances.
	Domain *string

	// The ARN for the Secrets Manager secret with the credentials for the user
	// joining the domain. Example:
	// arn:aws:secretsmanager:region:account-number:secret:myselfmanagedADtestsecret-123456
	DomainAuthSecretArn *string

	// The IPv4 DNS IP addresses of your primary and secondary Active Directory domain
	// controllers. Constraints:
	//   - Two IP addresses must be provided. If there isn't a secondary domain
	//   controller, use the IP address of the primary domain controller for both entries
	//   in the list.
	// Example: 123.124.125.126,234.235.236.237
	DomainDnsIps []string

	// The fully qualified domain name (FQDN) of an Active Directory domain.
	// Constraints:
	//   - Can't be longer than 64 characters.
	// Example: mymanagedADtest.mymanagedAD.mydomain
	DomainFqdn *string

	// The name of the IAM role to use when making API calls to the Directory Service.
	// This setting doesn't apply to RDS Custom DB instances.
	DomainIAMRoleName *string

	// The Active Directory organizational unit for your DB instance to join.
	// Constraints:
	//   - Must be in the distinguished name format.
	//   - Can't be longer than 64 characters.
	// Example: OU=mymanagedADtestOU,DC=mymanagedADtest,DC=mymanagedAD,DC=mydomain
	DomainOu *string

	// The list of logs that the new DB instance is to export to CloudWatch Logs. The
	// values in the list depend on the DB engine being used. For more information, see
	// Publishing Database Logs to Amazon CloudWatch Logs  (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom DB
	// instances.
	EnableCloudwatchLogsExports []string

	// Specifies whether to enable a customer-owned IP address (CoIP) for an RDS on
	// Outposts read replica. A CoIP provides local or external connectivity to
	// resources in your Outpost subnets through your on-premises network. For some use
	// cases, a CoIP can provide lower latency for connections to the read replica from
	// outside of its virtual private cloud (VPC) on your local network. For more
	// information about RDS on Outposts, see Working with Amazon RDS on Amazon Web
	// Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html)
	// in the Amazon RDS User Guide. For more information about CoIPs, see
	// Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing)
	// in the Amazon Web Services Outposts User Guide.
	EnableCustomerOwnedIp *bool

	// Specifies whether to enable mapping of Amazon Web Services Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping isn't
	// enabled. For more information about IAM database authentication, see IAM
	// Database Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom DB
	// instances.
	EnableIAMDatabaseAuthentication *bool

	// Specifies whether to enable Performance Insights for the read replica. For more
	// information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom DB
	// instances.
	EnablePerformanceInsights *bool

	// The amount of Provisioned IOPS (input/output operations per second) to
	// initially allocate for the DB instance.
	Iops *int32

	// The Amazon Web Services KMS key identifier for an encrypted read replica. The
	// Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or
	// alias name for the KMS key. If you create an encrypted read replica in the same
	// Amazon Web Services Region as the source DB instance or Multi-AZ DB cluster,
	// don't specify a value for this parameter. A read replica in the same Amazon Web
	// Services Region is always encrypted with the same KMS key as the source DB
	// instance or cluster. If you create an encrypted read replica in a different
	// Amazon Web Services Region, then you must specify a KMS key identifier for the
	// destination Amazon Web Services Region. KMS keys are specific to the Amazon Web
	// Services Region that they are created in, and you can't use KMS keys from one
	// Amazon Web Services Region in another Amazon Web Services Region. You can't
	// create an encrypted read replica from an unencrypted DB instance or Multi-AZ DB
	// cluster. This setting doesn't apply to RDS Custom, which uses the same KMS key
	// as the primary replica.
	KmsKeyId *string

	// The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale
	// the storage of the DB instance. For more information about this setting,
	// including limitations that apply to it, see Managing capacity automatically
	// with Amazon RDS storage autoscaling (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling)
	// in the Amazon RDS User Guide.
	MaxAllocatedStorage *int32

	// The interval, in seconds, between points when Enhanced Monitoring metrics are
	// collected for the read replica. To disable collection of Enhanced Monitoring
	// metrics, specify 0 . The default is 0 . If MonitoringRoleArn is specified, then
	// you must set MonitoringInterval to a value other than 0 . This setting doesn't
	// apply to RDS Custom DB instances. Valid Values: 0, 1, 5, 10, 15, 30, 60
	// Default: 0
	MonitoringInterval *int32

	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics
	// to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess .
	// For information on creating a monitoring role, go to To create an IAM role for
	// Amazon RDS Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html#USER_Monitoring.OS.IAMRole)
	// in the Amazon RDS User Guide. If MonitoringInterval is set to a value other
	// than 0, then you must supply a MonitoringRoleArn value. This setting doesn't
	// apply to RDS Custom DB instances.
	MonitoringRoleArn *string

	// Specifies whether the read replica is in a Multi-AZ deployment. You can create
	// a read replica as a Multi-AZ DB instance. RDS creates a standby of your replica
	// in another Availability Zone for failover support for the replica. Creating your
	// read replica as a Multi-AZ DB instance is independent of whether the source is a
	// Multi-AZ DB instance or a Multi-AZ DB cluster. This setting doesn't apply to RDS
	// Custom DB instances.
	MultiAZ *bool

	// The network type of the DB instance. Valid Values:
	//   - IPV4
	//   - DUAL
	// The network type is determined by the DBSubnetGroup specified for read replica.
	// A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6
	// protocols ( DUAL ). For more information, see  Working with a DB instance in a
	// VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html)
	// in the Amazon RDS User Guide.
	NetworkType *string

	// The option group to associate the DB instance with. If not specified, RDS uses
	// the option group associated with the source DB instance or cluster. For SQL
	// Server, you must use the option group associated with the source. This setting
	// doesn't apply to RDS Custom DB instances.
	OptionGroupName *string

	// The Amazon Web Services KMS key identifier for encryption of Performance
	// Insights data. The Amazon Web Services KMS key identifier is the key ARN, key
	// ID, alias ARN, or alias name for the KMS key. If you do not specify a value for
	// PerformanceInsightsKMSKeyId , then Amazon RDS uses your default KMS key. There
	// is a default KMS key for your Amazon Web Services account. Your Amazon Web
	// Services account has a different default KMS key for each Amazon Web Services
	// Region. This setting doesn't apply to RDS Custom DB instances.
	PerformanceInsightsKMSKeyId *string

	// The number of days to retain Performance Insights data. This setting doesn't
	// apply to RDS Custom DB instances. Valid Values:
	//   - 7
	//   - month * 31, where month is a number of months from 1-23. Examples: 93 (3
	//   months * 31), 341 (11 months * 31), 589 (19 months * 31)
	//   - 731
	// Default: 7 days If you specify a retention period that isn't valid, such as 94 ,
	// Amazon RDS returns an error.
	PerformanceInsightsRetentionPeriod *int32

	// The port number that the DB instance uses for connections. Valid Values:
	// 1150-65535 Default: Inherits the value from the source DB instance.
	Port *int32

	// When you are creating a read replica from one Amazon Web Services GovCloud (US)
	// Region to another or from one China Amazon Web Services Region to another, the
	// URL that contains a Signature Version 4 signed request for the
	// CreateDBInstanceReadReplica API operation in the source Amazon Web Services
	// Region that contains the source DB instance. This setting applies only to Amazon
	// Web Services GovCloud (US) Regions and China Amazon Web Services Regions. It's
	// ignored in other Amazon Web Services Regions. This setting applies only when
	// replicating from a source DB instance. Source DB clusters aren't supported in
	// Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions.
	// You must specify this parameter when you create an encrypted read replica from
	// another Amazon Web Services Region by using the Amazon RDS API. Don't specify
	// PreSignedUrl when you are creating an encrypted read replica in the same Amazon
	// Web Services Region. The presigned URL must be a valid request for the
	// CreateDBInstanceReadReplica API operation that can run in the source Amazon Web
	// Services Region that contains the encrypted source DB instance. The presigned
	// URL request must contain the following parameter values:
	//   - DestinationRegion - The Amazon Web Services Region that the encrypted read
	//   replica is created in. This Amazon Web Services Region is the same one where the
	//   CreateDBInstanceReadReplica operation is called that contains this presigned
	//   URL. For example, if you create an encrypted DB instance in the us-west-1 Amazon
	//   Web Services Region, from a source DB instance in the us-east-2 Amazon Web
	//   Services Region, then you call the CreateDBInstanceReadReplica operation in
	//   the us-east-1 Amazon Web Services Region and provide a presigned URL that
	//   contains a call to the CreateDBInstanceReadReplica operation in the us-west-2
	//   Amazon Web Services Region. For this example, the DestinationRegion in the
	//   presigned URL must be set to the us-east-1 Amazon Web Services Region.
	//   - KmsKeyId - The KMS key identifier for the key to use to encrypt the read
	//   replica in the destination Amazon Web Services Region. This is the same
	//   identifier for both the CreateDBInstanceReadReplica operation that is called
	//   in the destination Amazon Web Services Region, and the operation contained in
	//   the presigned URL.
	//   - SourceDBInstanceIdentifier - The DB instance identifier for the encrypted DB
	//   instance to be replicated. This identifier must be in the Amazon Resource Name
	//   (ARN) format for the source Amazon Web Services Region. For example, if you are
	//   creating an encrypted read replica from a DB instance in the us-west-2 Amazon
	//   Web Services Region, then your SourceDBInstanceIdentifier looks like the
	//   following example:
	//   arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115 .
	// To learn how to generate a Signature Version 4 signed request, see
	// Authenticating Requests: Using Query Parameters (Amazon Web Services Signature
	// Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
	// . If you are using an Amazon Web Services SDK tool or the CLI, you can specify
	// SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl
	// manually. Specifying SourceRegion autogenerates a presigned URL that is a valid
	// request for the operation that can run in the source Amazon Web Services Region.
	// SourceRegion isn't supported for SQL Server, because Amazon RDS for SQL Server
	// doesn't support cross-Region read replicas. This setting doesn't apply to RDS
	// Custom DB instances.
	PreSignedUrl *string

	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance. This setting doesn't apply to RDS Custom DB instances.
	ProcessorFeatures []types.ProcessorFeature

	// Specifies whether the DB instance is publicly accessible. When the DB cluster
	// is publicly accessible, its Domain Name System (DNS) endpoint resolves to the
	// private IP address from within the DB cluster's virtual private cloud (VPC). It
	// resolves to the public IP address from outside of the DB cluster's VPC. Access
	// to the DB cluster is ultimately controlled by the security group it uses. That
	// public access isn't permitted if the security group assigned to the DB cluster
	// doesn't permit it. When the DB instance isn't publicly accessible, it is an
	// internal DB instance with a DNS name that resolves to a private IP address. For
	// more information, see CreateDBInstance .
	PubliclyAccessible *bool

	// The open mode of the replica database: mounted or read-only. This parameter is
	// only supported for Oracle DB instances. Mounted DB replicas are included in
	// Oracle Database Enterprise Edition. The main use case for mounted replicas is
	// cross-Region disaster recovery. The primary database doesn't use Active Data
	// Guard to transmit information to the mounted replica. Because it doesn't accept
	// user connections, a mounted replica can't serve a read-only workload. You can
	// create a combination of mounted and read-only DB replicas for the same primary
	// DB instance. For more information, see Working with Oracle Read Replicas for
	// Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html)
	// in the Amazon RDS User Guide. For RDS Custom, you must specify this parameter
	// and set it to mounted . The value won't be set by default. After replica
	// creation, you can manage the open mode manually.
	ReplicaMode types.ReplicaMode

	// The identifier of the Multi-AZ DB cluster that will act as the source for the
	// read replica. Each DB cluster can have up to 15 read replicas. Constraints:
	//   - Must be the identifier of an existing Multi-AZ DB cluster.
	//   - Can't be specified if the SourceDBInstanceIdentifier parameter is also
	//   specified.
	//   - The specified DB cluster must have automatic backups enabled, that is, its
	//   backup retention period must be greater than 0.
	//   - The source DB cluster must be in the same Amazon Web Services Region as the
	//   read replica. Cross-Region replication isn't supported.
	SourceDBClusterIdentifier *string

	// The identifier of the DB instance that will act as the source for the read
	// replica. Each DB instance can have up to 15 read replicas, with the exception of
	// Oracle and SQL Server, which can have up to five. Constraints:
	//   - Must be the identifier of an existing MySQL, MariaDB, Oracle, PostgreSQL,
	//   or SQL Server DB instance.
	//   - Can't be specified if the SourceDBClusterIdentifier parameter is also
	//   specified.
	//   - For the limitations of Oracle read replicas, see Version and licensing
	//   considerations for RDS for Oracle replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.limitations.html#oracle-read-replicas.limitations.versions-and-licenses)
	//   in the Amazon RDS User Guide.
	//   - For the limitations of SQL Server read replicas, see Read replica
	//   limitations with SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.html#SQLServer.ReadReplicas.Limitations)
	//   in the Amazon RDS User Guide.
	//   - The specified DB instance must have automatic backups enabled, that is, its
	//   backup retention period must be greater than 0.
	//   - If the source DB instance is in the same Amazon Web Services Region as the
	//   read replica, specify a valid DB instance identifier.
	//   - If the source DB instance is in a different Amazon Web Services Region from
	//   the read replica, specify a valid DB instance ARN. For more information, see
	//   Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)
	//   in the Amazon RDS User Guide. This doesn't apply to SQL Server or RDS Custom,
	//   which don't support cross-Region replicas.
	SourceDBInstanceIdentifier *string

	// The AWS region the resource is in. The presigned URL will be created with this
	// region, if the PresignURL member is empty set.
	SourceRegion *string

	// Specifies the storage throughput value for the read replica. This setting
	// doesn't apply to RDS Custom or Amazon Aurora DB instances.
	StorageThroughput *int32

	// The storage type to associate with the read replica. If you specify io1 or gp3 ,
	// you must also include a value for the Iops parameter. Valid Values: gp2 | gp3 |
	// io1 | standard Default: io1 if the Iops parameter is specified. Otherwise, gp2 .
	StorageType *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []types.Tag

	// Whether to upgrade the storage file system configuration on the read replica.
	// This option migrates the read replica from the old storage file system layout to
	// the preferred layout.
	UpgradeStorageConfig *bool

	// Specifies whether the DB instance class of the DB instance uses its default
	// processor features. This setting doesn't apply to RDS Custom DB instances.
	UseDefaultProcessorFeatures *bool

	// A list of Amazon EC2 VPC security groups to associate with the read replica.
	// This setting doesn't apply to RDS Custom DB instances. Default: The default EC2
	// VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string

	// Used by the SDK's PresignURL autofill customization to specify the region the
	// of the client's request.
	destinationRegion *string

	noSmithyDocumentSerde
}

type CreateDBInstanceReadReplicaOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the operations CreateDBInstance ,
	// CreateDBInstanceReadReplica , DeleteDBInstance , DescribeDBInstances ,
	// ModifyDBInstance , PromoteReadReplica , RebootDBInstance ,
	// RestoreDBInstanceFromDBSnapshot , RestoreDBInstanceFromS3 ,
	// RestoreDBInstanceToPointInTime , StartDBInstance , and StopDBInstance .
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBInstanceReadReplicaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBInstanceReadReplica{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBInstanceReadReplica{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDBInstanceReadReplica"); err != nil {
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
	if err = addCreateDBInstanceReadReplicaPresignURLMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateDBInstanceReadReplicaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBInstanceReadReplica(options.Region), middleware.Before); err != nil {
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

func copyCreateDBInstanceReadReplicaInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateDBInstanceReadReplicaInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateDBInstanceReadReplicaInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCreateDBInstanceReadReplicaPreSignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CreateDBInstanceReadReplicaInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CreateDBInstanceReadReplicaInput type, got %T", params)
	}
	if input.PreSignedUrl == nil || len(*input.PreSignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PreSignedUrl, true, nil
}
func getCreateDBInstanceReadReplicaSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CreateDBInstanceReadReplicaInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CreateDBInstanceReadReplicaInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCreateDBInstanceReadReplicaPreSignedUrl(params interface{}, value string) error {
	input, ok := params.(*CreateDBInstanceReadReplicaInput)
	if !ok {
		return fmt.Errorf("expect *CreateDBInstanceReadReplicaInput type, got %T", params)
	}
	input.PreSignedUrl = &value
	return nil
}
func setCreateDBInstanceReadReplicadestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CreateDBInstanceReadReplicaInput)
	if !ok {
		return fmt.Errorf("expect *CreateDBInstanceReadReplicaInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}
func addCreateDBInstanceReadReplicaPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL: getCreateDBInstanceReadReplicaPreSignedUrl,

			GetSourceRegion: getCreateDBInstanceReadReplicaSourceRegion,

			CopyInput: copyCreateDBInstanceReadReplicaInputForPresign,

			SetDestinationRegion: setCreateDBInstanceReadReplicadestinationRegion,

			SetPresignedURL: setCreateDBInstanceReadReplicaPreSignedUrl,
		},
		Presigner: &presignAutoFillCreateDBInstanceReadReplicaClient{client: NewPresignClient(New(options))},
	})
}

type presignAutoFillCreateDBInstanceReadReplicaClient struct {
	client *PresignClient
}

// PresignURL is a middleware accessor that satisfies URLPresigner interface.
func (c *presignAutoFillCreateDBInstanceReadReplicaClient) PresignURL(ctx context.Context, srcRegion string, params interface{}) (*v4.PresignedHTTPRequest, error) {
	input, ok := params.(*CreateDBInstanceReadReplicaInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateDBInstanceReadReplicaInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = srcRegion
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	presignOptFn := WithPresignClientFromClientOptions(optFn)
	return c.client.PresignCreateDBInstanceReadReplica(ctx, input, presignOptFn)
}

func newServiceMetadataMiddleware_opCreateDBInstanceReadReplica(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDBInstanceReadReplica",
	}
}

// PresignCreateDBInstanceReadReplica is used to generate a presigned HTTP Request
// which contains presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignCreateDBInstanceReadReplica(ctx context.Context, params *CreateDBInstanceReadReplicaInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil {
		params = &CreateDBInstanceReadReplicaInput{}
	}
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)

	result, _, err := c.client.invokeOperation(ctx, "CreateDBInstanceReadReplica", params, clientOptFns,
		c.client.addOperationCreateDBInstanceReadReplicaMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}
