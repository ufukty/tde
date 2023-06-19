# TDE

## Start

-   Sourcing the `.source-me.sh` is required for each new shell. Which will set required environment vars, define functions, set aliases for commands and command completion for `commands` files. For the convenience, add below script to `~/.bash_profile`

    ```sh
    _source_me() {
        local START_DIR="$(pwd -P)"
        local ANCHOR="$START_DIR"
        while true; do
            if test -f "$ANCHOR/.source-me.sh"; then
                cd "$ANCHOR" && source "$ANCHOR/.source-me.sh"
            fi

            if test "$ANCHOR" == "/"; then
                break
            else
                ANCHOR="$(dirname "$ANCHOR")"
            fi
        done
        cd "$START_DIR"
    }
    _source_me
    _cd() {
        builtin cd "$@"
        _source_me "$@"
    }
    alias cd="_cd"
    ```

-   Use `.source-me-untracked.sh` for setting sensitive env-vars and anything device-specific.

## Application

## Platform

-   4 environments are set under `platform` folder: `debug`, `local`, `stage`, `prod`. Use `./commands <cmd>` where applicable environments.

## Dependencies

-   Application
    -   `Go` 1.20.5
    -   `stringer` (Go)
    -   `serdeser` (install from application/tools/serdeser)
    -   `REST Client` [vscode marketplace](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)
-   Platform
    -   Debug
        -   `Visual Studio Code`
    -   Local
        -   `postgresql` 15
    -   Stage
        -   `terraform` 1.5.0
        -   `ansible` 2.14.4
        -   `packer` 1.9.1
        -   `easyrsa` 3.1.15 (available thru `PATH`)
        -   `LibreSSL` 2.8.3
        -   `jq`
        -   `qr` (python)

## License

-   All rights reserved (c) Ufuktan Yıldırım. Dependencies may belong to others. Open source software used under various licenses including MIT, BSD etc.
