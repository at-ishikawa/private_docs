# Effective JavaScript

![](https://images-na.ssl-images-amazon.com/images/I/51W25NBDLQL._SL200_.jpg)

### Metadata

- Author: David Herman
- Full Title: Effective JavaScript
- Category: #books

### Highlights

- operating on their arguments directly as floating-point numbers, they implicitly convert them to 32-bit integers. ([Location 420](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=420))
- Since NaN is the only JavaScript value that is treated as unequal to itself, you can always test if a value is NaN by checking it for equality to itself: Click here to ([Location 560](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=560))
- strange consequence of this implicit wrapping is that you can set properties on primitive values with essentially no ([Location 713](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=713))
- Consider using lint tools to help check for unbound variables. ([Location 1197](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=1197))
- Variable definitions are not scoped to their nearest enclosing statement or block, but rather to their containing function. ([Location 1413](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=1413))
- That is, try...catch binds a caught exception to a variable that is scoped just to the catch block: ([Location 1466](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=1466))
- But in fact, it contains a reference to ([Location 1504](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=1504))
- Closures store their outer variables by reference, not by value. ([Location 1507](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=1507))
- Unlike a function declaration, a named function expression can’t be referred to externally by its internal name: ([Location 1593](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=1593))
- The real usefulness of named function expressions, though, is for debugging. ([Location 1617](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=1617))
- But recall that JavaScript is not block-scoped, so the inner f should be in scope for the whole body of test. A reasonable second guess would be ["local", "local"] and ["local"]. And in fact, some JavaScript environments behave this way. But not all of them! Others conditionally bind the inner f at runtime, based on whether its enclosing block is executed. ([Location 1720](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=1720))
- This helps catch accidental misuse of methods as plain functions by failing more quickly, since attempting to access properties of undefined immediately throws an error. ([Location 1932](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=1932))
- This allows you to fix any bugs in the logic just once, instead of having to hunt for every instance of the coding pattern spread throughout your program. ([Location 2062](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=2062))
- Use the call method to call ([Location 2139](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=2139))
- The apply method takes an array of arguments and calls the function as if each element of the array were an individual argument of the call. ([Location 2186](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=2186))
- This means that while we appear to be extracting obj["add"], we are actually extracting 17[25]! ([Location 2301](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=2301))
- As a consequence, it is much safer never to modify the arguments object. This is easy enough to avoid by first copying its elements to a real array. ([Location 2324](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=2324))
- Keep in mind that buffer.add.bind(buffer) creates a new function rather than modifying the buffer.add function. The new function behaves just like the old one, but with its receiver bound to buffer, while the old one remains unchanged. ([Location 2474](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=2474))
- This way, the result of calling User is an object that inherits from User.prototype, ([Location 2904](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=2904))
- this.separators = separators || [","]; ([Location 3101](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=3101))
- Programmers commonly use the variable name self for this pattern, ([Location 3172](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=3172))
- Yet another valid approach in ES5 is to use the callback function’s bind method, similar to the approach described in Item 25 ([Location 3177](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=3177))
- This syntax is equally convenient, but where Object.create is available, it is the more reliable approach. ([Location 3667](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=3667))
- "alice" in dict;        // false "bob" in dict;          // false "chris" in dict;        // false "toString" in dict;     // true "valueOf" in dict;      // true ([Location 3698](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=3698))
- This approach works regardless of whether its receiver has overridden its hasOwnProperty method: ([Location 3743](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=3743))
- Admittedly, this code is a mouthful. But this version has the distinct advantage of not polluting every other for...in loop over every other instance of Object. ([Location 4013](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=4013))
- This nondeterminism comes from the fact that the for...in loop may choose a different order of enumeration in different JavaScript environments (or even in different executions within the same JavaScript environment, at least in principle). ([Location 4141](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=4141))
- Remember that object property keys are always strings, ([Location 4198](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=4198))
- Instead, we have to extract a reference to the forEach method object and use its call method (see Item 20): ([Location 4408](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=4408))
- stateless APIs tend to be easier to learn and use, more self-documenting, and less error-prone. ([Location 4893](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=4893))
- To understand what any individual call to fillText does, you don’t have to understand all the modifications that precede it. ([Location 4933](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=4933))
- This last version of enable is an example of a more cautious style known as defensive programming, which attempts to defend against potential errors with additional checks. ([Location 5274](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=5274))
- Shifting by zero bits then has no effect on the integer value. ([Location 5308](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=5308))
- Method chaining for stateful APIs is sometimes known as the fluent style. (The term was coined by programmers simulating Smalltalk’s “method cascades”; a built-in syntax for calling multiple methods on a single object.) ([Location 5410](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=5410))
- The local tryNextURL function is recursive: Its implementation involves a call to itself. ([Location 5823](https://readwise.io/to_kindle?action=open&asin=B00AC1RP14&location=5823))
