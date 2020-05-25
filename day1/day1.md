Tehtävän teossa on käytetty apuna GO:n [dokumentaatiota.](https://golang.org/pkg/flag/)

# Kirjoita itse keksimäsi ohjelma, joka lukee käyttäjältä syötteen (flag) ja tulostaa käyttäjälle tekstiä.

Ohjelmani laskee käytäjän antamien syötteiden perusteella ajoneuvon polttoaineen kulutuksen, sekä polttoainekustannuksen per km/mile.

[Lähdekoodi](https://github.com/nikke845/Go-Programming-Course-2020/blob/master/day1/mpg.go)

## Ohjelman käyttö

Esimerkki:

    ./day1 --nationality European --distance 600 --fuel-used 33 --fuel-cost 1.3

Josta palautuu syötteenä tulos:

    Average consumption is 5.5 L/100km and cost per km is 0.0715 €

## Optiot (flagit)

```
Usage of ./day1:
  -distance float
    	Distance driven (default 600)
  -fuel-cost float
    	Cost of fuel per gallon or litre (default 1.3)
  -fuel-used float
    	Amount of fuel refueled (default 33)
  -nationality string
    	American or European style units i.e. gallons vs. litres (default "European")
```