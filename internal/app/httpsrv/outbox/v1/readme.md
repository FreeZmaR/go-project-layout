# Module layer

Слой модуля - это набор контроллеров(обработчиков), функции валидации, мидлвар, сервиса-провайдер и интерфейса ответа(респондер).

Формирования модуля происходит в файле `module.go` в директории модуля с использованием функции `fx.Module()`.
Для формирования зависимостей модуля используется пакет сервис-провайдер(`servprovider`), где реализуется конструкторы зависимостей слоев.

--- 

This is a module layer. It contains a set of controllers(handlers), validation functions, middlewares, service provider and response interface(responder).

The module is formed in the `module.go` file in the module directory using the `fx.Module()` function.
To form the module dependencies, the service provider package (`servprovider`) is used, where the constructors of layer dependencies are implemented.

