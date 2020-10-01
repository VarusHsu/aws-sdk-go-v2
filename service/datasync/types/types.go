// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Represents a single entry in a list of agents. AgentListEntry returns an array
// that contains a list of agents when the ListAgents () operation is called.
type AgentListEntry struct {

	// The name of the agent.
	Name *string

	// The Amazon Resource Name (ARN) of the agent.
	AgentArn *string

	// The status of the agent.
	Status AgentStatus
}

// The subnet and the security group that DataSync uses to access target EFS file
// system. The subnet must have at least one mount target for that file system. The
// security group that you provide needs to be able to communicate with the
// security group on the mount target in the subnet specified.
type Ec2Config struct {

	// The Amazon Resource Names (ARNs) of the security groups that are configured for
	// the Amazon EC2 resource.
	//
	// This member is required.
	SecurityGroupArns []*string

	// The ARN of the subnet and the security group that DataSync uses to access the
	// target EFS file system.
	//
	// This member is required.
	SubnetArn *string
}

// Specifies which files, folders and objects to include or exclude when
// transferring files from source to destination.
type FilterRule struct {

	// A single filter string that consists of the patterns to include or exclude. The
	// patterns are delimited by "|" (that is, a pipe), for example: /folder1|/folder2
	Value *string

	// The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN
	// rule type.
	FilterType FilterType
}

// Represents a single entry in a list of locations. LocationListEntry returns an
// array that contains a list of locations when the ListLocations () operation is
// called.
type LocationListEntry struct {

	// The Amazon Resource Name (ARN) of the location. For Network File System (NFS) or
	// Amazon EFS, the location is the export path. For Amazon S3, the location is the
	// prefix path that you want to mount and use as the root of the location.
	LocationArn *string

	// Represents a list of URLs of a location. LocationUri returns an array that
	// contains a list of locations when the ListLocations () operation is called.
	// Format: TYPE://GLOBAL_ID/SUBDIR. TYPE designates the type of location. Valid
	// values: NFS | EFS | S3. GLOBAL_ID is the globally unique identifier of the
	// resource that backs the location. An example for EFS is us-east-2.fs-abcd1234.
	// An example for Amazon S3 is the bucket name, such as myBucket. An example for
	// NFS is a valid IPv4 address or a host name compliant with Domain Name Service
	// (DNS). SUBDIR is a valid file system path, delimited by forward slashes as is
	// the *nix convention. For NFS and Amazon EFS, it's the export path to mount the
	// location. For Amazon S3, it's the prefix path that you mount to and treat as the
	// root of the location.
	LocationUri *string
}

// Represents the mount options that are available for DataSync to access an NFS
// location.
type NfsMountOptions struct {

	// The specific NFS version that you want DataSync to use to mount your NFS share.
	// If the server refuses to use the version specified, the sync will fail. If you
	// don't specify a version, DataSync defaults to AUTOMATIC. That is, DataSync
	// automatically selects a version based on negotiation with the NFS server.
	// <p>You can specify the following NFS versions:</p> <ul> <li> <p> <b> <a
	// href="https://tools.ietf.org/html/rfc1813">NFSv3</a> </b> - stateless protocol
	// version that allows for asynchronous writes on the server.</p> </li> <li> <p>
	// <b> <a href="https://tools.ietf.org/html/rfc3530">NFSv4.0</a> </b> - stateful,
	// firewall-friendly protocol version that supports delegations and pseudo
	// filesystems.</p> </li> <li> <p> <b> <a
	// href="https://tools.ietf.org/html/rfc5661">NFSv4.1</a> </b> - stateful protocol
	// version that supports sessions, directory delegations, and parallel data
	// processing. Version 4.1 also includes all features available in version 4.0.</p>
	// </li> </ul>
	Version NfsVersion
}

// A list of Amazon Resource Names (ARNs) of agents to use for a Network File
// System (NFS) location.
type OnPremConfig struct {

	// ARNs)of the agents to use for an NFS location.
	//
	// This member is required.
	AgentArns []*string
}

