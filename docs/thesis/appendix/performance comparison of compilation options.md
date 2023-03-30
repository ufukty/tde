## Serial compilation

```shell
for i in {1..100}; do
    go install ./cmd/client
done

real    0m12.518s
user    0m18.260s
sys     0m25.231s
```

```shell
for i in {1..100}; do
    go build -o /dev/null ./cmd/client
done

real    0m13.358s
user    0m18.472s
sys     0m25.719s
```

```shell
for i in {1..100}; do
    go build -o "$tempdir/$i" ./cmd/client
done

real    0m29.041s
user    0m38.632s
sys     0m30.056s
```

## Parallel compilation

```shell
for i in {1..100}; do
    go install ./cmd/client &
done

wait $(jobs -p)

real    0m5.076s
user    0m17.326s
sys     0m21.251s
```

```shell
for i in {1..100}; do
    go build -o /dev/null ./cmd/client &
done

wait $(jobs -p)

real    0m4.074s
user    0m7.016s
sys     0m8.842s
```

```shell
for i in {1..100}; do
    go build -o "$tempdir/$i" ./cmd/client &
done

wait $(jobs -p)

real    0m8.997s
user    0m41.531s
sys     0m26.942s
```

## Run executables

Running previously compiled binaries

```shell
cd "/var/folders/pr/0xpk709n5fdgbrbbdxhjcqv40000gn/T/tmp.P8NOU3Jc"
for i in {1..100}; do
    chmod +x "$i"
    "./$i" >>log.txt 2>&1 &
done

wait $(jobs -p)

real    0m0.237s
user    0m0.281s
sys     0m0.375s
```

Compile and run

```shell
for i in {1..100}; do
    go run ./cmd/client >/dev/null 2>&1 &
done

wait $(jobs -p)

real    0m16.620s
user    0m31.449s
sys     0m25.414s
```

## TL;DR

| Command                         | Seconds (Serial) | Seconds (Parallel) |
| ------------------------------- | ---------------- | ------------------ |
| `install` (overwrites previous) | 13               | 5                  |
| `build` & discard output        | 14               | 4                  |
| `build` & store output          | 29               | 9                  |

-   Time spent: `install` < `build -o /dev/null` < `build -o /tmp/$i`
-   `build` and execute from shell is faster than `run`
-   Parallelism reduces the time spent.
