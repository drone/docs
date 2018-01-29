+++
date = "2017-04-15T14:39:04+02:00"
title = "Installation Reference"
url = "installation-reference"

[menu.install]
  weight = 2
  identifier = "installation-reference"
  parent = "install_overview"
+++

{{% alert warn %}}
Some of these options may be experimental.

Options are auto-generated from the source code: https://github.com/drone/drone/tree/master/cmd
{{% /alert %}}

# Drone Server

This is the reference list of all environment variables available to configure Drone Server.

CLI FLAG <br><br> ENV VAR | DESCRIPTION <br><br> DEFAULT
--- | ---
`--debug--debug` <br><br> `$DRONE_DEBUG` | start the server in debug mode <br><br> ` `
`--server-host value` <br><br> `$DRONE_SERVER_HOST, $DRONE_HOST` | server host <br><br> ` `
`--server-addr value` <br><br> `$DRONE_SERVER_ADDR` | server address <br><br> `":8000" `
`--server-cert value` <br><br> `$DRONE_SERVER_CERT` | server ssl cert <br><br> ` `
`--server-key value` <br><br> `$DRONE_SERVER_KEY` | server ssl key <br><br> ` `
`--lets-encrypt--lets-encrypt` <br><br> `$DRONE_LETS_ENCRYPT` | lets encrypt enabled <br><br> ` `
`--quic--quic` <br><br> `$DRONE_QUIC` | start the server with quic enabled <br><br> ` `
`--admin value` <br><br> `$DRONE_ADMIN` | list of admin users <br><br> ` `
`--orgs value` <br><br> `$DRONE_ORGS` | list of approved organizations <br><br> ` `
`--open--open` <br><br> `$DRONE_OPEN` | open user registration <br><br> ` `
`--repo-config value` <br><br> `$DRONE_REPO_CONFIG` | file path for the drone config <br><br> `".drone.yml" `
`--session-expires value` <br><br> `$DRONE_SESSION_EXPIRES` | Set the session expiration time default 72h <br><br> `72h0m0s `
`--escalate value` <br><br> `$DRONE_ESCALATE` |  <br><br> `"plugins/docker", "plugins/gcr", "plugins/ecr" `
`--volume value` <br><br> `$DRONE_VOLUME` |  <br><br> ` `
`--network value` <br><br> `$DRONE_NETWORK` |  <br><br> ` `
`--agent-secret value` <br><br> `$DRONE_AGENT_SECRET, $DRONE_SECRET` | agent secret passcode <br><br> ` `
`--secret-service value` <br><br> `$DRONE_SECRET_ENDPOINT` | secret plugin endpoint <br><br> ` `
`--registry-service value` <br><br> `$DRONE_REGISTRY_ENDPOINT` | registry plugin endpoint <br><br> ` `
`--gating-service value` <br><br> `$DRONE_GATEKEEPER_ENDPOINT` | gated build endpoint <br><br> ` `
`--driver value` <br><br> `$DRONE_DATABASE_DRIVER, $DATABASE_DRIVER` | database driver <br><br> `"sqlite3" `
`--datasource value` <br><br> `$DRONE_DATABASE_DATASOURCE, $DATABASE_CONFIG` | database driver configuration string <br><br> `"drone.sqlite" `
`--limit-mem-swap value` <br><br> `$DRONE_LIMIT_MEM_SWAP` |  <br><br> `0 `
`--limit-mem value` <br><br> `$DRONE_LIMIT_MEM` |  <br><br> `0 `
`--limit-shm-size value` <br><br> `$DRONE_LIMIT_SHM_SIZE` |  <br><br> `0 `
`--limit-cpu-quota value` <br><br> `$DRONE_LIMIT_CPU_QUOTA` |  <br><br> `0 `
`--limit-cpu-shares value` <br><br> `$DRONE_LIMIT_CPU_SHARES` |  <br><br> `0 `
`--limit-cpu-set value` <br><br> `$DRONE_LIMIT_CPU_SET` |  <br><br> ` `
`--github--github` <br><br> `$DRONE_GITHUB` | github driver is enabled <br><br> ` `
`--github-server value` <br><br> `$DRONE_GITHUB_URL` | github server address <br><br> `"https://github.com" `
`--github-context value` <br><br> `$DRONE_GITHUB_CONTEXT` | github status context <br><br> `"continuous-integration/drone" `
`--github-client value` <br><br> `$DRONE_GITHUB_CLIENT` | github oauth2 client id <br><br> ` `
`--github-secret value` <br><br> `$DRONE_GITHUB_SECRET` | github oauth2 client secret <br><br> ` `
`--github-scope value` <br><br> `$DRONE_GITHUB_SCOPE` | github oauth scope <br><br> `"repo", "repo:status", "user:email", "read:org" `
`--github-git-username value` <br><br> `$DRONE_GITHUB_GIT_USERNAME` | github machine user username <br><br> ` `
`--github-git-password value` <br><br> `$DRONE_GITHUB_GIT_PASSWORD` | github machine user password <br><br> ` `
`--github-merge-ref--github-merge-ref` <br><br> `$DRONE_GITHUB_MERGE_REF` | github pull requests use merge ref <br><br> ` `
`--github-private-mode--github-private-mode` <br><br> `$DRONE_GITHUB_PRIVATE_MODE` | github is running in private mode <br><br> ` `
`--github-skip-verify--github-skip-verify` <br><br> `$DRONE_GITHUB_SKIP_VERIFY` | github skip ssl verification <br><br> ` `
`--gogs--gogs` <br><br> `$DRONE_GOGS` | gogs driver is enabled <br><br> ` `
`--gogs-server value` <br><br> `$DRONE_GOGS_URL` | gogs server address <br><br> `"https://github.com" `
`--gogs-git-username value` <br><br> `$DRONE_GOGS_GIT_USERNAME` | gogs service account username <br><br> ` `
`--gogs-git-password value` <br><br> `$DRONE_GOGS_GIT_PASSWORD` | gogs service account password <br><br> ` `
`--gogs-private-mode--gogs-private-mode` <br><br> `$DRONE_GOGS_PRIVATE_MODE` | gogs private mode enabled <br><br> ` `
`--gogs-skip-verify--gogs-skip-verify` <br><br> `$DRONE_GOGS_SKIP_VERIFY` | gogs skip ssl verification <br><br> ` `
`--gitea--gitea` <br><br> `$DRONE_GITEA` | gitea driver is enabled <br><br> ` `
`--gitea-server value` <br><br> `$DRONE_GITEA_URL` | gitea server address <br><br> `"https://try.gitea.io" `
`--gitea-git-username value` <br><br> `$DRONE_GITEA_GIT_USERNAME` | gitea service account username <br><br> ` `
`--gitea-git-password value` <br><br> `$DRONE_GITEA_GIT_PASSWORD` | gitea service account password <br><br> ` `
`--gitea-private-mode--gitea-private-mode` <br><br> `$DRONE_GITEA_PRIVATE_MODE` | gitea private mode enabled <br><br> ` `
`--gitea-skip-verify--gitea-skip-verify` <br><br> `$DRONE_GITEA_SKIP_VERIFY` | gitea skip ssl verification <br><br> ` `
`--bitbucket--bitbucket` <br><br> `$DRONE_BITBUCKET` | bitbucket driver is enabled <br><br> ` `
`--bitbucket-client value` <br><br> `$DRONE_BITBUCKET_CLIENT` | bitbucket oauth2 client id <br><br> ` `
`--bitbucket-secret value` <br><br> `$DRONE_BITBUCKET_SECRET` | bitbucket oauth2 client secret <br><br> ` `
`--gitlab--gitlab` <br><br> `$DRONE_GITLAB` | gitlab driver is enabled <br><br> ` `
`--gitlab-server value` <br><br> `$DRONE_GITLAB_URL` | gitlab server address <br><br> `"https://gitlab.com" `
`--gitlab-client value` <br><br> `$DRONE_GITLAB_CLIENT` | gitlab oauth2 client id <br><br> ` `
`--gitlab-secret value` <br><br> `$DRONE_GITLAB_SECRET` | gitlab oauth2 client secret <br><br> ` `
`--gitlab-git-username value` <br><br> `$DRONE_GITLAB_GIT_USERNAME` | gitlab service account username <br><br> ` `
`--gitlab-git-password value` <br><br> `$DRONE_GITLAB_GIT_PASSWORD` | gitlab service account password <br><br> ` `
`--gitlab-skip-verify--gitlab-skip-verify` <br><br> `$DRONE_GITLAB_SKIP_VERIFY` | gitlab skip ssl verification <br><br> ` `
`--gitlab-private-mode--gitlab-private-mode` <br><br> `$DRONE_GITLAB_PRIVATE_MODE` | gitlab is running in private mode <br><br> ` `
`--gitlab-v3-api--gitlab-v3-api` <br><br> `$DRONE_GITLAB_V3_API` | gitlab is running the v3 api <br><br> ` `
`--stash--stash` <br><br> `$DRONE_STASH` | stash driver is enabled <br><br> ` `
`--stash-server value` <br><br> `$DRONE_STASH_URL` | stash server address <br><br> ` `
`--stash-consumer-key value` <br><br> `$DRONE_STASH_CONSUMER_KEY` | stash oauth1 consumer key <br><br> ` `
`--stash-consumer-rsa value` <br><br> `$DRONE_STASH_CONSUMER_RSA` | stash oauth1 private key file <br><br> ` `
`--stash-consumer-rsa-string value` <br><br> `$DRONE_STASH_CONSUMER_RSA_STRING` | stash oauth1 private key string <br><br> ` `
`--stash-git-username value` <br><br> `$DRONE_STASH_GIT_USERNAME` | stash service account username <br><br> ` `
`--stash-git-password value` <br><br> `$DRONE_STASH_GIT_PASSWORD` | stash service account password <br><br> ` `
`--stash-skip-verify--stash-skip-verify` <br><br> `$DRONE_STASH_SKIP_VERIFY` | stash skip ssl verification <br><br> ` `
`--coding--coding` <br><br> `$DRONE_CODING` | coding driver is enabled <br><br> ` `
`--coding-server value` <br><br> `$DRONE_CODING_URL` | coding server address <br><br> `"https://coding.net" `
`--coding-client value` <br><br> `$DRONE_CODING_CLIENT` | coding oauth2 client id <br><br> ` `
`--coding-secret value` <br><br> `$DRONE_CODING_SECRET` | coding oauth2 client secret <br><br> ` `
`--coding-scope value` <br><br> `$DRONE_CODING_SCOPE` | coding oauth scope <br><br> `"user", "project", "project:depot" `
`--coding-git-machine value` <br><br> `$DRONE_CODING_GIT_MACHINE` | coding machine name <br><br> `"git.coding.net" `
`--coding-git-username value` <br><br> `$DRONE_CODING_GIT_USERNAME` | coding machine user username <br><br> ` `
`--coding-git-password value` <br><br> `$DRONE_CODING_GIT_PASSWORD` | coding machine user password <br><br> ` `
`--coding-skip-verify--coding-skip-verify` <br><br> `$DRONE_CODING_SKIP_VERIFY` | coding skip ssl verification <br><br> ` `
`--license-key value` <br><br> `$DRONE_LICENSE` | drone license key <br><br> ` `
`--vault-addr value` <br><br> `$VAULT_ADDR` | vault address <br><br> ` `
`--vault-cacert value` <br><br> `$VAULT_CACERT` | vault ca cert <br><br> ` `
`--vault-capath value` <br><br> `$VAULT_CAPATH` | vault ca path <br><br> ` `
`--vault-client-cert value` <br><br> `$VAULT_CLIENT_CERT` | vault client cert <br><br> ` `
`--vault-skip-verify--vault-skip-verify` <br><br> `$VAULT_SKIP_VERIFY` | vault skip verify <br><br> ` `
`--vault-max-retries value` <br><br> `$VAULT_MAX_RETRIES` | vault max retries <br><br> `0 `
`--vault-token value` <br><br> `$VAULT_TOKEN` | vault token <br><br> ` `
`--vault-tls-server-name value` <br><br> `$VAULT_TLS_SERVER_NAME` | vault tls server name <br><br> ` `
`--repo-whitelist value` <br><br> `$DRONE_REPO_WHITELIST` | repoisitory whitelist <br><br> ` `
`--hook value` <br><br> `$DRONE_HOOK` | server hooks <br><br> ` `

