# Cryptopals solutions

[![builds.sr.ht status](https://builds.sr.ht/~dbalan/cryptopals.svg)](https://builds.sr.ht/~dbalan/cryptopals?)


Solutions to [cryptopals](https://cryptopals.com) challenges in [Go](https://golang.org).

## Running
Solutions are implemented as a testcases.

```bash
cd cryptopals
for set in "set1" "set2" "set3"; do
cd $set && go test -v && cd ..
done
```

