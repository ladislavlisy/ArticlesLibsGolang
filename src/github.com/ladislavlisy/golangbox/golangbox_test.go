package golangbox_test

import (
	. "github.com/ladislavlisy/golangbox"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LibraryCode", func() {

	Describe("Configuration", func() {
		It("should Should_Return_53_For_Articles_Length", func() {
			articlesAll := ConfigureArticles()
			Expect(len(articlesAll)).Should((Equal(53)))
		})
	})
})
