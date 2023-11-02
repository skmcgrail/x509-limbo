package main

//go:generate go run github.com/atombender/go-jsonschema/cmd/gojsonschema@latest -v -p main -o schema.go ../../limbo-schema.json

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

type testcaseResult string

const (
	validationKindClient = "CLIENT"
	validationKindServer = "SERVER"

	resultFailure testcaseResult = "FAILURE"
	resultSuccess testcaseResult = "SUCCESS"
	resultSkipped testcaseResult = "SKIPPED"
)

type result struct {
	ID      string         `json:"id"`
	Result  testcaseResult `json:"actual_result"`
	Context string         `json:"context"`
}

type results struct {
	Version uint     `json:"version"`
	Harness string   `json:"harness"`
	Results []result `json:"results"`
}

func main() {
	testCasePath := flag.String("testcases", "../../limbo.json", "testcases")
	resultsPath := flag.String("results", "./results.json", "results")
	flag.Parse()

	testcases, err := loadTestcases(*testCasePath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded testcases from %s\n", *testCasePath)

	resultsFile, err := os.Create(*resultsPath)
	if err != nil {
		panic(err)
	}
	resultsEncoder := json.NewEncoder(resultsFile)

	var (
		conform, nonconform, skip int
		outputResults             results
	)
	for _, tc := range testcases.Testcases {
		fmt.Printf("Running test %s ... ", tc.Id)
		r, err := evaluateTestcase(tc)

		var context string
		if r != testcaseResult(tc.ExpectedResult.(string)) {
			if r != resultSkipped {
				fmt.Printf("NON-CONFORMANT\n\terr=%s\n", err)
				nonconform++
			} else {
				fmt.Println("SKIPPED")
				skip++
			}

			if err != nil {
				context = err.Error()
			}
		} else {
			fmt.Println("CONFORMANT")
			conform++
		}

		outputResults.Results = append(outputResults.Results, result{
			ID:      tc.Id,
			Context: context,
			Result:  r,
		})
	}

	outputResults.Version = 1
	outputResults.Harness = fmt.Sprintf("gocryptox509-%s", runtime.Version())
	resultsEncoder.Encode(outputResults)

	fmt.Printf("done! conformant/nonconformant/skipped/total %d/%d/%d/%d.\n", conform, nonconform, skip, len(testcases.Testcases))
}

func loadTestcases(path string) (testcases LimboSchemaJson, err error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(contents, &testcases)
	return
}

func concatPEMCerts(certs []string) []byte {
	var buf bytes.Buffer
	for _, cert := range certs {
		buf.WriteString(cert)
	}
	return buf.Bytes()
}

func evaluateTestcase(testcase Testcase) (testcaseResult, error) {
	_ = spew.Dump

	var ts time.Time
	if testcase.ValidationTime == nil {
		ts = time.Now()
	} else {
		var err error
		ts, err = time.Parse(time.RFC3339, *testcase.ValidationTime)

		if err != nil {
			fmt.Printf("%s\n", err)
			return resultSkipped, errors.Wrap(err, "unable to parse testcase time as RFC3339")
		}
	}

	// TODO: Support testcases that constrain signature algorthms.
	if len(testcase.SignatureAlgorithms) != 0 {
		return resultSkipped, fmt.Errorf("signature algorithm checks not supported yet")
	}

	// TODO: Support testcases that constrain key usages.
	if len(testcase.KeyUsage) != 0 {
		return resultSkipped, fmt.Errorf("key usage checks not supported yet")
	}

	var ekus []x509.ExtKeyUsage
	if len(testcase.ExtendedKeyUsage) != 0 {
		extKeyUsagesMap := map[KnownEKUs]x509.ExtKeyUsage{
			KnownEKUsAnyExtendedKeyUsage: x509.ExtKeyUsageAny,
			KnownEKUsClientAuth:          x509.ExtKeyUsageClientAuth,
			KnownEKUsCodeSigning:         x509.ExtKeyUsageCodeSigning,
			KnownEKUsEmailProtection:     x509.ExtKeyUsageEmailProtection,
			KnownEKUsOCSPSigning:         x509.ExtKeyUsageOCSPSigning,
			KnownEKUsServerAuth:          x509.ExtKeyUsageServerAuth,
			KnownEKUsTimeStamping:        x509.ExtKeyUsageTimeStamping,
		}

		for _, elem := range testcase.ExtendedKeyUsage {
			expected_eku := KnownEKUs(elem.(string))
			ekus = append(ekus, extKeyUsagesMap[expected_eku])
		}
	}

	switch testcase.ValidationKind {
	case validationKindClient:
		return resultSkipped, fmt.Errorf("unimplemented validationKindClient")
	case validationKindServer:
		var dnsName string
		if peerName, ok := testcase.ExpectedPeerName.(map[string]interface{}); ok {
			if peerName["kind"] != "DNS" {
				return resultSkipped, fmt.Errorf("non-DNS peer name checks not supported yet")
			}
			dnsName = peerName["value"].(string)
		}
		roots, intermediates := x509.NewCertPool(), x509.NewCertPool()
		roots.AppendCertsFromPEM(concatPEMCerts(testcase.TrustedCerts))
		intermediates.AppendCertsFromPEM(concatPEMCerts(testcase.UntrustedIntermediates))

		peerAsPEM, rest := pem.Decode([]byte(testcase.PeerCertificate))
		if peerAsPEM == nil || peerAsPEM.Type != "CERTIFICATE" {
			return resultFailure, fmt.Errorf("unexpected data, expected cert: %+#v", *peerAsPEM)
		} else if len(rest) > 0 {
			return resultFailure, fmt.Errorf("peer certificate has %d trailing bytes", len(rest))
		}

		peer, err := x509.ParseCertificate(peerAsPEM.Bytes)
		if err != nil {
			err = errors.Wrap(err, "unable to parse ASN1 certificate from PEM")
			return resultFailure, err
		}

		opts := x509.VerifyOptions{
			DNSName:       dnsName,
			Intermediates: intermediates,
			Roots:         roots,
			CurrentTime:   ts,
			KeyUsages:     ekus,
		}
		chain, err := peer.Verify(opts)
		_ = chain

		var (
			expected = testcaseResult(testcase.ExpectedResult.(string))
			actual   testcaseResult
		)
		if err != nil {
			actual = resultFailure
		} else {
			actual = resultSuccess
		}

		if expected != actual {
			if err == nil {
				err = errors.New("chain built")
			}
			err = errors.Wrap(err, "validation")
		}
		return actual, err
	}

	return resultSkipped, errors.New("no result returned from evaulation")
}
