package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/danieljoos/wincred"
)

var (
	get             bool
	set             bool
	profile         string
	accessKeyId     string
	secretAccessKey string
)

func init() {
	flag.BoolVar(&get, "get", false, "retrieve credentials from the store")
	flag.BoolVar(&set, "set", false, "save or update credentials in the store")
	flag.StringVar(&profile, "profile", "", "aws cli profile name")
}
func main() {
	flag.Parse()

	if profile == "" {
		fmt.Println("--profile is required")
		os.Exit(1)
	}

	target := "awscli://" + profile

	if set {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("AccessKeyID: ")
		accessKeyId, _ = reader.ReadString('\n')
		accessKeyId = strings.TrimSuffix(accessKeyId, "\r\n")

		fmt.Print("SecretAccessKey: ")
		secretAccessKey, _ = reader.ReadString('\n')
		secretAccessKey = strings.TrimSuffix(secretAccessKey, "\r\n")

		credJson, err := json.MarshalIndent(struct {
			Version         int    `json:"Version"`
			AccessKeyId     string `json:"AccessKeyId"`
			SecretAccessKey string `json:"SecretAccessKey"`
		}{
			Version:         1,
			AccessKeyId:     accessKeyId,
			SecretAccessKey: secretAccessKey,
		}, "", "  ")
		if err != nil {
			fmt.Printf("failed to serialize aws credential: %s\n", err)
			os.Exit(1)
		}

		cred := wincred.NewGenericCredential(target)
		cred.CredentialBlob = credJson
		err = cred.Write()
		if err != nil {
			fmt.Printf("failed to write credential %s: %s\n", target, err)
			os.Exit(1)
		}
	}

	if get {
		cred, err := wincred.GetGenericCredential(target)
		if err != nil {
			fmt.Printf("failed to get credential %s: %s\n", target, err)
			os.Exit(1)
		}

		fmt.Println(string(cred.CredentialBlob))
		os.Exit(0)
	}
}
