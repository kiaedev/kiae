package project

import "strings"

func (proj *Project) GitHTTPSUrl() string {
	gitSSH := strings.Replace(proj.GitRepo, ":", "/", -1)
	return strings.Replace(gitSSH, "git@", "https://", -1)
}
