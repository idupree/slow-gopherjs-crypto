
This is a demo for a gopherjs speed issue.  `SignPKCS1v15` is far slower
in gopherjs.

### To build:

If you need to install gopherjs, then
`go get -u github.com/gopherjs/gopherjs`.

Then
```
cd speedtest
go build
gopherjs build
```

### To test speed:

in `speedtest`, after building,

```
time ./speedtest
time node speedtest.js
```

My results on an Intel i5-2500 CPU:

```
speedtest$ time ./speedtest

real    0m0.009s
user    0m0.004s
sys     0m0.004s

speedtest$ time node speedtest.js

real    0m3.254s
user    0m3.328s
sys     0m0.024s
```
Which is about 300 times worse in javascript.
Maybe we can do better.
It's this bad in-browser too, which affects things like
macgyver's ChromeStorageBackend
https://github.com/stripe/macgyver

### To test with a different RSA key

The demo key size is 2048 bits.

If you want to generate different keys (e.g. different sizes)
for testing, go into `makekeyforspeedtest`, edit `main.go`, and
`go run main.go`.  Then copy the contents of
`makekeyforspeedtest/key.pem` into the
value of `keyfile` in `speedtest/main.go`.
