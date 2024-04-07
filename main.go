package main

/* copy from https://github.com/git9527/aliyun-cdn-https-cert-updater
CHAGES

- change envirment reader to arguments
*/

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/denverdino/aliyungo/cdn"
)

func main() {
	var ACCESS_KEY_ID = flag.String("a", "", "Aliyun access key ID")
	var ACCESS_KEY_SECRET = flag.String("s", "", "Aliyun access key secret")
	var domain = flag.String("d", "", "Domain name")
	var certPath = flag.String("c", "", "Path to certificate file")
	var keyPath = flag.String("k", "", "Path to private key file")

	flag.Parse()

	if *ACCESS_KEY_ID == "" || *ACCESS_KEY_SECRET == "" || *domain == "" || *certPath == "" || *keyPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	var cert []byte
	var key []byte
	var err error
	if cert, err = os.ReadFile(*certPath); err != nil {
		panic(err)
	}
	if key, err = os.ReadFile(*keyPath); err != nil {
		panic(err)
	}

	// 生成一个不重复的证书名称
	var certName = "cert" + time.Now().Format("20060102150405")

	// 记录日志
	fmt.Println("time: ", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("update cert for domain: ", *domain)
	fmt.Println("certName: ", certName)

	client := cdn.NewClient(*ACCESS_KEY_ID, *ACCESS_KEY_SECRET)
	res, err := client.SetDomainServerCertificate(cdn.CertificateRequest{
		DomainName:              *domain,
		CertName:                certName,
		ServerCertificateStatus: "on",
		ServerCertificate:       string(cert),
		PrivateKey:              string(key),
	})

	fmt.Printf("res: %v, err: %v\n", res, err)
}
