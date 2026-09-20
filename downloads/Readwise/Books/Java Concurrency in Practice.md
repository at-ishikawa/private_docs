# Java Concurrency in Practice

![](https://images-na.ssl-images-amazon.com/images/I/51DWjuuzJoL._SL200_.jpg)

### Metadata

- Author: Tim Peierls, Joseph Bowbeer, Joshua Bloch, Brian Goetz, Doug Lea, David Holmes
- Full Title: Java Concurrency in Practice
- Category: #books

### Highlights

- In the absence of synchronization, the compiler, hardware, and runtime are allowed to take substantial liberties with the timing and ordering of actions, such as caching variables in registers or processor-local caches where they are temporarily (or even permanently) invisible to other threads. ([Location 768](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=768))
- When designing thread-safe classes, good object-oriented techniques—encapsulation, immutability, and clear specification of invariants—are your best friends. ([Location 890](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=890))
- mutexes (or mutual exclusion locks), which means that at most one thread may own the lock. ([Location 1083](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1083))
- Reentrancy means that locks are acquired on a per-thread rather than per-invocation basis. ([Location 1100](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1100))
- Every shared, mutable variable should be guarded by exactly one lock. Make it clear to maintainers which lock that is. ([Location 1135](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1135))
- Synchronization also has another significant, and subtle, aspect: memory visibility. ([Location 1225](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1225))
- There is no guarantee that operations in one thread will be performed in the order given by the program, as long as the reordering is not detectable from within that thread—even if the reordering is apparent to other threads. ([Location 1246](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1246))
- always use the proper synchronization whenever data is shared across threads. ([Location 1256](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1256))
- at least it sees a value that was actually placed there by some thread rather than some random value. This safety guarantee is called out-of-thin-air safety. ([Location 1279](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1279))
- To ensure that all threads see the most up-to-date values of shared mutable variables, the reading and writing threads must synchronize on a common lock. ([Location 1300](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1300))
- When a field is declared volatile, the compiler and runtime are put on notice that this variable is shared and that operations on it should not be reordered with other memory operations. ([Location 1303](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1303))
- Volatile variables are not cached in registers or in caches where they are hidden from other processors, so a read of a volatile variable always returns the most recent write by any thread. ([Location 1305](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1305))
- Good uses of volatile variables include ensuring the visibility of their own state, that of the object they refer to, or indicating that an important lifecycle event (such as initialization or shutdown) has occurred. ([Location 1319](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1319))
- Locking can guarantee both visibility and atomicity; volatile variables can only guarantee visibility. ([Location 1334](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1334))
- You can use volatile variables only when all the following criteria are met: Writes to the variable do not depend on its current value, or you can ensure that only a single thread ever updates the value; The variable does not participate in invariants with other state variables; and Locking is not required for any other reason while the variable is being accessed. ([Location 1335](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1335))
- There's nothing wrong with creating a thread in a constructor, but it is best not to start the thread immediately. ([Location 1386](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1386))
- Initializing an object reference from a static initializer; Storing a reference to it into a volatile field or AtomicReference; Storing a reference to it into a final field of a properly constructed object; or Storing a reference to it into a field that is properly guarded by a lock. ([Location 1585](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1585))
- ListHelper provides only the illusion of synchronization; the various list operations, while all synchronized, use different locks, which means that putIfAbsent is not atomic relative to other operations on the List. So there is no guarantee that another thread won't modify the list while putIfAbsent is executing. ([Location 1965](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=1965))
- Counting semaphores are used to control the number of activities that can access a certain resource or perform a given action at the same time [CPJ 3.4.1]. ([Location 2475](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=2475))
- Activities that do not support cancellation but still call interruptible blocking methods will have to call them in a loop, retrying when interruption is detected. ([Location 3180](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=3180))
- Avoid premature optimization. First make it right, then make it fast—if it is not already fast enough. ([Location 4459](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=4459))
- Compilers may generate instructions in a different order than the “obvious” one suggested by the source code, or store variables in registers instead of in memory; processors may execute instructions in parallel or out of order; caches may vary the order in which writes to variables are committed to main memory; and values stored in processor-local caches may not be visible to other processors. ([Location 6441](https://readwise.io/to_kindle?action=open&asin=B004V9OA84&location=6441))
