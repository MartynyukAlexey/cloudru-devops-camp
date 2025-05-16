## Task 02-ansible

Для этого задания я использую Virtual Box.

Создаю виртуальную машину

![creating VM](../.assets/02-ansible/vm_running.png)

Для доступа к ней с хоста добавляю bridged-adapter

![creating VM](../.assets/02-ansible/bridged_adapter.png)

Настраиваю сетевой интерфейс

```
sudo ip link set enp0s8 up
sudo dhclient enp0s8
ip a
```

Копирую ssh ключ с хоста и проверяю ssh соединение

![setting up network interface and ssh connectivity](../.assets/02-ansible/ssh_connection.png)

Проверяю, что ansible видит виртуалку

![test ansible connectivity](../.assets/02-ansible/ansible_test_conn.png)

Выполняю сценарий для запуска докера:
![installing docker with ansible](../.assets/02-ansible/install_docker.png)