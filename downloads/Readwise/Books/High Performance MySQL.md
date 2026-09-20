# High Performance MySQL

![](https://images-na.ssl-images-amazon.com/images/I/51rXPG3cVDL._SL200_.jpg)

### Metadata

- Author: Jeremy D. Zawodny and Derek J. Balling
- Full Title: High Performance MySQL
- Category: #books
- Document Tags: #Infrastructure

### Highlights

- The only hot spot in page locking is the last page in the table. If records are inserted there at regular intervals, the last page will be locked frequently. ([Location 658](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=658))
    - **Note:** The disadvantage to use a page lock with DMLs
- Row-level locking, as it’s commonly known, is available in MySQL’s InnoDB tables. InnoDB doesn’t use a simple row locking mechanism, however. Instead it uses row-level locking in conjunction with a multiversioning scheme, ([Location 663](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=663))
    - **Note:** Row level locking isn't simple in InnoDB. There is a multiversioning scheme.
- MySQL provides two transaction-safe storage engines: Berkeley DB (BDB) and InnoDB. ([Location 885](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=885))
    - **Note:** There is another storage engine Berkeley DB (BDB)
- Another option is to use a MyISAM Merge table. Rather than always logging to the same table, adjust the application to log to a table that contains the name or number of the month in its name, such as web_logs_2004_01 or web_logs_2004_jan. ([Location 982](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=982))
    - **Note:** MyISAM daily partition is called Merge table, and it's option for logging in MySQL
- A final possibility is simply to switch to using a table that has more granular locking than MyISAM does. Either BDB or InnoDB works well in this case. Non-MyISAM tables will generally use more CPU and disk space, but that may be a reasonable tradeoff in this case. Also, in the event of a crash, MyISAM tables may take quite a long time to check and repair while InnoDB tables should recover quickly. ([Location 988](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=988))
    - **Note:** Inno DB has
      * row locking though more spaces
      * repair easily when crash
- If you import into InnoDB or BDB, be sure to use the --no-autocommit option to disable AUTOCOMMIT mode. Otherwise each individual insert will be performed in its own transaction. ([Location 1053](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=1053))
    - **Note:** mysqldump optoin to import
- In fact, MySQL will only ever use one index per table per query — except for UNIONs.[3] This fact is important enough to say again: MySQL will only ever use one index per table per query. ([Location 1776](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=1776))
    - **Note:** The best practice of index usage
- when MySQL believes more than about 30% of the rows are likely matches, it will resort to a table scan rather than using the index. ([Location 2069](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=2069))
    - **Note:** 30% is a criteria whether index can be used or not
- If you’ve already covered the causes listed earlier and implemented the suggestions, it’s likely that you need to spread the I/O load more effectively. ([Location 3223](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=3223))
    - **Note:** To solve I/O bottlenecks except slow queries, using faster RAMs, using RAID, or separating hosts are options
- The slave performs a sanity check, comparing its result with the master’s. If the query failed on the slave but succeeded on the master, replication stops. The reverse is also true. If the query partially completed on the master but succeeds on the slave, the slave stops and complains. The slave updates the master.info file to reflect the new offset at which it is reading the master’s binary log. ([Location 3601](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=3601))
    - **Note:** Replication fails if the query cannot run on a slave but can run on a master
- These two threads divide the work in an effort to make sure the slave can always be as up to date as possible. The IO thread is concerned only with replicating queries from the master’s binary log. Rather than execute them, it records them into the slave’s relay log.[7] The SQL thread reads queries from the local relay log and executes them. ([Location 3621](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=3621))
    - **Note:** IO thread and SQL thread for a replication
- This solution isn’t foolproof. It’s possible for the IO thread to miss one or more queries if the master crashes before the thread has had a chance to read them. ([Location 3639](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=3639))
    - **Note:** The disadvantage of a replication
- There may be times when you don’t need to replicate everything from the master to the slave. In such situations you can use the various replication filtering options to control what is replicated. ([Location 3674](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=3674))
    - **Note:** Options not to replicate some tables
- The heartbeat principle is easy. At a fixed interval, say 20 seconds, a process on the master inserts a record with the latest timestamp into a table. On the slave, a corresponding process reads the most recent record every 20 seconds. Assuming that the system clocks on both machines are in sync, you can tell how far behind the slave is to within 20 seconds of accuracy. ([Location 3825](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=3825))
- As long as the software is working properly, it could be a hardware or driver problem. ([Location 4058](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=4058))
- Most web servers and web application servers accept all connections, process a request, respond, and then disconnect almost immediately.[1] They don’t perform any fancy authentication. In fact, most don’t even bother with a reverse lookup of an inbound IP address. In other words, the process of establishing the connection is very lightweight. The actual request and response process is typically lightweight too. ([Location 4291](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=4291))
    - **Note:** HTTP request and response is 
      * lightweight
      * Disconnect immediately
      But not MySQL
- MySQL’s network protocol doesn’t have a way to expose any hints to the load balancer. There are no URL parameters or cookies in which to store a session ID. A solution to this problem is to handle the partitioning of queries at the application level. You can split the 8 servers into 4 clusters of 2 servers each. Then you’d decide, in your application, whether a given query should go to cluster 1, 2, 3, or 4. You’ll see more of this shortly. ([Location 4321](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=4321))
    - **Note:** For load balancing, the only way to separate queries are partitioning tables.
- Load balancing works best when clients connect and disconnect frequently. That gives the load balancer the best chance of spreading the load evenly; otherwise the transparency is lost. ([Location 4332](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=4332))
    - **Note:** MySQL connection keeps and is reused across multiple connections
- You can also enforce this on the MySQL side by setting each server’s wait_timeout to a relatively low number. (This value tells MySQL how long a connection may remain idle before it is disconnected.) Doing so encourages sessions to be reestablished when needed, but the negative affects on the application side are minimal. Most MySQL APIs allow for automatic reconnection to the server any time you attempt to reuse a closed connection. If you make this change, consider also adjusting the thread_cache as described ([Location 4339](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=4339))
    - **Note:** Short connection time is the way to load balance MySQL connection easily
- A good health check also depends on your application needs and what’s most important. For example, on a nearly real-time dynamic web site like Yahoo! News, you might put more emphasis on replication. ([Location 4392](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=4392))
    - **Note:** Good health check isn't unique
- The problem was that several servers could be doing the health check at exactly the same time. If that happened, it was possible for all servers to believe that all other servers were healthy and proceed to declare themselves unhealthy. ([Location 4407](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=4407))
    - **Note:** The reason why load balancers didn't work as expected in Yahoo News
- After quite a bit of theorizing and poking around, someone thought to question the load-balancer configuration. It turned out that it was set on a least-connections scheduling algorithm. That clearly explained why a new machine was bombarded with new connections and rendered useless for several minutes. ([Location 4454](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=4454))
    - **Note:** At least connections sounded good, but not good in some real cases
- What is a disaster? For our purposes, a disaster is any event that causes significant portions of the data to be corrupted or unavailable. Some examples of disasters include the following: Hardware failure Software failure Accidental erasure of data[2] Stolen server Physically destroyed server ([Location 4628](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=4628))
    - **Note:** The types of a disater
- mytop does much of the hard work involved in summarizing MySQL performance data. There are three primary display modes in mytop. The default, thread view (or top view), closely resembles the Unix top command, as seen in Figure B-1. It produces a multiline summary at the top of the screen followed by a listing of threads in MySQL. The command view aggregates the data from MySQL’s Com_* command counters (see Appendix A), as seen in Figure B-2. Finally, Figure B-3 illustrates status view, which tracks all the other values in the output of SHOW STATUS. Like top, mytop refreshes the display periodically. The default refresh ([Location 6221](https://readwise.io/to_kindle?action=open&asin=B0026OR31W&location=6221))
    - **Note:** over view of mytop
