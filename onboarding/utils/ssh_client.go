package utils

import (
	"bytes"
	"errors"
	"fmt"

	"golang.org/x/crypto/ssh"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
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

func DsmainRestart() bool {
	dsaIp := "100.21.233.93"
	dsmainCMD := ("cnsrun -utility 'bardsmain -s' -commands " " -output", "cnsrun -utility 'bardsmain' -commands " " -output")
	secretKey := GetSecret()

	signer, err := ssh.ParsePrivateKey([]byte(secretKey))
	if err != nil {
		msg := fmt.Sprintf("Error while parsing the tpa secret string: %v", err)
		log.Error(msg)
		return msg
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
		msg := fmt.Sprintf("Error while setting up SSH client to the TPA_PRIMARY system ip: %s, error: %s", dsaIp, err)
		log.Error(msg)
		return msg
	}
	defer client.Close()
	for _, cmd := range dsmainCMD{
		connectErr, cmdErr, _ := runShellCmd(client, dsaIp, cmd)
		if connectErr != nil {
			log.Error("Error:%v", connectErr)
			return false
		} else if cmdErr != nil {
			msg := fmt.Sprintf("Unexpected error running the cmd '%s' to the TPA_PRIMARY system ip: %s, error: %s", cmd, dsaIp, cmdErr)
			log.Error(msg)
			return false
		}
		return true
	}
	return true
}
