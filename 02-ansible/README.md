## 02-ansible

Для этого задания я использую Virtual Box.

Создаю виртуальную машину:

![VM overview](../.assets/02-ansible/vm-overview.png)

Для доступа к ней с хоста добавляю host-only адаптер (в дополнение к NAT адаптеру по умолчанию; виртуальная машина получает доступ в интернет через NAT и доступ к хосту через host-only адаптеры):
![adding host-only adapter](../.assets/02-ansible/host-only-adapter.png)

Настраиваю сетевой интерфейс:

```
sudo ip link set enp0s8 up
sudo ip addr add 192.168.56.10/24 dev enp0s8
```

![configuring network](../.assets/02-ansible/configuring-network.png)

Копирую ssh ключ с хоста и проверяю ssh соединение:

![ssh connectivity](../.assets/02-ansible/ssh-access.png)

Выполняю сценарии:
1. configure_docker.yml

![running docker playbook](../.assets/02-ansible/run-docker-playbook.png)

2. configure_application.yml

![running application playbook](../.assets/02-ansible/run-application-playbook.png)

3. configure_nginx.yml

![running nginx playbook](../.assets/02-ansible/run-nginx-playbook.png)

Проверяю контейнеры:

![containers list](../.assets/02-ansible/containers-list.png)

Приложение в браузере:

![containers list](../.assets/02-ansible/page-in-browser.png)