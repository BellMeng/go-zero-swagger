# go-zero-swagger

`go-zero-swagger` is a Swagger/OpenAPI documentation generator for [go-zero](https://github.com/zeromicro/go-zero), inspired by [gin-swagger](https://github.com/swaggo/gin-swagger).

Unlike the official `goctl` documentation generator that parses `.api` files, this tool extracts documentation directly from Go source code annotations, allowing you to maintain your API definitions closer to your handlers—just like with Gin.

## 📦 [Installation](#installation)

```bash
go get -u github.com/BellMeng/go-zero-swagger
```

## 🤖 [Usage](#usage)

1. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format).
2. Run Swag to generate Swagger docs:

   ```bash
   swag init
   ```

3. Import the docs like this: I assume your project named `github.com/go-project-name/docs`.
   ```go
    import (
       docs "github.com/go-project-name/docs"
    )
   ```
4. Add Swagger docs to your API handler:

   ```go
    import (
        swaggerFiles "github.com/swaggo/files"

        goZeroSwagger "github.com/BellMeng/go-zero-swagger"

        "github.com/zeromicro/go-zero/rest"

        docs "github.com/go-project-name/docs"
    )
    ...
    server.AddRoute(rest.Route{
        Method:  http.MethodGet,
        Path:    "/swagger/:any",
        Handler: goZeroSwagger.WrapHandler(swaggerFiles.NewHandler()),
    })
    ...
   ```

5. Other configurations you can refer to [gin-swagger](https://github.com/swaggo/gin-swagger)
