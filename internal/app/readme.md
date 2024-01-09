# App layer

Уровень для описания подпрограмм приложения. 
Каждая поддпрограмма имеет свою директорию внутри данного слоя и являеться отдельным модулем.
В моем примере подпрограммой является `http` сервер, который имеет разные модули для работы `inbox` и `outbox` и их версионность.

Я люблю использовать подпрограммы ввиде интерфейса `app`, так как они позволяют разделить приложение на независимые части, которые могут быть перенесены или переиспользованны на разные проекты.

```go
type App interface {
   Run(ctx context.Context) error
   Stop(ctx context.Context) error
}
```

Поддпрограммы могут быть разных типов, например, `http` сервер, `grpc` сервер, `cron` задача, `cli` приложение и т.д.
Как их реализовывать остаеться за разработчиком.

В моем примере модули `inbox` и `outbox` для реализации использую `uber.fx`, где в файле `module.go` описываю версии модуля подпрограммы, а в файле `module_components.go` описываю компоненты для модуля.

Компоненты для модуля называю по принципу как они описанны в `uber.fx`, например: `ProvidePostgres`, `InvokeAppLifeCycle`.

--- 

Level for describing application subprograms.
Each subprogram has its own directory inside this layer and is a separate module.
In my example, a subprogram is an `http` server, which has different modules for `inbox` and `outbox` work and their versioning.

I like using subprograms as an `app` interface, as they allow you to divide the application into independent parts that can be transferred or reused for different projects.

```go
type App interface {
   Run(ctx context.Context) error
   Stop(ctx context.Context) error
}
```

Subprograms can be of different types, for example, `http` server, `grpc` server, `cron` task, `cli` application, etc.
How to implement them remains with the developer.

In my example, the `inbox` and `outbox` modules for implementation use `uber.fx`, where in the `module.go` file I describe the versions of the subprogram module, and in the `module_components.go` file I describe the components for the module.

I name the components for the module according to the principle as they are described in `uber.fx`, for example: `ProvidePostgres`, `InvokeAppLifeCycle`.
