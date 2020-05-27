Pohjana Tero Karvinen 2020: [Go Programming Course 2020](http://terokarvinen.com/2020/go-programming-course-2020-w22/)<br>
Tehtävän teossa on käytetty apuna Go:n [dokumentaatiota.](https://golang.org/doc/)

# Kirjastosta päivää. Tee ohjelma, joka käyttää ainakin kahta uutta kirjastoa. Jos ohjelmallasi on jokin käyttötarkoitus, se on mainiota. Muista ensin kokeilla kutakin kirjastoa erikseen. Käytä jotain uutta kirjastoa, eli ei pelkästään rand, fmt, time eikä string.

## Ohjelma

Ohjelmani hakee kovakoodatulle käyttäjälle palkkakauden perusteella kirjatut työpäivät.

Ohjelmassa käytin kahta minulle ennestään tuntematonta kirjastoa `bytes` ja `encoding/json`. Koska halusin Json kirjastolle Json dataa, jota ei ole vain kovakoodattu tiedostoon, päätin käyttää aikaisemmalla kurssilla tekemääni Java ohjelmaa, jonka [Rest rajapinta](https://work-hours-management-app.herokuapp.com/paycycles/user1/Joulukuu) autogeneroi Json dataa, jota ohjelmani käyttää tässä harjoituksessa hyödyksi.

Koska Rest rajapinnan tarjoama Json data on hakasulkujen sisällä jouduin käyttämään `bytes` kirjastoa hakasulkujen poistamiseksi.

Ohjelmaa tehdessä törmäsin haasteeseen tulostaa Json datan sisältä objekteja, joita en saanut yrityksistä huolimatta tulostettua haluamallani tavalla.

Apuna ohjelman json osan teossa käytin [tätä ohjetta.](https://gobyexample.com/json)

Apuna ohjelman `bytes` kirjaston käytössä käytin [tätä ohjetta.](https://www.geeksforgeeks.org/how-to-trim-a-slice-of-bytes-in-golang/)

[Lähdekoodi](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day3/json.go)

### Ohjelman käyttö

Esimerkki:

    ./day3 --paycycle Tammikuu

Josta palautuu testattavan sivun http status:

```
Palkkakauden nimi: Tammikuu
Työpäiviä tällä kaudella: 3
map[startingtime:8.00 endingtime:15.25 owner:user1 id:5 date:1.1.2019 workhours:6.5]
map[workhours:5.5 startingtime:7.00 endingtime:17.00 owner:user1 id:6 date:2.1.2019]
map[endingtime:16.00 owner:user1 id:8 date:1.1.2019 workhours:8 startingtime:8.00]

```

### Optiot (flagit)

```
Usage of ./day3:
  -paycycle string
    	Paycycle name (default "Joulukuu")
```

# Lähteet. Katso, että olet viitannut kaikissa tehtävävastauksissasi kaikkiin lähteisiin, joita olet käyttänyt. Kurssiin, tehtävänantoihin, käyttämääsi materiaaliin (GoByExample tms), muiden koodeihin, StackOverflown vastauksiin ja kaikkiin muihinkin lähteisiin.

Tarkistin kurssin kaikki kotitehtävä raporttini ja mielestäni olen viitannut näissä kaikkiin lähteisiin.

# Kaikki tai jotain. Valitse aihe tai pari vaihtoehtoa kurssin loppuprojektille. Palastele: mieti, miten voit ensin toteuttaa pienen, toimivan kokonaisuuden. Ja sen päälle vaihtoehtoja laajentaa.

## 1. Idea

Sää ohjelma. Ohjelma hakee jostain (näitä ilmeisesti on useita) avoimesta rajapinnasta Json muotoista dataa ja tulostaa sitä ohjelman käyttäjälle.

Laajennetaan antamalla parametrinä kaupunki / maa, josta ohjelma hakee sään.

Laajennetaan ohjelmaa tallentamaan haettu data (sää + datan kellonaika, jolloin data on haettu) tekstitiedoston uudelle riville.

Laajennetaan ohjelmaa tallentamaan dataa tiedostoon CSV muodolla, jatkokäsittelyn helpottamiseksi.