# This is the minio docker chart for the TNG project. 

> Tip: To find Api Key.
> 1) To get/generate your apikey, navigate to https://tbctdevops.jfrog.io/tbctdevops/webapp/#/profile.
> 2) If not generated, click on the generate icon
> 3) Click the copy to keyboard icon.

## Versions

Alpine : `3.10`   
Minio : `latest`

## Commands

**Pull :** `docker pull tbctdevops-docker-local-poc.jfrog.io/minio/minio:2.1.0-fea-066818583-7259fae`   
**Run :** `docker run -p 9000:9000 tbctdevops-docker-local-poc.jfrog.io/minio/minio:2.1.0-fea-066818583-7259fae server /data`

## Explore Further

* [Distributed MinIO Quickstart Guide](https://docs.min.io/docs/distributed-minio-quickstart-guide)
