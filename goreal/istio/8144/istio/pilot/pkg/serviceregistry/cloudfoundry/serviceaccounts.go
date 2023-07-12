// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudfoundry

import "istio.io/istio/pilot/pkg/model"

type serviceAccounts struct {
}

// NewServiceAccounts instantiates the Cloud Foundry service account interface
func NewServiceAccounts() model.ServiceAccounts {
	return &serviceAccounts{}
}

func (sa *serviceAccounts) GetIstioServiceAccounts(hostname model.Hostname, ports []string) []string {
	return nil
}
