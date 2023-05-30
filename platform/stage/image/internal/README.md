# Key features

-   iptables is enabled
    -   configured to apply configuration in `/etc/iptables/picarus-custom-firewall.v4` after each restart. that file can be replaced by derived images.
    -   Firewall is configurated for only internal network data transfers, all public communication requires to adjust firewall
    -   **SSH can only be reachable from internal network**
-   fail2ban configurated to allow requests coming from internal network regarless of the amount.
