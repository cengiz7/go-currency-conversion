package exchange

type ExchangeRate struct {
	Source string
	Target 	 string
	Rate float32
}

type Exchanges struct {
	Rates []ExchangeRate
}

func (exchange *Exchanges) CreateOrUpdateRate(rate ExchangeRate) {
	if exRate, found := exchange.FindExchangeRate(rate.Source, rate.Target); found {
		exRate.UpdateRate(rate)
	} else {
		exchange.NewRate(rate)
	}
}


func (exchange *Exchanges) FindExchangeRate(source, target string) (*ExchangeRate, bool){
	for idx, exRate := range exchange.Rates {
        if exRate.Source == source && exRate.Target == target {
            return (&exchange.Rates[idx]), true
        }
    }
    return &ExchangeRate{}, false
}

func (exchange *Exchanges) NewRate(rate ExchangeRate) {
	exchange.Rates = append(exchange.Rates, rate)
}

func (exchangeRate *ExchangeRate) UpdateRate(rate ExchangeRate) {
	exchangeRate.Rate = rate.Rate
}