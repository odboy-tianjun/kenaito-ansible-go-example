# go-ansible库应用

## thank you for go-ansible authors

https://github.com/apenella/go-ansible

## 执行结果参考

#### 执行成功

```json
{
  "Command": "ansible all  --args 'ping 127.0.0.1 -c 2' --inventory 192.168.0.1,192.168.0.2 --module-name command --connection local",
  "Success": true,
  "Data": {
    "192.168.0.1": {
      "changed": 1,
      "failures": 0,
      "ignored": 0,
      "ok": 1,
      "rescued": 0,
      "skipped": 0,
      "unreachable": 0
    },
    "192.168.0.2": {
      "changed": 1,
      "failures": 0,
      "ignored": 0,
      "ok": 1,
      "rescued": 0,
      "skipped": 0,
      "unreachable": 0
    }
  },
  "ErrMsg": "",
  "StartTime": "2025-07-03 21:17:26",
  "EndTime": "2025-07-03 21:17:28",
  "Duration": 1566
}
```

#### 执行失败

```json
{
  "Command": "ansible all  --args 'pingxxxx 127.0.0.1 -c 2' --inventory 192.168.0.1,192.168.0.2 --module-name command --connection local",
  "Success": false,
  "Data": null,
  "ErrMsg": "Error during command execution.\n Command executed: ansible all  --args 'pingxxxx 127.0.0.1 -c 2' --inventory 192.168.0.1,192.168.0.2 --module-name command --connection local\n\n Environment variables:\nANSIBLE_LOAD_CALLBACK_PLUGINS=1\nANSIBLE_STDOUT_CALLBACK=json\n\n\texit status 2",
  "StartTime": "2025-07-03 21:20:07",
  "EndTime": "2025-07-03 21:20:08",
  "Duration": 497
}
```