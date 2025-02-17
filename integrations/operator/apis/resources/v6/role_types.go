/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package v6

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/integrations/operator/apis/resources"
)

func init() {
	SchemeBuilder.Register(&TeleportRole{}, &TeleportRoleList{})
}

// TeleportRoleSpec defines the desired state of TeleportRole
type TeleportRoleSpec types.RoleSpecV6

// TeleportRoleStatus defines the observed state of TeleportRole
type TeleportRoleStatus struct {
	// Conditions represent the latest available observations of an object's state
	// +optional
	Conditions []metav1.Condition `json:"conditions"`
	// +optional
	TeleportResourceID int64 `json:"teleportResourceID"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TeleportRole is the Schema for the roles API
type TeleportRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TeleportRoleSpec   `json:"spec,omitempty"`
	Status TeleportRoleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TeleportRoleList contains a list of TeleportRole
type TeleportRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeleportRole `json:"items"`
}

func (r TeleportRole) ToTeleport() types.Role {
	return &types.RoleV6{
		Kind:    types.KindRole,
		Version: types.V6,
		Metadata: types.Metadata{
			Name:        r.Name,
			Labels:      r.Labels,
			Description: r.Annotations[resources.DescriptionKey],
		},
		Spec: types.RoleSpecV6(r.Spec),
	}
}

func init() {
	SchemeBuilder.Register(&TeleportRole{}, &TeleportRoleList{})
}

// Marshal serializes a spec into binary data.
func (spec *TeleportRoleSpec) Marshal() ([]byte, error) {
	return (*types.RoleSpecV6)(spec).Marshal()
}

// Unmarshal deserializes a spec from binary data.
func (spec *TeleportRoleSpec) Unmarshal(data []byte) error {
	return (*types.RoleSpecV6)(spec).Unmarshal(data)
}

// DeepCopyInto deep-copies one role spec into another.
// Required to satisfy runtime.Object interface.
func (spec *TeleportRoleSpec) DeepCopyInto(out *TeleportRoleSpec) {
	data, err := spec.Marshal()
	if err != nil {
		panic(err)
	}
	*out = TeleportRoleSpec{}
	if err = out.Unmarshal(data); err != nil {
		panic(err)
	}
}

// StatusConditions returns a pointer to Status.Conditions slice.
func (r *TeleportRole) StatusConditions() *[]metav1.Condition {
	return &r.Status.Conditions
}
