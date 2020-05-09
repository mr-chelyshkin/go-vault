# go-vault

API для [Vault](https://www.vaultproject.io/)

## Описание
Go-vault - пакет для простого использования [Vault](https://www.vaultproject.io/).
Возможности пакета:
* Получить токен
* Обновить токен
* Забрать секрет из Vault

Основные методы клиента Vault:
* NewBaseClient() - создание объекта клиента Vault с минимальными набором опций.
* NewCustomClient() - создание объекта клиента Vault с кофигурацией  vaultApi.
* Get() - забирает данные из Vault.

### Установка
```bash
go get gitlab.corp.mail.ru/go/internal_dev/go-vault
```

### Использование
```go
import vault "gitlab.corp.mail.ru/go/internal_dev/go-vault"

client,  err := vault.NewBaseClient("roleId", "secretId", nil)
secrets, err := client.Get("vault/url")
```

```go
import vault "gitlab.corp.mail.ru/go/internal_dev/go-vault"

clientOpt := &vault.ClientOptions{
    TokenFilePath: "/.token"
}

client,  err := vault.NewBaseClient("roleId", "secretId", clientOpt)
secrets, err := client.Get("vault/url")
```

```go
import vault "gitlab.corp.mail.ru/go/internal_dev/go-vault"

clientOpt := &vault.ClientOptions{
    TokenFilePath: "/.token"
}
clientApi := &vault.ClientApi{
    Host: "https://mail.ru"
}

client,  err := vault.NewBaseClient("roleId", "secretId", clientOpt, clientApi)
secrets, err := client.Get("vault/url")
```

### Настройки
```go
ClientOptions{
    TokenFilePath string // путь к токен файлу
    CertFilePath  string // путь к файлу с сертификатом
}

ApiOptions{
    Host string  // vault хост
    Port string  // порт
    
    Version    string // версия api
    AuthLink   string // ссылка для авторизации
    UpdateLink string // ссылка для обновления токена
    LookupLink string // ссылка для получения информации о токена
}
```

## Тестирование
Проект содержит unit тесты. Запустить можно так:
```bash
go test .
```

