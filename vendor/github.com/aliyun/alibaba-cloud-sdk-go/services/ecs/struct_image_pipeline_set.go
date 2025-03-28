package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ImagePipelineSet is a nested struct in ecs response
type ImagePipelineSet struct {
	CreationTime            string                       `json:"CreationTime" xml:"CreationTime"`
	DeleteInstanceOnFailure bool                         `json:"DeleteInstanceOnFailure" xml:"DeleteInstanceOnFailure"`
	InstanceType            string                       `json:"InstanceType" xml:"InstanceType"`
	InternetMaxBandwidthOut int                          `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	ImagePipelineId         string                       `json:"ImagePipelineId" xml:"ImagePipelineId"`
	VSwitchId               string                       `json:"VSwitchId" xml:"VSwitchId"`
	SystemDiskSize          int                          `json:"SystemDiskSize" xml:"SystemDiskSize"`
	Description             string                       `json:"Description" xml:"Description"`
	BaseImage               string                       `json:"BaseImage" xml:"BaseImage"`
	ResourceGroupId         string                       `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ImageName               string                       `json:"ImageName" xml:"ImageName"`
	BaseImageType           string                       `json:"BaseImageType" xml:"BaseImageType"`
	Name                    string                       `json:"Name" xml:"Name"`
	BuildContent            string                       `json:"BuildContent" xml:"BuildContent"`
	RepairMode              string                       `json:"RepairMode" xml:"RepairMode"`
	TestContent             string                       `json:"TestContent" xml:"TestContent"`
	ImageFamily             string                       `json:"ImageFamily" xml:"ImageFamily"`
	NvmeSupport             string                       `json:"NvmeSupport" xml:"NvmeSupport"`
	ToRegionIds             ToRegionIds                  `json:"ToRegionIds" xml:"ToRegionIds"`
	AddAccounts             AddAccounts                  `json:"AddAccounts" xml:"AddAccounts"`
	ImportImageOptions      ImportImageOptions           `json:"ImportImageOptions" xml:"ImportImageOptions"`
	AdvancedOptions         AdvancedOptions              `json:"AdvancedOptions" xml:"AdvancedOptions"`
	ImageOptions            ImageOptions                 `json:"ImageOptions" xml:"ImageOptions"`
	Tags                    TagsInDescribeImagePipelines `json:"Tags" xml:"Tags"`
}
