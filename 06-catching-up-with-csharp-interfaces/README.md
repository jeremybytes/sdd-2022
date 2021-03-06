# Catching Up with Interfaces in C#: What You Know is Probably Wrong

## Abstract

C# 8 brought new features to interfaces, including default implementation, access modifiers, and static members. We'll look at these new features, and see where they are useful and where they should be avoided. The world of interfaces has changed; the line between interfaces and abstract classes has blurred; and C# now has multiple inheritance. With some practical tips, "gotchas", and plenty of examples, we'll see how to use these features effectively (and safely) in our code.  

This repository contains slides and code samples for a live presentation. Details of the various topics can be found in the accompanying articles.

## Summary of Recommendations and Opinions

**Recommendation**  
* When calling interface members, use the interface type.
* Default implementations should only reference other interface members.
* Know the capabilities of your mocking framework

**Observation**  
* Default implementation is good for calculated properties. (And not much else for properties.)

**My Opinion**  
* If code inside an interface is complex enough that it requires this type of refactoring (private members), maybe it is not appropriate for it to be part of an interface.
* Stick with "public and "private" access modifiers.
* Compatability with libraries is a good reason to support default implementation (however, I would like to see the use limited as much as possible).
* Because default implementations should call existing members, it is difficult to come up with scenarios that are not trivial.
* Interface inheritance supports a more controlled way of extending an interface.
* I like the idea of mix-ins, but I wish that it was implemented with a different construct.

## Resources

**Presentation Recordings:**  
* [Tulsa .NET User Group](http://usergroup.tv/videos/whats-new-in-c-8-interfaces-and-how-to-use-them-effectively) - Recorded Jul 23, 2020  
* [Southeast Valley .NET User Group & North West Valley .NET User Group](https://www.youtube.com/watch?v=on-MIseN6gU) - Recorded Jun 24, 2020  


**Articles:**  
* [Abstract Classes vs. Interfaces in C# - What You Know is Probably Wrong](https://jeremybytes.blogspot.com/2020/10/abstract-classes-vs-interfaces-in-c.html)  
* [A Closer Look at C# 8 Interfaces](https://jeremybytes.blogspot.com/2019/09/a-closer-look-at-c-8-interfaces.html)  
* [Properties and Default Implementation](https://jeremybytes.blogspot.com/2019/09/c-8-interfaces-properties-and-default.html)  
* [Dangerous Assumptions in Default Implementation](https://jeremybytes.blogspot.com/2019/09/c-8-interfaces-dangerous-assumptions-in.html)  
* ["dynamic" and Default Implementation](https://jeremybytes.blogspot.com/2019/09/c-8-interfaces-dynamic-and-default.html)  
* [Unit Testing Default Implementation](https://jeremybytes.blogspot.com/2019/09/c-8-interfaces-unit-testing-default.html)  
* [Public, Private, and Protected Members](https://jeremybytes.blogspot.com/2019/11/c-8-interfaces-public-private-and.html)  
* [Static Members](https://jeremybytes.blogspot.com/2019/12/c-8-interfaces-static-members.html)  
* [Static Main -- Why Not?](https://jeremybytes.blogspot.com/2019/12/c-8-interfaces-static-main-why-not.html)  

**Microsoft Docs Samples**
* [Tutorial: Update interfaces with default interface methods in C# 8.0](https://docs.microsoft.com/en-us/dotnet/csharp/whats-new/tutorials/default-interface-methods-versions)
* [Tutorial: Mix functionality in when creating classes using interfaces with default interface methods](https://docs.microsoft.com/en-us/dotnet/csharp/whats-new/tutorials/mixins-with-default-interface-methods)  

**Additional Topics**  
* [IEnumerable, ISaveable, IDontGetIt: Understanding C# Interfaces](http://www.jeremybytes.com/Demos.aspx#INT)  
* [Explicit Interface Implementation](https://jeremybytes.blogspot.com/2012/03/explicit-interface-implementation.html)  
* [Updating an Interface Implementation](https://jeremybytes.blogspot.com/2012/03/updating-interface-implementation.html) (prior to C# 8)  
* [Misusing C#: Multiple Main() Methods](https://jeremybytes.blogspot.com/2020/06/misusing-c-multiple-main-methods.html)  
* [-main (C# Compiler Options)](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-options/advanced) (how to specify the entry point if you have multiple "static Main()" methods in an application)
