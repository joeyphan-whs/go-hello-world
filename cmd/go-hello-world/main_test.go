package main

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

// 50% of the time, this test will pass
var _ = DescribeTable("Fake Test",
	func() {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		Expect(r1.Intn(2)).To(Equal(1))
	},
	Entry("No Resources"),
)
