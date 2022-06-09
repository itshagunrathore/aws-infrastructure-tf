package utils

import (
	"bytes"
	"errors"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/utils"
	"golang.org/x/crypto/ssh"
)

func runShellCmd(client *ssh.Client, dsaIp string, cmd string) (connectErr error, cmdErr error, retStdout *string) {
	var buff bytes.Buffer
	var msg string

	session, err := client.NewSession()
	if err != nil {
		msg := fmt.Sprintf("Error while connecting to the TPA_PRIMARY system ip: %s, error: %s", dsaIp, err)
		log.Error(msg)
		return errors.New(msg), nil, nil
	}
	defer session.Close()

	session.Stdout = &buff

	err = session.Run(cmd)
	stdout := buff.String()
	msg = fmt.Sprintf("%v", stdout)
	fmt.Println(cmd)
	fmt.Println(msg)

	return nil, err, &stdout
}

func DsmainRestart(pogIp string, tenantId string, TPASystemId string, cloudPlatform string) bool {
	dsmainCMD := []string{"cnsrun -utility 'bardsmain -s' -commands \" \" -output", "cnsrun -utility 'bardsmain' -commands \" \" -output"}
	secretKey := utils.GetSecret(tenantId, TPASystemId, cloudPlatform)

	signer, err := ssh.ParsePrivateKey([]byte(secretKey))
	if err != nil {
		msg := fmt.Sprintf("Error while parsing the TPA_PRIMARY Secret String: %v", err)
		fmt.Println(msg)
		return false
	}

	config := &ssh.ClientConfig{
		User: "ec2-user",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Setup SSH Client
	client, err := ssh.Dial("tcp", dsaIp+":22", config)
	if err != nil {
		msg := fmt.Sprintf("Error while setting up SSH client to the TPA_PRIMARY system ip: %s, error: %s", pogIp, err)
		fmt.Println(msg)
		return false
	}
	defer client.Close()
	for _, cmd := range dsmainCMD {
		connectErr, cmdErr, _ := runShellCmd(client, dsaIp, cmd)
		if connectErr != nil {
			fmt.Printf("Error:%v", connectErr)
			return false
		} else if cmdErr != nil {
			msg := fmt.Sprintf("Unexpected error running the cmd '%s' to the TPA_PRIMARY system ip: %s, error: %s", cmd, dsaIp, cmdErr)
			fmt.Println(msg)
			return false
		}
		return true
	}

	return true
}
