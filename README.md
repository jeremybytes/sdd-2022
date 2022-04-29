# Software Design & Development (SDD 2022) Presentations  

All of the slides and code samples for Jeremy's presentations at SDD 2022 (16-20 May, 2022).  

## Topics

**Learn to Love Lambdas in C# (and LINQ, Too!)**  
Lambda expressions in C# can be confusing the first time you walk up to them. But once you get to know them, you’ll see that they are a great addition to your toolbox. Used properly, they can add elegance and simplicity to your code. And some .NET constructs (such as LINQ) lend themselves to lambda expressions. In addition, lambdas let us scope our variables more appropriately with captured variables. We’ll take a look at how lambda expressions work and see them in action. We’ll also see how LINQ can help us move from imperative programming to declarative programming (a gateway to functional-style programming).

* Slides: [learn-to-love-lambdas.pdf](./learn-to-love-lambdas.pdf)  
* Code: [01-learn-to-love-lambdas/](./01-learn-to-love-lambdas/)

---

**Get Func-y: Delegates in C#**  
Delegates are the gateway to functional programming. So let’s understand delegates and how we can change the way we program by using functions as parameters, return types, variables, and properties. In addition, we’ll see how the built in delegate types, Func and Action, are waiting to make our lives easier. By the time we’re done, we’ll see how delegates can add elegance, extensibility, and safety to our programming. And as a bonus, you’ll have a few functional programming concepts to take with you.  

* Slides: [get-funcy-delegates-in-csharp.pdf](./get-funcy-delegates-in-csharp.pdf)
* Code: [02-get-funcy-delegates-in-csharp/](./02-get-funcy-delegates-in-csharp/)

---

**A Tour of Go for the C# Developer**  
Learning other programming languages enhances our work in our primary language. From the perspective of a C# developer, the Go language (golang) has many interesting ideas. Go is opinionated on some things (such as where curly braces go and what items are capitalized). Declaring an unused variable causes a compile failure; the use of "blank identifiers" (or "discards" in C#) are common. Concurrency is baked right in to the language through goroutines and channels. Programming by exception is discouraged; it's actually called a "panic" in Go. Instead, errors are treated as states to be handled like any other data state. We'll explore these features (and others) by building an application that uses concurrent operations to get data from a service. These ideas make us think about the way we program and how we can improve our day-to-day work (in C# or elsewhere).  

* Slides: [a-tour-of-go-for-the-csharp-dev.pdf](./a-tour-of-go-for-the-csharp-dev.pdf)
* Code: [03-a-tour-of-go-for-the-csharp-dev/](./03-a-tour-of-go-for-the-csharp-dev/)

---

**Diving Deeper into Dependency Injection**  
You know the basics of dependency injection (DI). Now it's time to take a closer look at how DI patterns and other design patterns can help us use DI effectively. We'll look at implementations and uses for DI patterns including constructor injection, method injection, and property injection. In addition, we'll use other design patterns to add functionality to existing objects and to manage disposable dependencies. We'll leave with several practical ways to improve the functionality and testing of our code.

* Slides: [diving-deeper-into-dependency-injection.pdf](./diving-deeper-into-dependency-injection.pdf)
* Code: [04-diving-deeper-into-dependency-injection/](./04-diving-deeper-into-dependency-injection/)

---

**Better Parallel Code with C# Channels**  
Producer/consumer problems show up in a lot of programming scenarios, including data processing and machine learning. Channels were added to .NET Core 3.0 and give us a thread-safe way to communicate between producers and consumers, and we can run them all concurrently. In this presentation, we will explore channels by comparing parallel tasks with continuations to using a producer/consumer model. In the end, we'll have another tool in our toolbox to help us with concurrent programming.

* Slides: [better-parallel-code-with-csharp-channels.pdf](./better-parallel-code-with-csharp-channels.pdf)
* Code: [05-better-parallel-code-with-csharp-channels/](./05-better-parallel-code-with-csharp-channels/)

---

**Catching up with C# Interfaces: What You Know is Probably Wrong**

C# 8 brought new features to interfaces, including default implementation, access modifiers, and static members. We'll look at these new features, and see where they are useful and where they should be avoided. The world of interfaces has changed; the line between interfaces and abstract classes has blurred; and C# now has multiple inheritance. With some practical tips, "gotchas", and plenty of examples, we'll see how to use these features effectively (and safely) in our code.

* Slides: [catching-up-with-csharp-interfaces.pdf](./catching-up-with-csharp-interfaces.pdf)
* Code: [06-catching-up-with-csharp-interfaces/](./06-catching-up-with-csharp-interfaces/)
---