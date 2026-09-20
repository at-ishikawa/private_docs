# More Effective C++

![](https://m.media-amazon.com/images/I/51-oc3ApHxL._SY160.jpg)

### Metadata

- Author: Scott Meyers
- Full Title: More Effective C++
- Category: #books

### Highlights

- If you have to worry about things like this in your software, you're probably best off avoiding references entirely. ([Location 472](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=472))
- you'll almost always want operator[] to return a reference. ([Location 517](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=517))
- This operator is used to perform type conversions whose result is nearly always implementation-defined. ([Location 605](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=605))
- to allow Rational objects to be implicitly converted to doubles (which might be useful for mixed-mode arithmetic involving Rational objects), you might define class Rational like this: class Rational { public:   ...   operator double() const;      // converts Rational to };                              // double This function would be automatically invoked in contexts like this: Rational r(1, 2);               // r has the value 1/2 double d = 0.5 * r;             // converts r to a double,                                 // then does multiplication ([Location 855](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=855))
- One of those rules is that no sequence of conversions is allowed to contain more than one user-defined conversion (i.e., a call to a single-argument constructor or an implicit type conversion operator). ([Location 969](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=969))
- In particular, prefix forms return a reference, postfix forms return a const object. ([Location 1037](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1037))
- when in doubt, do as the ints do, ([Location 1062](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1062))
- C++ prohibits it for ints, but you must prohibit it yourself for classes you write. The easiest way to do this is to make the return type of postfix increment a const object. ([Location 1071](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1071))
- First, when a function call is made, all parameters must be evaluated, so when calling the functions operator&& and operator||, both parameters are evaluated. There is, in other words, no short circuit. Second, the language specification leaves undefined the order of evaluation of parameters to a function call, so there is no way of knowing whether expression1 or expression2 will be evaluated first. ([Location 1124](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1124))
- if you can overload the comma operator, what can't you overload? As it turns out, there are limits. You can't overload the following operators: .            .*           ::           ?: new          delete       sizeof       typeid static_cast  dynamic_cast const_cast   reinterpret_cast You can overload these: operator new        operator delete operator new[]      operator delete[] +     -      *      /     %      ^      &      |     ~ !     =      <      >     +=     -=     *=     /=    %= ^=    &=     |=     <<    >>     >>=    <<=    ==    != <=    >=     &&     ||    ++     --     ,      ->*   -> ()    [] ([Location 1158](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1158))
- This function returns a pointer to a Widget object that's constructed within the buffer passed to the function. ([Location 1233](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1233))
- If you want to create an object on the heap, use the new operator. It both allocates memory and calls a constructor for the object. If you only want to allocate memory, call operator new; no constructor will be called. If you want to customize the memory allocation that takes place when heap objects are created, write your own version of operator new and use the new operator; it will automatically invoke your custom version of operator new. If you want to construct an object in memory you've already got a pointer to, use placement new. ([Location 1256](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1256))
- function signals an exceptional condition by setting a status variable or returning an error code, there is no way to guarantee the function's caller will check the variable or examine the code. ([Location 1361](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1361))
- C++ destroys only fully constructed objects, and an object isn't fully constructed until its constructor has run to completion. ([Location 1533](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1533))
- Because C++ won't clean up after objects that throw exceptions during construction, you must design your constructors so that they clean up after themselves. ([Location 1564](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1564))
- control leaves a destructor due to an exception while another exception is active, C++ calls the terminate function. ([Location 1694](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=1694))
- (The terminology for proxy objects and classes is far from universal; objects of such classes are also sometimes known as surrogates.) ([Location 5200](https://readwise.io/to_kindle?action=open&asin=B004VSMDNY&location=5200))
