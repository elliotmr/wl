package main

import (
	"encoding/xml"
	"github.com/pkg/errors"
	"github.com/serenize/snaker"
	"strings"
	"text/template"
	"bufio"
	"bytes"
	"fmt"
)

type Description struct {
	Summary string `xml:"summary,attr"`
	Text string `xml:",chardata"`
}

type Request struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Since string `xml:"since,attr"`
	Description *Description `xml:"description"`
	Args []*Arg `xml:"arg"`
}

type Event struct {
	Name string `xml:"name,attr"`
	Since string `xml:"since,attr"`
	Description *Description `xml:"description"`
	Args []*Arg `xml:"arg"`
}

type Enum struct {
	Name string `xml:"name,attr"`
	Since string `xml:"since,attr"`
	Bitfield string `xml:"bitfield,attr"`
	Description *Description `xml:"description"`
	Entries []*Entry `xml:"entry"`
}

type Arg struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Summary string `xml:"summary,attr"`
	Interface string `xml:"interface,attr"`
	AllowNull string `xml:"allow-null,attr"`
	Enum string `xml:"enum,attr"`
	Description *Description `xml:"description"`
}

type Entry struct {
	Name string `xml:"name,attr"`
	Value string `xml:"value,attr"`
	Summary string `xml:"summary,attr"`
	Since string `xml:"since,attr"`
	Description *Description `xml:"description"`
}

type Interface struct {
	Name string`xml:"name,attr"`
	Version string `xml:"version,attr"`
	Description *Description `xml:"description"`
	Requests []*Request `xml:"request"`
	Events []*Event `xml:"event"`
	Enums []*Enum `xml:"enum"`
}

type Protocol struct {
	Name string `xml:"name,attr"`
	Copyright string `xml:"copyright"`
	Description *Description `xml:"description"`
	Interfaces []*Interface `xml:"interface"`
}

func parse(raw []byte) (*Protocol, error) {
	p := &Protocol{}
	err := xml.Unmarshal(raw, p)
	return p, errors.Wrap(err, "unable to parse xml")
}

func genTemplate(templateText string) *template.Template {
	funcMap := template.FuncMap{
		"ifname": InterfaceName,
		"camel": snaker.SnakeToCamel,
		"desc_to_comment": DescriptionToComment,
		"req_sig": ReqSignature,
		"req_ret_sig": ReqReturnSignature,
		"req_ret": ReqReturn,
	}

	return template.Must(template.New("wl").Funcs(funcMap).Parse(templateText))

}

func InterfaceName(name string) string {
	name = strings.TrimLeft(name, "wl_")
	return snaker.SnakeToCamel(name)
}

func DescriptionToComment(desc string) string {
	buf := &bytes.Buffer{}
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(desc)))
	for scanner.Scan() {
			buf.WriteString("// ")
			buf.Write(bytes.TrimSpace(scanner.Bytes()))
			buf.WriteString("\n")

	}
	return buf.String()
}

func ArgSignature(arg *Arg) string {
	name := snaker.SnakeToCamelLower(arg.Name)
	if name == "interface" {
		name = "iface"
	}
	buf := bytes.NewBufferString(name)
	buf.WriteString(" ")
	switch arg.Type {
	case "int":
		buf.WriteString("int32")
	case "uint", "fixed", "object":
		buf.WriteString("uint32")
	case "string":
		buf.WriteString("string")
	case "array":
		buf.WriteString("[]byte")
	default:
		return ""
	}
	return buf.String()
}

func ReqSignature(args []*Arg) string {
	argSigs := make([]string, 0)
	for _, arg := range args {
		newSig := ArgSignature(arg)
		if newSig != "" {
			argSigs = append(argSigs, newSig)
		}
	}
	return strings.Join(argSigs, ", ")
}

func ReqReturnSignature(args []*Arg) string {
	newTypeInterface := ""
	for _, arg := range args {
		if arg.Type == "new_id" {
			newTypeInterface = arg.Interface
			break
		}
	}
	if newTypeInterface == "" {
		return "error"
	}
	return fmt.Sprintf("(*%s, error)", InterfaceName(newTypeInterface))
}

func ReqReturn(args []*Arg) string {
	newTypeInterface := ""
	for _, arg := range args {
		if arg.Type == "new_id" {
			newTypeInterface = arg.Interface
			break
		}
	}
	if newTypeInterface == "" {
		return "nil"
	}
	return "nil, nil"
}

func main() {

}
