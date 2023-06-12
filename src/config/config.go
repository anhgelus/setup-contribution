package config

const Datas = ".local/share/setup-contribution/"

type Config struct {
	DisabledFolders []string
	Root            RootConfig
	Github          GithubConfig
}

type RootConfig struct {
	CodeOfConduct  string
	Contributing   string
	SecurityPolicy string
}

type GithubConfig struct {
	CodeOwners          string
	PullRequestTemplate string
	Templates           GithubTemplatesConfig
}

type GithubTemplatesConfig struct {
	BugReport      string
	FeatureRequest string
	OtherIssues    string
	Config         string
}