# Drone Agent

This is the reference list of all environment variables available to configure Drone Agent.

CLI FLAG <br><br> ENV VAR | DESCRIPTION <br><br> DEFAULT
--- | ---
`--server value` <br><br> `$DRONE_SERVER` | drone server address <br><br> `"localhost:9000" `
`--username value` <br><br> `$DRONE_USERNAME` | drone auth username <br><br> `"x-oauth-basic" `
`--password value` <br><br> `$DRONE_PASSWORD, $DRONE_SECRET` | drone auth password <br><br> ` `
`--debug--debug` <br><br> `$DRONE_DEBUG` | start the agent in debug mode <br><br> ` `
`--pretty--pretty` <br><br> `$DRONE_DEBUG_PRETTY` | enable pretty-printed debug output <br><br> ` `
`--nocolor--nocolor` <br><br> `$DRONE_DEBUG_NOCOLOR` | disable colored debug output <br><br> ` `
`--hostname value` <br><br> `$DRONE_HOSTNAME, $HOSTNAME` |  <br><br> ` `
`--platform value` <br><br> `$DRONE_PLATFORM` |  <br><br> `"linux/amd64" `
`--filter value` <br><br> `$DRONE_FILTER` | filter expression used to restrict builds by label <br><br> ` `
`--max-procs value` <br><br> `$DRONE_MAX_PROCS` |  <br><br> `1 `
`--healthcheck--healthcheck` <br><br> `$DRONE_HEALTHCHECK` | enables the healthcheck endpoint <br><br> ` `
`--keepalive-time value` <br><br> `$DRONE_KEEPALIVE_TIME` | after a duration of this time if the agent doesn't see any activity it pings the server to see if the transport is still alive <br><br> `0s `
`--keepalive-timeout value` <br><br> `$DRONE_KEEPALIVE_TIMEOUT` | after having pinged for keepalive check, the client waits for a duration of Timeout and if no activity is seen even after that the connection is closed. <br><br> `20s `
