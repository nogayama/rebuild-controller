// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/rebuildcontroller/v1alpha1.RebuildPolicy":       schema_pkg_apis_rebuildcontroller_v1alpha1_RebuildPolicy(ref),
		"./pkg/apis/rebuildcontroller/v1alpha1.RebuildPolicySpec":   schema_pkg_apis_rebuildcontroller_v1alpha1_RebuildPolicySpec(ref),
		"./pkg/apis/rebuildcontroller/v1alpha1.RebuildPolicyStatus": schema_pkg_apis_rebuildcontroller_v1alpha1_RebuildPolicyStatus(ref),
	}
}

func schema_pkg_apis_rebuildcontroller_v1alpha1_RebuildPolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RebuildPolicy is the Schema for the rebuildpolicies API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/rebuildcontroller/v1alpha1.RebuildPolicySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/rebuildcontroller/v1alpha1.RebuildPolicyStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/rebuildcontroller/v1alpha1.RebuildPolicySpec", "./pkg/apis/rebuildcontroller/v1alpha1.RebuildPolicyStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_rebuildcontroller_v1alpha1_RebuildPolicySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RebuildPolicySpec defines the desired state of RebuildPolicy",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_rebuildcontroller_v1alpha1_RebuildPolicyStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RebuildPolicyStatus defines the observed state of RebuildPolicy",
				Type:        []string{"object"},
			},
		},
	}
}