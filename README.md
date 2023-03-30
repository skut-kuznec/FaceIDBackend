# FaceID Backend

Проект - бэкэнд для обработки времени прихода и ухода с работы по лицу.
Фото лица делается из мобильного приложения.
Над проектом работают:

[Степан Волков](https://github.com/stepan2volkov),
[Стас Кузнецов](https://github.com/smart48ru)

Подробности смотреть в [Wiki](https://github.com/smart48ru/FaceIDBackend/wiki) проекта.

Конфигурационный файл ищется по пути `$./cfg/config.yaml`. Пример конфига:

```yaml
---
r:
  host: "127.0.0.1"
  port: 8080
  program_key: "ProgAuthKey"
  admin_key: "AdminAuthKey"
  read_time_out: 40
  write_time_out: 20
  read_head_time_out: 50
db:
  enable: false
  host: "127.0.0.1"
  port: 5432
  user: "devUser"
  password: "devUserPassword"
  dbname: "faceid"
  ssl_mode: false
log:
  level: "debug"
  to_file: false
  filename: "log/epg.log"
files:
  upload_dir: "upl"
```

Run application:
```shell
make run
```

Make Swagger documentation 
```shell
make swag
```

For more details see:

```shell
make
# or
make help
```

[![codecov](https://codecov.io/gh/smart48ru/FaceIDBackend/branch/main/graph/badge.svg?token=0QBDIQB1YN)](https://codecov.io/gh/smart48ru/FaceIDBackend)