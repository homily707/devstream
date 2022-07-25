# repo-scaffolding plugin

This plugin bootstraps a GitHub or GitLab repo with scaffolding code for a web application.

## Requirement

This plugin depends on the following environment variable:

- GITHUB_TOKEN

Set it before using this plugin.

If you don't know how to create this token, check out:

- [Creating a personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)

*Tips:*

- If you run `dtm delete`, the repo on GitHub will be completely removed.

- If the `Update` interface is called, the repo will be completely removed and recreated. 

- For the  `repo-scaffolding` plugin, we only need `repo`, `delete_repo` permission for the token.

## Usage

**Please note that all parameter is case-sensitive.**

```yaml
--8<-- "repo-scaffolding.yaml"
```

Replace the following from the config above:

### destination_repo

This configuration is used for the target repo, it includes the following config.

- `YOUR_GITHUB_USERNAME`
- `YOUR_GITHUB_ORGANIZATION_NAME`
- `YOUR_GITHUB_REPO_NAME`
- `YOUR_GITHUB_REPO_MAIN_BRANCH`

Currently, `owner`, `org`, and `repo` are mandatory, `branch` has the default value "main".

### source_repo

This configuration is used for the source scaffolding repoI(only supports GitHub for now). It includes the following config.

- `YOUR_TEMPLATE_REPO_ORG`
- `YOUR_TEMPLATE_REPO_NAME`

All the parameters in the example above are mandatory for now.

### repo_type

This configuration is used for destination_repo location, `gitlab` and `github` can be used for now. If you set it to "github", the generated repo will be pushed to GitHub. If you set it to "gitlab", the generated repo will be pushed to GitLab by your config.

### vars

This configuration is used for template render, It has default variables listed below:

```json
{
    "AppName": destination_repo.repo,
    "Repo": {
        "Name": destination_repo.repo,
        "Owner": destination_repo.owner
    }
}
```

## Examples 

### official scaffolding repo config

These repos are official scaffolding repo to use for `source_repo` config, You can use these repo directly or just create one for yourself.

| language | org | repo |
|  ----  | ----  |----  |
| Golang | devstream-io | dtm-scaffolding-golang |
| Java Spring | spring-guides | gs-spring-boot |


### Golang

```yaml
tools:
  - name: repo-scaffolding
    instanceID: golang-scaffolding
    options:
      destination_repo:
        owner: test_owner
        org: ""
        repo: dtm-test-golang
        branch: main
      repo_type: github
      source_repo:
        org: devstream-io
        repo: dtm-scaffolding-golang
      vars:
        ImageRepo: dtm-test/golang-repo
```

This config will create `dtm-test-golang` repo for user test_owner in GitHub, and the variable ImageRepo will be used for template render. 

### Java Spring

```yaml
tools:
  - name: repo-scaffolding
    instanceID: java-scaffolding
    options:
      destination_repo:
        owner: test_owner
        org: ""
        repo: dtm-test-java
        branch: main
      repo_type: github
      source_repo:
        org: spring-guides
        repo: gs-spring-boot
```

this config will create `dtm-test-java` repo for user test_owner in GitHub.

## Outputs

This plugin has three outputs:

- `owner`
- `repo`
- `repoURL`