# 生成swagger 文件
# $1 = api 文件名
goctl api plugin -p goctl-swagger="swagger -filename ./doc/swagger/$1.json" --api ./app/$1/cmd/api/desc/$1.api