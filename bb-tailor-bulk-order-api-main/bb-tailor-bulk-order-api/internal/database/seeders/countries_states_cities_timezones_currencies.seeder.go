package seeders

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/data/data_mappers"
	"git.bopconsultancy.com/tejasc/bb-tailor-bulk-order-api/internal/database/models"
	"gorm.io/gorm"
)

func SeedCountriesStatesCitiesTimezonesCurrencies(db *gorm.DB) {
	file, err := os.Open("data/files/countries-cities-states-USA.json")
	if err != nil {
		log.Fatalf("failed to open countries, states, cities and timezones json file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read countries, states, cities and timezones json file: %v", err)
	}

	var countryData data_mappers.Country
	if err := json.Unmarshal(byteValue, &countryData); err != nil {
		log.Fatalf("failed to unmarshal countries, states, cities and timezones json: %v", err)
	}

	country := models.Country{
		Iso2:        countryData.Iso2,
		Iso3:        countryData.Iso3,
		Name:        countryData.Name,
		NumericCode: countryData.NumericCode,
		PhoneCode:   countryData.PhoneCode,
		Tld:         countryData.Tld,
		Native:      countryData.Native,
		Region:      countryData.Region,
		SubRegion:   countryData.Subregion,
		RegionId:    countryData.RegionID,
		SubRegionId: countryData.SubregionID,
		Nationality: countryData.Nationality,
	}

	if err := db.FirstOrCreate(&country, models.Country{NumericCode: countryData.NumericCode}).Error; err != nil {
		log.Fatalf("failed to seed country: %v", err)
	}

	for _, timezone := range countryData.Timezones {
		timezoneModel := models.Timezone{
			Name:          timezone.ZoneName,
			GmtOffset:     timezone.GmtOffset,
			GmtOffsetName: timezone.GmtOffsetName,
			Abbreviation:  timezone.Abbreviation,
			TzName:        timezone.TzName,
		}

		if err := db.FirstOrCreate(&timezoneModel, models.Timezone{Name: timezone.ZoneName}).Error; err != nil {
			log.Fatalf("failed to seed timezone %s: %v", timezone.ZoneName, err)
		}

		query := `
			INSERT INTO countries_timezones (country_id, timezone_id)
			VALUES (?, ?)
		`
		if err := db.Exec(query, country.ID, timezoneModel.ID).Error; err != nil {
			log.Fatalf("failed to insert into countries_timezones: %v", err)
		}
	}

	var states []models.State
	for _, stateData := range countryData.States {
		state := models.State{
			Name:      stateData.Name,
			StateCode: stateData.StateCode,
			CountryID: country.ID,
		}
		states = append(states, state)
	}

	if err := db.Create(&states).Error; err != nil {
		log.Fatalf("failed to seed states: %v", err)
	}

	var cities []models.City
	for _, state := range states {
		for _, stateData := range countryData.States {
			if state.Name == stateData.Name {
				for _, cityData := range stateData.Cities {
					city := models.City{
						Name:    cityData.Name,
						StateID: state.ID,
					}
					cities = append(cities, city)
				}
			}
		}
	}

	if err := db.Create(&cities).Error; err != nil {
		log.Fatalf("failed to seed cities: %v", err)
	}

	file_cur, err := os.Open("data/files/currencies.json")
	if err != nil {
		log.Fatalf("failed to open currencies json file: %v", err)
	}
	defer file.Close()

	byteValueCur, err := io.ReadAll(file_cur)
	if err != nil {
		log.Fatalf("failed to read currencies json file: %v", err)
	}

	var currencies map[string]data_mappers.Currency
	if err := json.Unmarshal(byteValueCur, &currencies); err != nil {
		log.Fatalf("failed to unmarshal currencies json: %v", err)
	}

	for _, currency := range currencies {
		currencyModel := models.Currency{
			Code:          currency.Code,
			Name:          currency.Name,
			NamePlural:    currency.NamePlural,
			Symbol:        currency.Symbol,
			SymbolNative:  currency.SymbolNative,
			DecimalDigits: currency.DecimalDigits,
			Rounding:      currency.Rounding,
		}

		if err := db.FirstOrCreate(&currencyModel, models.Currency{Code: currency.Code}).Error; err != nil {
			log.Fatalf("failed to seed currency %s: %v", currency.Name, err)
		}

		query := `
			INSERT INTO countries_currencies (country_id, currency_id)
			VALUES (?, ?)
		`
		if err := db.Exec(query, country.ID, currencyModel.ID).Error; err != nil {
			log.Fatalf("failed to insert into countries_currencies: %v", err)
		}
	}

	fmt.Println(
		"seeding of countries, states, cities, timezones and currencies completed successfully!",
	)
}
