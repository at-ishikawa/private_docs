# Malicious Kubernetes Helm Charts Can Be Used to Steal Sensitive Information From Argo CD Deployments

![](https://readwise-assets.s3.amazonaws.com/static/images/article1.be68295a7e40.png)

### Metadata

- Author: apiiro.com
- Full Title: Malicious Kubernetes Helm Charts Can Be Used to Steal Sensitive Information From Argo CD Deployments
- Category: #articles
- URL: https://apiiro.com/blog/malicious-kubernetes-helm-charts-can-be-used-to-steal-sensitive-information-from-argo-cd-deployments/

### Highlights

- TL;DR
  Argo CD is a popular, open-source, Continuous Delivery (CD) platform that is used by thousands of organizations globally.
  A 0-day vulnerability, discovered by Apiiro’s Security Research team, allows malicious actors to load a Kubernetes Helm Chart YAML file to the vulnerability and “hop” from their application ecosystem to other applications’ data outside of the user’s scope.
  The actors can read and exfiltrate secrets, tokens, and other sensitive information residing on other applications.
  The impact of the attack includes privilege escalation, sensitive information disclosure, lateral movement attacks, and more.
  Although Argo CD contributors were aware of this weak point in 2019 and implemented an anti-path-traversal mechanism, a bug in the control allows for exploitation of this vulnerability.
