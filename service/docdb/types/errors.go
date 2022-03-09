// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The specified CIDR IP or Amazon EC2 security group isn't authorized for the
// specified security group. Amazon DocumentDB also might not be authorized to
// perform necessary actions on your behalf using IAM.
type AuthorizationNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
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
func (e *AuthorizationNotFoundFault) ErrorCode() string             { return "AuthorizationNotFound" }
func (e *AuthorizationNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// CertificateIdentifier doesn't refer to an existing certificate.
type CertificateNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *CertificateNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateNotFoundFault) ErrorCode() string             { return "CertificateNotFound" }
func (e *CertificateNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a cluster with the given identifier.
type DBClusterAlreadyExistsFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBClusterAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterAlreadyExistsFault) ErrorCode() string             { return "DBClusterAlreadyExistsFault" }
func (e *DBClusterAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBClusterIdentifier doesn't refer to an existing cluster.
type DBClusterNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBClusterNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterNotFoundFault) ErrorCode() string             { return "DBClusterNotFoundFault" }
func (e *DBClusterNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBClusterParameterGroupName doesn't refer to an existing cluster parameter
// group.
type DBClusterParameterGroupNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBClusterParameterGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterParameterGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterParameterGroupNotFoundFault) ErrorCode() string {
	return "DBClusterParameterGroupNotFound"
}
func (e *DBClusterParameterGroupNotFoundFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The cluster can't be created because you have reached the maximum allowed quota
// of clusters.
type DBClusterQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBClusterQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterQuotaExceededFault) ErrorCode() string             { return "DBClusterQuotaExceededFault" }
func (e *DBClusterQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a cluster snapshot with the given identifier.
type DBClusterSnapshotAlreadyExistsFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBClusterSnapshotAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterSnapshotAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterSnapshotAlreadyExistsFault) ErrorCode() string {
	return "DBClusterSnapshotAlreadyExistsFault"
}
func (e *DBClusterSnapshotAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBClusterSnapshotIdentifier doesn't refer to an existing cluster snapshot.
type DBClusterSnapshotNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBClusterSnapshotNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBClusterSnapshotNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBClusterSnapshotNotFoundFault) ErrorCode() string             { return "DBClusterSnapshotNotFoundFault" }
func (e *DBClusterSnapshotNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a instance with the given identifier.
type DBInstanceAlreadyExistsFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBInstanceAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBInstanceAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBInstanceAlreadyExistsFault) ErrorCode() string             { return "DBInstanceAlreadyExists" }
func (e *DBInstanceAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBInstanceIdentifier doesn't refer to an existing instance.
type DBInstanceNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBInstanceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBInstanceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBInstanceNotFoundFault) ErrorCode() string             { return "DBInstanceNotFound" }
func (e *DBInstanceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A parameter group with the same name already exists.
type DBParameterGroupAlreadyExistsFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBParameterGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBParameterGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBParameterGroupAlreadyExistsFault) ErrorCode() string {
	return "DBParameterGroupAlreadyExists"
}
func (e *DBParameterGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBParameterGroupName doesn't refer to an existing parameter group.
type DBParameterGroupNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBParameterGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBParameterGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBParameterGroupNotFoundFault) ErrorCode() string             { return "DBParameterGroupNotFound" }
func (e *DBParameterGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This request would cause you to exceed the allowed number of parameter groups.
type DBParameterGroupQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBParameterGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBParameterGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBParameterGroupQuotaExceededFault) ErrorCode() string {
	return "DBParameterGroupQuotaExceeded"
}
func (e *DBParameterGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBSecurityGroupName doesn't refer to an existing security group.
type DBSecurityGroupNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBSecurityGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSecurityGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSecurityGroupNotFoundFault) ErrorCode() string             { return "DBSecurityGroupNotFound" }
func (e *DBSecurityGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBSnapshotIdentifier is already being used by an existing snapshot.
type DBSnapshotAlreadyExistsFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBSnapshotAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSnapshotAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSnapshotAlreadyExistsFault) ErrorCode() string             { return "DBSnapshotAlreadyExists" }
func (e *DBSnapshotAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBSnapshotIdentifier doesn't refer to an existing snapshot.
type DBSnapshotNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBSnapshotNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSnapshotNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSnapshotNotFoundFault) ErrorCode() string             { return "DBSnapshotNotFound" }
func (e *DBSnapshotNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DBSubnetGroupName is already being used by an existing subnet group.
type DBSubnetGroupAlreadyExistsFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBSubnetGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupAlreadyExistsFault) ErrorCode() string             { return "DBSubnetGroupAlreadyExists" }
func (e *DBSubnetGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Subnets in the subnet group should cover at least two Availability Zones unless
// there is only one Availability Zone.
type DBSubnetGroupDoesNotCoverEnoughAZs struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBSubnetGroupDoesNotCoverEnoughAZs) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupDoesNotCoverEnoughAZs) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupDoesNotCoverEnoughAZs) ErrorCode() string {
	return "DBSubnetGroupDoesNotCoverEnoughAZs"
}
func (e *DBSubnetGroupDoesNotCoverEnoughAZs) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// DBSubnetGroupName doesn't refer to an existing subnet group.
type DBSubnetGroupNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBSubnetGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupNotFoundFault) ErrorCode() string             { return "DBSubnetGroupNotFoundFault" }
func (e *DBSubnetGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed number of subnet groups.
type DBSubnetGroupQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBSubnetGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetGroupQuotaExceededFault) ErrorCode() string             { return "DBSubnetGroupQuotaExceeded" }
func (e *DBSubnetGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed number of subnets in a subnet
// group.
type DBSubnetQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBSubnetQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBSubnetQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBSubnetQuotaExceededFault) ErrorCode() string             { return "DBSubnetQuotaExceededFault" }
func (e *DBSubnetQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The upgrade failed because a resource that the depends on can't be modified.
type DBUpgradeDependencyFailureFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *DBUpgradeDependencyFailureFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DBUpgradeDependencyFailureFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DBUpgradeDependencyFailureFault) ErrorCode() string             { return "DBUpgradeDependencyFailure" }
func (e *DBUpgradeDependencyFailureFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have reached the maximum number of event subscriptions.
type EventSubscriptionQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *EventSubscriptionQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EventSubscriptionQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EventSubscriptionQuotaExceededFault) ErrorCode() string {
	return "EventSubscriptionQuotaExceeded"
}
func (e *EventSubscriptionQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The GlobalClusterIdentifier already exists. Choose a new global cluster
// identifier (unique name) to create a new global cluster.
type GlobalClusterAlreadyExistsFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *GlobalClusterAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GlobalClusterAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GlobalClusterAlreadyExistsFault) ErrorCode() string {
	return "GlobalClusterAlreadyExistsFault"
}
func (e *GlobalClusterAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The GlobalClusterIdentifier doesn't refer to an existing global cluster.
type GlobalClusterNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *GlobalClusterNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GlobalClusterNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GlobalClusterNotFoundFault) ErrorCode() string             { return "GlobalClusterNotFoundFault" }
func (e *GlobalClusterNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The number of global clusters for this account is already at the maximum
// allowed.
type GlobalClusterQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *GlobalClusterQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GlobalClusterQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GlobalClusterQuotaExceededFault) ErrorCode() string {
	return "GlobalClusterQuotaExceededFault"
}
func (e *GlobalClusterQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed number of instances.
type InstanceQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InstanceQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InstanceQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InstanceQuotaExceededFault) ErrorCode() string             { return "InstanceQuotaExceeded" }
func (e *InstanceQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The cluster doesn't have enough capacity for the current operation.
type InsufficientDBClusterCapacityFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InsufficientDBClusterCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientDBClusterCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientDBClusterCapacityFault) ErrorCode() string {
	return "InsufficientDBClusterCapacityFault"
}
func (e *InsufficientDBClusterCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified instance class isn't available in the specified Availability Zone.
type InsufficientDBInstanceCapacityFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InsufficientDBInstanceCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientDBInstanceCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientDBInstanceCapacityFault) ErrorCode() string {
	return "InsufficientDBInstanceCapacity"
}
func (e *InsufficientDBInstanceCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// There is not enough storage available for the current action. You might be able
// to resolve this error by updating your subnet group to use different
// Availability Zones that have more storage available.
type InsufficientStorageClusterCapacityFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InsufficientStorageClusterCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientStorageClusterCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientStorageClusterCapacityFault) ErrorCode() string {
	return "InsufficientStorageClusterCapacity"
}
func (e *InsufficientStorageClusterCapacityFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The provided value isn't a valid cluster snapshot state.
type InvalidDBClusterSnapshotStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDBClusterSnapshotStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBClusterSnapshotStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBClusterSnapshotStateFault) ErrorCode() string {
	return "InvalidDBClusterSnapshotStateFault"
}
func (e *InvalidDBClusterSnapshotStateFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The cluster isn't in a valid state.
type InvalidDBClusterStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDBClusterStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBClusterStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBClusterStateFault) ErrorCode() string             { return "InvalidDBClusterStateFault" }
func (e *InvalidDBClusterStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified instance isn't in the available state.
type InvalidDBInstanceStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDBInstanceStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBInstanceStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBInstanceStateFault) ErrorCode() string             { return "InvalidDBInstanceState" }
func (e *InvalidDBInstanceStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The parameter group is in use, or it is in a state that is not valid. If you are
// trying to delete the parameter group, you can't delete it when the parameter
// group is in this state.
type InvalidDBParameterGroupStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDBParameterGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBParameterGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBParameterGroupStateFault) ErrorCode() string             { return "InvalidDBParameterGroupState" }
func (e *InvalidDBParameterGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The state of the security group doesn't allow deletion.
type InvalidDBSecurityGroupStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDBSecurityGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSecurityGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSecurityGroupStateFault) ErrorCode() string             { return "InvalidDBSecurityGroupState" }
func (e *InvalidDBSecurityGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The state of the snapshot doesn't allow deletion.
type InvalidDBSnapshotStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDBSnapshotStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSnapshotStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSnapshotStateFault) ErrorCode() string             { return "InvalidDBSnapshotState" }
func (e *InvalidDBSnapshotStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The subnet group can't be deleted because it's in use.
type InvalidDBSubnetGroupStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDBSubnetGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSubnetGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSubnetGroupStateFault) ErrorCode() string             { return "InvalidDBSubnetGroupStateFault" }
func (e *InvalidDBSubnetGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The subnet isn't in the available state.
type InvalidDBSubnetStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidDBSubnetStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDBSubnetStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDBSubnetStateFault) ErrorCode() string             { return "InvalidDBSubnetStateFault" }
func (e *InvalidDBSubnetStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Someone else might be modifying a subscription. Wait a few seconds, and try
// again.
type InvalidEventSubscriptionStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidEventSubscriptionStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEventSubscriptionStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEventSubscriptionStateFault) ErrorCode() string {
	return "InvalidEventSubscriptionState"
}
func (e *InvalidEventSubscriptionStateFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The requested operation can't be performed while the cluster is in this state.
type InvalidGlobalClusterStateFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidGlobalClusterStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidGlobalClusterStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidGlobalClusterStateFault) ErrorCode() string             { return "InvalidGlobalClusterStateFault" }
func (e *InvalidGlobalClusterStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You cannot restore from a virtual private cloud (VPC) backup to a non-VPC DB
// instance.
type InvalidRestoreFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidRestoreFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRestoreFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRestoreFault) ErrorCode() string             { return "InvalidRestoreFault" }
func (e *InvalidRestoreFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested subnet is not valid, or multiple subnets were requested that are
// not all in a common virtual private cloud (VPC).
type InvalidSubnet struct {
	Message *string

	noSmithyDocumentSerde
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

// The subnet group doesn't cover all Availability Zones after it is created
// because of changes that were made.
type InvalidVPCNetworkStateFault struct {
	Message *string

	noSmithyDocumentSerde
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

// An error occurred when accessing an KMS key.
type KMSKeyNotAccessibleFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *KMSKeyNotAccessibleFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSKeyNotAccessibleFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSKeyNotAccessibleFault) ErrorCode() string             { return "KMSKeyNotAccessibleFault" }
func (e *KMSKeyNotAccessibleFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource ID was not found.
type ResourceNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundFault) ErrorCode() string             { return "ResourceNotFoundFault" }
func (e *ResourceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded the maximum number of accounts that you can share a manual DB
// snapshot with.
type SharedSnapshotQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SharedSnapshotQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SharedSnapshotQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SharedSnapshotQuotaExceededFault) ErrorCode() string             { return "SharedSnapshotQuotaExceeded" }
func (e *SharedSnapshotQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed number of snapshots.
type SnapshotQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
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
func (e *SnapshotQuotaExceededFault) ErrorCode() string             { return "SnapshotQuotaExceeded" }
func (e *SnapshotQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon SNS has responded that there is a problem with the specified topic.
type SNSInvalidTopicFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SNSInvalidTopicFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SNSInvalidTopicFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SNSInvalidTopicFault) ErrorCode() string             { return "SNSInvalidTopic" }
func (e *SNSInvalidTopicFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You do not have permission to publish to the SNS topic Amazon Resource Name
// (ARN).
type SNSNoAuthorizationFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SNSNoAuthorizationFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SNSNoAuthorizationFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SNSNoAuthorizationFault) ErrorCode() string             { return "SNSNoAuthorization" }
func (e *SNSNoAuthorizationFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The SNS topic Amazon Resource Name (ARN) does not exist.
type SNSTopicArnNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SNSTopicArnNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SNSTopicArnNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SNSTopicArnNotFoundFault) ErrorCode() string             { return "SNSTopicArnNotFound" }
func (e *SNSTopicArnNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested source could not be found.
type SourceNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SourceNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SourceNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SourceNotFoundFault) ErrorCode() string             { return "SourceNotFound" }
func (e *SourceNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request would cause you to exceed the allowed amount of storage available
// across all instances.
type StorageQuotaExceededFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *StorageQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageQuotaExceededFault) ErrorCode() string             { return "StorageQuotaExceeded" }
func (e *StorageQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Storage of the specified StorageType can't be associated with the DB instance.
type StorageTypeNotSupportedFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *StorageTypeNotSupportedFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StorageTypeNotSupportedFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StorageTypeNotSupportedFault) ErrorCode() string             { return "StorageTypeNotSupported" }
func (e *StorageTypeNotSupportedFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The subnet is already in use in the Availability Zone.
type SubnetAlreadyInUse struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SubnetAlreadyInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetAlreadyInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetAlreadyInUse) ErrorCode() string             { return "SubnetAlreadyInUse" }
func (e *SubnetAlreadyInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided subscription name already exists.
type SubscriptionAlreadyExistFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SubscriptionAlreadyExistFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubscriptionAlreadyExistFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubscriptionAlreadyExistFault) ErrorCode() string             { return "SubscriptionAlreadyExist" }
func (e *SubscriptionAlreadyExistFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided category does not exist.
type SubscriptionCategoryNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SubscriptionCategoryNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubscriptionCategoryNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubscriptionCategoryNotFoundFault) ErrorCode() string             { return "SubscriptionCategoryNotFound" }
func (e *SubscriptionCategoryNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The subscription name does not exist.
type SubscriptionNotFoundFault struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SubscriptionNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubscriptionNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubscriptionNotFoundFault) ErrorCode() string             { return "SubscriptionNotFound" }
func (e *SubscriptionNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }