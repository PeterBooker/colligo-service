// +build !dev

package server

import (
	"github.com/mholt/certmagic"
)

func (s *Server) startHTTP() {

	/*
		magic := certmagic.New(certmagic.Config{
			Agreed: true,
			Email:  "mail@peterbooker.com",
			CA:     certmagic.LetsEncryptProductionCA,
		})

		err := magic.Manage([]string{"colligo.dev", "www.colligo.dev"})
		if err != nil {
			s.Logger.Fatalf("HTTP server failed to start: %s\n")
		}
	*/

	certmagic.Agreed = true
	certmagic.Email = "mail@peterbooker.com"
	certmagic.HTTPS([]string{"colligo.dev"}, s.Router)

}
