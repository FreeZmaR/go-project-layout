# Internal Layer

- [Description](#description)
- [Diagram](#diagram)
- [App SubLayer](#app-sublayer)
  - [Details](app/readme.md)
- [Domain SubLayer](#domain-sublayer)
  - [Details](domain/readme.md)
- [Use-Case SubLayer](#use-case-sublayer)
  - [Details](usecase/readme.md)
- [Repository SubLayer](#repository-sublayer)
  - [Details](repository/readme.md)
- [Lib SubLayer](#lib-sublayer)
  - [Details](lib/readme.md)
- [Storage SubLayer](#storage-sublayer)
  - [Details](storage/readme.md)

## Description
Этот уровень является основную *начинку* приложения, которая содержит бизнес-логику, модели, хранилища и т.д.
К донному слою имеют доступ только слой `CMD`.
Слой имеет подуровни, которые описаны ниже.

В своем пример я хотел показать, что процесс взаимодействия между слоями этого уровня по следующему принципу:
- `app` слой обращается к `usecase` процессам
- `usecase` процессы взаимодействют с `repository` слоем данных
- `repository` слой данных взаимодействует с `storage` слоем хранения данных
- слой `domain` является общим для всех и не имеет доступа к другим слоям, кроме слоя `lib`
- слой `lib` является так же как и `domain` общим для всех и не имеет доступа к другим слоям

---

This layer is the main *core* of the application which contains business logic, models, repositories, etc.
Only the `CMD` layer has access to this layer.
The layer has sublayers described below.

In my example, I wanted to show that the process of interaction between the layers of this layer according to the following principle:
- `app` layer refers to `usecase` processes
- `usecase` processes interact with the `repository` data layer
- `repository` data layer interacts with the `storage` data storage layer
- the `domain` layer is common to all and does not have access to other layers, except for the `lib` layer
- the `lib` layer is also common to all and does not have access to other layers


## Diagram

```mermaid
stateDiagram-v2
    state Main { 
        [*] --> app
        app --> usecase
        usecase --> repository
        repository --> storage
    }
    Main --> domain
    Main --> lib
    domain --> lib

```


## App SubLayer

Этот подуровень содержит описания подпрограмм, основное приложение - это набор поддпрограмм.
В моем примере подпрограммой является `http` сервер, который имеет разные модули. 
Например, модулю `inbox` и `outbox`.

Название директории для данного слоя - `app`.

Более подробное описание подуровня в директории [app](app/readme.md).

---

This sublayer contains descriptions of subprograms, the main application is a set of subprograms.
In my example, a subprogram is an `http` server that has different modules.
For example, the `inbox` and `outbox` modules.

The directory name for this layer is `app`.

More detailed description of the sublayer in the [app](app/readme.md) directory.

## Domain SubLayer

Этот подуровень содержит описания моделей доменной области приложения. В этом слое не должно быть никакой бизнес-логики,
только описания предметной части выраженной в коде. Например, модель `user` или `transaction`. 
Все что связанно с объектами принимающие участие в бизнес-процессе.
Я предпочитаюс максимольно задокументировать данную часть приложения, чтобы любой другой человек мог понять все детали.

Данный слой не имеет доступа к другим слоям приложения, кроме слоя `lib`.

Название директории для данного слоя - `domain`.

Более подробное описание подуровня в директории [domain](domain/readme.md).

---

This sublayer contains descriptions of domain models of the application. This layer should not have any business logic,
only descriptions of the subject area expressed in code. For example, the `user` or `transaction` model.
Everything related to objects participating in the business process.
I prefer to document this part of the application as much as possible so that any other person can understand all details.

This layer does not have access to other layers of the application, except for the `lib` layer.

The directory name for this layer is `domain`.

More detailed description of the sublayer in the [domain](domain/readme.md) directory.

## Use-Case SubLayer

Этот подуровень содержит описания бизнес-процессов приложения, сервисный слой. Как описать и выразить бизнес-процесс решать вам. 
В моем примере я использую интерфейсы, которые описывают бизнес-процессы.

Название директории для данного слоя - `usecase`.

Более подробное описание подуровня в директории [usecase](usecase/readme.md).

---

This sublayer contains descriptions of the business processes of the application, the service layer. How to describe and express the business process is up to you.
In my example, I use interfaces that describe business processes.

The directory name for this layer is `usecase`.

More detailed description of the sublayer in the [usecase](usecase/readme.md) directory.

## Repository SubLayer

Этот подуровень является сервисным слоям отвечающий за данные.
В моем примере я использую репозитории для отдельных `usecase` бизнес-процессов, так же я использую репозитории для отдельной модели если это потребуется.
Например, репозиторий для модели `user` или `transaction`, а для бизнес-процессов репозиторий `inbox` и `outbox`.
Так же как и в слое `domain` я предпочитаю детально документировать интейрфесы, которые описывают бизнес-процессы.

Название директории для данного слоя - `repository`.

Более подробное описание подуровня в директории [repository](repository/readme.md).

---

This sublayer is a service layer responsible for data.
In my example, I use repositories for individual `usecase` business processes, I also use repositories for a separate model if necessary.
For example, a repository for the `user` or `transaction` model, and for business processes the `inbox` and `outbox` repository.
As in the `domain` layer, I prefer to document in detail the interfaces that describe business processes.

The directory name for this layer is `repository`.

More detailed description of the sublayer in the [repository](repository/readme.md) directory.

## Lib SubLayer

Этот подуровень является слоем библиотек, которые используются в приложении. 
Данный слой может иметь доступ только к слои `config`

Название директории для данного слоя - `lib`.

Более подробное описание подуровня в директории [lib](lib/readme.md).

---

This sublayer is a layer of libraries used in the application.
This layer can only have access to the `config` layer

The directory name for this layer is `lib`.

More detailed description of the sublayer in the [lib](lib/readme.md) directory.

## Storage SubLayer

Этот подуровень является слоем хранилища данных. В нем описываются все хранилища используемые в приложении.

Название директории для данного слоя - `storage`.

Более подробное описание подуровня в директории [storage](storage/readme.md).

---

This sublayer is a data storage layer. It describes all the storages used in the application.

The directory name for this layer is `storage`.

More detailed description of the sublayer in the [storage](storage/readme.md) directory.