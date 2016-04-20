package rtinstances

import (
	"github.com/jfrogdev/jfrog-cli-go/missioncontrol/utils"
	"github.com/jfrogdev/jfrog-cli-go/utils/cliutils"
	"github.com/jfrogdev/jfrog-cli-go/utils/ioutils"
	"github.com/jfrogdev/jfrog-cli-go/utils/config"
	"encoding/json"
	"fmt"
)

func AddInstance(instanceName string, flags *AddInstanceFlags) {
	data := AddInstanceRequestContent{
		Name: 	  	 instanceName,
		Url : 	  	 flags.ArtifactoryInstanceDetails.Url,
		User: 	  	 flags.ArtifactoryInstanceDetails.User,
		Password: 	 flags.ArtifactoryInstanceDetails.Password,
		Description: flags.Description,
		Location: 	 flags.Location}
	requestContent, err := json.Marshal(data)
	if err != nil {
		cliutils.Exit(cliutils.ExitCodeError, "Failed to execute request. " + cliutils.GetDocumentationMessage())
	}
	missionControlUrl := flags.MissionControlDetails.Url + "api/v1/instances";
	HttpAdditionalData := utils.GetMissionControlHttpClientDetails(flags.MissionControlDetails)
	resp, body := ioutils.SendPost(missionControlUrl, requestContent, HttpAdditionalData)
	if resp.StatusCode != 200 {
		cliutils.Exit(cliutils.ExitCodeError, resp.Status + ". " + ioutils.ReadHttpMessage(body))
	}
	fmt.Println("Mission Control response: " + resp.Status)
}

type AddInstanceFlags struct {
	MissionControlDetails      *config.MissionControlDetails
	Description 			   string
	Location 				   string
	NodeId 					   string
	ArtifactoryInstanceDetails *utils.ArtifactoryInstanceDetails
}

type AddInstanceRequestContent struct {
	Name        string `json:"name,omitempty"`
	Url        	string `json:"url,omitempty"`
	User        string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Location 	string `json:"location,omitempty"`
	Description string `json:"description,omitempty"`
}
