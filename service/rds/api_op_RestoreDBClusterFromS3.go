// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon Aurora DB cluster from data stored in an Amazon S3 bucket.
// Amazon RDS must be authorized to access the Amazon S3 bucket and the data must
// be created using the Percona XtraBackup utility as described in  Migrating Data
// to an Amazon Aurora MySQL DB Cluster
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Migrating.html)
// in the Amazon Aurora User Guide. This action only restores the DB cluster, not
// the DB instances for that DB cluster. You must invoke the CreateDBInstance
// action to create DB instances for the restored DB cluster, specifying the
// identifier of the restored DB cluster in DBClusterIdentifier. You can create DB
// instances only after the RestoreDBClusterFromS3 action has completed and the DB
// cluster is available. For more information on Amazon Aurora, see  What Is Amazon
// Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) RestoreDBClusterFromS3(ctx context.Context, params *RestoreDBClusterFromS3Input, optFns ...func(*Options)) (*RestoreDBClusterFromS3Output, error) {
	stack := middleware.NewStack("RestoreDBClusterFromS3", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRestoreDBClusterFromS3Middlewares(stack)
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
	addOpRestoreDBClusterFromS3ValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterFromS3(options.Region), middleware.Before)
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
			OperationName: "RestoreDBClusterFromS3",
			Err:           err,
		}
	}
	out := result.(*RestoreDBClusterFromS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreDBClusterFromS3Input struct {

	// The target backtrack window, in seconds. To disable backtracking, set this value
	// to 0. Currently, Backtrack is only supported for Aurora MySQL DB clusters.
	// Default: 0 Constraints:
	//
	//     * If specified, this value must be set to a number
	// from 0 to 259,200 (72 hours).
	BacktrackWindow *int64

	// A DB subnet group to associate with the restored DB cluster. Constraints: If
	// supplied, must match the name of an existing DBSubnetGroup. Example:
	// mySubnetgroup
	DBSubnetGroupName *string

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that authorizes Amazon RDS to access the Amazon S3 bucket on your behalf.
	//
	// This member is required.
	S3IngestionRoleArn *string

	// A list of EC2 VPC security groups to associate with the restored DB cluster.
	VpcSecurityGroupIds []*string

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region. To see the time blocks available, see  Adjusting the Preferred
	// Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Constraints:
	//
	//     * Must be in the format
	// hh24:mi-hh24:mi.
	//
	//     * Must be in Universal Coordinated Time (UTC).
	//
	//     * Must
	// not conflict with the preferred maintenance window.
	//
	//     * Must be at least 30
	// minutes.
	PreferredBackupWindow *string

	// The version number of the database engine to use. To list all of the available
	// engine versions for aurora (for MySQL 5.6-compatible Aurora), use the following
	// command: aws rds describe-db-engine-versions --engine aurora --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for aurora-mysql (for MySQL 5.7-compatible Aurora), use the following command:
	// aws rds describe-db-engine-versions --engine aurora-mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for aurora-postgresql, use the following command: aws rds
	// describe-db-engine-versions --engine aurora-postgresql --query
	// "DBEngineVersions[].EngineVersion" Aurora MySQL Example: 5.6.10a,
	// 5.6.mysql_aurora.1.19.2, 5.7.12, 5.7.mysql_aurora.2.04.5 Aurora PostgreSQL
	// Example: 9.6.3, 10.7
	EngineVersion *string

	// The name of the DB cluster to create from the source data in the Amazon S3
	// bucket. This parameter isn't case-sensitive. Constraints:
	//
	//     * Must contain
	// from 1 to 63 letters, numbers, or hyphens.
	//
	//     * First character must be a
	// letter.
	//
	//     * Can't end with a hyphen or contain two consecutive
	// hyphens.
	//
	// Example: my-cluster1
	//
	// This member is required.
	DBClusterIdentifier *string

	// The identifier for the database engine that was backed up to create the files
	// stored in the Amazon S3 bucket. Valid values: mysql
	//
	// This member is required.
	SourceEngine *string

	// A value that indicates whether to copy all tags from the restored DB cluster to
	// snapshots of the restored DB cluster. The default is not to copy them.
	CopyTagsToSnapshot *bool

	// The version of the database that the backup files were created from. MySQL
	// versions 5.5, 5.6, and 5.7 are supported. Example: 5.6.40
	//
	// This member is required.
	SourceEngineVersion *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []*types.Tag

	// The name of the database engine to be used for the restored DB cluster. Valid
	// Values: aurora, aurora-postgresql
	//
	// This member is required.
	Engine *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string

	// Specify the Active Directory directory ID to restore the DB cluster in. The
	// domain must be created prior to this operation. For Amazon Aurora DB clusters,
	// Amazon RDS can use Kerberos Authentication to authenticate users that connect to
	// the DB cluster. For more information, see Kerberos Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html)
	// in the Amazon Aurora User Guide.
	Domain *string

	// A value that indicates whether the restored DB cluster is encrypted.
	StorageEncrypted *bool

	// A value that indicates that the restored DB cluster should be associated with
	// the specified option group. Permanent options can't be removed from an option
	// group. An option group can't be removed from a DB cluster once it is associated
	// with a DB cluster.
	OptionGroupName *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region, occurring on a random day of the week. To see the time blocks available,
	// see  Adjusting the Preferred Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// The port number on which the instances in the restored DB cluster accept
	// connections. Default: 3306
	Port *int32

	// The AWS KMS key identifier for an encrypted DB cluster. The KMS key identifier
	// is the Amazon Resource Name (ARN) for the KMS encryption key. If you are
	// creating a DB cluster with the same AWS account that owns the KMS encryption key
	// used to encrypt the new DB cluster, then you can use the KMS key alias instead
	// of the ARN for the KM encryption key. If the StorageEncrypted parameter is
	// enabled, and you do not specify a value for the KmsKeyId parameter, then Amazon
	// RDS will use your default encryption key. AWS KMS creates the default encryption
	// key for your AWS account. Your AWS account has a different default encryption
	// key for each AWS Region.
	KmsKeyId *string

	// The prefix for all of the file names that contain the data used to create the
	// Amazon Aurora DB cluster. If you do not specify a SourceS3Prefix value, then the
	// Amazon Aurora DB cluster is created by using all of the files in the Amazon S3
	// bucket.
	S3Prefix *string

	// A value that indicates that the restored DB cluster should be associated with
	// the specified CharacterSet.
	CharacterSetName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool

	// The database name for the restored DB cluster.
	DatabaseName *string

	// The name of the master user for the restored DB cluster. Constraints:
	//
	//     *
	// Must be 1 to 16 letters or numbers.
	//
	//     * First character must be a letter.
	//
	//
	// * Can't be a reserved word for the chosen database engine.
	//
	// This member is required.
	MasterUsername *string

	// A list of Availability Zones (AZs) where instances in the restored DB cluster
	// can be created.
	AvailabilityZones []*string

	// The name of the DB cluster parameter group to associate with the restored DB
	// cluster. If this argument is omitted, default.aurora5.6 is used. Constraints:
	//
	//
	// * If supplied, must match the name of an existing DBClusterParameterGroup.
	DBClusterParameterGroupName *string

	// The name of the Amazon S3 bucket that contains the data used to create the
	// Amazon Aurora DB cluster.
	//
	// This member is required.
	S3BucketName *string

	// The number of days for which automated backups of the restored DB cluster are
	// retained. You must specify a minimum value of 1. Default: 1 Constraints:
	//
	//     *
	// Must be a value from 1 to 35
	BackupRetentionPeriod *int32

	// The password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@". Constraints: Must contain
	// from 8 to 41 characters.
	//
	// This member is required.
	MasterUserPassword *string

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// <p>For more information, see <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html">
	// IAM Database Authentication</a> in the <i>Amazon Aurora User Guide.</i> </p>
	EnableIAMDatabaseAuthentication *bool

	// The list of logs that the restored DB cluster is to export to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. For more information,
	// see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide.
	EnableCloudwatchLogsExports []*string
}

type RestoreDBClusterFromS3Output struct {

	// Contains the details of an Amazon Aurora DB cluster. This data type is used as a
	// response element in the DescribeDBClusters, StopDBCluster, and StartDBCluster
	// actions.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRestoreDBClusterFromS3Middlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterFromS3{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterFromS3{}, middleware.After)
}

func newServiceMetadataMiddleware_opRestoreDBClusterFromS3(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RestoreDBClusterFromS3",
	}
}
