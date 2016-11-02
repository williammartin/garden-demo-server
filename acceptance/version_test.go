package acceptance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = FDescribe("Version", func() {

	var (
		session *gexec.Session
	)

	BeforeEach(func() {
		session = execServer("version")
	})

	It("exits with a zero exit code", func() {
		Eventually(session).Should(gexec.Exit(0))
	})

	It("reports the semver", func() {
		Eventually(session.Out).Should(gbytes.Say("0.0.1"))
	})
})
