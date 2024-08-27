# api file path: ./app/usercenter/cmd/api/desc/usercenter.api
# api go file path : ./app/usercenter/cmd/api
# $1 name
goctl api go  --api  ./app/$1/cmd/api/desc/$1.api --dir ./app/$1/cmd/api --home ./develop/.goctl --style go_zero