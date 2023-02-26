### DEFER

> In Go, defer is a keyword used to schedule a function call to be executed just before the function that contains the defer statement returns.
```
defer functionCall(arguments)
```

> It works like first in last out.
> all defers are scheduled and last one is executed immediately and first come statement with defer excuted in last
