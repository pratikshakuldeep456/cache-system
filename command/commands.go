package command

import (
	"bufio"
	"fmt"
	"os"
	"pratikshakuldeep456/cache-system/pkg"
	"strings"

	"github.com/urfave/cli/v2"
)

func StartCacheSystem(ctx *cli.Context) error {

	fmt.Println("starting the cache system")
	//choose cache eviction type
	// 1- lfu 2 - lru
	//user input
	fmt.Println(" choose cache eviction type ")
	fmt.Println("1 : LRU")
	fmt.Println("2 : LFU")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("input is", input)

	fmt.Println("enter key:value to store in cache")

	keyValue, err := reader.ReadString('\n')
	keyValue = keyValue[:(len(keyValue) - 1)]

	if err != nil {
		fmt.Println(err)
		return err
	}

	parts := strings.Split(keyValue, ":")
	fmt.Println(parts)

	var evictAlgo pkg.EvictionAlgo
	switch input {
	case "1":
		evictAlgo = &pkg.LRU{}
	case "2":
		evictAlgo = &pkg.LFU{}
	}
	cache := pkg.InitCache(evictAlgo)
	cache.AddCache(parts[1], parts[0])

	return nil
}
