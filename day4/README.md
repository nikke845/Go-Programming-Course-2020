Pohjana Tero Karvinen 2020: [Go Programming Course 2020](http://terokarvinen.com/2020/go-programming-course-2020-w22/)<br>
Tehtävän teossa on käytetty apuna Go:n [dokumentaatiota.](https://golang.org/doc/)

# Oma ohjelma: WeatherApp

Kurssin viimeinen tehtävä on oma ohjelma, jossa on tarkoitus käyttää kurssilla opeteltuja taitoja, sekä syventää Go kielen osaamista.

Ohjelman aiheeksi valitsin [edellisen päivän tehtävässä](https://github.com/nikke845/Go-Programming-Course-2020/tree/master/day3#kaikki-tai-jotain-valitse-aihe-tai-pari-vaihtoehtoa-kurssin-loppuprojektille-palastele-mieti-miten-voit-ensin-toteuttaa-pienen-toimivan-kokonaisuuden-ja-sen-p%C3%A4%C3%A4lle-vaihtoehtoja-laajentaa) ideoidun sää ohjelman.

WeatherApin tarkoitus on kertoa käyttäjälle tämän hetkinen sää (lämpötila, sateen mahdollisuus, tuulen nopeus, kuvaus säästä), sekä antaa mahdollisuus sään lokittamiseen vapaavalintaiseen tiedostoon CSV muodossa.

![alt text](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/weatherapp.png "Ohjelman suoritus ilman parametrejä")

[Lähdekoodi](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/weather.go)

[Binääri Linux järjestelmille](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/weather)

[Binääri Mac järjestelmille](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/weather-mac-64bit)

[Binääri Windows järjestelmille](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/weather-windows-64bit.exe)

[Lisenssi](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/LICENSE)

Aloitin ohjelman teon noin klo. 14 ja ohjelman ensimmäinen versio piti palauttaa saman päivän aikana ennen klo 20. Ensimmäinen versio ohjelmasta löytyy [täältä.](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/v1/weather.go) Ensimmäisessä versiossa ohjelman seuraavat ominaisuudet toimivat:

* Json datan hakeminen API:sta
* Datan esittäminen tulosteena
* Flageilla voi antaa koordinaateilla oman lokaation säätä varten
* Flagilla voi antaa valinnaisen tiedoston, johon dataa tallennetaan
* Datan tulostaminen tiedostoon uudelle riville, mikäli flagilla on annettu tiedosto, johon data tallennetaan

Testasin ensimmäisen version toimintaa, ja totesin sen toimivan odotetusti [kuva](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/v1/weather_v1.png).

Viimeiseen versioon lisäsin seuraavat ominaisuudet aikaisemman (v1) version ominaisuuksien lisäksi:

* Tiedostoon data tulostuu CSV formaatissa
* Tiedostoon tulostetaan aikaleima datan keräyshetkestä
* Mikäli ohjelmaan syötetään parametrinä tiedosto, jota ei ole olemassa ohjelma tekee tiedoston, ja lisää sen ensimmäiselle riville tietotyyppien selitykset.
* Korjasin pienen bugin sateen todennäköisyyden esittämisessä, ennen ohjelma näytti `0.01%` todennäköisyydeksi, vaikka sen olisi pitänyt olla `1%`
* Lisäsin API:n secret keyn tiedoston sisään, josta se luetaan ohjelmalle.

Ohjelman "viimeiseen" version sain valmiiksi saman päivän aikana noin klo 22.30.

## Sää ohjelman API

Sää ohjelma käyttää [Dark Sky](https://darksky.net/poweredby/):n API:a, josta säätiedot haetaan. Dark Sky:n sivuilta löytyy hyvät [ohjeet API:n käyttämiseen](https://darksky.net/dev/docs), HUOM. Apple on ostanut kyseisen yrityksen ja [uusia käyttäjiä ei ilmeisesti oteta palveluun.](https://blog.darksky.net/).

API kyselyn muoto on seuraava:

    https://api.darksky.net/forecast/apikey/latitude,longitude?units=auto&exclude=minutely,hourly,daily,alerts,flags

Esimerkkinä ohjelman oletus haku API:sta Helsingin säälle:

    https://api.darksky.net/forecast/apikey/60.192059,24.945831?units=si&exclude=minutely,hourly,daily,alerts,flags

## Ohjelman testaus

Testasin ohjelman kaikkia toimintoja. Ohjelmalle tehdyt testit löytyvät [kuvasta.](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/weatherapp_test.png)

## Ohjelman käyttö

Jotta ohjelmani pystyy suorittamaan tarvitaan [Dark Sky](https://darksky.net/poweredby/):n API:in avain. Ohjelma lukee avaimen `.env` tiedostosta, jonka pitää olla samassa hakemistossa suoritettavan binäärin kanssa. Avaimen sisältävässä `.env` tiedostossa ei saa olla ensimmäisen rivin lisäksi muita/tyhjiä rivejä.

Ohjelman voi suorittaa ilman parametrejä, jolloin se palauttaa Helsingin sen hetkisen sään:

```
$:~/kotitehtavat/Go-Programming-Course-2020/day4$ ./weather 
Current temperature: 11.32℃
Chance of rain: 0%
Current wind: 3.94㎧
Summary: Clear
-------------------
Powered by Dark Sky
https://darksky.net/poweredby/

```

Ohjelmaan voidaan syöttää parametreinä kaupungin sijainti. Koordinaatteja kaupungille voi hakea vaikka [tällä verkkosivustolla.](https://www.latlong.net/) Esimerkiksi Espoon sijainti on `latitude 60.205490` ja `longitude 24.655899`, joilla voi hakea Espoon sään:

```
$:~/kotitehtavat/Go-Programming-Course-2020/day4$ ./weather --latitude 60.205490 --longitude 24.655899
Current temperature: 11.49℃
Chance of rain: 0%
Current wind: 3.76㎧
Summary: Clear
-------------------
Powered by Dark Sky
https://darksky.net/poweredby/

```

Paramerinä voi myös syöttää vapaavalintaisen tiedoston, jolloin ohjelma tallentaa viimeiselle riville lämpötilatiedot CSV muodossa:

```
$:~/kotitehtavat/Go-Programming-Course-2020/day4$ ./weather --file test.csv
Current temperature: 11.16℃
Chance of rain: 0%
Current wind: 3.95㎧
Summary: Clear
-------------------
Powered by Dark Sky
https://darksky.net/poweredby/
$:~/kotitehtavat/Go-Programming-Course-2020/day4$ cat test.csv 
Time,Temperature,RainProbability,WindSpeed,Summary
"2020-05-28T23:14:32+03:00",11.16,0,3.95,"Clear"
```

### Optiot (flagit)

```
Usage of ./weather:
  -file string
    	Filename for optional data logging
  -latitude string
    	City's latitude (default "60.192059")
  -longitude string
    	City's longitude (default "24.945831")
```

### Kirjastot ja lähteet

Ohjelmassa käytin yhtä [ulkopuolista kirjastoa "percent.go"](https://github.com/dariubs/percent), jolla käänsin float64 arvon prosentiksi. Kirjaston lisenssi (MIT) [löytyy täältä.](https://github.com/dariubs/percent/blob/master/license)

Ohjelman teossa käytin pääasiassa lähteinä [Go:n omaa dokumentaatiota](https://golang.org/doc/), mutta käytin myös muutamaa verkosta löytynyttä koodi esimerkkiä:

* Käytin [tätä esimerkkiä](https://flaviocopes.com/golang-remove-new-line-from-readstring/) newlinen poistamiseen .env tiedostosta readApiKey() funktiossa.
* Käytin [tätä esimerkkiä](https://stackoverflow.com/questions/53749097/how-to-append-to-a-file-if-it-exists-else-create-a-new-and-write-to-it) uuden tiedoston tekemiseen.
* Käytin [tätä esimerkkiä](https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go) tiedoston olemassaolon tarkistamiseen.
* Käytin [tätä esimerkkiä](https://golangbot.com/write-files/) tiedoston kirjoittamiseen.
* Käytin [tätä esimerkkiä](https://gobyexample.com/reading-files) tiedoston lukemiseen.
* Käytin [tätä esimerkkiä](https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/) Json datan lukemiseen.
* Käytin [tätä esimerkkiä](https://gobyexample.com/json) Json datan lukemiseen.
* Käytin [tätä esimerkkiä](https://gobyexample.com/string-formatting) stringien formatoinnissa.