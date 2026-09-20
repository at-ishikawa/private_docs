# Effective Java

![](https://images-na.ssl-images-amazon.com/images/I/51wl8cINKYL._SL200_.jpg)

### Metadata

- Author: Joshua Bloch
- Full Title: Effective Java
- Category: #books

### Highlights

- In cases where a class seems to require multiple constructors with the same signature, replace the constructors with static factory methods and carefully chosen names to highlight their differences. ([Location 399](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=399))
- a single-element enum type is the best way to implement a singleton. ([Location 655](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=655))
- prefer primitives to boxed primitives, and watch out for unintentional autoboxing. ([Location 751](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=751))
- never do anything time-critical in a finalizer. ([Location 836](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=836))
- There is no way to extend an instantiable class and add a value component while preserving the equals contract, ([Location 1038](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=1038))
- it is now legal for an overriding method's return type to be a subclass of the overridden method's return type. ([Location 1400](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=1400))
    - **Tags:** #programming
- the clone architecture is incompatible with normal use of final fields referring to mutable objects, ([Location 1430](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=1430))
- a class that is designed for inheritance (Item 17) overrides clone, the overriding method should mimic the behavior of Object.clone: it should be declared protected, it should be declared to throw CloneNotSupportedException, and the class should not implement Cloneable. ([Location 1477](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=1477))
- the hard-liners are correct when it comes to public classes: if a class is accessible outside its package, provide accessor methods, ([Location 1729](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=1729))
- Technically it's not delegation unless the wrapper object passes itself to the wrapped object ([Location 2004](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=2004))
- The cast method is the dynamic analog of Java's cast operator. It simply checks that its argument is an instance of the type represented by the Class ([Location 3146](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=3146))
- It provides nothing in the way of type safety and little in the way of convenience. ([Location 3224](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=3224))
- If the int associated with an enum constant is changed, its clients must be recompiled. ([Location 3232](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=3232))
- @Target(ElementType.METHOD) meta-annotation indicates that the Test annotation is legal only on method declarations: ([Location 3642](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=3642))
- A safe, conservative policy is never to export two overloadings with the same number of parameters. ([Location 4082](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=4082))
- It is no longer necessary to use the HTML <code> or <tt> tags in doc comments: the Javadoc {@code} tag is preferable because it eliminates the need to escape HTML metacharacters. ([Location 4316](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=4316))
- as it computes the limit of the array index only once. ([Location 4468](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=4468))
- use BigDecimal, int, or long for monetary calculations. ([Location 4595](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=4595))
- To achieve acceptable performance, use a StringBuilder in place of a String ([Location 4750](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=4750))
- Performance suffers. ([Location 4823](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=4823))
- If it is impossible to prevent exceptions from lower layers, the next best thing is to have the higher layer silently work around these exceptions, insulating the caller of the higher-level method from lower-level problems. ([Location 5251](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=5251))
- If an exception is thrown by many methods in a class for the same reason, it is acceptable to document the exception in the class's documentation comment ([Location 5284](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=5284))
- synchronization has no effect unless both read and write operations are synchronized. ([Location 5442](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=5442))
- While the volatile modifier performs no mutual exclusion, it guarantees that any thread that reads the field will see the most recently written value: ([Location 5447](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=5447))
- The problem is that the increment operator (++) is not atomic. ([Location 5460](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=5460))
- follow the advice in Item 47 and use the class AtomicLong, ([Location 5469](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=5469))
- inside a synchronized region, do not invoke a method that is designed to be overridden, or one provided by a client in the form of a function object (Item 21). ([Location 5493](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=5493))
- If the class has invariants that would be violated if its instance fields were initialized to their default values (zero for integral types, false for boolean, and null for object reference types), you must add this readObjectNoData method to the class: ([Location 6039](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6039))
- not to implement Serializable. ([Location 6046](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6046))
- Note that the variables that store the object's state (x and y) can't be final, as they are set by the initialize method: ([Location 6058](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6058))
- Even if you decide that the default serialized form is appropriate, you often must provide a readObject method to ensure invariants and security. ([Location 6115](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6115))
- Even if all instance fields are transient, invoking defaultWriteObject affects the serialized form, resulting in greatly enhanced flexibility. ([Location 6164](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6164))
- you must impose any synchronization on object serialization that you would impose on any other method that reads the entire state of the object. ([Location 6200](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6200))
- When an object is deserialized, it is critical to defensively copy any field containing an object reference that a client must not possess. ([Location 6304](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6304))
- would you feel comfortable adding a public constructor that took as parameters the values for each nontransient field in the object and stored the values in the fields with no validation whatsoever? ([Location 6327](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6327))
- For classes with object reference fields that must remain private, defensively copy each object in such a field. ([Location 6339](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6339))
- if you depend on readResolve for instance control, all instance fields with object reference types must be declared transient. ([Location 6366](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6366))
- it need not do any consistency checking or defensive copying. ([Location 6435](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6435))
- consider the serialization proxy pattern whenever you find yourself having to write a readObject or writeObject method on a class that is not extendable by its clients. This pattern is perhaps the easiest way to robustly serialize objects with nontrivial invariants. ([Location 6495](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6495))
- Java™ Object Serialization Specification. ([Location 6570](https://readwise.io/to_kindle?action=open&asin=B000WJOUPA&location=6570))
