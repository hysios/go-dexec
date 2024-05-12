package dexec_test

import (
	"fmt"
	"log"

	"github.com/fsouza/go-dockerclient"
	"github.com/hysios/go-dexec"
)

func ExampleCmd_Output() {
	cl, _ := docker.NewClient("unix:///var/run/docker.sock")
	d := dexec.Docker{cl}

	m, _ := dexec.ByCreatingContainer(docker.CreateContainerOptions{
		Config: &docker.Config{Image: "busybox"}})

	cmd := d.Command(m, "echo", `I am running inside a container!`)
	b, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)

	// Output: I am running inside a container!
}
