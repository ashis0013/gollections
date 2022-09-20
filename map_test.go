package gollections_test

import (
	"fmt"

	. "github.com/ashis0013/gollections"
)

var _ = Describe("tests for map utilities", func() {
    var hashMap map[int]string

    BeforeEach(func() {
        hashMap = map[int]string{
            1: "Hello",
            2: "World",
        }
    })

    Context("Entries()", func() {
        It("Should return a list of key value pair", func() {
            Expect(Entries(hashMap)).Should(Equal([]*Pair[int, string]{{1, "Hello"}, {2, "World"}}))
            Expect(Entries[int, string](nil)).Should(Equal([]*Pair[int, string]{}))
        })
    })

    Context("Keys()", func() {
        It("Should return a list of all the keys", func() {
            Expect(Keys(hashMap)).Should(BeEquivalentTo([]int{1, 2}))
            Expect(Keys[int, string](nil)).Should(Equal([]int{}))
        })
    })

    Context("Values()", func() {
        It("Should return a list of all the values", func() {
            Expect(Values(hashMap)).Should(BeEquivalentTo([]string{"Hello", "World"}))
            Expect(Values[int, string](nil)).Should(Equal([]string{}))
        })
    })

    Context("GetOrDefault()", func() {
        It("Should return the value corresponding to the key else default values", func() {
            Expect(GetOrDefault(hashMap, 1, "bruh")).Should(Equal("Hello"))
            Expect(GetOrDefault(hashMap, 3, "bruh")).Should(Equal("bruh"))
            Expect(GetOrDefault(nil, 3, "bruh")).Should(Equal("bruh"))
        })
    })

    Context("FilterKeys()", func() {
        It("Should return a map after filtering the keys based on predicate", func() {
            newMap := FilterKeys(hashMap, func(x int) bool { return x%2==0 })
            Expect(len(newMap)).Should(Equal(1))
            Expect(newMap[2]).Should(Equal("World"))

            newMap = FilterKeys[int, string](nil, func(x int) bool { return x%2==0 })
            Expect(len(newMap)).Should(Equal(0))

            newMap = FilterKeys(hashMap, nil)
            Expect(len(newMap)).Should(Equal(0))
        })
    })

    Context("FlatMap()", func() {
        It("Should transform the map into a slice", func() {
            Expect(FlatMap(hashMap, func(x int, s string) bool { return false })).Should(Equal([]bool{false, false}))
            Expect(FlatMap(nil, func(x int, s string) bool { return false })).Should(Equal([]bool{}))
            Expect(FlatMap[int, string, bool](hashMap, nil)).Should(Equal([]bool{}))
        })
    })

    Context("ForEachEntry()", func() {
        It("Should run operation for each entry in map", func() {
            dummy := 0
            dummyStr := ""

            ForEachEntry(hashMap, nil)

            Expect(dummy).Should(Equal(0))
            Expect(dummyStr).Should(Equal(""))

            ForEachEntry(hashMap, func(x int, s string) {
                dummy += x
                dummyStr = fmt.Sprintf("%s%s", dummyStr, s)
            })

            Expect(dummy).Should(Equal(3))
            Expect(dummyStr).Should(Equal("HelloWorld"))
        })
    })
})
