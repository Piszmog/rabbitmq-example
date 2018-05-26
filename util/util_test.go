package util_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "go.uber.org/zap"
    "github.com/Piszmog/rabbitmq-example/util"
)

var _ = Describe("Util", func() {
    Context("initially", func() {
        var logger zap.Logger
        It("is empty", func() {
            Expect(logger).Should(BeEquivalentTo(zap.Logger{}))
        })
    })

    Context("when a logger is created", func() {
        logger := util.CreateLogger()
        It("is not nil", func() {
            Expect(logger).ShouldNot(BeNil())
        })
    })
})
