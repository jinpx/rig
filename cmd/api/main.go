package main

import (
	"fmt"

	"rig/pkg/utils/u_go"
	"rig/pkg/utils/u_snowflake"
)

func main() {
	fmt.Println("api start.")

	u_go.Go(func() {
		fmt.Println("123")
	})

	var count = 10
	for i := 0; i < count; i++ {
		fmt.Println(u_snowflake.SnowflakeIdStr())
	}

}
