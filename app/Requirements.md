Requirements needed at rollout:

-   Product:
    -   Pay-as-you-go
    -   Invitation system
-   Development:
    -   Adaptability >> Velocity >>>> Performance
    -   Using with resources like VPS, VPC easy-to-find amongst different cloud providers.
    -   Debugging support. Mimicks the real-word deployment (1 instance/service).
    -   Easy to add new services/endpoints.
    -   Life-cycle-long consistency between services mostly with code reuse.

Outcomes:

-   Service discovery between services with built-in load balancing.
-   Kind of microservices:
    -   with: per-service database, per-service VPS, load-balancers between services
    -   without: multi-repo, independent deployment, multi-language
-   Runners are multi-tenant (for paying users)
-   Internal network closed outside. Internal service can not be reached outside.

Long term requirements to consider:

-   User-uploaded code will only be executed in isolated environment.
-   Horizontal scalability without service interruption
-   Afiliate
