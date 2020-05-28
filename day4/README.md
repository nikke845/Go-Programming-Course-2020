Pohjana Tero Karvinen 2020: [Go Programming Course 2020](http://terokarvinen.com/2020/go-programming-course-2020-w22/)<br>
Tehtävän teossa on käytetty apuna Go:n [dokumentaatiota.](https://golang.org/doc/)

# Oma ohjelma

Kurssin viimeinen tehtävä on oma ohjelma, jossa on tarkoitus käyttää kurssilla opeteltuja taitoja, sekä syventää Go kielen osaamista.

Ohjelman aiheeksi valitsin [edellisen päivän tehtävässä](https://github.com/nikke845/Go-Programming-Course-2020/tree/master/day3#kaikki-tai-jotain-valitse-aihe-tai-pari-vaihtoehtoa-kurssin-loppuprojektille-palastele-mieti-miten-voit-ensin-toteuttaa-pienen-toimivan-kokonaisuuden-ja-sen-p%C3%A4%C3%A4lle-vaihtoehtoja-laajentaa) ideoidun sää ohjelman.

Aloitin ohjelman teon noin klo. 14 ja ohjelman ensimmäinen versio piti palauttaa saman päivän aikana ennen klo 20. Ensimmäinen versio ohjelmasta löytyy [täältä.](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/v1/weather.go) Ensimmäisessä versiossa ohjelman seuraavat ominaisuudet toimivat:

* Json datan hakeminen API:sta
* Datan esittäminen tulosteena
* Flageilla voi antaa koordinaateilla oman lokaation säätä varten
* Flagilla voi antaa valinnaisen tiedoston, johon dataa tallennetaan
* Datan tulostaminen tiedostoon uudelle riville, mikäli flagilla on annettu tiedosto, johon data tallennetaan

Testasin ensimmäisen version toimintaa, ja totesin sen toimivan odotetusti [kuva](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day4/v1/weather_v1.png).

Ohjelman tekoa sai jatkaa seuraavan aamun ohjelmien esitystilaisuuteen asti.

## Sää ohjelman API

Sää ohjelma käyttää [Dark Sky](https://darksky.net/poweredby/):n API:a, josta säätiedot haetaan. Dark Sky:n sivuilta löytyy hyvät [ohjeet API:n käyttämiseen](https://darksky.net/dev/docs), HUOM. Apple on ostanut kyseisen yrityksen ja [uusia käyttäjiä ei ilmeisesti oteta palveluun.](https://blog.darksky.net/).

API kyselyn muoto on seuraava:

    https://api.darksky.net/forecast/apikey/latitude,longitude?units=auto&exclude=minutely,hourly,daily,alerts,flags

Esimerkkinä ohjelman oletus haku API:sta Helsingin säälle:

    https://api.darksky.net/forecast/apikey/60.192059,24.945831?units=si&exclude=minutely,hourly,daily,alerts,flags

