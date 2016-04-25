package utils

 import (
	 "github.com/jfrogdev/jfrog-cli-go/utils/config"
	 "github.com/jfrogdev/jfrog-cli-go/utils/ioutils"
 )

func GetMissionControlHttpClientDetails(missionControlDetails *config.MissionControlDetails) ioutils.HttpClientDetails {
	return ioutils.HttpClientDetails{
		User:     missionControlDetails.User,
		Password: missionControlDetails.Password,
		Headers:  map[string]string{"Content-Type": "application/json"}}
}

type LicenseRequestContent struct {
	Name      string `json:"instanceName,omitempty"`
	BucketKey string `json:"bucketKey,omitempty"`
	NodeID 	  string `json:"nodeId,omitempty"`
}

type ArtifactoryInstanceDetails struct {
	Url      string `json:"url,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string	`json:"instanceName,omitempty"`
}

