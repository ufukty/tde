2023.06.22

# Images

## Image Hierarchy

```
base                        user, utilities, fail2ban, basic security
├── vpn                     openvpn, easy-rsa, argon2
└── internal                firewall, accessible with internal network
    ├── gateway             allows :8080 on firewall
    ├── application-to-db   systemd service, logging, certs, tunnel with database
    ├── application         systemd service, logging, certs
    ├── combined            application + postgres
    │   └── compiler        go
    └── database            postgres, tunnel with application
```
