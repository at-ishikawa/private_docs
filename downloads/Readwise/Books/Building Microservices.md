# Building Microservices

![](https://images-na.ssl-images-amazon.com/images/I/51e6hCWFZNL._SL200_.jpg)

### Metadata

- Author: Sam Newman
- Full Title: Building Microservices
- Category: #books

### Highlights

- The risk of a service getting newly implemented fault tolerance wrong is high if it could impact more of the system. ([Location 597](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=597))
- I’ve seen this approach lead to disaster when these server-side endpoints become thick layers with too much behavior. ([Location 1578](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=1578))
- You could resolve this by exposing batch APIs to make reporting easier. For example, our customer service could allow you to pass a list of customer IDs to it to retrieve them in batches, or may even expose an interface that lets you page through all the customers. ([Location 2025](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=2025))
- A common example of this is the smoke test suite, a collection of tests designed to be run against newly deployed software to confirm that the deployment worked. ([Location 3069](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=3069))
- In my opinion, monitoring is one area where standardization is incredibly important. ([Location 3386](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=3386))
- Any organization that designs a system (defined more broadly here than just information systems) will inevitably produce a design whose structure is a copy of the organization’s communication structure. ([Location 3874](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=3874))
- In many situations, the feature team is a reaction to traditional IT organizations where team structure is aligned around technical boundaries. ([Location 3964](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=3964))
- Even if we’d had the timeouts on the pool set correctly, we were also sharing a single HTTP connection pool for all outbound requests. This meant that one slow service could exhaust the number of available workers all by itself, even if everything else was healthy. ([Location 4227](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=4227))
- use data pumps or event data pumps to consolidate data across multiple services for reporting purposes. ([Location 4974](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=4974))
- Greenfield development is also quite challenging. It isn’t just that the domain is also likely to be new; it’s that it is much easier to chunk up something you have than something you don’t! So again, consider starting monolithic first and break things out when you’re stable. ([Location 5020](https://readwise.io/to_kindle?action=open&asin=B00T3N7XB4&location=5020))
