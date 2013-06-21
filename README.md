# go-have-fun

### db-fun

Ensure Redis is installed, [homembrew](http://mxcl.github.io/homebrew/) `brew install redis` should work on Mac OS X.

In your first process, get Redis up and running.

```bash
$ redis-server
```

If you'd like to monitor the action happening in Redis, open yet another process and execute:

```bash
$ redis-cli
> MONITOR
```

Next, have some fun!

```bash
$ go run db-fun
```
This shows some simple concurrency hitting the DB.

## Epilogue

More fun to come!