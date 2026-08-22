myapp/
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   ├── auth/
│   └── service/
├── pkg/
│   └── utils/
├── go.mod
├── configs/
└── scripts/

Refer : https://github.com/golang-standards/project-layout

At the root of your project we have:
go.mod → declares your module and dependencies
go.sum → hashes of those dependencies


one place for app code (/internal)
one place for reusable libraries (/pkg)
one place for executables (/cmd)

/cmd :
Contains the application entry points — the code that produces one or more executables.
Example:
cmd/
  webserver/
    main.go
  worker/
    main.go

/internal

This is where private code goes — logic that’s internal to your app and not intended to be reused externally.
Go enforces that anything under an internal directory cannot be imported by other modules outside your project


/pkg
Holds public libraries — code that others can import.



/api	API definitions — e.g., OpenAPI/Swagger files
/configs	Configuration templates
/scripts	Useful scripts (build, CI, deploy)
/web	Web assets (HTML, CSS, JS)
/docs	Documentation files
/test	External test applications and test data
/examples	Example code showing how to use your project
/tools	Supporting utilities used during development
/deployments	Deployment configs (Kubernetes, Docker Compose)
/assets	Images, icons, other files








