package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	// VersionString is the version of the CLI
	VersionString string = "No version provided"
	// ProfileFlag is the AWS crendentials profile passed in
	ProfileFlag string
	// CredentialsFile is a boolean for the credentials provider logic
	CredentialsFile bool
	// RegionFlag is the AWS Region to target for the execution
	RegionFlag string
	// Retain is the number of versions to retain excluding $LATEST
	Retain int8
	// Verbose is to enable debug output
	Verbose bool
	// DryRun is to enable a preview of what an actual execution would do
	DryRun bool
	// LambdaListFile points a file that contains a listof Lambdas to delete
	LambdaListFile string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&RegionFlag, "region", "r", "", "Specify the desired AWS region to target.")
	rootCmd.PersistentFlags().StringVarP(&ProfileFlag, "profile", "p", "", "Specify the AWS profile to leverage for authentication.")
	rootCmd.PersistentFlags().StringVarP(&LambdaListFile, "listFile", "l", "", "Specify a file containing Lambdas to delete.")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Set to true to enable debugging (bool)")
	rootCmd.PersistentFlags().BoolVarP(&CredentialsFile, "enableSharedCredentials", "s", false, "Leverages the default ~/.aws/credentials file (bool)")
	rootCmd.PersistentFlags().BoolVarP(&DryRun, "dryrun", "d", false, "Executes a dry run (bool)")

	// Establish logging default
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "01/02/06",
		FullTimestamp:   true,
	})
	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
}

// Execute is the main execution function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if Verbose {
			log.WithFields(log.Fields{
				"package":  "cmd",
				"file":     "root.go",
				"function": "Execute",
				"error":    err,
				"data":     nil,
			}).Fatal("Error executing the CLI!")
		} else {
			log.Fatal(err.Error())
		}
	}
}
