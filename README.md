# Overview
JFrog CLI is a compact and smart client that provides a simple interface that automates access to Artifactory and Bintray through their respective REST APIs. By using the JFrog CLI, you can greatly simplify your automation scripts making them more readable and easier to maintain. Several features of the JFrog CLI makes your scripts more efficient and reliable:

- Multi-threaded upload and download of artifacts make builds run faster
- Checksum optimization reduces redundant file transfers
- Wildcards and regular expressions give you an easy way to collect all the artifacts you wish to upload or download.
- "Dry run" gives you a preview of file transfer operations before you actually run them

# Download and Installation

You can get the executable directly from the [JFrog CLI Download Page](https://www.jfrog.com/getcli/), or you can download the source files from this GitHub project and build it yourself.

## Building the Executable

JFrog CLI is written in the [Go programming language](https://golang.org/), so to build the CLI yourself, you first need to have Go installed and configured on your machine.

### Setup Go

To download and install `Go`, please refer to the [Go documentation](https://golang.org/doc/install).

Navigate to the directory where you want to create the jfrog-cli-go **<TBD Link>** project, and set the value of the GOPATH environment variable to the full path of this directory.

### Download and Build the CLI

To download the jfrog-cli-go project, execute the following command:
````
$ go get github.com/jfrogdev/jfrog-cli-go/...
````
Go will download and build the project on your machine. Once complete, you will find the JFrog CLI executable under your `$GOPATH/bin` directory.

### Integration tests
In order to execute the tests, run the following commands:
````
$ go test -v github.com/jfrogdev/jfrog-cli-go/jfrog
````

By default the tests are using the following settings:

* Artifactory url: http://localhost:8081/artifatory
* user: admin
* passowrd: password
* apikey: EMPTY

Those settings can be customized by using the --url, --user, --password and --apikey flags while executing the go test.<br/>
For example:
````
$ go test -v github.com/jfrogdev/jfrog-cli-go/jfrog --url=http://yourArtifactoryUrl/artifactory --user=user --password=password --apikey=apikey
````
* Running the tests will create two repositories: jfrog-cli-tests-repo and jfrog-cli-tests-repo1.<br/>
  By the end of the tests the content of those repositories will be deleted.


# Using JFrog CLI with Artifactory and Bintray
JFrog CLI can be used for quick and easy file management with both Artifactory and Bintray, and has a dedicated set of commands for each product. To learn how to use JFrog CLI, please refer to the relevant documentation through the corresponding link below: 
* [Using JFrog CLI with Artifactory](https://www.jfrog.com/confluence/display/RTF/JFrog+CLI)
* [Using JFrog CLI with Bintray](https://bintray.com/docs/usermanual/cli/cli_jfrogcli.html)
