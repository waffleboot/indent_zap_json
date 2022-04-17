translate

```json
{"level":"error","ts":1650191591.372684,"caller":"wrap/main.go:95","msg":"error was","error":"wrap: cause","errorVerbose":"cause\nmain.run1\n\t/Users/yangand/Workspace/wrap/main.go:100\nmain.run2\n\t/Users/yangand/Workspace/wrap/main.go:104\nmain.main\n\t/Users/yangand/Workspace/wrap/main.go:94\nruntime.main\n\t/Users/yangand/.asdf/installs/golang/1.17/go/src/runtime/proc.go:255\nruntime.goexit\n\t/Users/yangand/.asdf/installs/golang/1.17/go/src/runtime/asm_amd64.s:1581\nwrap\nmain.run2\n\t/Users/yangand/Workspace/wrap/main.go:105\nmain.main\n\t/Users/yangand/Workspace/wrap/main.go:94\nruntime.main\n\t/Users/yangand/.asdf/installs/golang/1.17/go/src/runtime/proc.go:255\nruntime.goexit\n\t/Users/yangand/.asdf/installs/golang/1.17/go/src/runtime/asm_amd64.s:1581","stacktrace":"main.main\n\t/Users/yangand/Workspace/wrap/main.go:95\nruntime.main\n\t/Users/yangand/.asdf/installs/golang/1.17/go/src/runtime/proc.go:255"}
```
to
```json
{
  "caller": "wrap/main.go:96",
  "error": "wrap: cause",
  "errorVerbose": [
    "cause",
    "main.run1/Users/yangand/Workspace/wrap/main.go:101",
    "main.run2/Users/yangand/Workspace/wrap/main.go:105",
    "main.main/Users/yangand/Workspace/wrap/main.go:95",
    "wrap",
    "main.run2/Users/yangand/Workspace/wrap/main.go:106",
    "main.main/Users/yangand/Workspace/wrap/main.go:95"
  ],
  "level": "error",
  "msg": "error was",
  "stacktrace": [
    "main.main/Users/yangand/Workspace/wrap/main.go:96"
  ],
  "ts": 1650191659.478837
}
```