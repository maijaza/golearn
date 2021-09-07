package manipulate

import (
	_ "os"
	"fmt"
    "net/url"
	"net"
	"crypto/sha1"
	b64 "encoding/base64"
)

func ExPanic(){
	// panic("a problem")
     
	// f, err := os.Create("defer.txt")
    // if err != nil {
    //     panic(err)
    // }

	// fmt.Fprintln(f, "data")

	// defer f.Close()
}

func UrlParser(){
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme)

    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}

func ExSha1(){
	s := "sha1 this string"

    h := sha1.New()

    h.Write([]byte(s))

    bs := h.Sum(nil)

    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}

func ExBase64(){

    data := "abc123!?$*&()'-=@~"

    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))	
}