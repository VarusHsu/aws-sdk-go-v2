// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An object that describes a managed permission associated with a resource share.
type AssociatedPermission struct {

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the associated managed permission.
	Arn *string

	// Indicates whether the associated resource share is using the default version of
	// the permission.
	DefaultVersion *bool

	// Indicates what features are available for this resource share. This parameter
	// can have one of the following values:
	//   - STANDARD – A resource share that supports all functionality. These resource
	//   shares are visible to all principals you share the resource share with. You can
	//   modify these resource shares in RAM using the console or APIs. This resource
	//   share might have been created by RAM, or it might have been CREATED_FROM_POLICY
	//   and then promoted.
	//   - CREATED_FROM_POLICY – The customer manually shared a resource by attaching
	//   a resource-based policy. That policy did not match any existing managed
	//   permissions, so RAM created this customer managed permission automatically on
	//   the customer's behalf based on the attached policy document. This type of
	//   resource share is visible only to the Amazon Web Services account that created
	//   it. You can't modify it in RAM unless you promote it. For more information, see
	//   PromoteResourceShareCreatedFromPolicy .
	//   - PROMOTING_TO_STANDARD – This resource share was originally
	//   CREATED_FROM_POLICY , but the customer ran the
	//   PromoteResourceShareCreatedFromPolicy and that operation is still in progress.
	//   This value changes to STANDARD when complete.
	FeatureSet PermissionFeatureSet

	// The date and time when the association between the permission and the resource
	// share was last updated.
	LastUpdatedTime *time.Time

	// The version of the permission currently associated with the resource share.
	PermissionVersion *string

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of a resource share associated with this permission.
	ResourceShareArn *string

	// The resource type to which this permission applies.
	ResourceType *string

	// The current status of the association between the permission and the resource
	// share. The following are the possible values:
	//   - ATTACHABLE – This permission or version can be associated with resource
	//   shares.
	//   - UNATTACHABLE – This permission or version can't currently be associated with
	//   resource shares.
	//   - DELETING – This permission or version is in the process of being deleted.
	//   - DELETED – This permission or version is deleted.
	Status *string

	noSmithyDocumentSerde
}

// Describes a principal for use with Resource Access Manager.
type Principal struct {

	// The date and time when the principal was associated with the resource share.
	CreationTime *time.Time

	// Indicates the relationship between the Amazon Web Services account the
	// principal belongs to and the account that owns the resource share:
	//   - True – The two accounts belong to same organization.
	//   - False – The two accounts do not belong to the same organization.
	External *bool

	// The ID of the principal that can be associated with a resource share.
	Id *string

	// The date and time when the association between the resource share and the
	// principal was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of a resource share the principal is associated with.
	ResourceShareArn *string

	noSmithyDocumentSerde
}

// A structure that represents the background work that RAM performs when you
// invoke the ReplacePermissionAssociations operation.
type ReplacePermissionAssociationsWork struct {

	// The date and time when this asynchronous background task was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the managed permission that this background task is replacing.
	FromPermissionArn *string

	// The version of the managed permission that this background task is replacing.
	FromPermissionVersion *string

	// The unique identifier for the background task associated with one
	// ReplacePermissionAssociations request.
	Id *string

	// The date and time when the status of this background task was last updated.
	LastUpdatedTime *time.Time

	// Specifies the current status of the background tasks for the specified ID. The
	// output is one of the following strings:
	//   - IN_PROGRESS
	//   - COMPLETED
	//   - FAILED
	Status ReplacePermissionAssociationsWorkStatus

	// Specifies the reason for a FAILED status. This field is present only when there
	// status is FAILED .
	StatusMessage *string

	// The ARN of the managed permission that this background task is associating with
	// the resource shares in place of the managed permission and version specified in
	// fromPermissionArn and fromPermissionVersion .
	ToPermissionArn *string

	// The version of the managed permission that this background task is associating
	// with the resource shares. This is always the version that is currently the
	// default for this managed permission.
	ToPermissionVersion *string

	noSmithyDocumentSerde
}

