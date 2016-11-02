package acceptance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = FDescribe("Help", func() {

	var (
		session *gexec.Session
	)

	BeforeEach(func() {
		session = execServer("--help")
	})

	It("exits with a non-zero exit code", func() {
		Eventually(session).Should(gexec.Exit(1))
	})

	It("reports the usage", func() {
		Eventually(session.Err).Should(gbytes.Say("Usage:"))
	})
})
