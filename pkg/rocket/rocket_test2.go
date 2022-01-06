package rocket

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rocket", func() {

	var (
		rocket *Rocket
	)

	BeforeEach(func() {
		rocket = NewRocket("Saturn-V", "Nasa", 5, 25000)
	})

	Describe("Igniting a rocket", func() {
		It("with fuel succeeds without error", func() {
			rocket.AddFuel(100)
			err := rocket.Ignite()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("without fuel errors", func() {
			err := rocket.Ignite()
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("A rocket", func() {
		Context("which has been fuelled and ignited", func() {
			JustBeforeEach(func() {
				rocket.AddFuel(100)
				err := rocket.Ignite()
				if err != nil {
					panic("ignition failed")
				}
			})

			It("can throttle up successfully if it's resulting speed is less or equal to its max speed", func() {
				speed, err := rocket.ThrottleUp(100)
				Expect(speed).Should(Equal(100))
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("can throttle down successfully as long as it's resulting speed does not drop below 1000", func() {
				_, _ = rocket.ThrottleUp(2000)
				speed, err := rocket.ThrottleDown(100)
				Expect(speed).Should(Equal(1900))
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("will throw an error if it trys to throttle up and it's resulting speed exceeds its max speed", func() {
				_, _ = rocket.ThrottleUp(25000)
				_, err := rocket.ThrottleUp(1)

				Expect(err).Should(HaveOccurred())
			})

			It("will throw an error if it trys to throttle down and it's resulting speed is less than 1000", func() {
				_, _ = rocket.ThrottleUp(2000)
				_, err := rocket.ThrottleDown(1500)

				Expect(err).Should(HaveOccurred())
			})

			It("tracks it's current speed", func() {
				_, _ = rocket.ThrottleUp(2000)
				_, _ = rocket.ThrottleDown(5)
				_, _ = rocket.ThrottleDown(10)
				_, _ = rocket.ThrottleUp(200)

				Expect(rocket.CurrentSpeed()).Should(Equal(2185))
			})
		})

		Context("which has been fuelled but not ignited", func() {
			JustBeforeEach(func() {
				rocket.AddFuel(100)
			})

			It("throws an error when it attempts to throttle up", func() {
				_, err := rocket.ThrottleUp(100)
				Expect(err).Should(HaveOccurred())
			})

			It("throws an error when it attempts to throttle down", func() {
				_, err := rocket.ThrottleDown(100)
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
