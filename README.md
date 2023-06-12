# setup-contribution

Linux CLI Application generating contributions guidelines.

I do not plan to make this application compatible with other systems.
If you want to do it, develop it on your own, open a PR and I'll gladly merge it.

## Usage

This CLI Application is very simple.

1. Use `setup-contribution` to set up your default configuration
2. Use `setup-contribution string` to set up your {string} configuration
3. Use `setup-contriubtion install` to install the data

The data is stored in `~/.local/share/setup-contribution/` and the config file is located at 
`~/.local/share/setup-contribution/config.toml`.

To create a new configuration, just copy the `~/.local/share/setup-contribution/default/` to 
`~/.local/share/setup-contribution/{string}/` where string is the name of your new configuration.
:warning: not working with space!

The default config file is:
```toml
DisabledFolders = [""]

[Root]
CodeOfConduct = "CODE_OF_CONDUCT"
Contributing = "CONTRIBUTING"
SecurityPolicy = "SECURITY"

[Github]
CodeOwners = "CODEOWNERS"
PullRequestTemplate = "templates/pull_request_template"

[Github.Templates]
BugReport = "templates/bug_report"
FeatureRequest = "templates/feature_request"
OtherIssues = "templates/other_issues"
Config = "templates/config"
```

If you want to disable a configuration, just add its name in the list `DisabledFolders`.

Every option excepted `DisabledFolders` is the name of your file.
For example, if I want to rename the file `CODE_OF_CONDUCT.md` by `fooBar.md`, I must change `CODE_OF_CONDUCT` to `fooBar`
for the option `CodeOfConduct` inside `Root`.
:warning: You cannot edit the extension of the files!

## Technologies

- Go 1.20
- pelletier/go-toml v2
