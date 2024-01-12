# Repository Layer

## RU
Этот подуровень является сервисным слоям отвечающий за работу с хранилищем.
В моем примере я использую репозитории для отдельных `usecase` бизнес-процессов, которые взаимодействую с интрейфесами других репозиторий.
Остальные репозитории используються для отдельных моделей или агрегатов в зависимости от типа хранилища и их название отражает тип хранилища.
Например, репозиторий: `User`, `UserCahce` и `Transaction`, `TransactionCasche` для моделей `user` и `transaction`, интерфейсы реализуют структуры `UserPG`, `UserCascheRD`, `TransactionPG`, `TransactionCacheRD`, где префикс `PG` - **postgres** и `RD` - **redis**.
Так же как и в слое `domain` я предпочитаю детально документировать интейрфесы, которые описывают бизнес-процессы.

Репозиторий обращаються к конкретному хранилищу, которое описывается в `storage` слое.
В слое `storage` реализуют функции для определенной логики, это могут быть как `sql` запросы, так и описание *транфармации данных для транспартировки*(dto).
В моем примере для того, что бы сохранить данные транзакции в `postgres` - вызываю функцию из пакета `pg` для репозитория `TransactionPG`, который находиться в слое `storage`. 


```go
func (rp TransactionPG) Create(ctx context.Context, transaction *model.Transaction) error {
	return pg.CreateTransaction(ctx, rp.db, transaction)
}
```

В функции `CreateTransaction` описывается логика сохранения данных в `postgres`. Так же функция принимает вторым аргументом интерфейс работы с хранилищем, в под интерфейсом может быть как объект коннекта, так и объект транзакции, что именно передавать решает репозиторий.

```go
func CreateTransaction(
	ctx context.Context,
	conn postgres.Connect,
	transaction *model.Transaction,
) error {
	sql := `
INSERT INTO transactions (
	id,
	payment_method,
	amount,
	currency,
	status,
	description,
	user_from_id,
	user_to_id,
	created_at,
 	updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	return conn.Exec(
		ctx,
		sql,
		transaction.ID,
		transaction.PaymentMethod,
		transaction.Amount,
		transaction.Currency,
		transaction.Status,
		transaction.Description,
		transaction.UserFromID,
		transaction.UserToID,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)
}
```
---

## EN

This sublayer is a service layer responsible for working with the repository.
In my example, I use repositories for individual `usecase` business processes that interact with the interfaces of other repositories.
The rest of the repositories are used for individual models or aggregates depending on the type of repository and their name reflects the type of repository.
For example, the repository: `User`, `UserCahce` and `Transaction`, `TransactionCasche` for the `user` and `transaction` models, the interfaces implement the `UserPG`, `UserCascheRD`, `TransactionPG`, `TransactionCacheRD` structures, where the prefix `PG` - **postgres** and `RD` - **redis**.
As in the `domain` layer, I prefer to document in detail the interfaces that describe business processes.

The repository refers to a specific repository, which is described in the `storage` layer.
In the `storage` layer, functions are implemented for a specific logic, these can be both `sql` queries and a description of *data transformation for transportation* (dto).
In my example, in order to save transaction data in `postgres`, I call the function from the `pg` package for the` TransactionPG` repository, which is located in the` storage` layer.

```go
func (rp TransactionPG) Create(ctx context.Context, transaction *model.Transaction) error {
	return pg.CreateTransaction(ctx, rp.db, transaction)
}
```

The `CreateTransaction` function describes the logic for saving data in `postgres`. The function also takes the interface of working with the repository as the second argument, the interface can be both a connection object and a transaction object, which one to pass is decided by the repository.

```go
func CreateTransaction(
	ctx context.Context,
	conn postgres.Connect,
	transaction *model.Transaction,
) error {
	sql := `
INSERT INTO transactions (
	id,
	payment_method,
	amount,
	currency,
	status,
	description,
	user_from_id,
	user_to_id,
	created_at,
 	updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	return conn.Exec(
		ctx,
		sql,
		transaction.ID,
		transaction.PaymentMethod,
		transaction.Amount,
		transaction.Currency,
		transaction.Status,
		transaction.Description,
		transaction.UserFromID,
		transaction.UserToID,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)
}
```