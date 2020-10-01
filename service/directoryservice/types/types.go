// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Represents a named directory attribute.
type Attribute struct {

	// The name of the attribute.
	Name *string

	// The value of the attribute.
	Value *string
}

// Information about the certificate.
type Certificate struct {

	// The common name for the certificate.
	CommonName *string

	// The identifier of the certificate.
	CertificateId *string

	// The state of the certificate.
	State CertificateState

	// The date and time when the certificate will expire.
	ExpiryDateTime *time.Time

	// Describes a state change for the certificate.
	StateReason *string

	// The date and time that the certificate was registered.
	RegisteredDateTime *time.Time
}

// Contains general information about a certificate.
type CertificateInfo struct {

	// The date and time when the certificate will expire.
	ExpiryDateTime *time.Time

	// The identifier of the certificate.
	CertificateId *string

	// The common name for the certificate.
	CommonName *string

	// The state of the certificate.
	State CertificateState
}

// Contains information about a computer account in a directory.
type Computer struct {

	// The identifier of the computer.
	ComputerId *string

	// An array of Attribute () objects containing the LDAP attributes that belong to
	// the computer account.
	ComputerAttributes []*Attribute

	// The computer name.
	ComputerName *string
}

// Points to a remote domain with which you are setting up a trust relationship.
// Conditional forwarders are required in order to set up a trust relationship with
// another domain.
type ConditionalForwarder struct {

	// The IP addresses of the remote DNS server associated with RemoteDomainName. This
	// is the IP address of the DNS server that your conditional forwarder points to.
	DnsIpAddrs []*string

	// The fully qualified domain name (FQDN) of the remote domains pointed to by the
	// conditional forwarder.
	RemoteDomainName *string

	// The replication scope of the conditional forwarder. The only allowed value is
	// Domain, which will replicate the conditional forwarder to all of the domain
	// controllers for your AWS directory.
	ReplicationScope ReplicationScope
}

// Contains information for the ConnectDirectory () operation when an AD Connector
// directory is being created.
type DirectoryConnectSettings struct {

	// A list of subnet identifiers in the VPC in which the AD Connector is created.
	//
	// This member is required.
	SubnetIds []*string

	// A list of one or more IP addresses of DNS servers or domain controllers in the
	// on-premises directory.
	//
	// This member is required.
	CustomerDnsIps []*string

	// The user name of an account in the on-premises directory that is used to connect
	// to the directory. This account must have the following permissions:
	//
	//     * Read
	// users and groups
	//
	//     * Create computer objects
	//
	//     * Join computers to the
	// domain
	//
	// This member is required.
	CustomerUserName *string

	// The identifier of the VPC in which the AD Connector is created.
	//
	// This member is required.
	VpcId *string
}

// Contains information about an AD Connector directory.
type DirectoryConnectSettingsDescription struct {

	// The security group identifier for the AD Connector directory.
	SecurityGroupId *string

	// The IP addresses of the AD Connector servers.
	ConnectIps []*string

	// A list of subnet identifiers in the VPC that the AD Connector is in.
	SubnetIds []*string

	// The identifier of the VPC that the AD Connector is in.
	VpcId *string

	// The user name of the service account in the on-premises directory.
	CustomerUserName *string

	// A list of the Availability Zones that the directory is in.
	AvailabilityZones []*string
}

