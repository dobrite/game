install postgres 9.3 and contrib

```
sudo -i -u postgres
createuser --interactive
game
y
createdb game
psql
alter user game with password 'badgam4';
\q
logout
```

add the line in pg_hba.conf above `local all all peer`
```
local game game md5
```

restart postgres for the changes to take effect

```
git clone https://github.com/dobrite/game
npm install
gulp
```

[install go](http://golang.org/doc/install)

game currently uses godep to manage dependencies and their versions:

```
go get github.com/tools/godep
godep restore
goose up
go run main.go
```

point your browser to [localhost:3000](http://localhost:3000)

![game](http://i.imgur.com/2ZNUEFO.png "game")
