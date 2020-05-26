Tehtävän teossa on käytetty apuna Go:n [dokumentaatiota.](https://golang.org/doc/)

# LinWinMac! Käännä ohjelma kolmelle alustalle: Windows, Linux, Mac. Tee staattinen käännös niin, että ohjelma toimii ilman mitään riippuvuuksia tai kirjastoja. Testaa ohjelman toiminta ainakin joillain näistä alustoista ja ota ruutukaappaukset.

[Lähdekoodi](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day2/osa1/hello.go)

## Käännökset eri alustoille (win, mac, linux)

Löysin mielenkiintoisen ["dokumentaation"](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63), jossa listataan Go:n kääntäjän tukemia arkkitehtuureja ja käyttöjärjestelmiä.

Avataan komentoa, jolla käännetään lähdekoodi binääriksi linux tietokoneelle, joka käyttää "modernia" prosessoria:

    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o hello hello.go

`CGO_ENABLED=0` Tässä ilmeisesti kirjastot sisällytetään binääriin.<br>
`GOARCH=amd64` Arkkitehtuuriksi on valittu amd64, eli 64bittiset "modernit" prosessorit.<br>
`GOOS=linux` Lähdekoodi käännetään binääriksi linux järjestelmille.<br>
`go build` Komento, jolla Go aloittaa kääntämisen.<br>
`-o hello` Flagi, jolla kerrotaan kääntäjälle, että lähdekoodi käännetään binääriksi `hello` -tiedostoon, joka tulee käännöksen jälkeen sijaitsemaan samassa hakemistossa `hello.go` lähdekoodin kanssa.<br>
`hello.go` Kerrotaan kääntäjälle käännettävän lähdekooditiedoston nimi (sijaintina sen hetkinen työhakemisto).<br>

Kaikki alla olevat käännökset on tehty Ubuntu 18.04 LTS käyttöjärjestelmä versiolla linux tietokoneella.

### Linux

[Käännetty binääri (64bit)](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day2/osa1/hello)

Käännetään amd64bit arkkitehtuurille:

    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o hello hello.go

Ajetaan koodi:

    ./hello

Koodin ajaminen binäärinä toimii:

    Hello World!

Voidaan todeta, että lähdekoodi kääntyi 64bittiseksi binääriksi:

    file hello

Tulos kertoo, että binääri on käännetty 64bit muotoon:

    hello: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, with debug_info, not stripped

Käännetään vielä linuxille 32bit muotoon:

    CGO_ENABLED=0 GOARCH=386 GOOS=linux go build -o hello hello.go

Tarkistetaan binäärin muoto:

    file hello

Muotona 32bit:

    hello: ELF 32-bit LSB executable, Intel 80386, version 1 (SYSV), statically linked, with debug_info, not stripped

Ajetaan binääri:

    ./hello

Ajo onnistui:

    Hello World!

### Windows

[Käännetty binääri (64bit)](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day2/osa1/hello-windows-64bit.exe)

Käännetään amd64bit arkkitehtuurille:

    CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o hello-windows-64bit.exe hello.go

Tämän suoritusta en testannut, mutta katsotaan `file` komennolla binäärin tietoja:

    file hello-windows-64bit.exe

Ainakin tämän perusteella voisi olettaa, että binäärin pystyisi suorittamaan:

    hello-windows-64bit.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows

### Mac

[Käännetty binääri (64bit)](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day2/osa1/hello-mac-64bit)

Käännetään amd64bit arkkitehtuurille:

    CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o hello-mac-64bit hello.go

Tarkistetaan binääri `file` komennolla:

    file hello-mac-64bit

Tulos:

    hello-mac-64bit: Mach-O 64-bit x86_64 executable, flags:<NOUNDEFS>

Binäärin siirron jälkeen mac koneelle ajetaan binääri:

    ./hello-mac-64bit

Ajo toimi oletetusti:

![alt text](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day2/osa1/mac_run.jpg "Binäärin suoritus terminaalilla")

Binäärin voi myös suorittaa graafisessa käyttöliittymässä tuplaklikkaamalla tiedostoa:

![alt text](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day2/osa1/doubleclick.jpg "Binäärin suoritus tuplaklikkaamalla")<br>

# Kirjastoja kohti. Kirjoita ohjelma, joka käyttää kahta uutta ominaisuutta tai kirjastoa Go by Example-kirjasta.

[Go by Example](https://gobyexample.com/)

## 1. Ohjelma

Ohjelma testaa käyttäjän antaman verkko-osoitteen statuksen.

Apuna ohjelman teossa käytin [tätä ohjetta.](https://gobyexample.com/http-clients)

[Lähdekoodi](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day2/osa2/get-http-status.go)

### Ohjelman käyttö

Esimerkki:

    ./osa2 --url https://google.com

Josta palautuu testattavan sivun http status:

    https://google.com status: 200 OK

### Optiot (flagit)

```
Usage of ./osa2:
  -url string
    	Website url for response status test (default "https://vuorivirta.fi/")
```

## 2. Ohjelmaan testi

Apuna testin teossa käytin [tätä ohjetta.](https://gobyexample.com/testing)

[Lähdekoodi](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day2/osa2/get-http-status_test.go)

### Testin ajo

Testiin on kovakoodattu testattavaksi osoitteeksi `https://google.com/`.

Esimerkki:

    go test -v

Josta palautuu testin tulos:

```
=== RUN   TestCheckConnection
https://google.com/ status: 200 OK
--- PASS: TestCheckConnection (0.21s)
PASS
ok  	_/home/nikke/kotitehtavat/Go-Programming-Course-2020/day2/osa2	0.217s
```

Mikäli vaihtaa kovakoodatun osoitteen esimerkiksi `https://terokarvinen.com/` (certti testaushetkellä ummessa) antaa testi seuraavan tuloksen:

```
=== RUN   TestCheckConnection
--- FAIL: TestCheckConnection (30.00s)
panic: Get https://terokarvinen.com/: dial tcp 139.162.131.217:443: i/o timeout [recovered]
	panic: Get https://terokarvinen.com/: dial tcp 139.162.131.217:443: i/o timeout

goroutine 5 [running]:
testing.tRunner.func1(0xc4201280f0)
	/usr/lib/go-1.10/src/testing/testing.go:742 +0x29d
panic(0x67f1e0, 0xc4200934a0)
	/usr/lib/go-1.10/src/runtime/panic.go:502 +0x229
_/home/nikke/kotitehtavat/Go-Programming-Course-2020/day2/osa2.checkConnection(0x6ce05c, 0x19, 0x0, 0x0)
	/home/nikke/kotitehtavat/Go-Programming-Course-2020/day2/osa2/get-http-status.go:22 +0x1a8
_/home/nikke/kotitehtavat/Go-Programming-Course-2020/day2/osa2.TestCheckConnection(0xc4201280f0)
	/home/nikke/kotitehtavat/Go-Programming-Course-2020/day2/osa2/get-http-status_test.go:6 +0x3a
testing.tRunner(0xc4201280f0, 0x6db820)
	/usr/lib/go-1.10/src/testing/testing.go:777 +0xd0
created by testing.(*T).Run
	/usr/lib/go-1.10/src/testing/testing.go:824 +0x2e0
exit status 2
FAIL	_/home/nikke/kotitehtavat/Go-Programming-Course-2020/day2/osa2	30.006s

```