package acceptance_test

import (
	"net/url"

	"github.com/gorilla/websocket"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Serve", func() {

	var (
		session *gexec.Session
		args    []string
	)

	BeforeEach(func() {
		args = []string{"serve"}
	})

	JustBeforeEach(func() {
		session = execServer(args...)
	})

	AfterEach(func() {
		session.Terminate()
		Eventually(session).Should(gexec.Exit())
	})

	It("log that it is serving requests on default port", func() {
		Eventually(session.Out).Should(gbytes.Say("Serving requests on port 8080"))
	})

	It("responds to the opening of a websocket", func() {
		serverURL := url.URL{Scheme: "ws", Host: "localhost:8080"}
		Eventually(func() error { return dial(serverURL) }).Should(Succeed())
	})

	Context("when provided a port to listen on", func() {
		BeforeEach(func() {
			args = append(args, "-p", "9080")
		})

		It("log that it is serving requests on the provided port", func() {
			Eventually(session.Out).Should(gbytes.Say("Serving requests on port 9080"))
		})

		It("responds to the opening of a websocket", func() {
			serverURL := url.URL{Scheme: "ws", Host: "localhost:9080"}
			Eventually(func() error { return dial(serverURL) }).Should(Succeed())
		})
	})
})

func dial(url url.URL) error {
	client, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		return err
	}
	defer client.Close()
	return nil
}
