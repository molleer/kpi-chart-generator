package utils

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Contribution struct {
	Email      string
	Insertions int64
	Deletions  int64
}

func GetLogs() ([]Contribution, error) {
	cmd := exec.Command("git", "log", "--compact-summary")
	logs, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return parseLogs(string(logs)), nil
}

func parseLogs(logs string) []Contribution {
	regCommit, _ := regexp.Compile("commit [0-9 + a-z]*\n")
	regAuthor, _ := regexp.Compile("Author: .* <.*>")
	regEmail, _ := regexp.Compile("<.*>")
	regChange, _ := regexp.Compile("[0-9]* file.*?\n")
	regInserts, _ := regexp.Compile("[0-9]*? insertion(s|)")
	regDeletes, _ := regexp.Compile("[0-9]*? deletion(s|)")
	regNumbers, _ := regexp.Compile("[0-9]*")
	contribs := []Contribution{}

	commits := regCommit.Split(logs, -1)

	for i := 0; i < len(commits); i++ {
		contrib := Contribution{}

		authorLine := regAuthor.Find([]byte(commits[i]))
		contrib.Email = strings.TrimSuffix(strings.TrimPrefix(string(regEmail.Find(authorLine)), "<"), ">")

		change := regChange.Find([]byte(commits[i]))
		if change == nil {
			continue
		}
		inserts := regInserts.Find(change)
		deletes := regDeletes.Find(change)

		if inserts != nil {
			contrib.Insertions, _ = strconv.ParseInt(string(regNumbers.Find(inserts)), 10, 32)
		}

		if deletes != nil {
			contrib.Deletions, _ = strconv.ParseInt(string(regNumbers.Find(deletes)), 10, 32)
		}

		contribs = append(contribs, contrib)
	}
	return contribs
}
