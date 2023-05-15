#!/bin/bash

# ------------------------------------------------------------- #
# Documentation
# ------------------------------------------------------------- #

# # Add below text into scripts:
# # shell_commons.sh file will be copied into same directory if the script
# # is running by a provisioner (vagrant, packer, etc) in guest/remote server
# [ -f shell_commons.sh ] && . shell_commons.sh || . ../../companions/shell_commons.sh

# ------------------------------------------------------------- #
# Variables
# ------------------------------------------------------------- #

# `EXECUTION_MODE` valid values:
#   - `ALL`           : Development purpose, maybe choose for vagrant's provisioners
#   - `GOLDEN_IMAGE`  : Choose for packer's provisioners
#   - `DEPLOYMENT`    : Choose for terraform's provisioners
EXECUTION_MODE="${EXECUTION_MODE:-"ALL"}"

# ------------------------------------------------------------- #
# Constants
# ------------------------------------------------------------- #

colr_def="\033[39m" # default
colr_bla="\033[30m" # black
colr_red="\033[31m" # red
colr_gre="\033[32m" # green
colr_yel="\033[33m" # yellow
colr_blu="\033[34m" # blue
colr_mag="\033[35m" # magenta
colr_cya="\033[36m" # cyan

colr_dgr="\033[90m" # d_gray
colr_lgr="\033[37m" # l_gray
colr_lre="\033[91m" # l_red
colr_lgr="\033[92m" # l_green
colr_lye="\033[93m" # l_yellow
colr_lbl="\033[94m" # l_blue
colr_lma="\033[95m" # l_magenta
colr_lcy="\033[96m" # l_cyan
colr_whi="\033[97m" # white

colr_bol="\033[1m" #  # bold
colr_dim="\033[2m" # dim
colr_und="\033[4m" # underlined
colr_bli="\033[5m" # blink
colr_inv="\033[7m" # inverted
colr_hid="\033[8m" # hidden

colr_rst_all="\033[0m"                 #  # rst_all
colr_rst_bold_bright="\033[1m\033[21m" # rst_bold_bright
colr_rst_dim="\033[2m\033[22m"         # rst_dim
colr_rst_underline="\033[4m\033[24m"   # rst_underline
colr_rst_blink="\033[5m\033[25m"       # rst_blink
colr_rst_reverse="\033[7m\033[27m"     # rst_reverse
colr_rst_hidden="\033[8m\033[28m"      # rst_hidden

separator_string=""
term_columns="$(tput cols 2>/dev/null)"
left_char="${term_columns:-120}"
while [[ $left_char -gt 0 ]]; do
    separator_string="${separator_string}─"
    left_char=$(($left_char - 1))
done
unset term_columns
unset left_char

# ------------------------------------------------------------- #
# Functions
# ------------------------------------------------------------- #

function error() {
    echo -e "${colr_rst_bold_bright}${colr_lre}${separator_string}${colr_rst_all}" >&2
    echo -e "${colr_rst_bold_bright}${colr_lre}$ECHO_PREFIX [ERROR] $1${colr_rst_all}" >&2
    echo -e "${colr_rst_bold_bright}${colr_lre}${separator_string}${colr_rst_all}" >&2
    exit 1
}

function log_error() {
    error "$@"
}

function info() {
    echo -e "${colr_rst_bold_bright}${colr_lma}$ECHO_PREFIX [INFO] $1${colr_rst_all}"
}

function log_info() {
    info "$@"
}

function log_notice() {
    echo -e "${colr_rst_bold_bright}${colr_lye}$ECHO_PREFIX [NOTICE] $1${colr_rst_all}"
}

function starting() {
    echo -e "${colr_rst_bold_bright}${colr_lbl}$ECHO_PREFIX [TASK:STARTING] $1${colr_rst_all}"
}

function log_start() {
    starting "$@"
}

function checkpoint() {
    echo -e "${colr_rst_bold_bright}${colr_lcy}$ECHO_PREFIX [TASK:DONE] $1${colr_rst_all}"
}

function log_checkpoint() {
    checkpoint "$@"
}

function retry() {
    count=0
    until $@; do
        count=$((count + 1))
        if [[ $count -le 300 ]]; then
            info "Attempted to run $1, but it's failed for $count times, now trying again..." && sleep 2
        else
            error "Seems like $1 is busy right now, please try again later."
        fi
    done
}

# Takes 2 parameters:
#   - `$1` Name of the function that is desited to run
#   - `$2` (optional): Execution preference. It either be set or not. Valid values are: `GOLDEN_IMAGE`, `DEPLOYMENT`
function execute_task() {
    starting "$@" && $@ && checkpoint "$@"
}

function execute_task_in_golden_image() {
    if [[ "$EXECUTION_MODE" == "ALL" || "$EXECUTION_MODE" == "GOLDEN_IMAGE" ]]
    then
        starting "$@" && $@ && checkpoint "$@"
    else
        log_notice "Passing \"$@\" because of EXECUTION_MODE"
    fi
}

function execute_task_in_deployment() {
    if [[ "$EXECUTION_MODE" == "ALL" || "$EXECUTION_MODE" == "DEPLOYMENT" ]]
    then
        starting "$@" && $@ && checkpoint "$@"
    else
        log_notice "Passing \"$@\" because of EXECUTION_MODE"
    fi
}

function assert_sudo() {
    if [[ "$EUID" > 0 ]]; then error "You need to run this script as root user (or with sudo)"; fi
}

function remove_password_change_requirement() {
    info "remove password change requirement to root"
    sed --in-place -E 's/root:(.*):0:0:(.*):/root:\1:18770:0:\2:/g' /etc/shadow
}

function wait_cloud_init() {
    cloud-init status --wait
}

export DEBIAN_FRONTEND=noninteractive
