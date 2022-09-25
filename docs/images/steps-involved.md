## Run PRS in k8s as a stand alone pod
- First launching a k8s cluster with minikube
    - `minikube start --vm-driver=docker`  
- Running PRS in the cluster as a stand alone pod
    - `kubectl run product-review --image=shashwat623/product-reviews:1.0 --port=8080`
- accessing the APIs by port-forwarding to the running pod
    - `kubectl port-forward product-review 8080:8080`

## Deploy a Replica-Set of PRS using manifest file.
- scale the PRS to 2 replicas
    -   ```
        apiVersion: apps/v1
        kind: ReplicaSet
        metadata:
        name: product-reviews
        labels:
            app: product-reviews
            tier: product-reviews
        spec:
        # modify replicas according to your case
        replicas: 2
        selector:
            matchLabels:
            tier: product-reviews
        template:
            metadata:
            labels:
                tier: product-reviews
            spec:
            containers:
            - name: product-reviews-container
                image: shashwat623/product-reviews:1.0
        ```

- Add a cluster-ip service and access the PRS by port-forwarding on the svc object.
    -  ```
        apiVersion: v1
        kind: Service
        metadata:
        name: product-review-service
        spec:
        selector:
            tier: product-reviews
        ports:
        - name: product-review-service-port
            protocol: TCP
            port: 80
            targetPort: 8080

        ```
    - Command to port forward the service `kubectl port-forward service/product-review-service 8080:80`

- Stern on the logs of the pods and ensure requests are getting routed to different pods.
    * Tailing the pods filtered by tier=product-reviews label selector across all namespaces `stern --all-namespaces -l tier=product-reviews`

## Use Deployment and run both the services using manifest files.

- ``` 
    apiVersion: apps/v1
    kind: Deployment
    metadata:
    name: product-details-deployment
    labels:
        app: product-details
    spec:
    replicas: 1
    selector:
        matchLabels:
        app: product-details
    template:
        metadata:
        labels:
            app: product-details
        spec:
        containers:
        - name: product-details-container
            image: shashwat623/product-details:1.0
            ports:
            - containerPort: 9090
            env:
            - name: REVIEW_SVC_HOST
            value: http://product-review-service-wrong:8080
    ---
    apiVersion: v1
    kind: Service
    metadata:
    name: product-details-service
    spec:
    selector:
        app: product-details
    type: LoadBalancer    
    ports:
        - protocol: TCP
        port: 8081
        targetPort: 9090
        nodePort: 30000
    ```

- ```
    apiVersion: apps/v1
    kind: Deployment
    metadata:
    name: product-review-deployment
    labels:
        app: product-review
    spec:
    replicas: 1
    selector:
        matchLabels:
        app: product-review
    template:
        metadata:
        labels:
            app: product-review
        spec:
        containers:
        - name: product-review-container
            image: shashwat623/product-reviews:1.0
            ports:
            - containerPort: 8080
    ---
    apiVersion: v1
    kind: Service
    metadata:
    name: product-review-service
    spec:
    selector:
        app: product-review
    ports:
        - protocol: TCP
        port: 8080
        targetPort: 8080
    ```

- Rollout restart the deployment after setting wrong value for env variable.

- Now Rollback to previous version and ensure things are working as expected
    - ` kubectl rollout undo deployment/product-details-deployment --to-revision=1`

    
            


        
