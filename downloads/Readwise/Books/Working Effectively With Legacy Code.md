# Working Effectively With Legacy Code

![](https://images-na.ssl-images-amazon.com/images/I/518yKmNefUL._SL200_.jpg)

### Metadata

- Author: Michael Feathers
- Full Title: Working Effectively With Legacy Code
- Category: #books

### Highlights

- I like to call them Edit and Pray and Cover and Modify. ([Location 448](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=448))
- Testing done this way is really “testing to attempt to show correctness.” Although that is a good goal, tests can also be used in a very different way. We can do “testing to detect change.” ([Location 467](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=467))
- Unit testing is one of the most important components in legacy code work. System-level regression tests are great, but small, localized tests are invaluable. They can give you feedback as you develop and allow you to refactor with much more safety. ([Location 505](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=505))
- unit test that takes 1/10th of a second to run is a slow unit test. ([Location 549](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=549))
- The first thing to notice is that, given a choice, it is always safer to have tests around the changes that we make. ([Location 571](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=571))
- All of these problems are dependency problems. When classes depend directly on things that are hard to use in a test, they are hard to modify and hard to work with. ([Location 589](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=589))
- Identify change points. 2. Find test points. 3. Break dependencies. 4. Write tests. 5. Make changes and refactor. ([Location 626](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=626))
- Every seam has an enabling point. Let’s look at the definition of a seam again: ([Location 950](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=950))
- Separation is often a reason to use a link seam. ([Location 1019](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=1019))
- It goes to each test case class and creates a set of objects, one for each test method. That is a large set of objects, but it isn’t so bad if those objects haven’t allocated what they need yet. By placing code in setUp to create what we need just when we need it, we save quite a bit on resources. In addition, by delaying setUp, we can also run it at a time when we can detect and report any problems that might happen during setup. ([Location 1240](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=1240))
- (Those issues had been a big issue in CppUnit. Nearly every day I received e-mail from people who couldn’t use templates or the standard library, or who had exceptions with their C++ compiler.) ([Location 1269](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=1269))
- Framework for Integrated Tests (FIT) ([Location 1291](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=1291))
- you have to make a change to a class right now, try instantiating the class in a test harness. ([Location 1357](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=1357))
- What are the downsides of Sprout Method? For one thing, when you use it, in effect you essentially are saying that you are giving up on the source method and its class for the moment. You aren’t going to get it under test, and you aren’t going to make it better—you are just going to add some new functionality in a new method. ([Location 1430](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=1430))
- we use this technique repeatedly and we don’t pay attention to some key aspects of our design, it starts to degrade rapidly. ([Location 2092](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=2092))
- This is a classic example of a Liskov Substitution Principle (LSP) violation. Objects of subclasses should be substitutable for objects of their superclasses throughout our code. If they aren’t we could have silent errors in our code. ([Location 2204](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=2204))
- Here’s the issue: When we override concrete methods as we did when we overrode the getFromAddress of MessageForwarder in AnonymousMessageForwarder, we could be changing the meaning of some of the code that uses MessageFowarders. ([Location 2213](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=2213))
- Pass Null is a very handy technique in some languages. It works well in Java and C# and in just about every language that throws an exception when null references are used at runtime. ([Location 2374](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=2374))
- The Null Object Pattern is a way of avoiding the use of null in programs. ([Location 2386](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=2386))
- Good design is testable, and design that isn’t testable is bad. ([Location 2933](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=2933))
- Command/Query Separation is a design principle first described by Bertrand Meyer. ([Location 3140](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=3140))
- The only way that users of the InMemoryDirectory class can sense effects is through the getElementCount and getElement methods. If we can write tests at those methods, it appears that we should be able to cover all of the effects of our change. ([Location 3413](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=3413))
- While higher-level tests are an important tool, they shouldn’t be a substitute for unit tests. Instead, they should be a first step toward getting unit tests in place. ([Location 3600](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=3600))
- Would it be okay to write tests at only one of those classes and not the other? The key question to ask is, “If I break this method, will I be able to sense it in this place?” ([Location 3744](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=3744))
- In general, automated tests should specify a goal that we’d like to fulfill or attempt to preserve behavior that is already there. ([Location 3807](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=3807))
- Good refactoring tools check each refactoring that you attempt and disallow ones that they can’t perform safely. ([Location 5698](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=5698))
- Your bias should be toward making changes that you feel more confident in rather than changes that give you the best structure. ([Location 6225](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=6225))
- The disadvantage is that you end up with a lot of code that has to know whether it is dealing with an interface. ([Location 6879](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=6879))
- Replacing singletons is just a little more work. Add a static setter to the singleton to replace the instance, and then make the constructor protected. You can then subclass the singleton, create a fresh object, and pass it to the setter. ([Location 7047](https://readwise.io/to_kindle?action=open&asin=B005OYHF0A&location=7047))