// Contains information about an AWS Directory Service directory.
type DirectoryDescription struct {

	// A DirectoryConnectSettingsDescription () object that contains additional
	// information about an AD Connector directory. This member is only present if the
	// directory is an AD Connector directory.
	ConnectSettings *DirectoryConnectSettingsDescription

	// A directory share request that is sent by the directory owner to the directory
	// consumer. The request includes a typed message to help the directory consumer
	// administrator determine whether to approve or reject the share invitation.
	ShareNotes *string

	// The current stage of the directory.
	Stage DirectoryStage

	// The method used when sharing a directory to determine whether the directory
	// should be shared within your AWS organization (ORGANIZATIONS) or with any AWS
	// account by sending a shared directory request (HANDSHAKE).
	ShareMethod ShareMethod

	// The directory size.
	Type DirectoryType

	// The status of the RADIUS MFA server connection.
	RadiusStatus RadiusStatus

	// The short name of the directory.
	ShortName *string

	// The alias for the directory. If no alias has been created for the directory, the
	// alias is the directory identifier, such as d-XXXXXXXXXX.
	Alias *string

	// Additional information about the directory stage.
	StageReason *string

	// The date and time that the stage was last updated.
	StageLastUpdatedDateTime *time.Time

	// Describes the AWS Managed Microsoft AD directory in the directory owner account.
	OwnerDirectoryDescription *OwnerDirectoryDescription

	// The description for the directory.
	Description *string

	// A DirectoryVpcSettingsDescription () object that contains additional information
	// about a directory. This member is only present if the directory is a Simple AD
	// or Managed AD directory.
	VpcSettings *DirectoryVpcSettingsDescription

	// A RadiusSettings () object that contains information about the RADIUS server
	// configured for this directory.
	RadiusSettings *RadiusSettings

	// The fully qualified name of the directory.
	Name *string

	// The directory identifier.
	DirectoryId *string

	// The directory size.
	Size DirectorySize

	// Specifies when the directory was created.
	LaunchTime *time.Time

	// The access URL for the directory, such as http://.awsapps.com. If no alias has
	// been created for the directory,  is the directory identifier, such as
	// d-XXXXXXXXXX.
	AccessUrl *string

	// The edition associated with this directory.
	Edition DirectoryEdition

	// The desired number of domain controllers in the directory if the directory is
	// Microsoft AD.
	DesiredNumberOfDomainControllers *int32

	// The IP addresses of the DNS servers for the directory. For a Simple AD or
	// Microsoft AD directory, these are the IP addresses of the Simple AD or Microsoft
	// AD directory servers. For an AD Connector directory, these are the IP addresses
	// of the DNS servers or domain controllers in the on-premises directory to which
	// the AD Connector is connected.
	DnsIpAddrs []*string

	// Current directory status of the shared AWS Managed Microsoft AD directory.
	ShareStatus ShareStatus

	// Indicates if single sign-on is enabled for the directory. For more information,
	// see EnableSso () and DisableSso ().
	SsoEnabled *bool
}

// Contains directory limit information for a Region.
type DirectoryLimits struct {

	// Indicates if the cloud directory limit has been reached.
	CloudOnlyDirectoriesLimitReached *bool

	// Indicates if the connected directory limit has been reached.
	ConnectedDirectoriesLimitReached *bool

	// The maximum number of cloud directories allowed in the Region.
	CloudOnlyDirectoriesLimit *int32

	// Indicates if the AWS Managed Microsoft AD directory limit has been reached.
	CloudOnlyMicrosoftADLimitReached *bool

	// The maximum number of connected directories allowed in the Region.
	ConnectedDirectoriesLimit *int32

	// The current number of connected directories in the Region.
	ConnectedDirectoriesCurrentCount *int32

	// The current number of cloud directories in the Region.
	CloudOnlyDirectoriesCurrentCount *int32

	// The maximum number of AWS Managed Microsoft AD directories allowed in the
	// region.
	CloudOnlyMicrosoftADLimit *int32

	// The current number of AWS Managed Microsoft AD directories in the region.
	CloudOnlyMicrosoftADCurrentCount *int32
}

// Contains VPC information for the CreateDirectory () or CreateMicrosoftAD ()
// operation.
type DirectoryVpcSettings struct {

	// The identifier of the VPC in which to create the directory.
	//
	// This member is required.
	VpcId *string

	// The identifiers of the subnets for the directory servers. The two subnets must
	// be in different Availability Zones. AWS Directory Service creates a directory
	// server and a DNS server in each of these subnets.
	//
	// This member is required.
	SubnetIds []*string
}

// Contains information about the directory.
type DirectoryVpcSettingsDescription struct {

	// The identifiers of the subnets for the directory servers.
	SubnetIds []*string

	// The list of Availability Zones that the directory is in.
	AvailabilityZones []*string

	// The identifier of the VPC that the directory is in.
	VpcId *string

	// The domain controller security group identifier for the directory.
	SecurityGroupId *string
}

// Contains information about the domain controllers for a specified directory.
type DomainController struct {

	// A description of the domain controller state.
	StatusReason *string

	// Identifier of the subnet in the VPC that contains the domain controller.
	SubnetId *string

	// The date and time that the status was last updated.
	StatusLastUpdatedDateTime *time.Time

	// Specifies when the domain controller was created.
	LaunchTime *time.Time

	// The IP address of the domain controller.
	DnsIpAddr *string

	// The status of the domain controller.
	Status DomainControllerStatus

	// Identifier of the directory where the domain controller resides.
	DirectoryId *string

	// The Availability Zone where the domain controller is located.
	AvailabilityZone *string

	// The identifier of the VPC that contains the domain controller.
	VpcId *string

	// Identifies a specific domain controller in the directory.
	DomainControllerId *string
}

