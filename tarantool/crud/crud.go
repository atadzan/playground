package customCrud

import (
	"github.com/tarantool/go-tarantool"
	"log"
)

type Tarantool struct {
	conn *tarantool.Connection
}

func NewTarantool(conn *tarantool.Connection) *Tarantool {
	return &Tarantool{
		conn: conn,
	}
}

func (t *Tarantool) InsertData(space string, tuple []interface{}) (*tarantool.Response, error) {
	resp, err := t.conn.Insert(space, tuple)
	if err != nil {
		log.Println("can't insert input data. Error:", err.Error())
		return nil, err
	}
	return resp, nil
}

func (t *Tarantool) GetData(space string, index interface{}, offset, limit uint32, key interface{}) (*tarantool.Response, error) {
	resp, err := t.conn.Select(space, index, offset, limit, tarantool.IterEq, key)
	if err != nil {
		log.Println("can't select data. Error: ", err.Error())
		return nil, err
	}
	return resp, nil
}

func (t *Tarantool) UpdateData(space string, index, key, opts interface{}) (*tarantool.Response, error) {
	resp, err := t.conn.Update(space, index, key, opts)
	if err != nil {
		log.Println("can't update data. Error: ", err.Error())
		return nil, err
	}
	return resp, nil
}

func (t *Tarantool) ReplaceData(space string, tuple interface{}) (*tarantool.Response, error) {
	resp, err := t.conn.Replace(space, tuple)
	if err != nil {
		log.Println("can't replace data. Error: ", err.Error())
		return nil, err
	}
	return resp, nil

}

func (t *Tarantool) UpsertData(space string, tuple, opts interface{}) (*tarantool.Response, error) {
	resp, err := t.conn.Upsert(space, tuple, opts)
	if err != nil {
		log.Println("can't upsert. Error:", err.Error())
		return nil, err
	}
	return resp, nil
}

func (t *Tarantool) DeleteData(space string, index, key interface{}) (*tarantool.Response, error) {
	resp, err := t.conn.Delete("tester", "primary", []interface{}{3})
	if err != nil {
		log.Println("can't delete. Error: ", err.Error())
		return nil, err
	}
	return resp, nil
}
