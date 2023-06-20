#!/bin/bash

# TODO: Harden Postgres https://www.cybertec-postgresql.com/en/postgresql-security-things-to-avoid-in-real-life/

PS4="\033[36m$(realpath --relative-to="$(dirname "$WORKSPACE")" "$(pwd -P)")\033[32m/\$(basename \"\${BASH_SOURCE}\"):\${LINENO}\033[0m\033[33m\${FUNCNAME[0]:+/\${FUNCNAME[0]}():}\033[0m "
set -o xtrace
# ------------------------------------------------------------- #
# Include
# ------------------------------------------------------------- #

PROVISIONER_FILES="/home/$SUDO_USER/provisioner-files"
cd "$PROVISIONER_FILES"
. utilities.sh
. secrets.sh

# ------------------------------------------------------------- #
# Variables
# ------------------------------------------------------------- #

# A Linux and Postgres user will be created with this name, in addition to a Postgres Database
POSTGRES_USER="${POSTGRES_USER:?"POSTGRES_USER is required"}"

POSTGRES_REGULAR_USER_PASSWD_HASHED="${POSTGRES_REGULAR_USER_PASSWD_HASHED:?"POSTGRES_REGULAR_USER_PASSWD_HASHED is required"}"

# ------------------------------------------------------------- #
# Tasks
# ------------------------------------------------------------- #

function install-postgresql() {
    # Source: https://www.postgresql.org/download/linux/ubuntu/
    # info "Create the file repository configuration"
    sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
    # info "Import the repository signing key"
    wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
    retry apt-get update
    retry apt-get install -y postgresql-15
}

function configure-postgresql() {
    # info "creating directory: /mnt/server/archivedir"
    mkdir -p "/mnt/server/archivedir"
    chown postgres:postgres "/etc/postgresql/13/main/conf.d/tde.conf"
}

function postgresql-create-user-and-database() {
    # Source: https://www.digitalocean.com/community/tutorials/how-to-install-postgresql-on-ubuntu-20-04-quickstart
    adduser --disabled-password --gecos "" "$POSTGRES_USER"
    sudo -u postgres createuser --createdb --no-superuser --no-createrole --no-replication "$POSTGRES_USER"
    sudo -u "$POSTGRES_USER" createdb "$POSTGRES_USER"
}

function restart-postgresql() {
    systemctl restart postgresql
}

function change-postgresql-password() {
    # source: https://subscription.packtpub.com/book/big_data_and_business_intelligence/9781789537581/1/ch01lvl1sec18/changing-your-password-securely
    sudo -u "$POSTGRES_USER" psql -c "ALTER USER \"$POSTGRES_USER\" PASSWORD '$POSTGRES_REGULAR_USER_PASSWD_HASHED';"
    # sudo -u $POSTGRES_USER psql -c "ALTER USER \"$POSTGRES_USER\" WITH PASSWORD '$POSTGRES_REGULAR_USER_PASSWD_CLEAR_TEXT';" # clear text version of password
    rm "/home/$POSTGRES_USER/.psql_history"
}

function configure-ssh() {
    # info "configuring sshd to allow logins to user $POSTGRES_USER"
    sed -r "s;^AllowUsers (.*);AllowUsers \1 $POSTGRES_USER;" --in-place /etc/ssh/sshd_config
    # info "creating ~/.ssh directory, adjusting ownership and privileges"
    mkdir -p "/home/$POSTGRES_USER/.ssh"
    chown -R $POSTGRES_USER:$POSTGRES_USER "/home/$POSTGRES_USER"
    chmod 700 "/home/$POSTGRES_USER"
    # info "deploy application server's public key to created user's authorized_keys"
    touch "/home/$POSTGRES_USER/.ssh/authorized_keys"
    chown -R $POSTGRES_USER:$POSTGRES_USER "/home/$POSTGRES_USER"
    chmod 700 "/home/$POSTGRES_USER"
    cat "$PROVISIONER_FILES/ssh-app-db.pub" >>"/home/$POSTGRES_USER/.ssh/authorized_keys"
    # info "restart sshd.service"
    systemctl restart sshd
}

# ------------------------------------------------------------- #
# Main
# ------------------------------------------------------------- #

with-echo assert_sudo
with-echo restart_journald
with-echo remove_password_change_requirement
with-echo wait_cloud_init
with-echo apt_update

with-echo deploy_provisioner_files

with-echo install-postgresql
with-echo configure-postgresql
with-echo postgresql-create-user-and-database
with-echo restart-postgresql
with-echo change-postgresql-password
with-echo configure-ssh
