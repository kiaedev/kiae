package project

import (
	"fmt"
	"strings"
)

func (proj *Project) Alias() string {
	return proj.Id[18:]
}

func (proj *Project) NameID() string {
	return fmt.Sprintf("%s-%s", proj.Name, proj.Alias())
}

func (proj *Project) GitHTTPSUrl() string {
	gitSSH := strings.Replace(proj.GitRepo, ":", "/", -1)
	return strings.Replace(gitSSH, "git@", "https://", -1)
}
