## 02-ansible

Для этого задания я использую Virtual Box.

Выбраный алгоритм балансировки: round_robin (дефолтный алгоритм).
Почему: 

1. ip hash (ip_hash) не не требуется в stateless приложении (используется для поддержания длительных сессий)
2. weighted (server ... weight=3) не подходит, так как инстансы приложения одинаковы (подошел бы для балансировки между несколькими серверами с разными параметрами CPU/RAM)
3. least connection (least_conn) не подходит, так как все запросы к приложению выполняются "быстро" (подошел бы если бы некоторые запросы могли бы выполняться значительно дольше остальных)

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

P.S. Мы использовали anible в университете для конфигурации микротиков (виртуалки с ОС микротика объединялись в VPN сеть и конфигурировались плейбуками). При необходимости могу открыть соответствующие репозитории