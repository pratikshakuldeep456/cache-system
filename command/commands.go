package command

import (
	"bufio"
	"fmt"
	"os"
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

	limit := 10
	fmt.Println("input is", input, "and limit is ", limit)

	//cache limit

	fmt.Println("enter key:value to store in cache")

	keyValue, err := reader.ReadString('\n')
	keyValue = keyValue[:(len(keyValue) - 1)]

	if err != nil {
		fmt.Println(err)
		return err
	}

	parts := strings.Split(keyValue, ":")
	fmt.Println(parts)
	// write
	//get
	//implement cache limit

	myMap := make(map[string]int)
	myMap[parts[0]] = myMap[parts[1]]

	//compare map size with limit

	if len(myMap) > limit {
		//call choosen algo of cache policy

		//lfu lru howw??
	}

	return nil
}
