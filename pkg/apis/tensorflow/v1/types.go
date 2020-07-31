// Copyright 2020 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	common "github.com/kubeflow/common/pkg/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=amltfjob

// Represents a AmlTFJob resource.
type AmlTFJob struct {
	// Standard Kubernetes type metadata.
	metav1.TypeMeta `json:",inline"`

	// Standard Kubernetes object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired state of the TFJob.
	// +optional
	Spec TFJobSpec `json:"spec,omitempty"`

	// Most recently observed status of the TFJob.
	// Populated by the system.
	// Read-only.
	// +optional
	Status common.JobStatus `json:"status,omitempty"`
}

// TFJobSpec is a desired state description of the TFJob.
type TFJobSpec struct {
	// Specifies the duration (in seconds) since startTime during which the job can remain active
	// before it is terminated. Must be a positive integer.
	// This setting applies only to pods where restartPolicy is OnFailure or Always.
	// +optional
	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty"`

	// Number of retries before marking this job as failed.
	// +optional
	BackoffLimit *int32 `json:"backoffLimit,omitempty"`

	// SuccessPolicy defines the policy to mark the TFJob as succeeded.
	// Default to "", using the default rules.
	// +optional
	SuccessPolicy *SuccessPolicy `json:"successPolicy,omitempty"`

	// Defines the policy for cleaning up pods after the TFJob completes.
	// Defaults to Running.
	// +optional
	CleanPodPolicy *common.CleanPodPolicy `json:"cleanPodPolicy,omitempty"`

	// Defines the TTL for cleaning up finished TFJobs (temporary
	// before kubernetes adds the cleanup controller).
	// It may take extra ReconcilePeriod seconds for the cleanup, since
	// reconcile gets called periodically.
	// Defaults to infinite.
	// +optional
	TTLSecondsAfterFinished *int32 `json:"ttlSecondsAfterFinished,omitempty"`

	// A map of TFReplicaType (type) to ReplicaSpec (value). Specifies the TF cluster configuration.
	// For example,
	//   {
	//     "PS": ReplicaSpec,
	//     "Worker": ReplicaSpec,
	//   }
	TFReplicaSpecs map[TFReplicaType]*common.ReplicaSpec `json:"tfReplicaSpecs"`

	// A switch to enable dynamic worker
	EnableDynamicWorker bool `json:"enableDynamicWorker,omitempty"`
}

// TFReplicaType is the type for TFReplica. Can be one of: "Chief"/"Master" (semantically equivalent),
// "Worker", "PS", or "Evaluator".
type TFReplicaType common.ReplicaType

const (
	// TFReplicaTypePS is the type for parameter servers of distributed TensorFlow.
	TFReplicaTypePS TFReplicaType = "PS"

	// TFReplicaTypeWorker is the type for workers of distributed TensorFlow.
	// This is also used for non-distributed TensorFlow.
	TFReplicaTypeWorker TFReplicaType = "Worker"

	// TFReplicaTypeChief is the type for chief worker of distributed TensorFlow.
	// If there is "chief" replica type, it's the "chief worker".
	// Else, worker:0 is the chief worker.
	TFReplicaTypeChief TFReplicaType = "Chief"

	// TFReplicaTypeMaster is the type for master worker of distributed TensorFlow.
	// This is similar to chief, and kept just for backwards compatibility.
	TFReplicaTypeMaster TFReplicaType = "Master"

	// TFReplicaTypeEval is the type for evaluation replica in TensorFlow.
	TFReplicaTypeEval TFReplicaType = "Evaluator"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=amltfjobs

// AmlTFJobList is a list of TFJobs.
type AmlTFJobList struct {
	// Standard type metadata.
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// List of TFJobs.
	Items []AmlTFJob `json:"items"`
}
