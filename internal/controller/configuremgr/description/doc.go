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

// Package confdescription defines the configuration informantion of service application
package confdescription

// Doc has config info of each services
type Doc struct {
	Version struct {
		ConfVersion string
	}
	ServiceInfo struct {
		ServiceName        string
		ExecutableFileName string
		AllowedRequester   []string
	}
	// Using this structure is an interesting idea that could be used in the future. See PRs: #20, #383 for quick recovery.
	// ScoringMethod struct {
	// 	LibFile      string
	// 	FunctionName string
	// }
	ResourceType struct {
		IntervalTimeMs int
		MaxCount       int
	}
}
