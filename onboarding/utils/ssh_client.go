package utils

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/utils"
	"golang.org/x/crypto/ssh"
)

func runShellCmd(client *ssh.Client, pogIP string, cmd string) (connectErr error, cmdErr error, retStdout *string) {
	var buff bytes.Buffer
	var msg string

	session, err := client.NewSession()
	if err != nil {
		msg := fmt.Sprintf("Error while connecting to the TPA_PRIMARY system ip: %s, error: %s", pogIP, err)
		log.Error(msg)
		return errors.New(msg), nil, nil
	}
	defer session.Close()

	session.Stdout = &buff
	log.Info("Executing command:\t" + cmd)
	err = session.Run(cmd)
	stdout := buff.String()
	msg = fmt.Sprintf("%v", stdout)
	log.Info(msg)
	return nil, err, &stdout
}

func DsmainRestart(pogIP string, tenantId string, tpaSystemId string, cloudPlatform string, region string) bool {
	dsmainCMD := []string{"cnsrun -utility 'bardsmain -s' -commands \" \" -output", "cnsrun -utility 'bardsmain' -commands \" \" -output"}
	sshSecretName := fmt.Sprintf("pod-tenant-%s-%s_sshkey", tenantId, tpaSystemId)
	secretKey, err := utils.GetSecret(sshSecretName, region, cloudPlatform)
	if err != nil {
		log.Error(err)
		return false
	}

	signer, err := ssh.ParsePrivateKey([]byte(secretKey))
	if err != nil {
		msg := fmt.Sprintf("Error while parsing the TPA_PRIMARY Secret String: %v", err)
		log.Error(msg)
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
	client, err := ssh.Dial("tcp", pogIP+":22", config)
	if err != nil {
		msg := fmt.Sprintf("Error while setting up SSH client to the TPA_PRIMARY system ip: %s, error: %s", pogIP, err)
		log.Error(msg)
		return false
	}
	defer client.Close()
	for _, cmd := range dsmainCMD {
		connectErr, cmdErr, _ := runShellCmd(client, pogIP, cmd)
		if connectErr != nil {
			log.Info("Error:%v", connectErr)
			return false
		} else if cmdErr != nil {
			msg := fmt.Sprintf("Unexpected error running the cmd '%s' to the TPA_PRIMARY system ip: %s, error: %s", cmd, pogIP, cmdErr)
			log.Error(msg)
			return false
		}
		// return true
		time.Sleep(15 * time.Second)
	}

	return true
}
