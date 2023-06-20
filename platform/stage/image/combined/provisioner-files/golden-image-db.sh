#!/bin/bash

# TODO: Harden Postgres https://www.cybertec-postgresql.com/en/postgresql-security-things-to-avoid-in-real-life/

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
    chown postgres:postgres "/etc/postgresql/15/main/conf.d/tde.conf"
}

function postgresql-create-user-and-database() {
    # Source: https://www.digitalocean.com/community/tutorials/how-to-install-postgresql-on-ubuntu-20-04-quickstart
    adduser --disabled-password --gecos "" "$POSTGRES_USER"
    sudo -u "$POSTGRES_USER" createuser --createdb --no-superuser --no-createrole --no-replication "$POSTGRES_USER"
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

# ------------------------------------------------------------- #
# Main
# ------------------------------------------------------------- #

assert_sudo
restart_journald
remove_password_change_requirement
wait_cloud_init
apt_update

deploy_provisioner_files

install-postgresql
configure-postgresql
postgresql-create-user-and-database
restart-postgresql
change-postgresql-password
