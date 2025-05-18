## 03-kubernetes

Для тестирования манифестов я использую minikube.

### Raw manifests

Создаю кластер и разворачиваю приложение:

```
minikube start --cpus=4 --memory=8192

minikube addons enable ingress

kubectl apply -f ./namespace.yaml

kubectl create secret docker-registry dockerhub-secret \
  --namespaces-echo-server-ns \
  --docker-server=docker.io \
  --docker-username=martynyukalexey \
  --docker-password=<password> \
  --docker-email=MartynyukAlexey05@gmail.com

kubectl apply -f ./deployment.yaml
kubectl apply -f ./service.yaml
kubectl apply -f ./ingress.yaml

minikube ip
```

![spin up minikube cluster and apply raw manifests](../.assets/03-kubernetes/raw-manifests/applying-manifests.png)

Проверяю, что все работает:

![list of pods](../.assets/03-kubernetes/raw-manifests/pods-list.png)

С помощью ```minikube ip``` определяю адрес кластера. 
Добавляю в /etc/hosts соответствующий маппинг:

![list of pods](../.assets/03-kubernetes/raw-manifests/etc-hosts.png)

Открываю приложение в браузере:

![list of pods](../.assets/03-kubernetes/raw-manifests/page-in-browser.png)

### Helm chart

Создаю кластер и разворачиваю приложение:

```
minikube start --cpus=4 --memory=8192

minikube addons enable ingress

helm install echo-server . \
  --set image.credentials.username=martynyukalexey \
  --set image.credentials.password=*password* \
  --set image.credentials.email=MartynyukAlexey05@gmail.com
```

![spin up minikube cluster and apply helm chart](../.assets/03-kubernetes/helm-chart/applying-chart.png)

Проверяю, что все работает:

![list of pods](../.assets/03-kubernetes/helm-chart/pods-list.png)

Запись в /etc/hosts осталась после работы с сырыми манифестами (cluster ip не изменился).

Открываю приложение в браузере:

![list of pods](../.assets/03-kubernetes/helm-chart/page-in-browser.png)
