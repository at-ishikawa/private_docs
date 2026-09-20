# Site Reliability Engineering

![](https://images-na.ssl-images-amazon.com/images/I/51XswOmuLqL._SL200_.jpg)

### Metadata

- Author: Betsy Beyer, Chris Jones, Jennifer Petoff, and Niall Richard Murphy
- Full Title: Site Reliability Engineering
- Category: #books

### Highlights

- What exactly is Site Reliability Engineering, as it has come to be defined at Google? My explanation is simple: SRE is what happens when you ask a software engineer to design an operations team. ([Location 337](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=337))
- In general, an SRE team is responsible for the availability, latency, performance, efficiency, change management, monitoring, emergency response, and capacity planning of their service(s). ([Location 396](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=396))
- SRE’s goal is no longer “zero outages”; rather, SREs and product developers aim to spend the error budget getting maximum feature velocity. ([Location 439](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=439))
- Extreme reliability comes at a cost: maximizing stability limits how fast new features can be developed and how quickly products can be delivered to users, and dramatically increases their cost, which in turn reduces the numbers of features a team can afford to offer. ([Location 763](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=763))
- For example, a batch process that extracts, transforms, and inserts the contents of one of our customer databases into a data warehouse to enable further analysis may be set to run periodically. Using a request success rate defined in terms of records successfully and unsuccessfully processed, we can calculate a useful availability metric despite the fact that the batch system does not run constantly. ([Location 819](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=819))
- some SRE teams focus only on high percentile values, on the grounds that if the 99.9th percentile behavior is good, then the typical experience is certainly going to be. ([Location 1128](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=1128))
- The four golden signals of monitoring are latency, traffic, errors, and saturation. If you can only measure four metrics of your user-facing system, focus on these four. ([Location 1441](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=1441))
- Signals that are collected, but not exposed in any prebaked dashboard nor used by any alert, are candidates for removal. ([Location 1506](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=1506))
- We review statistics about page frequency (usually expressed as incidents per shift, where an incident might be composed of a few related pages) in quarterly reports with management, ensuring that decision makers are kept up to date on the pager load and overall health of their teams. ([Location 1579](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=1579))
- We have built APIs for systems when no API was available from the vendor. ([Location 1665](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=1665))
- More broadly, in this view, automation is “meta-software” — software to act on software. ([Location 1680](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=1680))
- Therefore, the context for our automation is often automation to manage the lifecycle of systems, not their data: for example, deployments of a service in a new cluster. ([Location 1689](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=1689))
- “At the end of the day, our job is to keep agility and stability in balance in the system.” ([Location 2181](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=2181))
- Production, SREs develop a particularly intimate ([Location 2364](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=2364))
- As described in Chapter 6, teams send their page-worthy alerts to their on-call rotation and their important but subcritical alerts to their ticket queues. ([Location 2588](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=2588))
- Typical values are 5 minutes for user-facing or otherwise highly time-critical services, and 30 minutes for less time-sensitive systems. ([Location 2702](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=2702))
- Ideally, symptoms of operational overload should be measurable, so that the goals can be quantified (e.g., number of daily tickets < 5, paging events per shift < 2). ([Location 2801](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=2801))
- Sometimes a single abnormal condition can generate several alerts, so it’s important to regulate the alert fan-out by ensuring that related alerts are grouped together by the monitoring or alerting system. ([Location 2806](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=2806))
- Google also has a company-wide annual disaster recovery event called DiRT (Disaster Recovery Training) ([Location 2831](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=2831))
- Testing is one of the most profitable investments engineers can make to improve the reliability of their product. ([Location 4179](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=4179))
- Mandate that batch client jobs use a separate set of batch proxy backend tasks that do nothing but forward requests to the underlying backends and hand their responses back to the clients in a controlled way. Therefore, instead of “batch client → backend,” you have “batch client → batch proxy → backend.” In this case, when the very large job starts, only the batch proxy job suffers, shielding the actual backends (and higher-priority clients). Effectively, the batch proxy acts like a fuse. Another advantage of using the proxy is that it typically reduces the number of connections against the backend, which can improve the load balancing against the backend (e.g., the proxy tasks can use bigger subsets and probably have a better view of the state of the backend tasks). ([Location 5253](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=5253))
- one effective approach is to return an HTTP 503 (service unavailable) to any incoming request when there are more than a given number of client requests in flight. ([Location 5433](https://readwise.io/to_kindle?action=open&asin=B01DCPXKZ6&location=5433))
