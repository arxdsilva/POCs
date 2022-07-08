package main

import (
	"fmt"
	"time"

	"github.com/90poe/service-chassis/postgres"
)

var query = `INSERT INTO companies (id, account_id, name, archived, parent_company_id, address1, address2, address3, address4, address5, city, state, postal_code, country_iso_code, email, phone, last_updated_by, last_updated_at, web_address1, web_address2, operates_worldwide, short_name)
VALUES ('2156c37d-ad7b-4b80-a348-60a85484af72', '8a396ee1-a16b-4cc4-8b54-d3de465b8fc8', 'TestCompanySeed', false, NULL, '123', '124', '', '', '', '444', '', '1234', 'UA', 'email@gmail.com', '', '', '0001-01-01 00:00:00.000000', '', '', false, 'shortName'),
('bf3004a6-1648-431b-92c3-82a8b5f2967e', '8a396ee1-a16b-4cc4-8b54-d3de465b8fc8', 'TestCompanySeedArchived', true, '2156c37d-ad7b-4b80-a348-60a85484af72', '123', '124', '', '', '', '444', '', '1234', 'UA', 'email@gmail.com', '', '', '0001-01-01 00:00:00.000000', '', '', false, 'shortName1')`

func main() {
	db, err := postgres.ConnectClientWithRetry(&postgres.ClientConfig{
		Host:     "0.0.0.0",
		Port:     5432,
		Db:       "cm",
		User:     "user",
		Password: "password",
		Ssl:      false,
	},
		5*time.Second,
		3,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Exec(query).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(err)
}
