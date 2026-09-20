# System Design Interview – An Insider&#39;s Guide

![](https://m.media-amazon.com/images/I/61qzVb-+eAL._SY160.jpg)

### Metadata

- Author: Alex Xu and Sahn Lam
- Full Title: System Design Interview – An Insider&#39;s Guide
- Category: #books

### Highlights

- Are there any other things we need to consider? ([Location 2733](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=2733))
- As we can see, the average reservation transaction per second (TPS) is not high. ([Location 2748](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=2748))
- A relational database works well with read-heavy and write less frequently workflows. This is because the number of users who visit the hotel website/apps is a few orders of magnitude higher than those who actually make reservations. NoSQL databases are generally optimized for writes and the relational database works well enough for read-heavy workflow. ([Location 2803](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=2803))
- A user actually reserves a type of room in a given hotel instead of a specific room. ([Location 2818](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=2818))
- Room numbers are given when the guest checks in and not at the time of the reservation. ([Location 2820](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=2820))
- Due to these limitations, we do not recommend pessimistic locking for the reservation system. ([Location 3000](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=3000))
- A database validation check is put in place; the next version value should exceed the current version value by 1. The transaction aborts if the validation fails and the user tries again from step 2. ([Location 3009](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=3009))
- Optimistic locking is usually faster than pessimistic locking because we do not lock the database. However, the performance of optimistic locking drops dramatically when concurrency is high. ([Location 3011](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=3011))
- These clients have to retry. In the subsequent round of retries, there is only one successful client again, and the rest have to retry. Although the end result is correct, repeated retries cause a very unpleasant user experience. ([Location 3015](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=3015))
- Optimistic locking is a good option for a hotel reservation system since the QPS for reservations is usually not high. ([Location 3022](https://readwise.io/to_kindle?action=open&asin=B0CR977BQH&location=3022))
