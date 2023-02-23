# Background (TDD-related topics)

## Test driven development

Test driven development are implemented by developers with writing so-called tests before actually implementing the program. Those tests can be divided into a couple of classes. Most popularly adopted and well known examples are unit tests and integration tests. Unit test is when a component tested by itself in isolation from the components it supposed to interact in active deployment. So, when a unit test failed, smallest component can be addressed immediately by tracking lately-turned-to-fail tests. Other class is when systems that consists by more than one components deployed into real or mock (a.k.a test, development) environment and tested for higher-level requirements than unit tests. To understand relation between test types, thinking about that system definition may help: A system consists by components and its behavior emerges by the interactions of components, also each component can be seen as sub-systems that consists by sub-components smaller than itself. Thus, unit test are for smallest components, and at each organizational-level from that to the higher, different type of test tools and techniques are considered.

Advantages of TDD:

-   Reducing cost of maintenance mistakes due to employee resign / work force circulation.
-   Increased security and robustness overall.
-   Increased robustness against future workers making mistakes due to lack of prior knowledge about and incorrect removals of existent structures in code.

### Requirements of commercialization

#### Changing requirements of projects

Software projects (service/product) mature with time and the increasing diffusion area on target group. Early adopters of innovations are the people who are expected to be the most educated/experienced ones amongst the whole target group. Through diffusion process, it is expected to feedbacks and requests raised by customers differantiate with the diffusion phase. Thus it is unavoidable that requirements for a particular product or a service will change with time; and software development process should

#### Work force circulation

Companies can not rely their employees on staying in the company until a software project completes its lifecycle. Employees come and go, but products should keep get fixed for newly discovered requirements, and get updated for new features customers request without effecting by if employees who worked on first phases of project still in the company. Latest workforce should be able to work on the existing codebase to add new features without the fear of breaking previously working features that implemented by previous workers. Following by test driven approach, key requirements that the software expected to satisfy are added into the codebase as tests at the moment those requirements raised and their reasoning known by employees of that time. Codebase is supposed to get tested for all tests at any time a possibly breaking-change occurred, or before and after every deployment by the CI/CD pipeline.
