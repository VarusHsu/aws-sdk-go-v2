// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// A resource that failed to be added to or removed from a group.
type FailedResource struct {

	// The error code associated with the failure.
	ErrorCode *string

	// The error message text associated with the failure.
	ErrorMessage *string

	// The ARN of the resource that failed to be added or removed.
	ResourceArn *string
}

// A resource group that contains AWS resources. You can assign resources to the
// group by associating either of the following elements with the group:
//
//     *
// ResourceQuery () - Use a resource query to specify a set of tag keys and values.
// All resources in the same AWS Region and AWS account that have those keys with
// the same values are included in the group. You can add a resource query when you
// create the group.
//
//     * GroupConfiguration () - Use a service configuration to
// associate the group with an AWS service. The configuration specifies which
// resource types can be included in the group.
type Group struct {

	// The name of the resource group.
	//
	// This member is required.
	Name *string

	// The ARN of the resource group.
	//
	// This member is required.
	GroupArn *string

	// The description of the resource group.
	Description *string
}

// A service configuration associated with a resource group. The configuration
// options are determined by the AWS service that defines the Type, and specifies
// which resources can be included in the group. You can add a service
// configuration when you create the group.
type GroupConfiguration struct {

	// The current status of an attempt to update the group configuration.
	Status GroupConfigurationStatus

	// If present, the new configuration that is in the process of being applied to the
	// group.
	ProposedConfiguration []*GroupConfigurationItem

	// If present, the reason why a request to update the group configuration failed.
	FailureReason *string

	// The configuration currently associated with the group and in effect.
	Configuration []*GroupConfigurationItem
}

// An item in a group configuration. A group configuration can have one or more
// items.
type GroupConfigurationItem struct {

	// Specifies the type of group configuration item. Each item must have a unique
	// value for type. You can specify the following string values:
	//
	//     *
	// AWS::EC2::CapacityReservationPool For more information about EC2 capacity
	// reservation groups, see Working with capacity reservation groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/capacity-reservations-using.html#create-cr-group)
	// in the EC2 Users Guide.
	//
	//     * AWS::ResourceGroups::Generic - Supports
	// parameters that configure the behavior of resource groups of any type.
	//
	// This member is required.
	Type *string

	// A collection of parameters for this group configuration item.
	Parameters []*GroupConfigurationParameter
}

// A parameter for a group configuration item.
type GroupConfigurationParameter struct {

	// The name of the group configuration parameter. You can specify the following
	// string values:
	//
	//     * For configuration item type
	// AWS::ResourceGroups::Generic:
	//
	//         * allowed-resource-types Specifies the
	// types of resources that you can add to this group by using the GroupResources ()
	// operation.
	//
	//     * For configuration item type
	// AWS::EC2::CapacityReservationPool:
	//
	//         * None - This configuration item
	// type doesn't support any parameters.
	//
	//     For more information about EC2
	// capacity reservation groups, see Working with capacity reservation groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/capacity-reservations-using.html#create-cr-group)
	// in the EC2 Users Guide.
	//
	// This member is required.
	Name *string

	// The values of for this parameter. You can specify the following string value:
	//
	//
	// * For item type allowed-resource-types: the only supported parameter value is
	// AWS::EC2::CapacityReservation.
	Values []*string
}

// A filter collection that you can use to restrict the results from a List
// operation to only those you want to include.
type GroupFilter struct {

	// The name of the filter. Filter names are case-sensitive.
	//
	// This member is required.
	Name GroupFilterName

	// One or more filter values. Allowed filter values vary by group filter name, and
	// are case-sensitive.
	//
	// This member is required.
	Values []*string
}

// The unique identifiers for a resource group.
type GroupIdentifier struct {

	// The name of the resource group.
	GroupName *string

	// The ARN of the resource group.
	GroupArn *string
}

