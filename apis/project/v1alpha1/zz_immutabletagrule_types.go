// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ImmutableTagRuleInitParameters struct {

	// (Boolean) Specify if the rule is disable or not. Defaults to false
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// (String) For the repositories excluding.
	// Use doublestar pattern for path matching.
	RepoExcluding *string `json:"repoExcluding,omitempty" tf:"repo_excluding,omitempty"`

	// (String) For the repositories matching.
	// Use doublestar pattern for path matching.
	RepoMatching *string `json:"repoMatching,omitempty" tf:"repo_matching,omitempty"`

	// (String) For the tag excluding.
	// Use doublestar pattern for path matching.
	TagExcluding *string `json:"tagExcluding,omitempty" tf:"tag_excluding,omitempty"`

	// (String) For the tag matching.
	// Use doublestar pattern for path matching.
	TagMatching *string `json:"tagMatching,omitempty" tf:"tag_matching,omitempty"`
}

type ImmutableTagRuleObservation struct {

	// (Boolean) Specify if the rule is disable or not. Defaults to false
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The project id of which you would like to apply this policy.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (String) For the repositories excluding.
	// Use doublestar pattern for path matching.
	RepoExcluding *string `json:"repoExcluding,omitempty" tf:"repo_excluding,omitempty"`

	// (String) For the repositories matching.
	// Use doublestar pattern for path matching.
	RepoMatching *string `json:"repoMatching,omitempty" tf:"repo_matching,omitempty"`

	// (String) For the tag excluding.
	// Use doublestar pattern for path matching.
	TagExcluding *string `json:"tagExcluding,omitempty" tf:"tag_excluding,omitempty"`

	// (String) For the tag matching.
	// Use doublestar pattern for path matching.
	TagMatching *string `json:"tagMatching,omitempty" tf:"tag_matching,omitempty"`
}

type ImmutableTagRuleParameters struct {

	// (Boolean) Specify if the rule is disable or not. Defaults to false
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// (String) The project id of which you would like to apply this policy.
	// +crossplane:generate:reference:type=github.com/globallogicuki/provider-harbor/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// (String) For the repositories excluding.
	// Use doublestar pattern for path matching.
	// +kubebuilder:validation:Optional
	RepoExcluding *string `json:"repoExcluding,omitempty" tf:"repo_excluding,omitempty"`

	// (String) For the repositories matching.
	// Use doublestar pattern for path matching.
	// +kubebuilder:validation:Optional
	RepoMatching *string `json:"repoMatching,omitempty" tf:"repo_matching,omitempty"`

	// (String) For the tag excluding.
	// Use doublestar pattern for path matching.
	// +kubebuilder:validation:Optional
	TagExcluding *string `json:"tagExcluding,omitempty" tf:"tag_excluding,omitempty"`

	// (String) For the tag matching.
	// Use doublestar pattern for path matching.
	// +kubebuilder:validation:Optional
	TagMatching *string `json:"tagMatching,omitempty" tf:"tag_matching,omitempty"`
}

// ImmutableTagRuleSpec defines the desired state of ImmutableTagRule
type ImmutableTagRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImmutableTagRuleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ImmutableTagRuleInitParameters `json:"initProvider,omitempty"`
}

// ImmutableTagRuleStatus defines the observed state of ImmutableTagRule.
type ImmutableTagRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImmutableTagRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImmutableTagRule is the Schema for the ImmutableTagRules API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,harbor}
type ImmutableTagRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImmutableTagRuleSpec   `json:"spec"`
	Status            ImmutableTagRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImmutableTagRuleList contains a list of ImmutableTagRules
type ImmutableTagRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImmutableTagRule `json:"items"`
}

// Repository type metadata.
var (
	ImmutableTagRule_Kind             = "ImmutableTagRule"
	ImmutableTagRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImmutableTagRule_Kind}.String()
	ImmutableTagRule_KindAPIVersion   = ImmutableTagRule_Kind + "." + CRDGroupVersion.String()
	ImmutableTagRule_GroupVersionKind = CRDGroupVersion.WithKind(ImmutableTagRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ImmutableTagRule{}, &ImmutableTagRuleList{})
}
