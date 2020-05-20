package v1

import (
	coreV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//HelixSaga describes a HelixSaga resource
type HelixSaga struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	metav1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec HelixSagaSpec `json:"spec"`

	Status HelixSagaSpecStatus `json:"status"`
}

//HelixSagaSpec is the spec for a HelixSaga resource
type HelixSagaSpec struct {
	ConfigMapName       string              `json:"config_map_name"`
	VersionSpec         NginxPhpFpmSpec     `json:"version_spec"`
	ApiSpec             NginxPhpFpmSpec     `json:"api_spec"`
	GameSpec            NginxPhpFpmSpec     `json:"game_spec"`
	PayNotifySpec       NginxPhpFpmSpec     `json:"pay_notify_spec"`
	GmtSpec             NginxPhpFpmSpec     `json:"gmt_spec"`
	FriendSpec          PhpSwooleSpec       `json:"friend_spec"`
	QueueSpec           PhpSwooleSpec       `json:"queue_spec"`
	RankSpec            PhpSwooleSpec       `json:"rank_spec"`
	ChatSpec            PhpWorkermanSpec    `json:"chat_spec"`
	HeartSpec           PhpWorkermanSpec    `json:"heart_spec"`
	CampaignSpec        CampaignSpec        `json:"campaign_spec"`
	GuildWarSpec        GuildWarSpec        `json:"guild_war_spec"`
	AppNotificationSpec AppNotificationSpec `json:"app_notification_spec"`
}

//NginxPhpFpmSpec is the sub spec for a HelixSaga resource
type NginxPhpFpmSpec struct {
	// Name of the container specified as a DNS_LABEL.
	// Each container in a pod must have a unique name (DNS_LABEL).
	// Cannot be updated.
	Name string `json:"name" protobuf:"bytes,1,rep,name=name"`
	// Replicas is the number of desired replicas.
	// This is a pointer to distinguish between explicit zero and unspecified.
	// Defaults to 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	// +optional
	Replicas *int32 `json:"replicas,omitempty" protobuf:"bytes,2,rep,name=replicas"`
	// Docker image name.
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// This field is optional to allow higher level config management to default or override
	// container images in workload controllers like Deployments and StatefulSets.
	// +optional
	Image string `json:"image,omitempty" protobuf:"bytes,3,opt,name=image"`
	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
	// If specified, these secrets will be passed to individual puller implementations for them to use. For example,
	// in the case of docker, only DockerConfig type secrets are honored.
	// More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	ImagePullSecrets []coreV1.LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,4,rep,name=imagePullSecrets"`
	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	Env []coreV1.EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,5,rep,name=env"`
	// Pod volumes to mount into the container's filesystem.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=mountPath
	// +patchStrategy=merge
	VolumeMounts []coreV1.VolumeMount `json:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,6,rep,name=volumeMounts"`
}

//PhpSwooleSpec is the sub spec for a HelixSaga resource
type PhpSwooleSpec struct {
	Name             string `json:"name"`
	Replicas         *int32 `json:"replicas"`
	Image            string `json:"image"`
	ImagePullSecrets string `json:"imagePullSecrets"`
}

//PhpWorkermanSpec is the sub spec for a HelixSaga resource
type PhpWorkermanSpec struct {
	Name                   string `json:"name"`
	RegisterReplicas       *int32 `json:"register_replicas"`
	GatewayReplicas        *int32 `json:"gateway_replicas"`
	BusinessWorkerReplicas *int32 `json:"business_worker_replicas"`
	Replicas               *int32 `json:"replicas"`
	Image                  string `json:"image"`
	ImagePullSecrets       string `json:"imagePullSecrets"`
}

//Campaign is the sub spec for a HelixSaga resource
type CampaignSpec struct {
	Name             string `json:"name"`
	Replicas         *int32 `json:"replicas"`
	Image            string `json:"image"`
	ImagePullSecrets string `json:"imagePullSecrets"`
}

//GuildWarSpec is the sub spec for a HelixSaga resource
type GuildWarSpec struct {
	Name             string `json:"name"`
	RegisterReplicas *int32 `json:"register_replicas"`
	GatewayReplicas  *int32 `json:"gateway_replicas"`
	Image            string `json:"image"`
	ImagePullSecrets string `json:"imagePullSecrets"`
}

//AppNotificationSpec is the sub spec for a HelixSaga resource
type AppNotificationSpec struct {
	Name             string `json:"name"`
	DispatchReplicas *int32 `json:"dispatch_replicas"`
	LogicReplicas    *int32 `json:"logic_replicas"`
	Image            string `json:"image"`
	ImagePullSecrets string `json:"imagePullSecrets"`
}

// HelixSagaStatus is the status for a HelixSaga resource
type HelixSagaSpecStatus struct {
	VersionStatus         CommonStatus `json:"version_status"`
	ApiStatus             CommonStatus `json:"api_status"`
	GameStatus            CommonStatus `json:"game_status"`
	PayNotifyStatus       CommonStatus `json:"pay_notify_status"`
	GmtStatus             CommonStatus `json:"gmt_status"`
	FriendStatus          CommonStatus `json:"friend_status"`
	QueueStatus           CommonStatus `json:"queue_status"`
	RankStatus            CommonStatus `json:"rank_status"`
	ChatStatus            CommonStatus `json:"chat_status"`
	HeartStatus           CommonStatus `json:"heart_status"`
	CampaignStatus        CommonStatus `json:"campaign_status"`
	GuildWarStatus        CommonStatus `json:"guild_war_status"`
	AppNotificationStatus CommonStatus `json:"app_notification_status"`
}

//CommonStatus is the sub status for a HelixSaga resource
type CommonStatus struct {
	// The generation observed by the deployment controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`

	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas,omitempty" protobuf:"varint,2,opt,name=replicas"`

	// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	// +optional
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty" protobuf:"varint,3,opt,name=updatedReplicas"`

	// Total number of ready pods targeted by this deployment.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty" protobuf:"varint,7,opt,name=readyReplicas"`

	// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	// +optional
	AvailableReplicas int32 `json:"availableReplicas,omitempty" protobuf:"varint,4,opt,name=availableReplicas"`

	// Total number of unavailable pods targeted by this deployment. This is the total number of
	// pods that are still required for the deployment to have 100% available capacity. They may
	// either be pods that are running but not yet available or pods that still have not been created.
	// +optional
	UnavailableReplicas int32 `json:"unavailableReplicas,omitempty" protobuf:"varint,5,opt,name=unavailableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//HelixSagaList is a list of HelixSaga resources
type HelixSagaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []HelixSaga `json:"items"`
}
