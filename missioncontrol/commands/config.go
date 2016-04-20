package commands


import (
	"github.com/jfrogdev/jfrog-cli-go/utils/ioutils"
	"github.com/jfrogdev/jfrog-cli-go/utils/config"
	"github.com/jfrogdev/jfrog-cli-go/utils/cliutils"
	"fmt"
)

func GetConfig() *config.MissionControlDetails {
	return config.ReadMissionControlConf()
}

func ShowConfig() {
	details := config.ReadMissionControlConf()
	if details.Url != "" {
		fmt.Println("Url: " + details.Url)
	}
	if details.User != "" {
		fmt.Println("User: " + details.User)
	}
	if details.Password != "" {
		fmt.Println("Password: ***")
	}
}

func ClearConfig() {
	config.SaveMissionControlConf(new(config.MissionControlDetails))
}

func Config(details, defaultDetails *config.MissionControlDetails, interactive bool) *config.MissionControlDetails {
	if details == nil {
		details = new(config.MissionControlDetails)
	}
	if interactive {
		if defaultDetails == nil {
			defaultDetails = config.ReadMissionControlConf()
		}
		if details.Url == "" {
			ioutils.ScanFromConsole("Mission Control URL", &details.Url, defaultDetails.Url)
		}
		ioutils.ReadCredentialsFromConsole(details, defaultDetails)
	}
	details.Url = cliutils.AddTrailingSlashIfNeeded(details.Url)
	config.SaveMissionControlConf(details)
	return details
}

type ConfigFlags struct {
	MissionControlDetails *config.MissionControlDetails
	Interactive           bool
}