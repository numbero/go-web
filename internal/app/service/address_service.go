package service

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type City struct {
	Name       string `yaml:"name"`
	PostalCode string `yaml:"postal_code"`
}

type Province struct {
	Name   string `yaml:"name"`
	Code   string `yaml:"code"`
	Cities []City `yaml:"cities"`
}

type Country struct {
	Name      string     `yaml:"name"`
	Code      string     `yaml:"code"`
	Provinces []Province `yaml:"provinces"`
}

type AddressData struct {
	Country Country `yaml:"country"`
}

func LoadAddresses() {
	data, err := os.ReadFile("config/address.yaml")
	if err != nil {
		log.Fatalf("读取 YAML 失败: %v", err)
	}

	var addr AddressData
	if err := yaml.Unmarshal(data, &addr); err != nil {
		log.Fatalf("解析 YAML 失败: %v", err)
	}

	// 打印结构化地址
	fmt.Printf("🌍 国家: %s (%s)\n", addr.Country.Name, addr.Country.Code)
	for _, province := range addr.Country.Provinces {
		fmt.Printf("  📍 省份: %s (%s)\n", province.Name, province.Code)
		for _, city := range province.Cities {
			fmt.Printf("    🏙️ 城市: %s, 邮编: %s\n", city.Name, city.PostalCode)
		}
	}
}
