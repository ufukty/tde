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
# Database related tasks
# ------------------------------------------------------------- #

function remove-outdated-postgres() {
    apt-get remove --purge postgresql postgresql-*
    rm -rfv /var/lib/postgresql
    rm -rfv /etc/postgresql
    deluser --remove-home postgres
    delgroup postgres
    find / -iname '*postgres*'
}

function install-postgresql() {
    # Source: https://www.postgresql.org/download/linux/ubuntu/
    sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
    wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
    retry apt-get update
    retry apt-get install -y postgresql-15
}

function configure-postgresql() {
    mkdir -p "/mnt/server/archivedir"
    chown -R postgres:postgres "/etc/postgresql/15/main"
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
    rm -rfv "/home/$POSTGRES_USER/.psql_history"
}

# ---------------------------------------------------------------------------- #
# Application related tasks
# ---------------------------------------------------------------------------- #

function deploy-tls-certificates() {
    # "moving files"
    mv "$PROVISIONER_FILES/tde-non-specific.crt" "/etc/ssl/certs/tde-non-specific.crt"
    mv "$PROVISIONER_FILES/tde-non-specific.key" "/etc/ssl/private/tde-non-specific.key"
    # "adjust permissions"
    chmod 755 "/etc/ssl/certs/tde-non-specific.crt"
    chmod 755 "/etc/ssl/private/tde-non-specific.key"
    # "adjust ownership"
    chown root:root "/etc/ssl/certs/tde-non-specific.crt"
    chown root:root "/etc/ssl/private/tde-non-specific.key"
}

function configure-logging() {
    chown -R syslog:root "/var/log/tde.d"
    chown syslog:root "/var/log/syslog"
    chmod 640 "/var/log/syslog"
    systemctl restart rsyslog
}

# ------------------------------------------------------------- #
# Main
# ------------------------------------------------------------- #

assert_sudo
restart_journald
remove_password_change_requirement
wait_cloud_init
apt_update

remove-outdated-postgres
install-postgresql
deploy_provisioner_files

configure-postgresql
postgresql-create-user-and-database
restart-postgresql
change-postgresql-password

deploy-tls-certificates
configure-logging