// Information about SNS topic and AWS Directory Service directory associations.
type EventTopic struct {

	// The SNS topic ARN (Amazon Resource Name).
	TopicArn *string

	// The name of an AWS SNS topic the receives status messages from the directory.
	TopicName *string

	// The date and time of when you associated your directory with the SNS topic.
	CreatedDateTime *time.Time

	// The topic registration status.
	Status TopicStatus

	// The Directory ID of an AWS Directory Service directory that will publish status
	// messages to an SNS topic.
	DirectoryId *string
}

// IP address block. This is often the address block of the DNS server used for
// your on-premises domain.
type IpRoute struct {

	// Description of the address block.
	Description *string

	// IP address block using CIDR format, for example 10.0.0.0/24. This is often the
	// address block of the DNS server used for your on-premises domain. For a single
	// IP address use a CIDR address block with /32. For example 10.0.0.0/32.
	CidrIp *string
}

// Information about one or more IP address blocks.
type IpRouteInfo struct {

	// Description of the IpRouteInfo ().
	Description *string

	// Identifier (ID) of the directory associated with the IP addresses.
	DirectoryId *string

	// IP address block in the IpRoute ().
	CidrIp *string

	// The status of the IP address block.
	IpRouteStatusMsg IpRouteStatusMsg

	// The date and time the address block was added to the directory.
	AddedDateTime *time.Time

	// The reason for the IpRouteStatusMsg.
	IpRouteStatusReason *string
}

// Contains general information about the LDAPS settings.
type LDAPSSettingInfo struct {

	// Describes a state change for LDAPS.
	LDAPSStatusReason *string

	// The date and time when the LDAPS settings were last updated.
	LastUpdatedDateTime *time.Time

	// The state of the LDAPS settings.
	LDAPSStatus LDAPSStatus
}

// Represents a log subscription, which tracks real-time data from a chosen log
// group to a specified destination.
type LogSubscription struct {

	// The date and time that the log subscription was created.
	SubscriptionCreatedDateTime *time.Time

	// Identifier (ID) of the directory that you want to associate with the log
	// subscription.
	DirectoryId *string

	// The name of the log group.
	LogGroupName *string
}

// Describes the directory owner account details that have been shared to the
// directory consumer account.
type OwnerDirectoryDescription struct {

	// Information about the status of the RADIUS server.
	RadiusStatus RadiusStatus

	// IP address of the directory’s domain controllers.
	DnsIpAddrs []*string

	// A RadiusSettings () object that contains information about the RADIUS server.
	RadiusSettings *RadiusSettings

	// Identifier of the directory owner account.
	AccountId *string

	// Identifier of the AWS Managed Microsoft AD directory in the directory owner
	// account.
	DirectoryId *string

	// Information about the VPC settings for the directory.
	VpcSettings *DirectoryVpcSettingsDescription
}

// Contains information about a Remote Authentication Dial In User Service (RADIUS)
// server.
type RadiusSettings struct {

	// Not currently used.
	UseSameUsername *bool

	// The amount of time, in seconds, to wait for the RADIUS server to respond.
	RadiusTimeout *int32

	// The port that your RADIUS server is using for communications. Your on-premises
	// network must allow inbound traffic over this port from the AWS Directory Service
	// servers.
	RadiusPort *int32

	// Required for enabling RADIUS on the directory.
	SharedSecret *string

	// An array of strings that contains the IP addresses of the RADIUS server
	// endpoints, or the IP addresses of your RADIUS server load balancer.
	RadiusServers []*string

	// The maximum number of times that communication with the RADIUS server is
	// attempted.
	RadiusRetries *int32

	// The protocol specified for your RADIUS endpoints.
	AuthenticationProtocol RadiusAuthenticationProtocol

	// Not currently used.
	DisplayLabel *string
}

// Information about a schema extension.
type SchemaExtensionInfo struct {

	// A description of the schema extension.
	Description *string

	// The identifier of the directory to which the schema extension is applied.
	DirectoryId *string

	// The reason for the SchemaExtensionStatus.
	SchemaExtensionStatusReason *string

	// The identifier of the schema extension.
	SchemaExtensionId *string

	// The current status of the schema extension.
	SchemaExtensionStatus SchemaExtensionStatus

	// The date and time that the schema extension was completed.
	EndDateTime *time.Time

	// The date and time that the schema extension started being applied to the
	// directory.
	StartDateTime *time.Time
}

