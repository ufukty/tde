IPV4_ADDRESS_PRIVATE="${IPV4_ADDRESS_PRIVATE:?"Required"}"

set -e # stop script execution when a command fails

ssh-keygen -R "$IPV4_ADDRESS_PRIVATE"
ssh-keyscan "$IPV4_ADDRESS_PRIVATE" >>~/.ssh/known_hosts

