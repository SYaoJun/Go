package main

import (
	"fmt"
	"log"

	"github.com/Connor1996/badger"
)

func main() {
	opts := badger.DefaultOptions
	opts.Dir = "./qq"
	opts.ValueDir = opts.Dir
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	updates := make(map[string]string)
	updates["hello"] = "world"
	updates["yaojun"] = "1995"
	// txn := db.NewTransaction(true)
	// badger写入数据
	for k, v := range updates {
		err := db.Update(func(txn *badger.Txn) error {
			err := txn.Set([]byte([]byte(k)), []byte(v))
			return err
		})
		if err != nil {
			fmt.Println("write failed")
		}
	}
	// badger读出数据
	_ = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("yaojun"))
		if err != nil {
			return err
		}
		val, err := item.Value()
		if err != nil {
			return err
		}
		fmt.Printf("The answer is: %s\n", val)
		return nil
	})

	defer db.Close()
}
