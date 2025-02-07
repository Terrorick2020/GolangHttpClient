package main

import (
	"fmt"
	"log"
	"my-coincapio/coincapio"
	"time"
)

func main() {
	client, err := coincapio.NewClient( time.Second * 10 )
	if err != nil {
		log.Fatal(err)
	}

	var assets []coincapio.Asset
	if assets, err = client.GetAssets(); err != nil {
		log.Fatal(err)
	}

	for _, item := range assets {
		fmt.Println(item)
	}

	var asset coincapio.Asset
	if asset, err = client.GetAsset("bitcoin"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(asset)
}