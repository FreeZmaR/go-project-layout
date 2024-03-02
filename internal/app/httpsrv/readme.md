# Sub App layer

Данный слой содержит в себе подпрограммы, которые могут быть использованы в основном приложении.
Подпрограммы могут быть как встроенными, так и внешними. Встроенные подпрограммы находятся внутри основного приложения, все подпрограммы могут включать отдельные модули.

Подпрограмма в приведенном примере реализуют работу `http` сервера. У подпрограммы есть два модуля, которые могут быть использованные в месте или отдельно.
Эти модули `inbox` и `outbox` реализуют входящие и исходящие запросы соответственно и инкапсулируют в себе сборку под свои задачи.

---

This layer contains sub-programs that can be used in the main application.
Sub-program can be both built-in and external. Built-in sub-applications are located inside the main application, all sub-applications can include separate modules.

The sub-program in the example provided implements the operation of the `http` server. The sub-program has two modules that can be used together or separately.
These modules `inbox` and `outbox` implement incoming and outgoing requests, respectively, and encapsulate the building for their needs.

## Sub-program structure

### App.go

Является базовым типом подпрограммы, отвечает за правильный запуск и остановку подпрограммы.
Подпрограмма использует внутри себя контроллер отвечающий за работу подпрограмм - `fxutils.Runner`.
Каждая подпрограмма должна реализовать методы `Start(ctx context.Context) error` и `Stop(ctx context.Context) error`.


--- 

This type is the base type of the sub-program, responsible for the correct start and stop of the sub-program.
The sub-program uses a controller inside it that is responsible for the work of the sub-program - `fxutils.Runner`.
Each sub-program must implement the methods `Start(ctx context.Context) error` and `Stop(ctx context.Context) error`.

### Module.go

Данный файл отвечает за формирование модулей подпрограммы и самой подпрограммы.
В файле реализованны функции-конструкторы для определенной версии подпрограммы, так и для версий модулей.

Для реализации приложения на базе подпрограммы используется тип `fxutils.App`, который контролирует запуск приложения.
Для реализации модулей используется тип `fx.Module`, который интегрирует все нужные зависимости в подпрограмму.

---

This file is responsible for forming the modules of the sub-program and the sub-program itself.
The file implements constructor functions for a specific version of the sub-program, as well as for module versions.

To implement an application based on a sub-program, the type `fxutils.App` is used, which controls the launch of the application.
To implement modules, the type `fx.Module` is used, which integrates all the necessary dependencies into the sub-program.

### Module_provider.go

Данный файл является провайдером зависимостей для подпрограммы.
Файл делить на 3 основные функции, которые отвечают за создание и конфигурацию зависимостей.

 - `Provider` - функция, которая создает и конфигурирует зависимости для подпрограммы.
 - `Invoke` - функция, которая выполняют логику для зависимостей подпрограммы.
 - `Decorator` - функция, которая выполняет дополнительную конфигурацию для зависимостей подпрограммы.

---

This file is a dependency provider for the sub-program.
The file is divided into 3 main functions that are responsible for creating and configuring dependencies.

 - `Provider` - a function that creates and configures dependencies for the sub-program.
 - `Invoke` - a function that performs logic for sub-program dependencies.
 - `Decorator` - a function that performs additional configuration for sub-program dependencies.