// Represents the options that are available to control the behavior of a
// StartTaskExecution () operation. Behavior includes preserving metadata such as
// user ID (UID), group ID (GID), and file permissions, and also overwriting files
// in the destination, data integrity verification, and so on. A task has a set of
// default options associated with it. If you don't specify an option in
// StartTaskExecution (), the default value is used. You can override the defaults
// options on each task execution by specifying an overriding Options value to
// StartTaskExecution ().
type Options struct {

	// A value that determines whether a data integrity verification should be
	// performed at the end of a task execution after all data and metadata have been
	// transferred. For more information, see create-task () Default value:
	// POINT_IN_TIME_CONSISTENT. ONLY_FILES_TRANSFERRED (recommended): Perform
	// verification only on files that were transferred.  <p>POINT_IN_TIME_CONSISTENT:
	// Scan the entire source and entire destination at the end of the transfer to
	// verify that source and destination are fully synchronized. This option isn't
	// supported when transferring to S3 Glacier or S3 Glacier Deep Archive storage
	// classes.</p> <p>NONE: No additional verification is done at the end of the
	// transfer, but all data transmissions are integrity-checked with checksum
	// verification during the transfer.</p>
	VerifyMode VerifyMode

	// The group ID (GID) of the file's owners. Default value: INT_VALUE. This
	// preserves the integer value of the ID. INT_VALUE: Preserve the integer value of
	// user ID (UID) and GID (recommended). NONE: Ignore UID and GID.
	Gid Gid

	// A value that specifies whether files in the destination that don't exist in the
	// source file system should be preserved. This option can affect your storage
	// cost. If your task deletes objects, you might incur minimum storage duration
	// charges for certain storage classes. For detailed information, see
	// using-storage-classes () in the AWS DataSync User Guide. Default value:
	// PRESERVE. PRESERVE: Ignore such destination files (recommended). REMOVE: Delete
	// destination files that aren’t present in the source.
	PreserveDeletedFiles PreserveDeletedFiles

	// A value that determines whether tasks should be queued before executing the
	// tasks. If set to ENABLED, the tasks will be queued. The default is ENABLED. If
	// you use the same agent to run multiple tasks, you can enable the tasks to run in
	// series. For more information, see queue-task-execution ().
	TaskQueueing TaskQueueing

	// A value that determines which users or groups can access a file for a specific
	// purpose such as reading, writing, or execution of the file. Default value:
	// PRESERVE. PRESERVE: Preserve POSIX-style permissions (recommended). NONE: Ignore
	// permissions. AWS DataSync can preserve extant permissions of a source location.
	PosixPermissions PosixPermissions

	// TransferMode has two values: CHANGED and ALL. CHANGED performs an "incremental"
	// or "delta sync", it compares file modification time between source and
	// destination to determine which files need to be transferred. ALL skips
	// destination inventory and transfers all files discovered on the source.
	TransferMode TransferMode

	// A value that limits the bandwidth used by AWS DataSync. For example, if you want
	// AWS DataSync to use a maximum of 1 MB, set this value to 1048576 (=1024*1024).
	BytesPerSecond *int64

	// A value that determines whether files at the destination should be overwritten
	// or preserved when copying files. If set to NEVER a destination file will not be
	// replaced by a source file, even if the destination file differs from the source
	// file. If you modify files in the destination and you sync the files, you can use
	// this value to protect against overwriting those changes. Some storage classes
	// have specific behaviors that can affect your S3 storage cost. For detailed
	// information, see using-storage-classes () in the AWS DataSync User Guide.
	OverwriteMode OverwriteMode

	// A file metadata value that shows the last time a file was accessed (that is,
	// when the file was read or written to). If you set Atime to BEST_EFFORT, DataSync
	// attempts to preserve the original Atime attribute on all source files (that is,
	// the version before the PREPARING phase). However, Atime's behavior is not fully
	// standard across platforms, so AWS DataSync can only do this on a best-effort
	// basis. Default value: BEST_EFFORT. BEST_EFFORT: Attempt to preserve the per-file
	// Atime value (recommended). NONE: Ignore Atime. If Atime is set to BEST_EFFORT,
	// Mtime must be set to PRESERVE. If Atime is set to NONE, Mtime must also be NONE.
	Atime Atime

	// The user ID (UID) of the file's owner. Default value: INT_VALUE. This preserves
	// the integer value of the ID. INT_VALUE: Preserve the integer value of UID and
	// group ID (GID) (recommended). NONE: Ignore UID and GID.
	Uid Uid

	// A value that determines the type of logs that DataSync publishes to a log stream
	// in the Amazon CloudWatch log group that you provide. For more information about
	// providing a log group for DataSync, see CloudWatchLogGroupArn
	// (https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateTask.html#DataSync-CreateTask-request-CloudWatchLogGroupArn).
	// If set to OFF, no logs are published. BASIC publishes logs on errors for
	// individual files transferred, and TRANSFER publishes logs for every file or
	// object that is transferred and integrity checked.
	LogLevel LogLevel

	// A value that indicates the last time that a file was modified (that is, a file
	// was written to) before the PREPARING phase. Default value: PRESERVE. PRESERVE:
	// Preserve original Mtime (recommended) NONE: Ignore Mtime. If Mtime is set to
	// PRESERVE, Atime must be set to BEST_EFFORT. If Mtime is set to NONE, Atime must
	// also be set to NONE.
	Mtime Mtime

	// A value that determines whether AWS DataSync should preserve the metadata of
	// block and character devices in the source file system, and recreate the files
	// with that device name and metadata on the destination. AWS DataSync can't sync
	// the actual contents of such devices, because they are nonterminal and don't
	// return an end-of-file (EOF) marker. Default value: NONE. NONE: Ignore special
	// devices (recommended). PRESERVE: Preserve character and block device metadata.
	// This option isn't currently supported for Amazon EFS.
	PreserveDevices PreserveDevices
}

