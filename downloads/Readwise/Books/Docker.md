# Docker

![](https://images-na.ssl-images-amazon.com/images/I/51IBjhz7VpL._SL200_.jpg)

### Metadata

- Author: Karl Matthias, Sean P. Kane
- Full Title: Docker
- Category: #books

### Highlights

- These tools generally provide the simplest way to get into production with Docker. In this category are tools like: New Relic’s Centurion Spotify’s Helios Ansible’s Docker tooling ([Location 2472](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2472))
- We have a post-commit hook that triggers a build on each commit, so that job is kicked off on the build server. ([Location 2556](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2556))
- Because Docker’s link mechanism is limited to working on a single host, Compose is best for things like development and testing rather than production. ([Location 2605](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2605))
- using the docker top command, you can see the process list as your container understands it. ([Location 2617](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2617))
- We’ll chop the output to just the part we care about: ([Location 2669](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2669))
- Whether it be systemd, upstart, runit, supervisor, or your own homegrown tools, this allows you to treat containers atomically even when they contain more than one process. ([Location 2746](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2746))
- How can you tell what images are actually underneath the one your container is running on? docker history does just that. ([Location 2798](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2798))
- But underneath that is a directory on the host’s disk that is dedicated to the container. Usually this is in /var/lib/docker/containers. ([Location 2821](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2821))
- One of Docker’s major strengths is its ability to abstract away the underlying hardware and operating system so that your application is not constrained to any particular host or environment. ([Location 2873](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2873))
- If you have your own private cloud, you can leverage a tool like Docker Swarm to deploy containers easily across a large pool of Docker hosts, or use the community tool Centurion or Helios to quickly facilitate multi-host deployments. ([Location 2880](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=2880))
- But that service assigns only a single container to an Amazon instance, which means that it’s not ideal for short-lived or lightweight containers. ([Location 3095](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=3095))
- The first thing we need to do is start a cluster in the container service. We’ll then push our containers into the cluster once it’s up and running. ([Location 3149](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=3149))
- As we mentioned, it is also possible to take an existing Docker host within your EC2 environment and make it compatible with the EC2 Container Service. To do this, you need to connect to the EC2 instance and ensure that you are running Docker version 1.3.3 or greater, and then deploy the Amazon ECS Container Agent to the local Docker host with the proper environment variable configured for your setup, as shown here: ([Location 3158](https://readwise.io/to_kindle?action=open&asin=B00ZGRS4XM&location=3158))
