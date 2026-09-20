# CentOS High Availability

![](https://m.media-amazon.com/images/I/81KCmpIR5SL._SY160.jpg)

### Metadata

- Author: Mitja Resman
- Full Title: CentOS High Availability
- Category: #books

### Highlights

- It is good practice for the cluster members to have the same hardware and specifications, which means that the cluster computers consist of components from the same manufacturer and likely have the same resource specifications. ([Location 305](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=305))
- It became really popular very quickly, mostly due to the fact that it significantly lowers the cost of server maintenance, boosts the server life cycle, and facilitates management of the servers. ([Location 365](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=365))
- A special protocol called Totem Single-ring Ordering and Membership (TOTEM) was developed to provide cluster group communication. TOTEM provides cluster messaging and membership services. The cluster messaging service is the communication process among cluster members. This communication process is called a heartbeat and the formation cluster members adopt is called a ring. ([Location 433](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=433))
    - **Note:** Protocol to communicate inside a cluster
- This communication process is called a heartbeat and the formation cluster members adopt is called a ring. ([Location 436](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=436))
- Quorum is the minimum number of cluster member votes required to perform a cluster operation. ([Location 466](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=466))
- This decision is not up an individual cluster member; it has to be voted and agreed upon. ([Location 469](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=469))
- The quorum disk is very useful and also required in an even-node cluster configuration. ([Location 473](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=473))
- The quorum disk is a shared storage disk to which all the cluster members have access. ([Location 476](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=476))
- Split brain is a condition in the cluster where the cluster is split into two sides, each side thinking that the other side is dead or inactive. In this situation, each side proceeds to take over the resources as if the other side no longer exists. ([Location 480](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=480))
- Before any of the split sides of the cluster takes over the resources, the fencing mechanism fences the other side by issuing a configured action of reboot or power-off on the cluster nodes, thus making sure that the cluster nodes on the other side are really dead. ([Location 482](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=482))
- The cluster resource manager provides high availability of cluster services by detecting and recovering cluster service failure from the cluster nodes and cluster resource failures. ([Location 498](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=498))
- There are two types of fencing: resource-level fencing and node-level fencing. Resource-level fencing is the way the cluster disables a specific or problematic node to access a specific resource. This is usually used on shared storage resources, where a problematic node is denied access to certain shared storage. More commonly used is node-level fencing. Node-level fencing is a way to make sure a node does not run any resources at all. ([Location 507](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=507))
- Node fencing is also called STONITH, which stands for Shoot The Other Node In The Head. ([Location 511](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=511))
- The Corosync cluster engine is an open source project that was derived from the OpenAIS project and was announced in July 2008. ([Location 556](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=556))
- RGManager development is known to be slower and therefore less mutable, which can also be a good thing when it comes to the stability of the software. ([Location 564](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=564))
- The cluster nodes must be configured to communicate via an IP address and also the Fully Qualified Domain Name (FQDN) domain names, which requires proper configuration of DNS servers. It is also highly recommended to configure NTP date and time synchronization for all the cluster nodes. ([Location 583](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=583))
- It is known to be good practice that the DNS and NTP servers are in the same local network area as the cluster nodes, thus reducing the possibility of network failure causing problems in the DNS resolution or NTP synchronization. ([Location 586](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=586))
- Change the NM_CONTROLLED line to no. This specifies that the Network Manager service cannot control the network interface. ([Location 632](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=632))
- Change the BOOTPROTO line to none. This specifies that no boot protocol is used for this interface, since this interface has a static IP address assigned to it. ([Location 634](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=634))
- You must also configure FQDN for the cluster nodes by editing the /etc/sysconfig/network file and changing the HOSTNAME line to the FQDN of your cluster node. ([Location 659](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=659))
- You must add an iptables rule to allow UDP traffic on 5404 and 5405, and another rule to allow multicast traffic communication. ([Location 678](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=678))
- You should generate the Corosync encryption keys with the corosync-keygen command to increase cluster security and encrypt cluster communication traffic. ([Location 690](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=690))
- Pacemaker traffic on TCP port 2224, and reload the firewall daemon. ([Location 1574](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=1574))
- You must install the pacemaker software and the pcs shell on all cluster nodes. ([Location 1584](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=1584))
- There are six resource classes supported by Pacemaker, ([Location 1694](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=1694))
- The difference between the two is that the Ipaddr parameter uses the ifconfig command to create the interface and the IPaddr2 parameter uses the ip command to create the interface. ([Location 1714](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=1714))
- pcs resource describe ocf:heartbeat:IPaddr2 ([Location 1720](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=1720))
- Colocation is a rule binding two cluster resources to run on the same location. Usually more than one resource defines a cluster service and these resources must be running on the same cluster node. ([Location 1756](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=1756))
- To provide a working Apache web server cluster service, the ClusterIP resource must be started before the WebServer resource. We can achieve this by configuring a cluster resource order constraint with the following command: pcs constraint order set ClusterIP WebServer ([Location 1783](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=1783))
- Adding a new cluster node to the existing cluster configuration does not require any cluster service downtime. ([Location 1921](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=1921))
- Removing a cluster node from a cluster configuration does not require cluster service downtime, unless the cluster service is running on the cluster node you would like to remove. ([Location 1987](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=1987))
- The node fencing cluster feature is provided by fencing agents. Fencing agents are scripts that get executed when the cluster situation makes calls for them. ([Location 2027](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=2027))
- Quorum is the minimum number of cluster member votes required to perform a cluster operation. ([Location 2222](https://readwise.io/to_kindle?action=open&asin=B00WX1CWXC&location=2222))
