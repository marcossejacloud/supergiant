# Capacity Service

Compute resources for Kubernetes clusters can come in several different flavors: cloud-based virtual machines, cloud-based bare-metal servers, on-premise virtual machines, or on-premise bare-metal servers. In addition, each cloud service provider (e.g., AWS, GCE, etc.) offers its users tens of server types with different cpu, memory, networking, architecture, and more. Managing the selection, instantiation, and scaling of all these diverse server types manually can be a complicated task.

Supergiant's **Capacity Service** is a solution to this problem. The Capacity Service is Supergiant's way of abstracting servers from application management and deployment, allowing the user to focus only on how much CPU, RAM, and disc applications need -- not which flavor of a server to use or how to set it up for each application. If _containerization_ was similar to hardware-level virtualization (hypervisors partitioning big host machines in the cloud into multiple virtual machines), then the following analogy could be made. Without the Supergiant Capacity Service, running a container orchestration platform like Kubernetes forces DevOps engineers to manage two levels of capacity. That is, they must worry not only about container resource allocation but also about _server allocation_. It would be equally difficult to imagine AWS requiring its users not only to request new servers but also to request additional server capacity in each region and availability zone ahead of time!

Of course, that would be silly, because, after all, cloud servers are generally _on-demand_. That's what the Capacity Service does for Kubernetes pods. Pods can be provisioned on-demand without worrying about server capacity. Supergiant will handle creating nodes when pods are requesting more than the available capacity, and (gently) deleting nodes when sufficiently under capacity.