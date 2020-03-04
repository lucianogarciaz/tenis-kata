## Considerations and Notes 
## ``Tests``
#### Why I didn't use assertion in tests? 
> Go doesn't provide assertions. They are undeniably convenient, but our experience has been that programmers use them as a crutch to avoid thinking about proper error handling and reporting. 

**[more](https://golang.org/doc/faq#assertions)**

#### Why I wrote test cases in different packages?
When we write test cases inside a package, these test files are exposed to exported and non-exported members of the package (because they are in the same directory of the source files). So we are not sure, not whether or not these unit components are accessible to the end-user.
#### Object Mother
I rather prefer to use Object Mother instead of Builder Pattern, 
but always depend on what you are testing.

[Object Mother](https://martinfowler.com/bliki/ObjectMother.html)
#### Given-When-Then or Arrange, Act, Assert 🔋 
Apart from Acceptance test, it's nice to have organized Unit Test with these patterns

[Given-When-Then](https://martinfowler.com/bliki/GivenWhenThen.html)

## ``Patterns``

#### Tell don't ask
[Martin Fowler](https://martinfowler.com/bliki/TellDontAsk.html)
#### Remove Setters: Immutability
If you force to create classes with every attribute needed, you will (not) have fewer problems with unwanted states.  
If you need to set value after their creation (eg. you have optional attributes ), use it with semantics (accessors).
#### GetterEradicator:
[Getter Eradicator](https://martinfowler.com/bliki/GetterEradicator.html)

#### Named Constructor
It's a pattern that allows us to instantiate an object semantically. 

#### SOLID
[SRP](https://blog.cleancoder.com/uncle-bob/2014/05/08/SingleReponsibilityPrinciple.html)
> Imagine you took your car to a mechanic to fix a broken electric window. He calls you the next day saying it’s all fixed. When you pick up your car, you find the window works fine; but the car won’t start. 😂 

Nice article to read: [DIP](https://martinfowler.com/articles/dipInTheWild.html)


``To sum up:`` 
As you can see, all of the patterns that I try to use is to make my code more readable, understandable, simple as possible and above all *TESTABLE*. Because at the very end, this translates into productivity and software economics.
Obviously, working with patterns has pros and cons and always depends on the situation that you are going to use it and what is best for the product (company).

Some things that have inspired me: 
###### **Articles:**
[SOLID](https://scotch.io/bar-talk/s-o-l-i-d-the-first-five-principles-of-object-oriented-design)

[Design Patterns](https://refactoring.guru/design-patterns)

Patterns are agnostic, so: [Dessing Patterns php](https://designpatternsphp.readthedocs.io/en/latest/README.html)

[DIP](https://martinfowler.com/articles/dipInTheWild.html)
###### **Books:**
Buenos Vinos [DDD](https://leanpub.com/ddd-in-php)

Minimum viable [test](https://leanpub.com/minimumviabletests)
