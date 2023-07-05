# qzhub

### usage
ビルド
```shell
docker build -t qzhub .
```
初回実行
```shell
docker run --name qzhub -p 8080:80 -d -v $(pwd):/app qzhub
```
2回目以降の実行
```shell
docker start qzhub
```
