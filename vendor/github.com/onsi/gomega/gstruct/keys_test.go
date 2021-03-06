package gstruct_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Map", func() {
	allKeys := map[string]string{"A": "a", "B": "b"}
	missingKeys := map[string]string{"A": "a"}
	extraKeys := map[string]string{"A": "a", "B": "b", "C": "c"}
	emptyKeys := map[string]string{}

	It("should strictly match all keys", func() {
		m := MatchAllKeys(Keys{
			"B": Equal("b"),
			"A": Equal("a"),
		})
		Expect(allKeys).Should(m, "should match all keys")
		Expect(missingKeys).ShouldNot(m, "should fail with missing keys")
		Expect(extraKeys).ShouldNot(m, "should fail with extra keys")
		Expect(emptyKeys).ShouldNot(m, "should fail with empty keys")

		m = MatchAllKeys(Keys{
			"A": Equal("a"),
			"B": Equal("fail"),
		})
		Expect(allKeys).ShouldNot(m, "should run nested matchers")
	})

	It("should handle empty maps", func() {
		m := MatchAllKeys(Keys{})
		Expect(map[string]string{}).Should(m, "should handle empty maps")
		Expect(allKeys).ShouldNot(m, "should fail with extra keys")
	})

	It("should ignore missing keys", func() {
		m := MatchKeys(IgnoreMissing, Keys{
			"B": Equal("b"),
			"A": Equal("a"),
		})
		Expect(allKeys).Should(m, "should match all keys")
		Expect(missingKeys).Should(m, "should ignore missing keys")
		Expect(extraKeys).ShouldNot(m, "should fail with extra keys")
		Expect(emptyKeys).Should(m, "should match empty keys")
	})

	It("should ignore extra keys", func() {
		m := MatchKeys(IgnoreExtras, Keys{
			"B": Equal("b"),
			"A": Equal("a"),
		})
		Expect(allKeys).Should(m, "should match all keys")
		Expect(missingKeys).ShouldNot(m, "should fail with missing keys")
		Expect(extraKeys).Should(m, "should ignore extra keys")
		Expect(emptyKeys).ShouldNot(m, "should fail with empty keys")
	})

	It("should ignore missing and extra keys", func() {
		m := MatchKeys(IgnoreMissing|IgnoreExtras, Keys{
			"B": Equal("b"),
			"A": Equal("a"),
		})
		Expect(allKeys).Should(m, "should match all keys")
		Expect(missingKeys).Should(m, "should ignore missing keys")
		Expect(extraKeys).Should(m, "should ignore extra keys")
		Expect(emptyKeys).Should(m, "should match empty keys")

		m = MatchKeys(IgnoreMissing|IgnoreExtras, Keys{
			"A": Equal("a"),
			"B": Equal("fail"),
		})
		Expect(allKeys).ShouldNot(m, "should run nested matchers")
	})
})
