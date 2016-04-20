package rtinstances

import (
	"github.com/jfrogdev/jfrog-cli-go/missioncontrol/utils"
	"github.com/jfrogdev/jfrog-cli-go/utils/cliutils"
	"github.com/jfrogdev/jfrog-cli-go/utils/ioutils"
	"github.com/jfrogdev/jfrog-cli-go/utils/config"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

func AttachLic(instanceName string, flags *AttachLicFlags) {
	licensePath := flags.LicensePath
	prepareLicenseFile(licensePath)
	bucketId := flags.BucketId
	postContent := utils.LicenseRequestContent{
		Name: 	  	 instanceName,
		BucketKey:	 flags.BucketKey,
		NodeID:	     flags.NodeId}
	requestContent, err := json.Marshal(postContent)
	if err != nil {
		cliutils.Exit(cliutils.ExitCodeError, "Failed to marshal json. " + cliutils.GetDocumentationMessage())
	}
	missionControlUrl := flags.MissionControlDetails.Url + "api/v1/buckets/" + bucketId + "/licenses";
	httpClientDetails := utils.GetMissionControlHttpClientDetails(flags.MissionControlDetails)
	resp, body := ioutils.SendPost(missionControlUrl, requestContent, httpClientDetails)
	if resp.StatusCode != 200 {
		if licensePath != "" {
			os.Remove(licensePath)
		}
		cliutils.Exit(cliutils.ExitCodeError, resp.Status + ". " + ioutils.ReadHttpMessage(body))
	}
	fmt.Println("Mission Control response: " + resp.Status)
	if licensePath == "" {
		requestContent, err := json.Marshal(extractJsonValue(body))
		if err != nil{
			panic(err)
		}
		fmt.Println(string(requestContent))
	} else{
		licenseKey := getLicenseFromJson(body)
		saveLicense(licensePath, licenseKey)
	}
}

func getLicenseFromJson(body []byte) (licenseKey [] byte) {
	licenseKey = []byte(extractJsonValue(body).LicenseKey)
	return
}

func extractJsonValue(body []byte) Message {
	data := &Data{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	return data.Data
}

func prepareLicenseFile(filepath string) {
	saveLicense(filepath, nil)
}

func saveLicense(filepath string, content []byte){
	if filepath == "" {
		return
	}
	_, dir := ioutils.GetFileAndDirFromPath(filepath)
	isPathExists := ioutils.IsPathExists(dir)
	if !isPathExists {
		os.MkdirAll(dir, 0700)
	}
	err := ioutil.WriteFile(filepath, content, 0600)
	if err != nil {
		panic(err)
	}
}

type AttachLicFlags struct {
	MissionControlDetails *config.MissionControlDetails
	LicensePath 	      string
	NodeId 			      string
	BucketKey 			  string
	BucketId 			  string
}

type Message struct {
	LicenseKey string `json:"licenseKey,omitempty"`
}

type Data struct {
	Data Message
}
