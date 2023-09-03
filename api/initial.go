package api

import (
	"fmt"
	"strings"
	"sync"

	res "example.com/assignment/reservation/model"
	"github.com/spf13/viper"
)

var lock = &sync.Mutex{}

type singleton map[int]string

var tableInstant singleton
var seatPerTable int
var duplicateError = "The table instant is already initialized."

// initial table as defined
func InitTable(num int) res.Response {

	// Create singleton table instant
	if tableInstant == nil {
		lock.Lock()
		defer lock.Unlock()
		if tableInstant == nil {
			fmt.Println("Creating single instance now.")
			tableInstant = make(singleton, num)

			if tableInstant == nil {
				return res.Response{IsSuccess: false, Message: "Failed to create table instant"}
			}

		} else {
			fmt.Println(duplicateError)
			return res.Response{IsSuccess: false, Message: duplicateError}
		}
	} else {
		fmt.Println(duplicateError)
		return res.Response{IsSuccess: false, Message: duplicateError}
	}

	// Assign value
	for i := 0; i < num; i++ {
		tableInstant[i+1] = ""
	}

	// Read config data
	viper.SetConfigName("init")     // ชื่อ config file
	viper.AddConfigPath("./config") // ระบุ path ของ config file
	viper.AutomaticEnv()            // อ่าน value จาก ENV variable
	// แปลง _ underscore ใน env เป็น . dot notation ใน viper
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// อ่าน config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	value, ok := viper.Get("table.seatPerTabel").(int)
	seatPerTable = value

	if !ok {
		panic(fmt.Errorf("fatal error read data: table.seatPerTabel \n"))
	}

	//fmt.Printf("Init Table %d ", num)
	//bs, _ := json.Marshal(tableInstant)
	//fmt.Println(string(bs))
	return res.Response{IsSuccess: true}
}
