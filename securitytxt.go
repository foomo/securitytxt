package securitytxt

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// TODO: http handler
// generate security txt once on startup
// zero external deps
// digital signing
// PGP key

type Date struct {
	time.Time
}

//func (t *Date) UnmarshalText(text []byte) error {
//	tt, err := time.Parse("2006-01-02", string(text))
//	if err != nil {
//		tt, err = time.Parse(time.RFC3339, string(text))
//	}
//	*t = Date{tt}
//	return err
//}

type config struct {
	Expires string
	Comment string

	Contact            string
	Acknowledgments    string
	Canonical          string
	Encryption         string
	Hiring             string
	PreferredLanguages []string
	Policy             string
	CSAF               string
}

func getStringsEnv(key string) []string {
	return strings.Split(os.Getenv(key), " ")
}

func Handler() http.HandlerFunc {

	cfg := config{
		Comment:            os.Getenv("COMMENT"),
		Expires:            os.Getenv("EXPIRES"),
		Contact:            os.Getenv("CONTACT"),
		Acknowledgments:    os.Getenv("ACKNOWLEDGMENT"),
		Canonical:          os.Getenv("CANONICAL"),
		Encryption:         os.Getenv("ENCRYPTION"),
		Hiring:             os.Getenv("HIRING"),
		PreferredLanguages: getStringsEnv("PREFERRED_LANGUAGES"),
		Policy:             os.Getenv("POLICY"),
		CSAF:               os.Getenv("CSAF"),
	}

	data, err := createSecurityTxt(cfg)
	if err != nil {
		fmt.Println("could not load required values:", err)
		os.Exit(1)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(data)
	}
}

func createSecurityTxt(c config) ([]byte, error) {

	if len(c.Contact) == 0 {
		return nil, errors.New("contact must be provided")
	}

	var out []string

	field := func(prefix string, body string) string {
		return fmt.Sprintf("%s %s", prefix, body)
	}

	if len(c.Comment) > 0 {
		out = append(out, field("#", c.Comment))
	}

	out = append(out, field("Contact:", c.Contact))
	out = append(out, fmt.Sprintf("Expires: %s", c.Expires))
	out = append(out, field("Encryption:", c.Encryption))
	out = append(out, field("Acknowledgments:", c.Acknowledgments))
	if len(c.PreferredLanguages) > 0 {
		out = append(out, fmt.Sprintf("Preferred-Languages: %s", strings.Join(c.PreferredLanguages, " ")))
	}
	out = append(out, field("Canonical:", c.Canonical))
	out = append(out, field("Policy:", c.Policy))
	out = append(out, field("Hiring:", c.Hiring))

	if len(c.CSAF) > 0 {
		out = append(out, field("CSAF:", c.CSAF))
	}

	return []byte(strings.Join(out, "\n")), nil
}