// Details about the shared directory in the directory owner account for which the
// share request in the directory consumer account has been accepted.
type SharedDirectory struct {

	// Identifier of the shared directory in the directory consumer account. This
	// identifier is different for each directory owner account.
	SharedDirectoryId *string

	// A directory share request that is sent by the directory owner to the directory
	// consumer. The request includes a typed message to help the directory consumer
	// administrator determine whether to approve or reject the share invitation.
	ShareNotes *string

	// The date and time that the shared directory was created.
	CreatedDateTime *time.Time

	// The date and time that the shared directory was last updated.
	LastUpdatedDateTime *time.Time

	// The method used when sharing a directory to determine whether the directory
	// should be shared within your AWS organization (ORGANIZATIONS) or with any AWS
	// account by sending a shared directory request (HANDSHAKE).
	ShareMethod ShareMethod

	// Identifier of the directory owner account, which contains the directory that has
	// been shared to the consumer account.
	OwnerAccountId *string

	// Identifier of the directory in the directory owner account.
	OwnerDirectoryId *string

	// Current directory status of the shared AWS Managed Microsoft AD directory.
	ShareStatus ShareStatus

	// Identifier of the directory consumer account that has access to the shared
	// directory (OwnerDirectoryId) in the directory owner account.
	SharedAccountId *string
}

// Identifier that contains details about the directory consumer account.
type ShareTarget struct {

	// Type of identifier to be used in the Id field.
	//
	// This member is required.
	Type TargetType

	// Identifier of the directory consumer account.
	//
	// This member is required.
	Id *string
}

// Describes a directory snapshot.
type Snapshot struct {

	// The snapshot status.
	Status SnapshotStatus

	// The snapshot identifier.
	SnapshotId *string

	// The date and time that the snapshot was taken.
	StartTime *time.Time

	// The descriptive name of the snapshot.
	Name *string

	// The snapshot type.
	Type SnapshotType

	// The directory identifier.
	DirectoryId *string
}

// Contains manual snapshot limit information for a directory.
type SnapshotLimits struct {

	// The maximum number of manual snapshots allowed.
	ManualSnapshotsLimit *int32

	// The current number of manual snapshots of the directory.
	ManualSnapshotsCurrentCount *int32

	// Indicates if the manual snapshot limit has been reached.
	ManualSnapshotsLimitReached *bool
}

// Metadata assigned to a directory consisting of a key-value pair.
type Tag struct {

	// The optional value of the tag. The string value can be Unicode characters. The
	// string can contain only the set of Unicode letters, digits, white-space, '_',
	// '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	//
	// This member is required.
	Value *string

	// Required name of the tag. The string value can be Unicode characters and cannot
	// be prefixed with "aws:". The string can contain only the set of Unicode letters,
	// digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex:
	// "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	//
	// This member is required.
	Key *string
}

// Describes a trust relationship between an AWS Managed Microsoft AD directory and
// an external domain.
type Trust struct {

	// The Directory ID of the AWS directory involved in the trust relationship.
	DirectoryId *string

	// The reason for the TrustState.
	TrustStateReason *string

	// The date and time that the trust relationship was created.
	CreatedDateTime *time.Time

	// The trust relationship state.
	TrustState TrustState

	// Current state of selective authentication for the trust.
	SelectiveAuth SelectiveAuth

	// The date and time that the TrustState was last updated.
	StateLastUpdatedDateTime *time.Time

	// The Fully Qualified Domain Name (FQDN) of the external domain involved in the
	// trust relationship.
	RemoteDomainName *string

	// The unique ID of the trust relationship.
	TrustId *string

	// The date and time that the trust relationship was last updated.
	LastUpdatedDateTime *time.Time

	// The trust relationship direction.
	TrustDirection TrustDirection

	// The trust relationship type. Forest is the default.
	TrustType TrustType
}

// Identifier that contains details about the directory consumer account with whom
// the directory is being unshared.
type UnshareTarget struct {

	// Identifier of the directory consumer account.
	//
	// This member is required.
	Id *string

	// Type of identifier to be used in the Id field.
	//
	// This member is required.
	Type TargetType
}
