# Builder

- Some objects are simple and can be created in a single constructor call.  
- Other objects require a lot of ceremony to create.
- Having a factory function with 10 arguments is not productive.
- Instead, opt for piecewise (piece-by-piece) construction.
- Builder provides an API for constructing an object step-by-step.

### When piecewise object construction is complicated, provide an API for doing it succinctly.

## Objective of the Builder pattern
### A Builder design pattern tries to:
- Abstract complex creations so that object creation is separated from the object user.
- Create an object step by step by filling its fields and creating the embedded objects.
- Reuse the object creation algorithm between many objects.
