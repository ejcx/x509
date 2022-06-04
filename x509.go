package main

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "x509",
		Short: "Print information about the x509 cert of a url",
		Run: func(cmd *cobra.Command, args []string) {
			err := Run(cmd, args)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

type CoolThings struct {
	IsCA                        bool
	Version                     int
	SerialNumber                *big.Int
	Issuer                      string
	Subject                     string
	NotBefore, NotAfter         time.Time // Validity bounds.
	KeyUsage                    x509.KeyUsage
	DNSNames                    []string
	EmailAddresses              []string
	IPAddresses                 []net.IP
	URIs                        []*url.URL
	PermittedDNSDomainsCritical bool
	PermittedDNSDomains         []string
	ExcludedDNSDomains          []string
	PermittedIPRanges           []*net.IPNet
	ExcludedIPRanges            []*net.IPNet
	PermittedEmailAddresses     []string
	ExcludedEmailAddresses      []string
	PermittedURIDomains         []string
	ExcludedURIDomains          []string
}

func Run(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Usage: x509 <url>")
	}

	u, err := url.Parse(args[0])
	if err != nil {
		return err
	}

	// Set to default 443 if none specified.
	// Required by TLS Dial
	host := u.Host
	if !strings.Contains(u.Host, ":") {
		host += ":443"
	}
	conn, err := tls.Dial("tcp", host, &tls.Config{})
	if err != nil {
		return err
	}
	defer conn.Close()

	// Take only the last peer certificate
	connState := conn.ConnectionState()
	peer := connState.PeerCertificates[0]

	c := &CoolThings{
		DNSNames:                    peer.DNSNames,
		IsCA:                        peer.IsCA,
		Version:                     peer.Version,
		SerialNumber:                peer.SerialNumber,
		Issuer:                      StringPkixName(peer.Issuer),
		Subject:                     StringPkixName(peer.Subject),
		NotBefore:                   peer.NotBefore,
		NotAfter:                    peer.NotAfter,
		KeyUsage:                    peer.KeyUsage,
		EmailAddresses:              peer.EmailAddresses,
		IPAddresses:                 peer.IPAddresses,
		URIs:                        peer.URIs,
		PermittedDNSDomainsCritical: peer.PermittedDNSDomainsCritical,
		ExcludedDNSDomains:          peer.ExcludedDNSDomains,
		PermittedIPRanges:           peer.PermittedIPRanges,
		ExcludedIPRanges:            peer.ExcludedIPRanges,
		PermittedEmailAddresses:     peer.PermittedEmailAddresses,
		ExcludedEmailAddresses:      peer.ExcludedEmailAddresses,
		PermittedURIDomains:         peer.PermittedURIDomains,
		ExcludedURIDomains:          peer.ExcludedURIDomains,
	}
	// Overwrite Peers to be a shorter list.
	buf, err := json.MarshalIndent(c, " ", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(buf))
	return nil
}

func StringPkixName(n pkix.Name) string {
	s := ""
	s += strings.Join(n.Country, " ")
	s += strings.Join(n.Organization, " ")
	s += strings.Join(n.OrganizationalUnit, " ")
	s += strings.Join(n.StreetAddress, " ")
	s += strings.Join(n.PostalCode, " ")
	s += " " + n.SerialNumber
	s += " " + n.CommonName
	return s
}
func main() {
	RootCmd.Execute()
}
