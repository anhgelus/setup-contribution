package src

import (
	"fmt"
	"github.com/anhgelus/setup-contribution/src/config"
	"os"
	"path/filepath"
)

func Copy(conf config.Config, home string, folder string) error {
	basePath := filepath.Join(home, config.Datas, folder)

	err := os.MkdirAll(".github/ISSUE_TEMPLATE", 0764)
	if err != nil {
		return err
	}

	// Root
	err = mdCopy(basePath, "CODE_OF_CONDUCT", conf.Root.CodeOfConduct)
	if err != nil {
		return err
	}

	err = mdCopy(basePath, "CONTRIBUTING", conf.Root.Contributing)
	if err != nil {
		return err
	}

	err = mdCopy(basePath, "SECURITY", conf.Root.SecurityPolicy)
	if err != nil {
		return err
	}

	// .github
	err = sCopy(basePath, ".github/CODEOWNERS", conf.Github.CodeOwners, "")
	if err != nil {
		return err
	}

	err = mdCopy(basePath, ".github/pull_request_template", conf.Github.PullRequestTemplate)
	if err != nil {
		return err
	}

	// .github/ISSUE_TEMPLATE
	err = mdCopy(basePath, ".github/ISSUE_TEMPLATE/bug_report", conf.Github.Templates.BugReport)
	if err != nil {
		return err
	}

	err = mdCopy(basePath, ".github/ISSUE_TEMPLATE/feature_request", conf.Github.Templates.FeatureRequest)
	if err != nil {
		return err
	}

	err = mdCopy(basePath, ".github/ISSUE_TEMPLATE/other_issues", conf.Github.Templates.OtherIssues)
	if err != nil {
		return err
	}

	err = ymlCopy(basePath, ".github/ISSUE_TEMPLATE/config", conf.Github.Templates.Config)
	if err != nil {
		return err
	}

	return nil
}

func mdCopy(basePath string, name string, value string) error {
	return sCopy(basePath, name, value, "md")
}

func ymlCopy(basePath string, name string, value string) error {
	return sCopy(basePath, name, value, "yml")
}

func sCopy(basePath string, name string, value string, ext string) error {
	fmt.Printf("Copying %s\n", name)
	var c []byte
	var err error
	no := ext == ""
	if no {
		c, err = os.ReadFile(fmt.Sprintf("%s/%s", basePath, value))
	} else {
		c, err = os.ReadFile(fmt.Sprintf("%s/%s.%s", basePath, value, ext))
	}
	if err != nil {
		return err
	}
	if no {
		err = os.WriteFile(name, c, 0666)
		if err != nil {
			return err
		}
		return nil
	}
	err = os.WriteFile(fmt.Sprintf("%s.%s", name, ext), c, 0666)
	if err != nil {
		return err
	}
	return nil
}
