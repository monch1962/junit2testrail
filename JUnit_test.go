package main

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"

	"golang.org/x/net/html/charset"
)

//func TestFailure(t *testing.T) {
//	t.Fail()
//}

func TestReadJUnitConvertToJson(t *testing.T) {
	jUnitSampleXML := `
	<?xml version="1.0" encoding="UTF-8"?>
<testsuites>
        <testsuite tests="3" failures="2" time="0.155" name="main.go">
                <properties>
                        <property name="go.version" value="go1.15"></property>
                </properties>
                <testcase classname="main.go" name="TestGenerated" time="0.150">
                        <failure message="Failed" type="">    main_test.go:105: Running 1 tests concurrently</failure>
                </testcase>
                <testcase classname="main.go" name="TestGenerated/0_-_TestA" time="0.100"></testcase>
                <testcase classname="main.go" name="TestGenerated/1_-_TestB" time="0.050">
                        <failure message="Failed" type="">    main_test.go:65: GET /posts/2 HTTP/1.1&#xA;        Host: jsonplaceholder.typicode.com&#xA;        Abc: def&#xA;        &#xA;        &#xA;    main_test.go:96: Response code: 200&#xA;    main_test.go:97: Response headers: map[Access-Control-Allow-Credentials:[true] Age:[13479] Cache-Control:[max-age=43200] Cf-Cache-Status:[HIT] Cf-Int-Pingora-Origin-Digest:[{&#34;ext_ip&#34;:&#34;162.158.106.53&#34;,&#34;ext_port&#34;:28584,&#34;upstream_rtt&#34;:99}] Cf-Ray:[5d4b37359b370925-SEA] Cf-Request-Id:[0542e0d57b000009257fad4200000001] Connection:[keep-alive] Content-Type:[application/json; charset=utf-8] Date:[Fri, 18 Sep 2020 12:54:27 GMT] Etag:[W/&#34;116-jnDuMpjju89+9j7e0BqkdFsVRjs&#34;] Expect-Ct:[max-age=604800, report-uri=&#34;https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct&#34;] Expires:[-1] Pragma:[no-cache] Server:[cloudflare] Set-Cookie:[__cfduid=d6c78ed7359e3971ba8a261b243b107bc1600433667; expires=Sun, 18-Oct-20 12:54:27 GMT; path=/; domain=.typicode.com; HttpOnly; SameSite=Lax] Vary:[Origin, Accept-Encoding] Via:[1.1 vegur] X-Content-Type-Options:[nosniff] X-Powered-By:[Express] X-Ratelimit-Limit:[1000] X-Ratelimit-Remaining:[999] X-Ratelimit-Reset:[1599866916]]&#xA;    main_test.go:99: Response headers (JSON): {&#xA;          &#34;Access-Control-Allow-Credentials&#34;: [&#xA;            &#34;true&#34;&#xA;          ],&#xA;          &#34;Age&#34;: [&#xA;            &#34;13479&#34;&#xA;          ],&#xA;          &#34;Cache-Control&#34;: [&#xA;            &#34;max-age=43200&#34;&#xA;          ],&#xA;          &#34;Cf-Cache-Status&#34;: [&#xA;            &#34;HIT&#34;&#xA;          ],&#xA;          &#34;Cf-Int-Pingora-Origin-Digest&#34;: [&#xA;            &#34;{\&#34;ext_ip\&#34;:\&#34;162.158.106.53\&#34;,\&#34;ext_port\&#34;:28584,\&#34;upstream_rtt\&#34;:99}&#34;&#xA;          ],&#xA;          &#34;Cf-Ray&#34;: [&#xA;            &#34;5d4b37359b370925-SEA&#34;&#xA;          ],&#xA;          &#34;Cf-Request-Id&#34;: [&#xA;            &#34;0542e0d57b000009257fad4200000001&#34;&#xA;          ],&#xA;          &#34;Connection&#34;: [&#xA;            &#34;keep-alive&#34;&#xA;          ],&#xA;          &#34;Content-Type&#34;: [&#xA;            &#34;application/json; charset=utf-8&#34;&#xA;          ],&#xA;          &#34;Date&#34;: [&#xA;            &#34;Fri, 18 Sep 2020 12:54:27 GMT&#34;&#xA;          ],&#xA;          &#34;Etag&#34;: [&#xA;            &#34;W/\&#34;116-jnDuMpjju89+9j7e0BqkdFsVRjs\&#34;&#34;&#xA;          ],&#xA;          &#34;Expect-Ct&#34;: [&#xA;            &#34;max-age=604800, report-uri=\&#34;https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct\&#34;&#34;&#xA;          ],&#xA;          &#34;Expires&#34;: [&#xA;            &#34;-1&#34;&#xA;          ],&#xA;          &#34;Pragma&#34;: [&#xA;            &#34;no-cache&#34;&#xA;          ],&#xA;          &#34;Server&#34;: [&#xA;            &#34;cloudflare&#34;&#xA;          ],&#xA;          &#34;Set-Cookie&#34;: [&#xA;            &#34;__cfduid=d6c78ed7359e3971ba8a261b243b107bc1600433667; expires=Sun, 18-Oct-20 12:54:27 GMT; path=/; domain=.typicode.com; HttpOnly; SameSite=Lax&#34;&#xA;          ],&#xA;          &#34;Vary&#34;: [&#xA;            &#34;Origin, Accept-Encoding&#34;&#xA;          ],&#xA;          &#34;Via&#34;: [&#xA;            &#34;1.1 vegur&#34;&#xA;          ],&#xA;          &#34;X-Content-Type-Options&#34;: [&#xA;            &#34;nosniff&#34;&#xA;          ],&#xA;          &#34;X-Powered-By&#34;: [&#xA;            &#34;Express&#34;&#xA;          ],&#xA;          &#34;X-Ratelimit-Limit&#34;: [&#xA;            &#34;1000&#34;&#xA;          ],&#xA;          &#34;X-Ratelimit-Remaining&#34;: [&#xA;            &#34;999&#34;&#xA;          ],&#xA;          &#34;X-Ratelimit-Reset&#34;: [&#xA;            &#34;1599866916&#34;&#xA;          ]&#xA;        }&#xA;    main_test.go:101: Response body: {&#xA;          &#34;userId&#34;: 1,&#xA;          &#34;id&#34;: 2,&#xA;          &#34;title&#34;: &#34;qui est esse&#34;,&#xA;          &#34;body&#34;: &#34;est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla&#34;&#xA;        }&#xA;    main_test.go:114: Expected response code 201, received response code 200</failure>
                </testcase>
        </testsuite>
</testsuites>
	`
	r := strings.NewReader(jUnitSampleXML)
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	dec.Strict = false

	var doc Testsuites
	if err := dec.Decode(&doc); err != nil {
		t.Fatal(err)
	}
	_, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(string(b))
}
