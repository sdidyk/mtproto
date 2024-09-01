package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type nametype struct {
	name      string
	_type     string
	flag_name string
	flag_bit  int
}

type constuctor struct {
	id        string
	predicate string
	params    []nametype
	_type     string
}

func normalize(s string) string {
	x := []byte(s)
	for i, r := range x {
		if r == '.' {
			x[i] = '_'
		}
	}
	y := string(x)
	if y == "type" {
		return "_type"
	}
	if y == "range" {
		return "_range"
	}
	if y == "default" {
		return "_default"
	}
	return y
}

func main() {
	var err error
	var parsed interface{}

	// read json file from stdin
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	// parse json
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	err = d.Decode(&parsed)
	if err != nil {
		fmt.Println(err)
		return
	}

	// process constructors
	_order := make([]string, 0, 1000)
	_cons := make(map[string]constuctor, 1000)
	_types := make(map[string][]string, 1000)

	parsefunc := func(data []interface{}, kind string) {
		for _, data := range data {
			data := data.(map[string]interface{})

			// id
			idx, err := strconv.Atoi(data["id"].(string))
			if err != nil {
				fmt.Println(err)
				return
			}
			_id := fmt.Sprintf("0x%08x", uint32(idx))

			// predicate
			_predicate := normalize(data[kind].(string))

			if _predicate == "vector" {
				continue
			}

			// params
			_params := make([]nametype, 0, 16)
			params := data["params"].([]interface{})
			for _, params := range params {
				params := params.(map[string]interface{})
				_name := normalize(params["name"].(string))
				_type := normalize(params["type"].(string))

				var _inner_type string
				var _flag_name string
				var _flag_num int
				var _flag_bit int

				n, _ := fmt.Sscanf(_type, "flags_%d?%s", &_flag_bit, &_inner_type)
				if n == 2 {
					_flag_name = "flags"
					_type = _inner_type
				} else {
					n, _ := fmt.Sscanf(_type, "flags%d_%d?%s", &_flag_num, &_flag_bit, &_inner_type)
					if n == 3 {
						_flag_name = fmt.Sprintf("flags%d", _flag_num)
						_type = _inner_type
					} else {
						_flag_name = ""
					}
				}

				_params = append(_params, nametype{_name, _type, _flag_name, _flag_bit})
			}

			// type
			_type := normalize(data["type"].(string))

			_order = append(_order, _predicate)
			_cons[_predicate] = constuctor{_id, _predicate, _params, _type}
			if kind == "predicate" {
				_types[_type] = append(_types[_type], _predicate)
			}
		}
	}
	parsefunc(parsed.(map[string]interface{})["constructors"].([]interface{}), "predicate")
	parsefunc(parsed.(map[string]interface{})["methods"].([]interface{}), "method")

	// constants
	fmt.Print("package mtproto\nimport \"fmt\"\nconst (\n")
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("crc_%s = %s\n", c.predicate, c.id)
	}
	fmt.Print(")\n\n")

	// type structs
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("type TL_%s struct {\n", c.predicate)
		for _, t := range c.params {
			fmt.Printf("%s\t", t.name)
			s_comment := make([]string, 0)

			if t.flag_name != "" {
				s_comment = append(s_comment, fmt.Sprintf("(bit %s.%d)", t.flag_name, t.flag_bit))
			}

			switch t._type {
			case "true":
				fmt.Print("bool")
			case "int":
				fmt.Print("int32")
			case "long":
				fmt.Print("int64")
			case "string":
				fmt.Print("string")
			case "double":
				fmt.Print("float64")
			case "bytes":
				fmt.Print("[]byte")
			case "Vector<int>":
				fmt.Print("[]int32")
			case "Vector<long>":
				fmt.Print("[]int64")
			case "Vector<string>":
				fmt.Print("[]string")
			case "Vector<double>":
				fmt.Print("[]float64")
			case "!X":
				fmt.Print("TL")
			case "#":
				fmt.Print("int32")
			default:
				var inner1 string
				n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner1)
				if n == 1 {
					fmt.Print("[]TL")
					s_comment = append(s_comment, inner1[:len(inner1)-1])
				} else {
					fmt.Print("TL")
					s_comment = append(s_comment, t._type)
				}
			}
			if len(s_comment) > 0 {
				fmt.Printf(" // %s", strings.Join(s_comment[:], " | "))
			}
			fmt.Print("\n")
		}
		fmt.Print("}\n\n")
	}

	// encode funcs
	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("func (e TL_%s) encode() []byte {\n", c.predicate)
		fmt.Print("x := NewEncodeBuf(512)\n")
		fmt.Printf("x.UInt(crc_%s)\n", c.predicate)
		for _, t := range c.params {
			if t.flag_name != "" && t._type != "true" {
				fmt.Printf("if e.%s & (1<<%d) != 0 {\n", t.flag_name, t.flag_bit)
			}

			switch t._type {
			case "true":
				// void constructor
			case "int", "#":
				fmt.Printf("x.Int(e.%s)\n", t.name)
			case "long":
				fmt.Printf("x.Long(e.%s)\n", t.name)
			case "string":
				fmt.Printf("x.String(e.%s)\n", t.name)
			case "double":
				fmt.Printf("x.Double(e.%s)\n", t.name)
			case "bytes":
				fmt.Printf("x.StringBytes(e.%s)\n", t.name)
			case "Vector<int>":
				fmt.Printf("x.VectorInt(e.%s)\n", t.name)
			case "Vector<long>":
				fmt.Printf("x.VectorLong(e.%s)\n", t.name)
			case "Vector<string>":
				fmt.Printf("x.VectorString(e.%s)\n", t.name)
			case "!X":
				fmt.Printf("x.Bytes(e.%s.encode())\n", t.name)
			case "Vector<double>":
				panic(fmt.Sprintf("Unsupported %s", t._type))
			default:
				var inner string
				n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
				if n == 1 {
					fmt.Printf("x.Vector(e.%s)\n", t.name)
				} else {
					fmt.Printf("x.Bytes(e.%s.encode())\n", t.name)
				}
			}

			if t.flag_name != "" && t._type != "true" {
				fmt.Print("}\n")
			}
		}
		fmt.Print("return x.buf\n")
		fmt.Print("}\n\n")

	}

	// decode funcs
	fmt.Println(`
func (m *DecodeBuf) ObjectGenerated(constructor uint32) (g TL) {
	switch constructor {`)

	for _, key := range _order {
		c := _cons[key]
		fmt.Printf("case crc_%s:\n", c.predicate)

		fmt.Printf("r := TL_%s{}\n", c.predicate)
		for _, t := range c.params {
			if t.flag_name != "" {
				fmt.Printf("if r.%s & (1<<%d) != 0 {\n", t.flag_name, t.flag_bit)
			}

			fmt.Printf("r.%s = ", t.name)
			switch t._type {
			case "true":
				fmt.Print("true\n")
			case "int", "#":
				fmt.Print("m.Int()\n")
			case "long":
				fmt.Print("m.Long()\n")
			case "string":
				fmt.Print("m.String()\n")
			case "double":
				fmt.Print("m.Double()\n")
			case "bytes":
				fmt.Print("m.StringBytes()\n")
			case "Vector<int>":
				fmt.Print("m.VectorInt()\n")
			case "Vector<long>":
				fmt.Print("m.VectorLong()\n")
			case "Vector<string>":
				fmt.Print("m.VectorString()\n")
			case "!X":
				fmt.Print("m.Object()\n")
			case "Vector<double>":
				panic(fmt.Sprintf("Unsupported %s", t._type))
			default:
				var inner string
				n, _ := fmt.Sscanf(t._type, "Vector<%s", &inner)
				if n == 1 {
					fmt.Print("m.Vector()\n")
				} else {
					fmt.Print("m.Object()\n")
				}
			}

			if t.flag_name != "" {
				fmt.Print("}\n")
			}
		}

		fmt.Print("return r\n")
	}

	fmt.Println(`
	default:
		m.err = fmt.Errorf("Unknown constructor: \u002508x", constructor)
		return nil

	}

}`)

}
