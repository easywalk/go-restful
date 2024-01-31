# go-easywalk : RESTFul API를 제공합니다.

## install
```shell
go get -u github.com/easywalk/go-restful
```

## usage
```go
	// create File Service
	repo := repository.NewSimplyRepository[*model.File](db)
	svc := service.NewGenericService[*model.File](repo)

	// Gin router
	r := gin.Default()
	group := r.Group("/files")

	handler.NewHandler[*model.File](group, svc)
```

## entity Spec
```go
type SimplyEntityInterface interface {
	GetID() uuid.UUID
	SetID(id uuid.UUID)
	SetCreatedAt(t time.Time)
	SetUpdatedAt(t time.Time)
}
```

## example

```go
r := gin.Default()
group := r.Group("/files")

repo := repository.NewSimplyRepository[*model.File](db) // db is *gorm.DB
svc := service.NewGenericService[*model.File](repo) // svc is *service.GenericService[*model.File]
handler.NewHandler[*model.File](group, svc) // group is *gin.RouterGroup
r.Run()

```