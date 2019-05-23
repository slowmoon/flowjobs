#  CI/CD example with gitlab and drone 

###install the gitlab and drone  

#####first set environment variable HOSTADDR with your public network address or host address
```shell
export HOSTADDR=xx:xx:xx:xx
docker-compose -f drone-compose.yaml up -d 
```

###need to modify the gitlab property

```shell
external_url="your host url"

```
###### then login to the container gitlab-ce and restart the gitlab

```shell
docker exec -it gitlab-ce bash 
```

#### brower http://xxxxx:8090 addr to login into the gitlab


