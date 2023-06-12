package src

import (
	"github.com/anhgelus/setup-contribution/src/config"
	"os"
	"path/filepath"
)

func Install(home string, baseConf string) error {
	base := filepath.Join(home, config.Datas, "default")
	err := os.MkdirAll(base, 0764)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(base, "templates"), 0764)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(base, "config.toml"), []byte(baseConf), 0666)
	if err != nil {
		return err
	}

	files := []string{"CODE_OF_CONDUCT.md", "CONTRIBUTING.md", "SECURITY.md", "CODEOWNERS",
		"templates/pull_request_template", "templates/bug_report", "templates/feature_request", "templates/other_issues",
		"templates/config.yml"}
	for _, f := range files {
		err := create(base, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func create(base string, file string) error {
	f, err := os.Create(filepath.Join(base, file))
	if err != nil {
		return err
	}
	return f.Close()
}
