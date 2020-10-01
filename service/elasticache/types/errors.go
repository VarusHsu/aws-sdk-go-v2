// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// The customer has exceeded the allowed rate of API calls.
type APICallRateForCustomerExceededFault struct {
	Message *string
}

func (e *APICallRateForCustomerExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *APICallRateForCustomerExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *APICallRateForCustomerExceededFault) ErrorCode() string {
	return "APICallRateForCustomerExceededFault"
}
func (e *APICallRateForCustomerExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified Amazon EC2 security group is already authorized for the specified
// cache security group.
type AuthorizationAlreadyExistsFault struct {
	Message *string
}

func (e *AuthorizationAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AuthorizationAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AuthorizationAlreadyExistsFault) ErrorCode() string {
	return "AuthorizationAlreadyExistsFault"
}
func (e *AuthorizationAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified Amazon EC2 security group is not authorized for the specified
// cache security group.
type AuthorizationNotFoundFault struct {
	Message *string
}

func (e *AuthorizationNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AuthorizationNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AuthorizationNotFoundFault) ErrorCode() string             { return "AuthorizationNotFoundFault" }
func (e *AuthorizationNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a cluster with the given identifier.
type CacheClusterAlreadyExistsFault struct {
	Message *string
}

func (e *CacheClusterAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheClusterAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheClusterAlreadyExistsFault) ErrorCode() string             { return "CacheClusterAlreadyExistsFault" }
func (e *CacheClusterAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested cluster ID does not refer to an existing cluster.
type CacheClusterNotFoundFault struct {
	Message *string
}

func (e *CacheClusterNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheClusterNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheClusterNotFoundFault) ErrorCode() string             { return "CacheClusterNotFoundFault" }
func (e *CacheClusterNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A cache parameter group with the requested name already exists.
type CacheParameterGroupAlreadyExistsFault struct {
	Message *string
}

func (e *CacheParameterGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheParameterGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheParameterGroupAlreadyExistsFault) ErrorCode() string {
	return "CacheParameterGroupAlreadyExistsFault"
}
func (e *CacheParameterGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested cache parameter group name does not refer to an existing cache
// parameter group.
type CacheParameterGroupNotFoundFault struct {
	Message *string
}

func (e *CacheParameterGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheParameterGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheParameterGroupNotFoundFault) ErrorCode() string {
	return "CacheParameterGroupNotFoundFault"
}
func (e *CacheParameterGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the maximum number of
// cache security groups.
type CacheParameterGroupQuotaExceededFault struct {
	Message *string
}

func (e *CacheParameterGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheParameterGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheParameterGroupQuotaExceededFault) ErrorCode() string {
	return "CacheParameterGroupQuotaExceededFault"
}
func (e *CacheParameterGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// A cache security group with the specified name already exists.
type CacheSecurityGroupAlreadyExistsFault struct {
	Message *string
}

func (e *CacheSecurityGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheSecurityGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheSecurityGroupAlreadyExistsFault) ErrorCode() string {
	return "CacheSecurityGroupAlreadyExistsFault"
}
func (e *CacheSecurityGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested cache security group name does not refer to an existing cache
// security group.
type CacheSecurityGroupNotFoundFault struct {
	Message *string
}

func (e *CacheSecurityGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheSecurityGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheSecurityGroupNotFoundFault) ErrorCode() string {
	return "CacheSecurityGroupNotFoundFault"
}
func (e *CacheSecurityGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the allowed number of
// cache security groups.
type CacheSecurityGroupQuotaExceededFault struct {
	Message *string
}

func (e *CacheSecurityGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheSecurityGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheSecurityGroupQuotaExceededFault) ErrorCode() string {
	return "CacheSecurityGroupQuotaExceededFault"
}
func (e *CacheSecurityGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested cache subnet group name is already in use by an existing cache
// subnet group.
type CacheSubnetGroupAlreadyExistsFault struct {
	Message *string
}

func (e *CacheSubnetGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheSubnetGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheSubnetGroupAlreadyExistsFault) ErrorCode() string {
	return "CacheSubnetGroupAlreadyExistsFault"
}
func (e *CacheSubnetGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested cache subnet group is currently in use.
type CacheSubnetGroupInUse struct {
	Message *string
}

func (e *CacheSubnetGroupInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheSubnetGroupInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheSubnetGroupInUse) ErrorCode() string             { return "CacheSubnetGroupInUse" }
func (e *CacheSubnetGroupInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested cache subnet group name does not refer to an existing cache subnet
// group.
type CacheSubnetGroupNotFoundFault struct {
	Message *string
}

func (e *CacheSubnetGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheSubnetGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheSubnetGroupNotFoundFault) ErrorCode() string             { return "CacheSubnetGroupNotFoundFault" }
func (e *CacheSubnetGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the allowed number of
// cache subnet groups.
type CacheSubnetGroupQuotaExceededFault struct {
	Message *string
}

func (e *CacheSubnetGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheSubnetGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheSubnetGroupQuotaExceededFault) ErrorCode() string {
	return "CacheSubnetGroupQuotaExceededFault"
}
func (e *CacheSubnetGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The request cannot be processed because it would exceed the allowed number of
// subnets in a cache subnet group.
type CacheSubnetQuotaExceededFault struct {
	Message *string
}

func (e *CacheSubnetQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CacheSubnetQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CacheSubnetQuotaExceededFault) ErrorCode() string             { return "CacheSubnetQuotaExceededFault" }
func (e *CacheSubnetQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the allowed number of
// clusters per customer.
type ClusterQuotaForCustomerExceededFault struct {
	Message *string
}

func (e *ClusterQuotaForCustomerExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterQuotaForCustomerExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterQuotaForCustomerExceededFault) ErrorCode() string {
	return "ClusterQuotaForCustomerExceededFault"
}
func (e *ClusterQuotaForCustomerExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The Global Datastore name already exists.
type GlobalReplicationGroupAlreadyExistsFault struct {
	Message *string
}

func (e *GlobalReplicationGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GlobalReplicationGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GlobalReplicationGroupAlreadyExistsFault) ErrorCode() string {
	return "GlobalReplicationGroupAlreadyExistsFault"
}
func (e *GlobalReplicationGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The Global Datastore does not exist
type GlobalReplicationGroupNotFoundFault struct {
	Message *string
}

func (e *GlobalReplicationGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GlobalReplicationGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GlobalReplicationGroupNotFoundFault) ErrorCode() string {
	return "GlobalReplicationGroupNotFoundFault"
}
func (e *GlobalReplicationGroupNotFoundFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested cache node type is not available in the specified Availability
// Zone. For more information, see InsufficientCacheClusterCapacity
// (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/ErrorMessages.html#ErrorMessages.INSUFFICIENT_CACHE_CLUSTER_CAPACITY)
// in the ElastiCache User Guide.
type InsufficientCacheClusterCapacityFault struct {
	Message *string
}

func (e *InsufficientCacheClusterCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientCacheClusterCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientCacheClusterCapacityFault) ErrorCode() string {
	return "InsufficientCacheClusterCapacityFault"
}
func (e *InsufficientCacheClusterCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested Amazon Resource Name (ARN) does not refer to an existing resource.
type InvalidARNFault struct {
	Message *string
}

func (e *InvalidARNFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidARNFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidARNFault) ErrorCode() string             { return "InvalidARNFault" }
func (e *InvalidARNFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested cluster is not in the available state.
type InvalidCacheClusterStateFault struct {
	Message *string
}

func (e *InvalidCacheClusterStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCacheClusterStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCacheClusterStateFault) ErrorCode() string             { return "InvalidCacheClusterStateFault" }
func (e *InvalidCacheClusterStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The current state of the cache parameter group does not allow the requested
// operation to occur.
type InvalidCacheParameterGroupStateFault struct {
	Message *string
}

func (e *InvalidCacheParameterGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCacheParameterGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCacheParameterGroupStateFault) ErrorCode() string {
	return "InvalidCacheParameterGroupStateFault"
}
func (e *InvalidCacheParameterGroupStateFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The current state of the cache security group does not allow deletion.
type InvalidCacheSecurityGroupStateFault struct {
	Message *string
}

func (e *InvalidCacheSecurityGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCacheSecurityGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCacheSecurityGroupStateFault) ErrorCode() string {
	return "InvalidCacheSecurityGroupStateFault"
}
func (e *InvalidCacheSecurityGroupStateFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The Global Datastore is not available or in primary-only state.
type InvalidGlobalReplicationGroupStateFault struct {
	Message *string
}

func (e *InvalidGlobalReplicationGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidGlobalReplicationGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidGlobalReplicationGroupStateFault) ErrorCode() string {
	return "InvalidGlobalReplicationGroupStateFault"
}
func (e *InvalidGlobalReplicationGroupStateFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The KMS key supplied is not valid.
type InvalidKMSKeyFault struct {
	Message *string
}

func (e *InvalidKMSKeyFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidKMSKeyFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidKMSKeyFault) ErrorCode() string             { return "InvalidKMSKeyFault" }
func (e *InvalidKMSKeyFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Two or more incompatible parameters were specified.
type InvalidParameterCombinationException struct {
	Message *string
}

func (e *InvalidParameterCombinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterCombinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterCombinationException) ErrorCode() string {
	return "InvalidParameterCombinationException"
}
func (e *InvalidParameterCombinationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The value for a parameter is invalid.
type InvalidParameterValueException struct {
	Message *string
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string             { return "InvalidParameterValueException" }
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested replication group is not in the available state.
type InvalidReplicationGroupStateFault struct {
	Message *string
}

func (e *InvalidReplicationGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidReplicationGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidReplicationGroupStateFault) ErrorCode() string {
	return "InvalidReplicationGroupStateFault"
}
func (e *InvalidReplicationGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The current state of the snapshot does not allow the requested operation to
// occur.
type InvalidSnapshotStateFault struct {
	Message *string
}

func (e *InvalidSnapshotStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSnapshotStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSnapshotStateFault) ErrorCode() string             { return "InvalidSnapshotStateFault" }
func (e *InvalidSnapshotStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An invalid subnet identifier was specified.
type InvalidSubnet struct {
	Message *string
}

func (e *InvalidSubnet) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSubnet) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSubnet) ErrorCode() string             { return "InvalidSubnet" }
func (e *InvalidSubnet) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The VPC network is in an invalid state.
type InvalidVPCNetworkStateFault struct {
	Message *string
}

func (e *InvalidVPCNetworkStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidVPCNetworkStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidVPCNetworkStateFault) ErrorCode() string             { return "InvalidVPCNetworkStateFault" }
func (e *InvalidVPCNetworkStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The node group specified by the NodeGroupId parameter could not be found. Please
// verify that the node group exists and that you spelled the NodeGroupId value
// correctly.
type NodeGroupNotFoundFault struct {
	Message *string
}

func (e *NodeGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NodeGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NodeGroupNotFoundFault) ErrorCode() string             { return "NodeGroupNotFoundFault" }
func (e *NodeGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the maximum allowed
// number of node groups (shards) in a single replication group. The default
// maximum is 90
type NodeGroupsPerReplicationGroupQuotaExceededFault struct {
	Message *string
}

func (e *NodeGroupsPerReplicationGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NodeGroupsPerReplicationGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NodeGroupsPerReplicationGroupQuotaExceededFault) ErrorCode() string {
	return "NodeGroupsPerReplicationGroupQuotaExceededFault"
}
func (e *NodeGroupsPerReplicationGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The request cannot be processed because it would exceed the allowed number of
// cache nodes in a single cluster.
type NodeQuotaForClusterExceededFault struct {
	Message *string
}

func (e *NodeQuotaForClusterExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NodeQuotaForClusterExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NodeQuotaForClusterExceededFault) ErrorCode() string {
	return "NodeQuotaForClusterExceededFault"
}
func (e *NodeQuotaForClusterExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the allowed number of
// cache nodes per customer.
type NodeQuotaForCustomerExceededFault struct {
	Message *string
}

func (e *NodeQuotaForCustomerExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NodeQuotaForCustomerExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NodeQuotaForCustomerExceededFault) ErrorCode() string {
	return "NodeQuotaForCustomerExceededFault"
}
func (e *NodeQuotaForCustomerExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation was not performed because no changes were required.
type NoOperationFault struct {
	Message *string
}

func (e *NoOperationFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoOperationFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoOperationFault) ErrorCode() string             { return "NoOperationFault" }
func (e *NoOperationFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified replication group already exists.
type ReplicationGroupAlreadyExistsFault struct {
	Message *string
}

func (e *ReplicationGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicationGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicationGroupAlreadyExistsFault) ErrorCode() string {
	return "ReplicationGroupAlreadyExistsFault"
}
func (e *ReplicationGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The targeted replication group is not available.
type ReplicationGroupAlreadyUnderMigrationFault struct {
	Message *string
}

func (e *ReplicationGroupAlreadyUnderMigrationFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicationGroupAlreadyUnderMigrationFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicationGroupAlreadyUnderMigrationFault) ErrorCode() string {
	return "ReplicationGroupAlreadyUnderMigrationFault"
}
func (e *ReplicationGroupAlreadyUnderMigrationFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified replication group does not exist.
type ReplicationGroupNotFoundFault struct {
	Message *string
}

func (e *ReplicationGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicationGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicationGroupNotFoundFault) ErrorCode() string             { return "ReplicationGroupNotFoundFault" }
func (e *ReplicationGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The designated replication group is not available for data migration.
type ReplicationGroupNotUnderMigrationFault struct {
	Message *string
}

func (e *ReplicationGroupNotUnderMigrationFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicationGroupNotUnderMigrationFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicationGroupNotUnderMigrationFault) ErrorCode() string {
	return "ReplicationGroupNotUnderMigrationFault"
}
func (e *ReplicationGroupNotUnderMigrationFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// You already have a reservation with the given identifier.
type ReservedCacheNodeAlreadyExistsFault struct {
	Message *string
}

func (e *ReservedCacheNodeAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReservedCacheNodeAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReservedCacheNodeAlreadyExistsFault) ErrorCode() string {
	return "ReservedCacheNodeAlreadyExistsFault"
}
func (e *ReservedCacheNodeAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested reserved cache node was not found.
type ReservedCacheNodeNotFoundFault struct {
	Message *string
}

func (e *ReservedCacheNodeNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReservedCacheNodeNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReservedCacheNodeNotFoundFault) ErrorCode() string             { return "ReservedCacheNodeNotFoundFault" }
func (e *ReservedCacheNodeNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the user's cache node
// quota.
type ReservedCacheNodeQuotaExceededFault struct {
	Message *string
}

func (e *ReservedCacheNodeQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReservedCacheNodeQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReservedCacheNodeQuotaExceededFault) ErrorCode() string {
	return "ReservedCacheNodeQuotaExceededFault"
}
func (e *ReservedCacheNodeQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested cache node offering does not exist.
type ReservedCacheNodesOfferingNotFoundFault struct {
	Message *string
}

func (e *ReservedCacheNodesOfferingNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReservedCacheNodesOfferingNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReservedCacheNodesOfferingNotFoundFault) ErrorCode() string {
	return "ReservedCacheNodesOfferingNotFoundFault"
}
func (e *ReservedCacheNodesOfferingNotFoundFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified service linked role (SLR) was not found.
type ServiceLinkedRoleNotFoundFault struct {
	Message *string
}

func (e *ServiceLinkedRoleNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceLinkedRoleNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceLinkedRoleNotFoundFault) ErrorCode() string             { return "ServiceLinkedRoleNotFoundFault" }
func (e *ServiceLinkedRoleNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service update doesn't exist
type ServiceUpdateNotFoundFault struct {
	Message *string
}

func (e *ServiceUpdateNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUpdateNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUpdateNotFoundFault) ErrorCode() string             { return "ServiceUpdateNotFoundFault" }
func (e *ServiceUpdateNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a snapshot with the given name.
type SnapshotAlreadyExistsFault struct {
	Message *string
}

func (e *SnapshotAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotAlreadyExistsFault) ErrorCode() string             { return "SnapshotAlreadyExistsFault" }
func (e *SnapshotAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You attempted one of the following operations:
//
//     * Creating a snapshot of a
// Redis cluster running on a cache.t1.micro cache node.
//
//     * Creating a snapshot
// of a cluster that is running Memcached rather than Redis.
//
// Neither of these are
// supported by ElastiCache.
type SnapshotFeatureNotSupportedFault struct {
	Message *string
}

func (e *SnapshotFeatureNotSupportedFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotFeatureNotSupportedFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotFeatureNotSupportedFault) ErrorCode() string {
	return "SnapshotFeatureNotSupportedFault"
}
func (e *SnapshotFeatureNotSupportedFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested snapshot name does not refer to an existing snapshot.
type SnapshotNotFoundFault struct {
	Message *string
}

func (e *SnapshotNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotNotFoundFault) ErrorCode() string             { return "SnapshotNotFoundFault" }
func (e *SnapshotNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the maximum number of
// snapshots.
type SnapshotQuotaExceededFault struct {
	Message *string
}

func (e *SnapshotQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotQuotaExceededFault) ErrorCode() string             { return "SnapshotQuotaExceededFault" }
func (e *SnapshotQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested subnet is being used by another cache subnet group.
type SubnetInUse struct {
	Message *string
}

func (e *SubnetInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetInUse) ErrorCode() string             { return "SubnetInUse" }
func (e *SubnetInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested tag was not found on this resource.
type TagNotFoundFault struct {
	Message *string
}

func (e *TagNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagNotFoundFault) ErrorCode() string             { return "TagNotFoundFault" }
func (e *TagNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would cause the resource to have more
// than the allowed number of tags. The maximum number of tags permitted on a
// resource is 50.
type TagQuotaPerResourceExceeded struct {
	Message *string
}

func (e *TagQuotaPerResourceExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagQuotaPerResourceExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagQuotaPerResourceExceeded) ErrorCode() string             { return "TagQuotaPerResourceExceeded" }
func (e *TagQuotaPerResourceExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The TestFailover action is not available.
type TestFailoverNotAvailableFault struct {
	Message *string
}

func (e *TestFailoverNotAvailableFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TestFailoverNotAvailableFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TestFailoverNotAvailableFault) ErrorCode() string             { return "TestFailoverNotAvailableFault" }
func (e *TestFailoverNotAvailableFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