// A mapping of a query attached to a resource group that determines the AWS
// resources that are members of the group.
type GroupQuery struct {

	// The name of the resource group that is associated with the specified resource
	// query.
	//
	// This member is required.
	GroupName *string

	// The resource query that determines which AWS resources are members of the
	// associated resource group.
	//
	// This member is required.
	ResourceQuery *ResourceQuery
}

// A two-part error structure that can occur in ListGroupResources or
// SearchResources operations on CloudFormation stack-based queries. The error
// occurs if the CloudFormation stack on which the query is based either does not
// exist, or has a status that renders the stack inactive. A QueryError occurrence
// does not necessarily mean that AWS Resource Groups could not complete the
// operation, but the resulting group might have no member resources.
type QueryError struct {

	// A message that explains the ErrorCode value. Messages might state that the
	// specified CloudFormation stack does not exist (or no longer exists). For
	// CLOUDFORMATION_STACK_INACTIVE, the message typically states that the
	// CloudFormation stack has a status that is not (or no longer) active, such as
	// CREATE_FAILED.
	Message *string

	// Possible values are CLOUDFORMATION_STACK_INACTIVE and
	// CLOUDFORMATION_STACK_NOT_EXISTING.
	ErrorCode QueryErrorCode
}

// A filter name and value pair that is used to obtain more specific results from a
// list of resources.
type ResourceFilter struct {

	// The name of the filter. Filter names are case-sensitive.
	//
	// This member is required.
	Name ResourceFilterName

	// One or more filter values. Allowed filter values vary by resource filter name,
	// and are case-sensitive.
	//
	// This member is required.
	Values []*string
}

// The ARN of a resource, and its resource type.
type ResourceIdentifier struct {

	// The resource type of a resource, such as AWS::EC2::Instance.
	ResourceType *string

	// The ARN of a resource.
	ResourceArn *string
}

// The query that is used to define a resource group or a search for resources.
type ResourceQuery struct {

	// The query that defines a group or a search.
	//
	// This member is required.
	Query *string

	// The type of the query. You can use the following values:
	//
	//     *
	// CLOUDFORMATION_STACK_1_0: A JSON syntax that lets you specify a CloudFormation
	// stack ARN.
	//
	//     * TAG_FILTERS_1_0: A JSON syntax that lets you specify a
	// collection of simple tag filters for resource types and tags, as supported by
	// the AWS Tagging API ResourceTypeFilters parameter of the tagging:GetResources
	// (https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_GetResources.html#resourcegrouptagging-GetResources-request-ResourceTypeFilters)
	// operation. If you specify more than one tag key, only resources that match all
	// tag keys, and at least one value of each specified tag key, are returned in your
	// query. If you specify more than one value for a tag key, a resource matches the
	// filter if it has a tag key value that matches any of the specified values. For
	// example, consider the following sample query for resources that have two tags,
	// Stage and Version, with two values each:
	// [{"Key":"Stage","Values":["Test","Deploy"]},{"Key":"Version","Values":["1","2"]}]
	// The results of this query could include the following.
	//
	//         * An EC2
	// instance that has the following two tags: {"Key":"Stage","Value":"Deploy"}, and
	// {"Key":"Version","Value":"2"}
	//
	//         * An S3 bucket that has the following two
	// tags: {"Key":"Stage","Value":"Test"}, and {"Key":"Version","Value":"1"}
	//
	//     The
	// query would not include the following items in the results, however.
	//
	//         *
	// An EC2 instance that has only the following tag:
	// {"Key":"Stage","Value":"Deploy"}. The instance does not have all of the tag keys
	// specified in the filter, so it is excluded from the results.
	//
	//         * An RDS
	// database that has the following two tags: {"Key":"Stage","Value":"Archived"},
	// and {"Key":"Version","Value":"4"} The database has all of the tag keys, but none
	// of those keys has an associated value that matches at least one of the specified
	// values in the filter.
	//
	// This member is required.
	Type QueryType
}
