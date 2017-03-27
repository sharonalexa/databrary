// this is until https://github.com/lib/pq/pull/390 gets merged in

package netaddr

import (
	"bytes"
	"database/sql"
	"net"
	"testing"

	_ "github.com/lib/pq"
	"os"
)

type Fatalistic interface {
	Fatal(args ...interface{})
}

func openTestConn(t Fatalistic) *sql.DB {
	datname := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("PGSSLMODE")

	if datname == "" {
		os.Setenv("PGDATABASE", "pqgotest")
	}

	if sslmode == "" {
		os.Setenv("PGSSLMODE", "disable")
	}

	conn, err := sql.Open("postgres", "")
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func TestInet(t *testing.T) {
	db := openTestConn(t)
	defer db.Close()

	inet := Inet{}

	// Test scanning NULL values
	err := db.QueryRow("SELECT NULL::inet").Scan(&inet)
	if err != nil {
		t.Fatal(err)
	}
	if inet.Valid {
		t.Fatalf("expected null result")
	}

	// Test setting NULL values
	err = db.QueryRow("SELECT $1::inet", inet).Scan(&inet)
	if err != nil {
		t.Fatalf("re-query null value failed: %s", err.Error())
	}
	if inet.Valid {
		t.Fatalf("expected null result")
	}

	// test encoding in query params, then decoding during Scan
	testBidirectional := func(i Inet, label string) {
		err = db.QueryRow("SELECT $1::inet", i).Scan(&inet)
		if err != nil {
			t.Fatalf("re-query %s inet failed: %s", label, err.Error())
		}
		if !inet.Valid {
			t.Fatalf("expected non-null value, got null for %s", label)
		}
		if bytes.Compare(i.Inet, inet.Inet) != 0 {
			t.Fatalf("expected IP addresses to match, but did not for %s - %s %s", label, inet.Inet.String(), inet.Inet.String())
		}
	}

	testBidirectional(Inet{Inet: net.ParseIP("192.168.0.1"), Valid: true}, "Simple IPv4")
	testBidirectional(Inet{Inet: net.ParseIP("::1"), Valid: true}, "Loopback IPv6")
	testBidirectional(Inet{Inet: net.ParseIP("abcd:2345::"), Valid: true}, "Loopback IPv6")

	// Bad argument
	inet = Inet{}
	err = inet.Scan(456)
	if err == nil {
		t.Fatal("Expected error for non-byte[] argument to Scan")
	}

	inet = Inet{}
	err = inet.Scan([]byte(""))
	if err != nil {
		t.Fatalf("Unexpected error for empty string - %s", err.Error())
	}
	if inet.Valid {
		t.Fatalf("Unexpected not null for empty/non-IP string string")
	}
}
