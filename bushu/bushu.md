## 应用程序容器镜像制作
```dockerfile
FROM golang:1.18
ADD ./ /go/src/helloworld/
WORKDIR /go/src/helloworld
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:latest
ENV env1=env1value
ENV env2=env2value
MAINTAINER huanglingbo
WORKDIR /app/
COPY --from=0 /go/src/helloworld/app ./
EXPOSE 80
ENTRYPOINT ["./app"]
CMD ["--param1=p1","--param2=p2"]
```

```shell
docker build --no-cache -t hellloworld:1.0.0 .
docker tag helloworld:1.0.0 huanglingbo/hellloworld:1.0.0
docker push huanglingbo/hellloworld:1.0.0
```
## 定义k8s应用程序载体Pod
compose可以把多个容器绑定为一个组

但k8s中，Pod是最小的部署单元，一个Pod可以包含一个或多个容器

```shell
# 看看k8s支持的资源类型
kubectl api-resources
```

restartPolicy
    
    使用指数退避延迟重启（10s, 20s, 40s ...），最长延迟5分钟。
    Always表示容器退出后总是重启，适合长期运行的服务
    OnFailure表示容器以非零退出码退出(非正常退出)后重启，适合批处理任务
    Never表示容器退出后不重启，适合调试，测试环境

imagePullPolicy
    
    镜像拉取策略里镜像指的是镜像名+标签
    Always表示总是从镜像仓库拉取镜像，即使本地存在镜像，适合开发环境
    IfNotPresent表示如果本地存在镜像，则使用本地镜像，否则从镜像仓库拉取镜像，适合测试环境或者生产环境
    Never表示从镜像仓库拉取镜像，即使本地存在镜像，适合离线或者受控制的生产环境

k8s把cpu分成了1000份，每份称为1m，所以100m就是0.1核

```shell
vim myhellowoeld-pod.yaml
:set paste # 粘贴模式
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myhelloworld-pod-demo
  namespace: default
  labels:
    name: myhelloworld-pod-demo
    env: dev
spec:
  restartPolicy: Always
  containers:
    - name: myhelloworld
      image: huanglingbo/helloworld:1.0.0
      imagePullPolicy: IfNotPresent
      ports:
        - containerPort: 80
      command: ["./app"]
      args: ["--param1=k8s-p1","--param2=k8s-p2"]
      resources:
        limits:
          memory: "200Mi"
        requests:
          cpu: "100m"
          memory: "200Mi"
      env:
        - name: env1
          value: "k8s-env1"
        - name: env2
          value: "k8s-env2"
```

单独部署pod，不推荐，因为pod是最小的部署单元，没有副本，没有健康检查，没有负载均衡，没有自愈能力，所以使用Deployment

```shell
kubectl create -f myhelloworld-pod.yaml # 创建pod
kubectl get pods # 查看pod
```

健康检查：存活检查和就绪检查，存活检查用于检查容器是否存活，就绪检查用于检查容器是否准备好接收流量，如果存活检查失败，容器会被重启，如果就绪检查失败，容器不会接收流量

```yaml
# 存活检查
livenessProbe:
  httpGet:
    path: /healthz
    port: 80
  initialDelaySeconds: 3
  periodSeconds: 3
  
# 就绪检查
readinessProbe:
  httpGet:
    path: /ready
    port: 80
  initialDelaySeconds: 3
  periodSeconds: 3
```

```shell
kubectl apply -f myhelloworld-pod.yaml # 更新pod
kubectl logs -f myhelloworld-pod-demo # 查看pod日志
kubectl describe pod myhelloworld-pod-demo # 查看pod详情
kubectl delete pod myhelloworld-pod-demo # 删除pod
kubectl get pods # 查看pod
```

## 定义k8s应用程序载体Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myhelloworld-deployment
  labels:
    name: myapp
spec:
  replicas: 5
  selector:
    matchLabels:
      name: myapp
  template:
    metadata:
      labels:
        name: myapp
    spec:
      containers:
      - name: myhelloworld
        image: huanglingbo/helloworld:1.0.0
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 80
        livenessProbe:
          httpGet:
            path: /healthz
            port: 80
          initialDelaySeconds: 3
          periodSeconds: 3
        readinessProbe:
          httpGet:
            path: /ready
            port: 80
          initialDelaySeconds: 3
          periodSeconds: 3
        env:
          - name: env1
            value: "k8s-env1"
          - name: env2
            value: "k8s-env2"
          - name: env3
            value: "k8s-env3"
      - name: myredis
        image: redis
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 6379
        env:
          - name: env1
            value: "k8s-env1"
          - name: env2
            value: "k8s-env2"
          - name: env3
            value: "k8s-env3"
```

## service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: myapp-svc
spec:
  type: NodePort
  selector:
    name: myapp
  ports:
    - protocol: TCP
      name: http
      port: 8080
      targetPort: 80
      nodePort: 30001
    - protocol: TCP
      name: https
      port: 443
      targetPort: 80
    - protocol: TCP
      name: redis-tcp
      port: 6379
      targetPort: 6379
```

从不同级别去访问服务，要用不同的地址和端口

```shell
kubectl describe -f myhelloworld-deployment.yaml # 查看deployment详情
```