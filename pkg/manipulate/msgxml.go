package manipulate

import (
    "encoding/xml"
    "fmt"
)

type Plant struct{
	XMLName xml.Name `xml:"plant"`
    Id      int      `xml:"id,attr"`   //  attr for mask field id to attribute tag <xx id=1>
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
        p.Id, p.Name, p.Origin)
}

func ExXml(){
	coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}

	// xml marshall need indent
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }
    fmt.Println(p)
    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`     // can setup hierachy of xml like this
    }
    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}
    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
}