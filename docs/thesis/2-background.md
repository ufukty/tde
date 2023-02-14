## Background - Concepts and terms

### Test driven development

Test driven development are implemented by developers with writing so-called tests before actually implementing the program. Those tests can be divided into a couple of classes. Most popularly adopted and well known examples are unit tests and integration tests. Unit test is when a component tested by itself in isolation from the components it supposed to interact in active deployment. So, when a unit test failed, smallest component can be addressed immediately by tracking lately-turned-to-fail tests. Other class is when systems that consists by more than one components deployed into real or mock (a.k.a test, development) environment and tested for higher-level requirements than unit tests. To understand relation between test types, thinking about that system definition may help: A system consists by components and its behavior emerges by the interactions of components, also each component can be seen as sub-systems that consists by sub-components smaller than itself. Thus, unit test are for smallest components, and at each organizational-level from that to the higher, different type of test tools and techniques are considered.

Advantages of TDD:

-   Reducing cost of maintenance mistakes due to employee resign / work force circulation.
-   Increased security and robustness overall.
-   Increased robustness against future workers making mistakes due to lack of prior knowledge about and incorrect removals of existent structures in code.

#### Requirements of commercialization

##### Changing requirements of projects

Software projects (service/product) mature with time and the increasing diffusion area on target group. Early adopters of innovations are the people who are expected to be the most educated/experienced ones amongst the whole target group. Through diffusion process, it is expected to feedbacks and requests raised by customers differantiate with the diffusion phase. Thus it is unavoidable that requirements for a particular product or a service will change with time; and software development process should

##### Work force circulation

Companies can not rely their employees on staying in the company until a software project completes its lifecycle. Employees come and go, but products should keep get fixed for newly discovered requirements, and get updated for new features customers request without effecting by if employees who worked on first phases of project still in the company. Latest workforce should be able to work on the existing codebase to add new features without the fear of breaking previously working features that implemented by previous workers. Following by test driven approach, key requirements that the software expected to satisfy are added into the codebase as tests at the moment those requirements raised and their reasoning known by employees of that time. Codebase is supposed to get tested for all tests at any time a possibly breaking-change occurred, or before and after every deployment by the CI/CD pipeline.

### Evolutionary Algorithms and Genetic Programming

Developer doesn't need to know form of solution prior to development, being able to define the problem and test cases the solution should pass and fail is enough to start evolution.

### Koza's work on generating human competitive results with GP in the area of electric circuitry and symbolic regression

Used in symbolic regression like

### Program synyhesis with user-provided TDD

#### Benchmarks

-   String transformations
-   Table transformations
-   XML transformations
-   Pex4Fun

### Liability and responsibility in Software Development

Products creates responsibility and liability on producer and seller. Liability usually comes from law and contracts; and customer expectations raised by market players.

Software is frequently developed as a commercial product that solves a social (or derived) problem and creates responsibility and liability on the seller company. Responsibility comes from the need to maintain company reputation in target consumer demographic and the liability comes from promises that company explicitly gave to customers as compliences to certain standards or the requirement of complying with the law.

### Meaning of "production-ready"

Just like workers of companies' other product departments, software developers should be able to inspect the prior work as the team makes progress in development to spot unfitting progress to everchanging requirements.

The aim is to present the Genetic Programming within a tool to software developers for helping them to build their commercial purposed software. The tool eventually should have features like familiarity, easy-to-learn

Most of the companies develop CRUD programs. That is a software runs on a computer accessible to clients. Exposes an stable interface consists by serial of endpoints each creates/reads/updates/deletes a resource. Resources are abstract objects, meaningful for the domain: like users, posts, photos for a social media platform; or employees, customers, e-mails, calls for a customer relation management software. Developer teams from variety of companies operates on different sector, or industries; it is safe to say that most of the teams specialized on developing CRUD programs and their skill-sets are shaped around the requirements of developing CRUD programs.

CI/CD is trending topic for nearly 20 years. That is automating the cycle of integration and deployment which is expected to be a loop because of everchanging set of requirements and priorities that are faced by organizations produce software, in addition to a work force that consists by employees don't stay forever on same company and have different degrees of experiences on the field. Steps of CI/CD are mainly considered as code, build, test, release, deploy, operate, monitor and plan (for next round).

Familiarity: Most of the devs are experienced in building CRUD applications and their skill set don't cover the Genetic Programming or another EA variation. In order to enable those developers leveraging from advantages of GP, the method should be delivered as a tool that abstracts the internal mechanisms of evolution and only presents controls and output that they need.

For the sake of our will to make Genetic Programming widely used method for generating computer programs it is necessary to make sure a regular dev team of a standard organization that is hired for developing CRUD programs with popular programming languages of domain (eg. Java, C#, Go) should be able to
