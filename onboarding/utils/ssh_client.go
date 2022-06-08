package utils

import (
	"bytes"
	"errors"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func runShellCmd(client *ssh.Client, dsaIp string, cmd string) (connectErr error, cmdErr error, retStdout *string) {
	var buff bytes.Buffer
	var msg string

	// Open SSH session
	session, err := client.NewSession()
	if err != nil {
		msg := fmt.Sprintf("Error while connecting to the DSA system ip: %s, error: %s", dsaIp, err)
		fmt.Println(msg)
		return errors.New(msg), nil, nil
	}
	defer session.Close()

	// Set buffer for stdout return
	session.Stdout = &buff

	// Do not return failure on error. Let caller decide if the
	// error is alright or not
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
		msg := fmt.Sprintf("Error while parsing the DSA Secret String: %v", err)
		fmt.Println(msg)
		return msg
	}
	// ssh.ParseRawPrivateKey(byte(secretKey))

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
		msg := fmt.Sprintf("Error while setting up SSH client to the DSA system ip: %s, error: %s", dsaIp, err)
		fmt.Println(msg)
		return msg
	}
	defer client.Close()
	for _, cmd := range dsmainCMD{
		connectErr, cmdErr, _ := runShellCmd(client, dsaIp, cmd)
		if connectErr != nil {
			fmt.Printf("Error:%v", connectErr)
			return false
		} else if cmdErr != nil {
			msg := fmt.Sprintf("Unexpected error running the cmd '%s' to the DSA system ip: %s, error: %s", cmd, dsaIp, cmdErr)
			fmt.Println(msg)
			return false
		}
		return true
	}
	// connectErr, cmdErr, _ := runShellCmd(client, dsaIp, cmd)
	// if connectErr != nil {
	// 	fmt.Printf("Error:%v", connectErr)
	// 	return "Failed to run command"
	// } else if cmdErr != nil {
	// 	msg := fmt.Sprintf("Unexpected error running the cmd '%s' to the DSA system ip: %s, error: %s", cmd, dsaIp, cmdErr)
	// 	fmt.Println(msg)
	// 	return false
	// }
	return true
}
