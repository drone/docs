#!/bin/sh

# secrets
drone secret add --help    > secrets/data/drone_secret_add.out.txt
drone secret update --help > secrets/data/drone_secret_update.out.txt
drone secret info --help   > secrets/data/drone_secret_info.out.txt
drone secret ls --help     > secrets/data/drone_secret_ls.out.txt
drone secret rm --help     > secrets/data/drone_secret_rm.out.txt

# registries
drone registry add --help    > registry/data/drone_registry_add.out.txt
drone registry update --help > registry/data/drone_registry_update.out.txt
drone registry info --help   > registry/data/drone_registry_info.out.txt
drone registry ls --help     > registry/data/drone_registry_ls.out.txt
drone registry rm --help     > registry/data/drone_registry_rm.out.txt

# user
drone user add --help    > user/data/drone_user_add.out.txt
drone user info --help   > user/data/drone_user_info.out.txt
drone user ls --help     > user/data/drone_user_ls.out.txt
drone user rm --help     > user/data/drone_user_rm.out.txt

# repository
drone repo add --help      > repo/data/drone_repo_add.out.txt
drone repo update --help   > repo/data/drone_repo_update.out.txt
drone repo repair --help   > repo/data/drone_repo_repair.out.txt
drone repo chown --help    > repo/data/drone_repo_chown.out.txt
drone repo info --help     > repo/data/drone_repo_info.out.txt
drone repo ls --help       > repo/data/drone_repo_ls.out.txt
drone repo rm --help       > repo/data/drone_repo_rm.out.txt

# builds
drone build list --help   > build/data/drone_build_list.out.txt
drone build last --help    > build/data/drone_build_last.out.txt
drone build logs --help    > build/data/drone_build_logs.out.txt
drone build info --help    > build/data/drone_build_info.out.txt
drone build start --help   > build/data/drone_build_start.out.txt
drone build stop --help    > build/data/drone_build_stop.out.txt
drone build approve --help > build/data/drone_build_approve.out.txt
drone build decline --help > build/data/drone_build_decline.out.txt

# misc
drone deploy --help > misc/data/drone_deploy.out.txt
drone exec --help   > misc/data/drone_exec.out.txt
