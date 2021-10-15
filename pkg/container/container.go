package container

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func RunInsideContainer(containerName, volume string, commands []string) {
	fmt.Printf("Running inside %s ... \n", containerName)
	fmt.Println("The commands are: ", commands)

	createCgroup(containerName)

	syscall.Chroot("/home/aliakbar/Downloads/ubuntu-base-18.04.5-base-amd64")
	syscall.Chdir("/")

	syscall.Mount("proc", "proc", "proc", 0, "")

	syscall.Sethostname([]byte(containerName))

	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

	syscall.Unmount("proc", 0)

}

func RunContainer(containerName, volume string, commands []string) {
	fmt.Printf("Running container %s ... \n", containerName)
	fmt.Println("The commands are: ", commands)

	flags := []string{"runc", "--volume", volume, "--name", containerName}

	cmd := exec.Command("/proc/self/exe", append(flags, commands...)...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	cmd.Run()
	fmt.Println("after all ..")
}

func createCgroup(name string) {
	cgroupPath := "/sys/fs/cgroup/pids/" + name + "/"

	err := os.Mkdir(cgroupPath, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	must(ioutil.WriteFile(cgroupPath+"pids.max", []byte("20"), 0700))
	must(ioutil.WriteFile(cgroupPath+"notify_on_release", []byte("1"), 0700))
	must(ioutil.WriteFile(cgroupPath+"cgroup.procs", []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
