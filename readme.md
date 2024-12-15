# **Web Scraping application demo**

## 🔍 Overview

This project consists of two microservices:

1. **Scraper Service (Golang)**: Fetches HTML content from the `books.toscrape.com` website and sends the raw HTML to the Parser Service.

2. **Parser Service (Golang)**: Receives HTML content and extracts product information such as Name, Availability, UPC, Price (excl. tax), and Tax, storing the data in a SQLite database.

Both services are deployed in Kubernetes and communicate via gRPC.

---

## ⚙️ **Architecture**

- **Scraper Service**:
  - Scrapes paginated product lists from `https://books.toscrape.com`.
  - Sends HTML content to the Parser Service for data extraction.
  - Saves product data in a SQLite database.

- **Parser Service**:
  - Exposes a gRPC endpoint to receive HTML content.
  - Extracts product details from the raw HTML received from the Scraper Service.

### Deployment Diagram

``` plaintext
[Scraper Service] --> [Parser Service] --> [SQLite Database]
```

---

## 📄 **Deployment Files**

Deployment files are located in the `k8s` directory:

- **`scraper-deployment.yaml`**: Deploys the Scraper Service.
- **`parser-deployment.yaml`**: Deploys the Parser Service.
- **`parser-service.yaml`**: Exposes the Parser Service to other services.
- **`sqlite-pv.yaml`**: Persistent volume for storing the SQLite database.

---

## 🚀 **Deployment Instructions**

Helper scripts are provided in the `scripts` directory to simplify the deployment process:

1. **Ensure Kubernetes is Running Locally**:

    ``` powershell
    minikube start
    ```

2. **Build Docker Images**:

    ``` powershell
    ./scripts/build.ps1
    ```

3. **Deploy Services**:

    ``` powershell
    ./scripts/deploy.ps1
    ```

4. **Check Deployment Status**:

    ``` powershell
    kubectl get pods
    kubectl get pvc
    ```

5. **Access Logs**:

    ``` powershell
    kubectl logs <scraper-pod>
    kubectl logs <parser-pod>
    ```

---

## 🧪 **Testing the Services**

1. **Check Scraper Service Logs**:

    ``` powershell
    kubectl logs -l app=scraper
    ```

2. **Check Parser Service Logs**:

    ``` powershell
    kubectl logs -l app=parser
    ```

3. **Verify Database Contents**:

    Access the SQLite database stored in the persistent volume:

    ``` powershell
    kubectl exec -it <scraper-pod> -- sqlite3 /data/products.db
    ```

    Query the database:

    ``` sql
    SELECT * FROM products;
    ```

---

## 🛠️ **Possible Issues and Improvements**

### Issues

1. **Parser Service Timeout**

   - **Description**: The scraper service may experience timeouts when waiting for the parser service to respond, especially if the parsing operation takes a long time or the parser service is under heavy load.
   - **Solution**:
     - Increase the scraper’s request timeout configuration if needed.
     - Monitor the resource utilization of the parser service and scale it accordingly.

2. **HTML Page Structure Changes**

   - **Description**: The scraper service relies on a specific HTML structure to extract data. If the structure of the target website changes, the scraper may fail to extract data correctly or miss important information.
   - **Solution**:
     - Regularly monitor the target website for changes in its HTML structure.
     - Implement error handling to detect when parsing fails or yields unexpected results.
     - Use flexible parsing methods that can adapt to minor changes in the HTML structure.

3. **Resource Limits and Out-of-Memory (OOM) Errors**

   - **Description**: The scraper or parser service may be terminated due to exceeding memory or CPU limits.
   - **Solution**:
     - Set appropriate resource requests and limits in the deployment manifests.
     - Monitor the resource usage of the services using Kubernetes tools like `kubectl top pods` or Prometheus.

### Improvements

1. **Logging**:
   - Add structured logging and log levels (INFO, ERROR) to help with debugging.

2. **Database**:
   - Migrate from SQLite to a more scalable database like PostgreSQL for production use.

3. **Security**:
   - Implement TLS verification properly and avoid bypassing it in production environments.
