/*******************************************************************************
 * Copyright 2019 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

package requestervalidator

import (
	"errors"

	"github.com/lf-edge/edge-home-orchestration-go/internal/common/requestervalidator/requesterstore"
	"github.com/lf-edge/edge-home-orchestration-go/internal/db/bolt/common"
)

// IRequesterValidator provides interfaces for the requestervalidator
type IRequesterValidator interface {
	GetRequester(serviceName string) ([]string, error)
	StoreRequesterInfo(serviceName string, requesters []string)
	CheckRequester(serviceName, requester string) error
}

const notAllowedServiceExecution = "not allowed service execution"

// RequesterValidator structure
type RequesterValidator struct{}

// CheckRequester checks the requester's ability to execute the service
func (r RequesterValidator) CheckRequester(serviceName, requester string) error {
	stored, err := r.GetRequester(serviceName)
	if err != nil {
		return err
	}

	if allowed := common.HasElem(stored, requester); allowed {
		return nil
	}
	return errors.New(notAllowedServiceExecution)
}

// GetRequester gets servicename requester relation
func (RequesterValidator) GetRequester(serviceName string) ([]string, error) {
	return requesterstore.GetInstance().GetRequester(serviceName)
}

// StoreRequesterInfo stores info about requesters' servicename
func (RequesterValidator) StoreRequesterInfo(serviceName string, requesters []string) {
	requesterstore.GetInstance().StoreRequesterInfo(serviceName, requesters)
}
