# FaceID Backend

Проект - бэкэнд для обработки времени прихода и ухода с работы по лицу.
Фото лица делается из мобильного приложения.
Над проектом работают:

[Степан Волков](https://github.com/stepan2volkov),
[Стас Кузнецов](https://github.com/smart48ru)

Подробности смотреть в [Wiki](https://github.com/smart48ru/FaceIDAppFaceIDBackend/wiki) проекта.

Конфигурационный файл ищется по пути `$./cfg/config.yaml`. Пример конфига:

```yaml
api:
  host: "0.0.0.0"
  port: 8080
  read_timeout: 40s
  write_timeout: 20s
  read_head_timeout: 50s
  secret_key: secret-admin-key
log_level: "debug"
debug: true
image:
  upload_dir: "./upload"
  model_dir: "./models"
```

Run application:
```shell
make run
```

For more details see:
```shell
make
# or
make help
```

[![codecov](https://codecov.io/gh/smart48ru/FaceIDApp/branch/main/graph/badge.svg?token=0QBDIQB1YN)](https://codecov.io/gh/smart48ru/FaceIDApp)
