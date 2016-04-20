package commands

import (
	"encoding/json"
	"github.com/jfrogdev/jfrog-cli-go/bintray/utils"
	"github.com/jfrogdev/jfrog-cli-go/utils/cliutils"
	"github.com/jfrogdev/jfrog-cli-go/utils/config"
	"github.com/jfrogdev/jfrog-cli-go/utils/ioutils"
	"strings"
	"sync"
)

func DownloadVersion(versionDetails *utils.VersionDetails, flags *utils.DownloadFlags) {
	ioutils.CreateTempDirPath()
	defer ioutils.RemoveTempDir()

	if flags.BintrayDetails.User == "" {
		flags.BintrayDetails.User = versionDetails.Subject
	}
	path := BuildDownloadVersionUrl(versionDetails, flags.BintrayDetails)
	httpClientsDetails := utils.GetBintrayHttpClientDetails(flags.BintrayDetails)
	resp, body, _, _ := ioutils.SendGet(path, true, httpClientsDetails)
	if resp.StatusCode != 200 {
		cliutils.Exit(cliutils.ExitCodeError, resp.Status+". " + ioutils.ReadHttpMessage(body))
	}
	var results []VersionFilesResult
	err := json.Unmarshal(body, &results)
	cliutils.CheckError(err)

	downloadFiles(results, versionDetails, flags)
}

func BuildDownloadVersionUrl(versionDetails *utils.VersionDetails, bintrayDetails *config.BintrayDetails) string {
    return bintrayDetails.ApiUrl + "packages/" + versionDetails.Subject + "/" +
		versionDetails.Repo + "/" + versionDetails.Package + "/versions/" + versionDetails.Version + "/files"
}

func downloadFiles(results []VersionFilesResult, versionDetails *utils.VersionDetails,
	flags *utils.DownloadFlags) {

	size := len(results)
	var wg sync.WaitGroup
	for i := 0; i < flags.Threads; i++ {
		wg.Add(1)
		go func(threadId int) {
			logMsgPrefix := cliutils.GetLogMsgPrefix(threadId, false)
			for j := threadId; j < size; j += flags.Threads {
                pathDetails := &utils.PathDetails{
                    Subject: versionDetails.Subject,
                    Repo:    versionDetails.Repo,
                    Path:    results[j].Path}

				utils.DownloadBintrayFile(flags.BintrayDetails, pathDetails,
					flags, logMsgPrefix)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func CreateVersionDetailsForDownloadVersion(versionStr string) *utils.VersionDetails {
	parts := strings.Split(versionStr, "/")
	if len(parts) != 4 {
		cliutils.Exit(cliutils.ExitCodeError, "Argument format should be subject/repository/package/version. Got "+versionStr)
	}
	return utils.CreateVersionDetails(versionStr)
}

type VersionFilesResult struct {
	Path string
}
