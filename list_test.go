package gollections_test

import . "github.com/ashis0013/gollections"

var _ = Describe("Tests for list utilities", func() {

    var list []int

    BeforeEach(func() {
        list = []int{1, 2, 3, 4, 5}
    })

    Context("Filter())", func() {
        It("should filter as per predicate", func() {
            Expect(Filter(list, func(x int) bool { return x % 2 == 1})).Should(Equal([]int{1, 3, 5}))
            Expect(Filter(list, func(x int) bool {return false })).Should(Equal([]int{}))
            Expect(Filter(list, nil)).Should(Equal([]int{}))
        })
    })

    Context("FilterIndexed())", func() {
        It("should filter as per predicate", func() {
            Expect(FilterIndexed(list, func(i, x int) bool { return i > 1 || x % 2 == 1})).Should(Equal([]int{1, 3, 4, 5}))
            Expect(FilterIndexed(list, func(i, x int) bool {return false })).Should(Equal([]int{}))
            Expect(FilterIndexed(list, nil)).Should(Equal([]int{}))
        })
    })

    Context("Map())", func() {
        It("should Map as per transform function", func() {
            Expect(Map(list, func(x int) int { return x * x})).Should(Equal([]int{1, 4, 9, 16, 25}))
            Expect(Map(list, func(x int) int { return x })).Should(Equal(list))
            Expect(Map[int, int](list, nil)).Should(Equal([]int{}))
        })
    })

    Context("MapIndexed())", func() {
        It("should Map as per transform function", func() {
            Expect(MapIndexed(list, func(i, x int) int { return x - i })).Should(Equal([]int{1, 1, 1, 1, 1}))
            Expect(MapIndexed(list, func(i, x int) int { return x })).Should(Equal(list))
            Expect(MapIndexed[int, int](list, nil)).Should(Equal([]int{}))
        })
    })

    Context("All()", func() {
        It("Should return true if all the element satisfies predicate", func() {
            Expect(All(list, func(x int) bool {  return x > 0 })).Should(BeTrue())
            Expect(All(list, func(x int) bool { return x % 2 == 0 })).Should(BeFalse())
            Expect(All(list, nil)).Should(BeFalse())
        })
    })

    Context("Any()", func() {
        It("Should return true if all the element satisfies predicate", func() {
            Expect(Any(list, func(x int) bool { return x > 0 })).Should(BeTrue())
            Expect(Any(list, func(x int) bool { return x % 2 == 0 })).Should(BeTrue())
            Expect(Any(list, nil)).Should(BeFalse())
        })
    })

    Context("Associate()", func() {
        It("Should return a map bassed on transform", func() {
            Expect(Associate(list, func(x int) (int, string) { return x*x, "hello"})).Should(Equal(map[int]string{
                1: "hello",
                4: "hello",
                9: "hello",
                16: "hello",
                25: "hello",
            }))
            Expect(Associate[int, int, string](list, nil)).Should(Equal(map[int]string{}))
        })
    })

    Context("Contains()", func() {
        It("Should return true if target element is there in slice", func() {
            Expect(Contains(list, 1)).Should(BeTrue())
            Expect(Contains(list, 10)).Should(BeFalse())
            Expect(Contains([]int{}, 0)).Should(BeFalse())
        })
    })

    Context("Drop()", func() {
        It("Should drop the value at index, ignore if invalid", func() {
            Expect(Drop(list, 1)).Should(Equal([]int{1, 3, 4, 5}))
            Expect(Drop(list, -1)).Should(Equal(list))
            Expect(Drop(list, 6)).Should(Equal(list))
        })
    })

    Context("First()", func() {
        It("Should return first value that satisfies the predicate, error otherwise", func() {
            Expect(First(list, func(x int) bool { return x % 2 == 0 })).Should(Equal(2))
            _, err1 := First(list, func(x int) bool { return x < 0 })
            Expect(err1).ShouldNot(BeNil())
            _, err2 := First(list, nil)
            Expect(err2).ShouldNot(BeNil())
        })
    })

    Context("FirstOrDefault()", func() {
        It("Should return the first value the satisfies the predicate, default value otherwise", func() {
            Expect(FirstOrDefault(list, -1, func(x int) bool { return x % 2 == 0 })).Should(Equal(2))
            Expect(FirstOrDefault(list, -1, func(x int) bool { return x < 0 })).Should(Equal(-1))
            Expect(FirstOrDefault(list, -1, nil)).Should(Equal(-1))
        })
    })

    Context("Flatten", func() {
        It("Should flatten a 2d array", func() {
            matrix := [][]int{
                {1, 2, 3},
                {3, 4, 5},
            }
            Expect(Flatten(matrix)).Should(Equal([]int{1, 2, 3, 3, 4, 5}))
            Expect(Flatten[int](nil)).Should(BeNil())
        })
    })

    Context("Fold()", func() {
        It("Should reduce the list as per the operation", func() {
            Expect(Fold(list, 0, func(a, b int) int { return a+b })).Should(Equal(15))
            Expect(Fold(list, 0, nil)).Should(Equal(0))
        })
    })

    Context("ForEach()", func() {
        It("Should behave like ForEach :)", func() {
            length := 0
            ForEach(list, nil)
            Expect(length).Should(Equal(0))
            ForEach(list, func(x int) { length++ })
            Expect(length).Should(Equal(5))
        })
    })

    Context("ForEachIndexed()", func() {
        It("Should run callback that takes index as input", func() {
            indexArr := []int{}
            ForEachIndexed(list, func(i, x int) { indexArr = append(indexArr, i) })
            Expect(indexArr).Should(Equal([]int{0, 1, 2, 3, 4}))
        })
    })

    Context("GroupBy()", func() {
        It("Shold group by the selector function", func() {
            Expect(GroupBy(list, func(x int) int { return x % 2})).Should(Equal(map[int][]int{
                0: {2, 4},
                1: {1, 3, 5},
            }))
            Expect(GroupBy[int, int](list, nil)).Should(Equal(map[int][]int{}))
        })
    })

    Context("IndexOf()", func() {
        It("Should return index of target element, -1 if not fount", func() {
            Expect(IndexOf(list, 3)).Should(Equal(2))
            Expect(IndexOf(list, 8)).Should(Equal(-1))
        })
    })

    Context("Reversed()", func() {
        It("Should return the reverse of the list", func() {
            Expect(Reversed(list)).Should(Equal([]int{5, 4, 3, 2, 1}))
        })
    })

    Context("SubList()", func() {
        It("Should return valid sublist given a valid range", func() {
            Expect(SubList(list, 1, 3)).Should(Equal([]int{2, 3, 4}))
            Expect(SubList(list, -1, 3)).Should(Equal([]int{}))
        })
    })

    Context("Zip()", func() {
        It("Should zip two different sized array properly", func() {
            Expect(Zip(list, []string{"Hello", "World"})).Should(Equal([]*Pair[int, string]{{1, "Hello"}, {2, "World"}}))
            Expect(Zip(list, []bool{})).Should(Equal([]*Pair[int, bool]{}))
        })
    })
})
