# Internal Layer

Этот уровень является основную *начинку* приложения, которая содержит бизнес-логику, модели, хранилища и т.д.
К донному слою имеют доступ только слой `CMD`.
Слой имеет подуровни, которые описаны ниже.

---

This layer is the main *core* of the application which contains business logic, models, repositories, etc.
Only the `CMD` layer has access to this layer.
The layer has sublayers described below.

## App SubLayer

Этот подуровень содержит описания подпрограмм, основное приложение - это набор поддпрограмм.
В моем примере подпрограммой является `http` сервер, который имеет разные модули. 
Например, модулю `inbox` и `outbox`. Детальное описание модулей в директории слоя `app`.

Название директории для данного слоя - `app`.

---

This sublayer contains descriptions of subprograms, the main application is a set of subprograms.
In my example, a subprogram is an `http` server that has different modules.
For example, the `inbox` and `outbox` modules. A detailed description of the modules in the `app` layer directory.

The directory name for this layer is `app`.

## Domain SubLayer

Этот подуровень содержит описания моделей доменной области приложения. В этом слое не должно быть никакой бизнес-логики,
только описания предметной части выраженной в коде. Например, модель `user` или `transaction`. 
Все что связанно с объектами принимающие участие в бизнес-процессе.

Данный слой не имеет доступа к другим слоям приложения, кроме слоя `lib`.

Название директории для данного слоя - `domain`.

---

This sublayer contains descriptions of domain models of the application. This layer should not have any business logic,
only descriptions of the subject area expressed in code. For example, the `user` or `transaction` model.
Everything related to objects participating in the business process.

This layer does not have access to other layers of the application, except for the `lib` layer.

The directory name for this layer is `domain`.

## Use-Case SubLayer

Этот подуровень содержит описания бизнес-процессов приложения, сервисный слой. Как описать и выразить бизнес-процесс решать вам. 
В моем примере я использую интерфейсы, которые описывают бизнес-процессы.

Название директории для данного слоя - `usecase`.

---

This sublayer contains descriptions of the business processes of the application, the service layer. How to describe and express the business process is up to you.
In my example, I use interfaces that describe business processes.

The directory name for this layer is `usecase`.

## Repository SubLayer

Этот подуровень является сервисным слоям отвечающий за данные.
В моем примере я использую репозитории для отдельных `usecase` бизнес-процессов, так же я использую репозитории для отдельной модели если это потребуется.
Например, репозиторий для модели `user` или `transaction`, а для бизнес-процессов репозиторий `inbox` и `outbox`.

Название директории для данного слоя - `repository`.

---

This sublayer is a service layer responsible for data.
In my example, I use repositories for individual `usecase` business processes, I also use repositories for a separate model if necessary.
For example, a repository for the `user` or `transaction` model, and for business processes the `inbox` and `outbox` repository.

The directory name for this layer is `repository`.

## Lib SubLayer

Этот подуровень является слоем библиотек, которые используются в приложении. 
Данный слой может иметь доступ только к слои `config`

Название директории для данного слоя - `lib`.

---

This sublayer is a layer of libraries used in the application.
This layer can only have access to the `config` layer

The directory name for this layer is `lib`.

## Storage SubLayer

Этот подуровень является слоем хранилища данных. В нем описываются все хранилища используемые в приложении.
Детальное описания структуры подуровня в директории `storage`.

Название директории для данного слоя - `storage`.

---

This sublayer is a data storage layer. It describes all the storages used in the application.
A detailed description of the structure of the sublayer in the `storage` directory.

The directory name for this layer is `storage`.