// Describes a resource associated with a resource share in RAM.
type Resource struct {

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource.
	Arn *string

	// The date and time when the resource was associated with the resource share.
	CreationTime *time.Time

	// The date an time when the association between the resource and the resource
	// share was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource group. This value is available only if the resource is part of a
	// resource group.
	ResourceGroupArn *string

	// Specifies the scope of visibility of this resource:
	//   - REGIONAL – The resource can be accessed only by using requests that target
	//   the Amazon Web Services Region in which the resource exists.
	//   - GLOBAL – The resource can be accessed from any Amazon Web Services Region.
	ResourceRegionScope ResourceRegionScope

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share this resource is associated with.
	ResourceShareArn *string

	// The current status of the resource.
	Status ResourceStatus

	// A message about the status of the resource.
	StatusMessage *string

	// The resource type. This takes the form of: service-code : resource-code , and is
	// case-insensitive. For example, an Amazon EC2 Subnet would be represented by the
	// string ec2:subnet .
	Type *string

	noSmithyDocumentSerde
}

// Describes a resource share in RAM.
type ResourceShare struct {

	// Indicates whether principals outside your organization in Organizations can be
	// associated with a resource share.
	//   - True – the resource share can be shared with any Amazon Web Services
	//   account.
	//   - False – the resource share can be shared with only accounts in the same
	//   organization as the account that owns the resource share.
	AllowExternalPrincipals *bool

	// The date and time when the resource share was created.
	CreationTime *time.Time

	// Indicates what features are available for this resource share. This parameter
	// can have one of the following values:
	//   - STANDARD – A resource share that supports all functionality. These resource
	//   shares are visible to all principals you share the resource share with. You can
	//   modify these resource shares in RAM using the console or APIs. This resource
	//   share might have been created by RAM, or it might have been CREATED_FROM_POLICY
	//   and then promoted.
	//   - CREATED_FROM_POLICY – The customer manually shared a resource by attaching
	//   a resource-based policy. That policy did not match any existing managed
	//   permissions, so RAM created this customer managed permission automatically on
	//   the customer's behalf based on the attached policy document. This type of
	//   resource share is visible only to the Amazon Web Services account that created
	//   it. You can't modify it in RAM unless you promote it. For more information, see
	//   PromoteResourceShareCreatedFromPolicy .
	//   - PROMOTING_TO_STANDARD – This resource share was originally
	//   CREATED_FROM_POLICY , but the customer ran the
	//   PromoteResourceShareCreatedFromPolicy and that operation is still in progress.
	//   This value changes to STANDARD when complete.
	FeatureSet ResourceShareFeatureSet

	// The date and time when the resource share was last updated.
	LastUpdatedTime *time.Time

	// The name of the resource share.
	Name *string

	// The ID of the Amazon Web Services account that owns the resource share.
	OwningAccountId *string

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share
	ResourceShareArn *string

	// The current status of the resource share.
	Status ResourceShareStatus

	// A message about the status of the resource share.
	StatusMessage *string

	// The tag key and value pairs attached to the resource share.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes an association between a resource share and either a principal or a
// resource.
type ResourceShareAssociation struct {

	// The associated entity. This can be either of the following:
	//   - For a resource association, this is the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	//   of the resource.
	//   - For principal associations, this is one of the following:
	//   - The ID of an Amazon Web Services account
	//   - The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	//   of an organization in Organizations
	//   - The ARN of an organizational unit (OU) in Organizations
	//   - The ARN of an IAM role
	//   - The ARN of an IAM user
	AssociatedEntity *string

	// The type of entity included in this association.
	AssociationType ResourceShareAssociationType

	// The date and time when the association was created.
	CreationTime *time.Time

	// Indicates whether the principal belongs to the same organization in
	// Organizations as the Amazon Web Services account that owns the resource share.
	External *bool

	// The date and time when the association was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share.
	ResourceShareArn *string

	// The name of the resource share.
	ResourceShareName *string

	// The current status of the association.
	Status ResourceShareAssociationStatus

	// A message about the status of the association.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Describes an invitation for an Amazon Web Services account to join a resource
// share.
type ResourceShareInvitation struct {

	// The date and time when the invitation was sent.
	InvitationTimestamp *time.Time

	// The ID of the Amazon Web Services account that received the invitation.
	ReceiverAccountId *string

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the IAM user or role that received the invitation.
	ReceiverArn *string

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share
	ResourceShareArn *string

	// To view the resources associated with a pending resource share invitation, use
	// ListPendingInvitationResources .
	//
	// Deprecated: This member has been deprecated. Use ListPendingInvitationResources.
	ResourceShareAssociations []ResourceShareAssociation

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the invitation.
	ResourceShareInvitationArn *string

	// The name of the resource share.
	ResourceShareName *string

	// The ID of the Amazon Web Services account that sent the invitation.
	SenderAccountId *string

	// The current status of the invitation.
	Status ResourceShareInvitationStatus

	noSmithyDocumentSerde
}

// Information about a RAM managed permission.
type ResourceSharePermissionDetail struct {

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of this RAM managed permission.
	Arn *string

	// The date and time when the permission was created.
	CreationTime *time.Time

	// Specifies whether the version of the permission represented in this response is
	// the default version for this permission.
	DefaultVersion *bool

	// Indicates what features are available for this resource share. This parameter
	// can have one of the following values:
	//   - STANDARD – A resource share that supports all functionality. These resource
	//   shares are visible to all principals you share the resource share with. You can
	//   modify these resource shares in RAM using the console or APIs. This resource
	//   share might have been created by RAM, or it might have been CREATED_FROM_POLICY
	//   and then promoted.
	//   - CREATED_FROM_POLICY – The customer manually shared a resource by attaching
	//   a resource-based policy. That policy did not match any existing managed
	//   permissions, so RAM created this customer managed permission automatically on
	//   the customer's behalf based on the attached policy document. This type of
	//   resource share is visible only to the Amazon Web Services account that created
	//   it. You can't modify it in RAM unless you promote it. For more information, see
	//   PromoteResourceShareCreatedFromPolicy .
	//   - PROMOTING_TO_STANDARD – This resource share was originally
	//   CREATED_FROM_POLICY , but the customer ran the
	//   PromoteResourceShareCreatedFromPolicy and that operation is still in progress.
	//   This value changes to STANDARD when complete.
	FeatureSet PermissionFeatureSet

	// Specifies whether the version of the permission represented in this response is
	// the default version for all resources of this resource type.
	IsResourceTypeDefault *bool

	// The date and time when the permission was last updated.
	LastUpdatedTime *time.Time

	// The name of this permission.
	Name *string

	// The permission's effect and actions in JSON format. The effect indicates
	// whether the specified actions are allowed or denied. The actions list the
	// operations to which the principal is granted or denied access.
	Permission *string

	// The type of managed permission. This can be one of the following values:
	//   - AWS_MANAGED – Amazon Web Services created and manages this managed
	//   permission. You can associate it with your resource shares, but you can't modify
	//   it.
	//   - CUSTOMER_MANAGED – You, or another principal in your account created this
	//   managed permission. You can associate it with your resource shares and create
	//   new versions that have different permissions.
	PermissionType PermissionType

	// The resource type to which this permission applies.
	ResourceType *string

	// The current status of the association between the permission and the resource
	// share. The following are the possible values:
	//   - ATTACHABLE – This permission or version can be associated with resource
	//   shares.
	//   - UNATTACHABLE – This permission or version can't currently be associated with
	//   resource shares.
	//   - DELETING – This permission or version is in the process of being deleted.
	//   - DELETED – This permission or version is deleted.
	Status PermissionStatus

	// The tag key and value pairs attached to the resource share.
	Tags []Tag

	// The version of the permission described in this response.
	Version *string

	noSmithyDocumentSerde
}

// Information about an RAM permission.
type ResourceSharePermissionSummary struct {

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the permission you want information about.
	Arn *string

	// The date and time when the permission was created.
	CreationTime *time.Time

	// Specifies whether the version of the managed permission used by this resource
	// share is the default version for this managed permission.
	DefaultVersion *bool

	// Indicates what features are available for this resource share. This parameter
	// can have one of the following values:
	//   - STANDARD – A resource share that supports all functionality. These resource
	//   shares are visible to all principals you share the resource share with. You can
	//   modify these resource shares in RAM using the console or APIs. This resource
	//   share might have been created by RAM, or it might have been CREATED_FROM_POLICY
	//   and then promoted.
	//   - CREATED_FROM_POLICY – The customer manually shared a resource by attaching
	//   a resource-based policy. That policy did not match any existing managed
	//   permissions, so RAM created this customer managed permission automatically on
	//   the customer's behalf based on the attached policy document. This type of
	//   resource share is visible only to the Amazon Web Services account that created
	//   it. You can't modify it in RAM unless you promote it. For more information, see
	//   PromoteResourceShareCreatedFromPolicy .
	//   - PROMOTING_TO_STANDARD – This resource share was originally
	//   CREATED_FROM_POLICY , but the customer ran the
	//   PromoteResourceShareCreatedFromPolicy and that operation is still in progress.
	//   This value changes to STANDARD when complete.
	FeatureSet PermissionFeatureSet

	// Specifies whether the managed permission associated with this resource share is
	// the default managed permission for all resources of this resource type.
	IsResourceTypeDefault *bool

	// The date and time when the permission was last updated.
	LastUpdatedTime *time.Time

	// The name of this managed permission.
	Name *string

	// The type of managed permission. This can be one of the following values:
	//   - AWS_MANAGED – Amazon Web Services created and manages this managed
	//   permission. You can associate it with your resource shares, but you can't modify
	//   it.
	//   - CUSTOMER_MANAGED – You, or another principal in your account created this
	//   managed permission. You can associate it with your resource shares and create
	//   new versions that have different permissions.
	PermissionType PermissionType

	// The type of resource to which this permission applies. This takes the form of:
	// service-code : resource-code , and is case-insensitive. For example, an Amazon
	// EC2 Subnet would be represented by the string ec2:subnet .
	ResourceType *string

	// The current status of the permission.
	Status *string

	// A list of the tag key value pairs currently attached to the permission.
	Tags []Tag

	// The version of the permission associated with this resource share.
	Version *string

	noSmithyDocumentSerde
}

// Information about a shareable resource type and the Amazon Web Services service
// to which resources of that type belong.
type ServiceNameAndResourceType struct {

	// Specifies the scope of visibility of resources of this type:
	//   - REGIONAL – The resource can be accessed only by using requests that target
	//   the Amazon Web Services Region in which the resource exists.
	//   - GLOBAL – The resource can be accessed from any Amazon Web Services Region.
	ResourceRegionScope ResourceRegionScope

	// The type of the resource. This takes the form of: service-code : resource-code ,
	// and is case-insensitive. For example, an Amazon EC2 Subnet would be represented
	// by the string ec2:subnet .
	ResourceType *string

	// The name of the Amazon Web Services service to which resources of this type
	// belong.
	ServiceName *string

	noSmithyDocumentSerde
}

// A structure containing a tag. A tag is metadata that you can attach to your
// resources to help organize and categorize them. You can also use them to help
// you secure your resources. For more information, see Controlling access to
// Amazon Web Services resources using tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html)
// . For more information about tags, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
// in the Amazon Web Services General Reference Guide.
type Tag struct {

	// The key, or name, attached to the tag. Every tag must have a key. Key names are
	// case sensitive.
	Key *string

	// The string value attached to the tag. The value can be an empty string. Key
	// values are case sensitive.
	Value *string

	noSmithyDocumentSerde
}

// A tag key and optional list of possible values that you can use to filter
// results for tagged resources.
type TagFilter struct {

	// The tag key. This must have a valid string value and can't be empty.
	TagKey *string

	// A list of zero or more tag values. If no values are provided, then the filter
	// matches any tag with the specified key, regardless of its value.
	TagValues []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
