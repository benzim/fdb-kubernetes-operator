/*
 * process_class.go
 *
 * This source file is part of the FoundationDB open source project
 *
 * Copyright 2021 Apple Inc. and the FoundationDB project authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package internal

import (
	fdbtypes "github.com/FoundationDB/fdb-kubernetes-operator/api/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ProcessClassFromLabels extracts the ProcessClass label from the metav1.ObjectMeta.Labels map
func ProcessClassFromLabels(cluster *fdbtypes.FoundationDBCluster, labels map[string]string) fdbtypes.ProcessClass {
	return fdbtypes.ProcessClass(labels[cluster.GetProcessClassLabel()])
}

// GetProcessClassFromMeta fetches the process class from an object's metadata.
func GetProcessClassFromMeta(cluster *fdbtypes.FoundationDBCluster, metadata v1.ObjectMeta) fdbtypes.ProcessClass {
	return ProcessClassFromLabels(cluster, metadata.Labels)
}
