# Diving Deeper into Dependency Injection

## Abstract  

You know the basics of dependency injection (DI). Now it's time to take a closer look at how DI patterns and other design patterns can help us use DI effectively. We'll look at implementations and uses for DI patterns including constructor injection, method injection, and property injection. In addition, we'll use other design patterns to add functionality to existing objects and to manage disposable dependencies. We'll leave with several practical ways to improve the functionality and testing of our code.  

Code samples and slides are included in this repository.

## Topics + Code  

The following connects the topics with the sample code.  

**Constructor Injection**  
* [PeopleController.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer/Controllers/PeopleController.cs) - Web application controller
* [PeopleViewerWindow.xaml.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer.View/PeopleViewerWindow.xaml.cs) - Desktop application main window
* [PeopleViewModel.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer.Presentation/PeopleViewModel.cs) - Desktop application view model

**Property Injection**  
* [CSVReader.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PersonDataReader.CSV/CSVReader.cs) - CSV File data reader  

**Method Injection**  
* [MainWindow.xaml.cs](/04-diving-deeper-into-dependency-injection/MethodInjection/PeopleViewer/MainWindow.xaml.cs) - person.ToString(selectedFormatter)  

**Read-Only / Guard Clause**  
* [PeopleViewModel.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer.Presentation/PeopleViewModel.cs) - Desktop application view model
* [PeopleViewerWindow.xaml.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer.View/PeopleViewerWindow.xaml.cs) - Desktop application main window

**Decorators**
* [CachingReader.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PersonDataReader.Decorators/CachingReader.cs) - Local cache decorator  
* [ExceptionLoggingReader.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PersonDataReader.Decorators/ExceptionLoggingReader.cs) - Exception logging decorator  
* [RetryReader.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PersonDataReader.Decorators/RetryReader.cs) - Retry decorator  
* [Program.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer/Program.cs) - Web application composition (with decorators)  
* [App.xaml.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer.Desktop/App.xaml.cs) - Desktop application composition (with decorators)  

**Proxy / IDisposable**  
* [SQLReaderProxy.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PersonDataReader.SQL/SQLReaderProxy.cs) - Proxy to wrap IDisposable SQL Reader  

**Static Dependencies**  
* [ITimeProvider.cs](/04-diving-deeper-into-dependency-injection/StaticDependencies/HouseControl.Library/Schedules/ITimeProvider.cs) - Wrapper interface
* [CurrentTimeProvider.cs](/04-diving-deeper-into-dependency-injection/StaticDependencies/HouseControl.Library/Schedules/CurrentTimeProvider.cs) - Default that uses "DateTimeOffset.Now"  
* [ScheduleHelper.cs](/04-diving-deeper-into-dependency-injection/StaticDependencies/HouseControl.Library/Schedules/ScheduleHelper.cs) - Static property that uses current time by default  
* [ScheduleHelperTests.cs](/04-diving-deeper-into-dependency-injection/StaticDependencies/HouseControl.Library.Test/ScheduleHelperTests.cs) - Tests that swap out time provider for deterministic tests

**Configuration Strings**  
* [ServiceReader.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PersonDataReader.Service/ServiceReader.cs) - Service reader that has 'string' dependency
* [ServiceReaderUri.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PersonDataReader.Service/ServiceReaderUri.cs) - Strongly-typed wrapper around 'string'
* [Program.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer/Program.cs) - Web application composition (with ServiceReaderUri)  
* [App.xaml.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer.Desktop/App.xaml.cs) - Desktop application composition (with ServiceReaderUri)  
* [App.xaml.cs](/04-diving-deeper-into-dependency-injection/MainDemo/PeopleViewer.Desktop.Ninject/App.xaml.cs) - Desktop application composition using Ninject container  

## Resources

**DI Patterns**  
* [Dependency Injection: The Property Injection Pattern](http://jeremybytes.blogspot.com/2014/01/dependency-injection-property-injection.html)  
* [Property Injection: Simple vs. Safe](http://jeremybytes.blogspot.com/2015/06/property-injection-simple-vs-safe.html)  
* [Dependency Injection: The Service Locator Pattern](http://jeremybytes.blogspot.com/2013/04/dependency-injection-service-locator.html)  

**Decorators and Async Interfaces**
* [Async Interfaces, Decorators, and .NET Standard](https://jeremybytes.blogspot.com/2019/01/more-di-async-interfaces-decorators-and.html)  
* [Async Interfaces](https://jeremybytes.blogspot.com/2019/01/more-di-async-interfaces.html)  
* [Adding Retry with the Decorator Pattern](https://jeremybytes.blogspot.com/2019/01/more-di-adding-retry-with-decorator.html)  
* [Unit Testing Async Methods](https://jeremybytes.blogspot.com/2019/01/more-di-unit-testing-async-methods.html)  
* [Adding Exception Logging with the Decorator Pattern](https://jeremybytes.blogspot.com/2019/01/more-di-adding-exception-logging-with.html)  
* [Adding a Client-Side Cache with the Decorator Pattern](https://jeremybytes.blogspot.com/2019/01/more-di-adding-client-side-cache-with.html)  
* [The Real Power of Decorators -- Stacking Functionality](https://jeremybytes.blogspot.com/2019/01/more-di-real-power-of-decorators.html)  

**Challenges**  
* Static Objects: [Mocking Current Time with a Simple Time Provider](https://jeremybytes.blogspot.com/2015/01/mocking-current-time-with-time-provider.html)  

**Related Topics**
* Session: [DI Why? Getting a Grip on Dependency Injection](http://www.jeremybytes.com/Demos.aspx#DI)
* Pluralsight: [Getting Started with Dependency Injection in .NET](https://app.pluralsight.com/library/courses/using-dependency-injection-on-ramp/table-of-contents) 

More information at [http://www.jeremybytes.com](http://www.jeremybytes.com)  