package main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"math/rand"
	"testing"
)

type Tuple struct {
	_msgpack struct{} `msgpack:",asArray"`
	Key      string
	Value    int32
}

func BenchmarkSetRandomTntParallel(b *testing.B) {
	opts := tarantool.Opts{
		User: "admin",
	}
	pconn2, err := tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		b.Fatal(err)
	}
	b.RunParallel(func(pb *testing.PB) {
		var tuple Tuple
		for pb.Next() {
			tuple.Key = fmt.Sprintf("bench-%d", rand.Int31())
			tuple.Value = rand.Int31()
			_, err := pconn2.Replace("kv", tuple)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
