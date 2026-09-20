# Clean Code

![](https://images-na.ssl-images-amazon.com/images/I/51d1qVhmAmL._SL200_.jpg)

### Metadata

- Author: Robert C. Martin
- Full Title: Clean Code
- Category: #books
- Document Tags: #Programming

### Highlights

- When constructors are overloaded, use static factory methods with names that describe the arguments. ([Location 1017](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=1017))
- So, another way to know that a function is doing more than “one thing” is if you can extract another function from it with a name that is not merely a restatement of its implementation ([Location 1257](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=1257))
- The ideal number of arguments for a function is zero (niladic). Next comes one (monadic), followed closely by two (dyadic). Three arguments (triadic) should be avoided where possible. More than three (polyadic) requires very special justification—and then shouldn’t be used anyway. ([Location 1357](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=1357))
- Flag arguments are ugly. Passing a boolean into a function is a truly terrible practice. ([Location 1389](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=1389))
    - **Tags:** #favorite
- In general output arguments should be avoided. If your function must change the state of something, have it change the state of its owning object. ([Location 1484](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=1484))
- Functions should either do something or answer something, but not both. Either your function should change the state of an object, or it should return some information about that object. Doing both often leads to confusion. ([Location 1486](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=1486))
- When you use exceptions rather than error codes, then new exceptions are derivatives of the exception class. They can be added without forcing any recompilation or redeployment. ([Location 1545](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=1545))
- So if you find yourself wanting to mark your closing braces, try to shorten your functions instead. ([Location 1960](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=1960))
- It appears to be possible to build significant systems (FitNesse is close to 50,000 lines) out of files that are typically 200 lines long, with an upper limit of 500. ([Location 2165](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2165))
- If one function calls another, they should be vertically close, and the caller should be above the callee, if at all possible. ([Location 2276](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2276))
- This is because the function and its arguments are closely related. ([Location 2348](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2348))
- Another use for white space is to accentuate the precedence of operators. ([Location 2350](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2350))
- Active Records are special forms of DTOs. ([Location 2631](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2631))
- This means that a change at a low level of the software can force signature changes on many higher levels. The changed modules must be rebuilt and redeployed, even though nothing they care about changed. ([Location 2730](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2730))
- Encapsulation is broken because all functions in the path of a throw must know about details of that low-level exception. ([Location 2737](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2737))
- In fact, wrapping third-party APIs is a best practice. ([Location 2774](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2774))
- This is called the SPECIAL CASE PATTERN [Fowler]. You create a class or configure an object so that it handles a special case for you. When you do, the client code doesn’t have to deal with exceptional behavior. That behavior is encapsulated in the special case object. ([Location 2799](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2799))
- Fortunately, Java has Collections.emptyList(), and it returns a predefined immutable list that we can use for this purpose: ([Location 2829](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2829))
- Because this is the case, the rational approach is to forbid passing null by default. ([Location 2857](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2857))
- If you use a boundary interface like Map, keep it inside the class, or close family of classes, where it is used. ([Location 2908](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2908))
- we could write some tests to explore our understanding of the third-party code. Jim Newkirk calls such tests learning tests. ([Location 2917](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=2917))
- What makes a clean test? Three things. Readability, readability, and readability. Readability is perhaps even more important in unit tests than it is in production code. ([Location 3086](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3086))
- The BUILD-OPERATE-CHECK2 pattern is made obvious by the structure of these tests. ([Location 3148](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3148))
- There are things that you might never do in a production environment that are perfectly fine in a test environment. ([Location 3212](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3212))
- Self-Validating The tests should have a boolean output. Either they pass or fail. You should not have to read through a log file to tell whether the tests pass. ([Location 3278](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3278))
- In this case it is because despite its small number of methods, SuperDashboard has too many responsibilities. ([Location 3364](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3364))
- We should also be able to write a brief description of the class in about 25 words, without using the words “if,” “and,” “or,” or “but.” ([Location 3369](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3369))
- The Single Responsibility Principle (SRP)2 states that a class or module should have one, and only one, reason to change. ([Location 3374](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3374))
    - **Tags:** #programming, #favorite
- A class in which each variable is used by each method is maximally cohesive. ([Location 3409](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3409))
- It supports the SRP. It also supports another key OO class design principle known as the Open-Closed Principle, or OCP:4 Classes should be open for extension but closed for modification. ([Location 3632](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3632))
- First, consider that construction is a very different process from use. ([Location 3694](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3694))
- The separation of concerns is one of the oldest and most important design techniques in our craft. ([Location 3701](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3701))
- This is the LAZY INITIALIZATION/EVALUATION idiom, and it has several merits. We don’t incur the overhead of construction unless we actually use the object, and our startup times can be faster as a result. We also ensure that null is never returned. ([Location 3705](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3705))
- We can’t compile without resolving these dependencies, even if we never actually use an object of this type at runtime! ([Location 3709](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3709))
- Because we have construction logic mixed in with normal runtime processing, we should test all execution paths (for example, the null test and its block). ([Location 3712](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3712))
- A powerful mechanism for separating construction from use is Dependency Injection (DI), the application of Inversion of Control (IoC) to dependency management. ([Location 3746](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3746))
- We often forget that it is also best to postpone decisions until the last possible moment. ([Location 3997](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=3997))
- According to Kent, a design is “simple” if it follows these rules: • Runs all the tests • Contains no duplication • Expresses the intent of the programmer • Minimizes the number of classes and methods ([Location 4066](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=4066))
- Recommendation: Avoid using more than one method on a shared object. ([Location 4345](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=4345))
- Get your nonthreaded code working first. ([Location 4384](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=4384))
- Recommendation: Do not try to chase down nonthreading bugs and threading bugs at the same time. Make sure your code works outside of threads. ([Location 4398](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=4398))
- Allow the number of threads to be easily tuned. Consider allowing it to change while the system is running. Consider allowing self-tuning based on throughput and system utilization. ([Location 4408](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=4408))
- In all cases the code under test was known to be incorrect. This just reinforced the fact that different operating systems have different threading policies, each of which impacts the code’s execution. Multithreaded code behaves differently in different environments.16 You should run your tests in every potential deployment environment. 16. Did you know that the threading model in Java does not guarantee preemptive threading? Modern OS’s support preemptive threading, so you get that “for free.” Even so, it not guaranteed by the JVM. Recommendation: Run your threaded code on all target platforms early and often. ([Location 4415](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=4415))
- Multithreaded code behaves differently in different environments. ([Location 4417](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=4417))
- Recommendation: Run your threaded code on all target platforms early and often. ([Location 4420](https://readwise.io/to_kindle?action=open&asin=B001GSTOAM&location=4420))
