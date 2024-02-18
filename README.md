# y74h1116/go-gin-libs
## page library
- The page library is useful to linkage between handler and view.  
- usage  
  - refer to [examples/login.go](examples/login.go)  
  - import github.com/y74h1116/go-gin-libs/page and gin and go-playground/validator  
    ```Go  
    import (
      "net/http"

      "github.com/y74h1116/go-gin-libs/page"

      "github.com/gin-gonic/gin"
      "github.com/go-playground/validator/v10"
    )
    ```  
  - initialize page object  
    ```Go  
    var loginValidationMessages = map[string]map[string]string{
        "Email": {
            "required": "メールアドレスを入力してください。",
        },
        "Password": {
            "required": "パスワードを入力してください。",
            "custom_password_characters": "パスワードに利用できない文字が含まれています。",
        },
    }

  	pageLogin := page.NewPage("ログイン", loginValidationMessages)
    ```  
  - error handling  
    ```Go  
    // handler call SetFormErrorsByValidator and SetOld, then pass viewParam to ctx.HTML()
    if err := ctx.ShouldBind(&loginForm); err != nil {
        viewParam := login.pageLogin.CreateViewParam()
        login.pageLogin.SetFormErrorsByValidator(viewParam, err.(validator.ValidationErrors))
        login.pageLogin.SetOld(viewParam, loginForm)
        ctx.HTML(http.StatusBadRequest, "login/index", viewParam)
        return
    }

    // view can refer formErrors and old
    メールアドレス<br>
    {{ if .formErrors.Email }}<span class="red_text">{{.formErrors.Email}}</span><br>{{ end }}
    <input type="text" name="email" required class="input_text" value="{{.old.Email}
      placeholder="xxxx@xxx.xx" size="60" min="5" max="255">
    ```  
## sample
- usage  
  ```  
  $ cd go-gin-libs/examples

  $ docker compose build

  $ docker run -it --rm -v $(pwd):/home/go -v $(pwd)/go/pkg:/go/pkg --workdir /home/go  examples-example  go mod tidy

  $ docker compose up

  # access http://localhost:8080 with a web browser
  ```  