// The VPC endpoint, subnet, and security group that an agent uses to access IP
// addresses in a VPC (Virtual Private Cloud).
type PrivateLinkConfig struct {

	// The ID of the VPC endpoint that is configured for an agent. An agent that is
	// configured with a VPC endpoint will not be accessible over the public internet.
	VpcEndpointId *string

	// The Amazon Resource Names (ARNs) of the subnets that are configured for an agent
	// activated in a VPC or an agent that has access to a VPC endpoint.
	SubnetArns []*string

	// The private endpoint that is configured for an agent that has access to IP
	// addresses in a PrivateLink
	// (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-service.html). An
	// agent that is configured with this endpoint will not be accessible over the
	// public internet.
	PrivateLinkEndpoint *string

	// The Amazon Resource Names (ARNs) of the security groups that are configured for
	// the EC2 resource that hosts an agent activated in a VPC or an agent that has
	// access to a VPC endpoint.
	SecurityGroupArns []*string
}

// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
// role that is used to access an Amazon S3 bucket.  <p>For detailed information
// about using such a role, see Creating a Location for Amazon S3 in the <i>AWS
// DataSync User Guide</i>.</p>
type S3Config struct {

	// The Amazon S3 bucket to access. This bucket is used as a parameter in the
	// CreateLocationS3 () operation.
	//
	// This member is required.
	BucketAccessRoleArn *string
}

// Represents the mount options that are available for DataSync to access an SMB
// location.
type SmbMountOptions struct {

	// The specific SMB version that you want DataSync to use to mount your SMB share.
	// If you don't specify a version, DataSync defaults to AUTOMATIC. That is,
	// DataSync automatically selects a version based on negotiation with the SMB
	// server.
	Version SmbVersion
}

// Represents a single entry in a list of AWS resource tags. TagListEntry returns
// an array that contains a list of tasks when the ListTagsForResource () operation
// is called.
type TagListEntry struct {

	// The value for an AWS resource tag.
	Value *string

	// The key for an AWS resource tag.
	//
	// This member is required.
	Key *string
}

// Represents a single entry in a list of task executions. TaskExecutionListEntry
// returns an array that contains a list of specific invocations of a task when
// ListTaskExecutions () operation is called.
type TaskExecutionListEntry struct {

	// The Amazon Resource Name (ARN) of the task that was executed.
	TaskExecutionArn *string

	// The status of a task execution.
	Status TaskExecutionStatus
}

// Describes the detailed result of a TaskExecution operation. This result includes
// the time in milliseconds spent in each phase, the status of the task execution,
// and the errors encountered.
type TaskExecutionResultDetail struct {

	// The total time in milliseconds that AWS DataSync spent in the PREPARING phase.
	PrepareDuration *int64

	// The status of the PREPARING phase.
	PrepareStatus PhaseStatus

	// The total time in milliseconds that AWS DataSync took to transfer the file from
	// the source to the destination location.
	TotalDuration *int64

	// Errors that AWS DataSync encountered during execution of the task. You can use
	// this error code to help troubleshoot issues.
	ErrorCode *string

	// The total time in milliseconds that AWS DataSync spent in the VERIFYING phase.
	VerifyDuration *int64

	// Detailed description of an error that was encountered during the task execution.
	// You can use this information to help troubleshoot issues.
	ErrorDetail *string

	// The total time in milliseconds that AWS DataSync spent in the TRANSFERRING
	// phase.
	TransferDuration *int64

	// The status of the VERIFYING Phase.
	VerifyStatus PhaseStatus

	// The status of the TRANSFERRING Phase.
	TransferStatus PhaseStatus
}

// Represents a single entry in a list of tasks. TaskListEntry returns an array
// that contains a list of tasks when the ListTasks () operation is called. A task
// includes the source and destination file systems to sync and the options to use
// for the tasks.
type TaskListEntry struct {

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string

	// The name of the task.
	Name *string

	// The status of the task.
	Status TaskStatus
}

// Specifies the schedule you want your task to use for repeated executions. For
// more information, see Schedule Expressions for Rules
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
type TaskSchedule struct {

	// A cron expression that specifies when AWS DataSync initiates a scheduled
	// transfer from a source to a destination location.
	//
	// This member is required.
	ScheduleExpression *string
}
