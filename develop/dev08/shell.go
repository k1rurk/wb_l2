package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mitchellh/go-ps"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

type execError struct {
}

func (e *execError) Error() string {
	return "incorrect usage command: exec command [args]"
}

type shell struct {
	pipeMode   bool
	pipeBuffer string
}

func mainRun() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	sh := &shell{
		pipeMode:   false,
		pipeBuffer: "",
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		currDir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("[%s]:%s$ ", usr.Username, currDir)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		text = strings.TrimSuffix(text, "\n")

		if text == "exit" {
			break
		}
		err = sh.lookCommand(text)
		if err != nil {
			log.Fatalln(err)
		}

	}
}

func (s *shell) lookCommand(userInput string) error {
	commands := strings.Split(userInput, "|")
	var output string
	var err error
	if len(commands) > 1 {
		s.pipeMode = true
		for _, command := range commands {
			if s.pipeBuffer, err = s.execCommand(command + s.pipeBuffer); err != nil {
				return err
			}
		}
		s.pipeMode = false
		output = s.pipeBuffer
	} else {
		if output, err = s.execCommand(commands[0] + s.pipeBuffer); err != nil {
			return err
		}
	}

	fmt.Println(output)
	return nil
}

func (s *shell) execCommand(command string) (string, error) {
	cmd := strings.Fields(command)

	switch cmd[0] {
	case "ps":
		output, err := runPs()
		if err != nil {
			return "", nil
		}
		return output, nil
	case "echo":
		return strings.Join(cmd[1:], " "), nil
	case "kill":
		return "", runKill(cmd[1:])
	case "pwd":
		if dir, err := os.Getwd(); err != nil {
			return "", err
		} else {
			return dir, nil
		}
	case "cd":
		if len(cmd) != 2 {
			return "", errors.New("incorrect syntax cd")
		}
		err := os.Chdir(cmd[1])
		return "", err
	case "exec":
		return runExecute(cmd[1:])
	default:
		return "", errors.New("unknown command " + cmd[0])
	}
}

func runExecute(args []string) (string, error) {
	var output string
	if len(args) > 0 {
		cmd, err := exec.Command("powershell", args...).Output()
		if err != nil {
			return "", err
		}
		output = string(cmd)
	} else {
		return "", &execError{}
	}
	return output, nil
}

func runKill(args []string) error {
	for _, strPid := range args {
		kill := exec.Command("taskkill", "/T", "/F", "/PID", strPid)
		err := kill.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func runPs() (string, error) {
	prs, err := ps.Processes()
	if err != nil {
		return "", err
	}
	var res string
	// map ages
	for x := range prs {
		res += fmt.Sprintf("%d\t%s\n", prs[x].Pid(), prs[x].Executable())
		// do os.* stuff on the pid
	}

	return res, nil
}
