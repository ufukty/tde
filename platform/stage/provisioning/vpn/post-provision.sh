IPV4_ADDRESS_PRIVATE="${IPV4_ADDRESS_PRIVATE:?"Required"}"

ssh-keygen -R "$IPV4_ADDRESS_PRIVATE"
ssh-keyscan "$IPV4_ADDRESS_PRIVATE" >>~/.ssh/known_hosts
