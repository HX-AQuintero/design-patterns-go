# SOLID Design Principles

Introduced by Robert C. Martin.
## **S**ingle Responsibility Principle (SRP)
States that a type/class/constructor should have one primary responsibility.  
*Separation of concerns:*  
- Different concerns that the system solves have to reside in different constructs/structs.  
- Break into packages.  

## **O**pen-Closed Principle (OCP)  
States that types should be open for extension, closed for modification.  
Uses *Specification Pattern.*
## **L**iskov Substitution Principle (LSP)  
States that if something takes a base class and works correctly with it,
then it should also work correctly with derived class.  
It is not really that applicable in Go, because Go doesn't has inheritance.  
However, there is a variation.
## **I**nterface Segregation Principle (ISP)  
States that you should not put too much into a single interface.  
It makes sense to break up the interface into several smaller ones.  
Grabs the needed interfaces. Avoids extra members in those interfaces.  
## **D**ependency Inversion Principle (DIP)  
High-level modules should not depend on low-level modules.  
Both should depend on abstractions.  