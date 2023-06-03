# Clean-Code-Architecture
Clean Architecture is designed to separate concerns by organizing code into several layers with a very explicit rule which enables us to create a testable and maintainable project.

## Advantages of Clean Code Architecture
1. UI independent: changing the UI should not affect other parts of the projects.
2. Framework independent: it doesn’t matter what library you’re using. frameworks will be used as tools and will not force any business rules.
3. Database independent: the project shouldn’t care or depend on the chosen database. moreover, we would want to choose our database as late as possible in order to maintain flexibility
4. Testable: a system that is testable is a system that you have confidence working with — therefor it’s a system that is easy to extend or maintain.
5. Strict on it’s dependency rules: source code dependencies can only point INWARDS, meaning that an inner circle will know nothing about an outer circle.

### Entities
1. Entities encapsulate the most general and high-level business rules.
They are the least likely to change when something external changes.
2. The entities do not depend on any of the other circles.

### Use Case:
1. The use case layer contains application specific business rules.
2. The use case layer only depends on the entities. change in the entities will require a change in the use case layer, but changes to other layers won’t.

### Interface Adapters:
Interface adapters layer contains a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web.

### Repository
Here we write and handle the data base operations.

### Handlers 
Here we do the data management. Here we recieve the data from the user and give them the output.

### DB
Data Base connection and migartion is been done in this section

### DI
This will hold the diffrent Dependencies which are injected and Wire package has been used.

### Utils
This section handles the utility functions such as password hashing, email creation etc.
