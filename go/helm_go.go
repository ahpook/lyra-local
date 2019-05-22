package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/lyraproj/servicesdk/lang/go/lyra"
)

type helmIn struct {
	Name      string
	Chart     string
	Overrides []string
	Namespace *string
}

type helmOut struct {
	Output string
}

func helmInstall(in helmIn) helmOut {
	log := hclog.Default()
	namespace := "default"
	if in.Namespace != nil {
		namespace = *in.Namespace
	}
	args := []string{
		"upgrade",
		"--force",
		"--install",
		in.Name,
		in.Chart,
		"--namespace",
		namespace,
	}
	if len(in.Overrides) > 0 {
		args = append(args, "--set")
		x := strings.Join(in.Overrides, ",")
		args = append(args, x)
	}
	cmd := exec.Command("helm", args...)

	log.Debug("about to run command", "cmd", cmd)

	out, err := cmd.CombinedOutput()
	output := fmt.Sprintf("%s", out)
	if err != nil {
		panic(fmt.Errorf("error running helm cmd %v \n error is %v \n output is %v", cmd, err, output))
	}

	return helmOut{Output: output}
}

func main() {
	lyra.Serve(`helm_go`, nil, &lyra.Action{Do: helmInstall})
}
