## gotils

### builtin APIs

#### slice APIs

1. make slice
```golang
import "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"

func makeSlice() {
    // before
    var slice = make([]string, 0, 0)

    // after
    var slice = slice.Make[string](0,0)
    println("slice: ", slice) // []

    var slice = slice.From[string]()
    println("slice: ", slice) // []

    var slice = slice.From("")
    println("slice: ", slice) // [""]
}
```

2. collect from slice
```golang
import "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"

func collectFromSlice() {
    var fruits = ["apple", "banana", "orange", "peach"]
    // before
    var bigFruits []string
    for _, fruit := range fruits {
        bigFruits = append(bigFruits, strings.ToUpper(fruit))
    }
    // after
    bigFruits = slice.Collect(fruits, func(fruit string) string { return strings.ToUpper(fruit)}).Slice()
}
```

```golang
import "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"

func collectControlFromSlice() error {
    var numbers = ["1", "2", "3", "four"]
    // before
    var intNumbers []int
    for _, number := range numbers {
        intNumber, err = strconv.Atoi(number)
        if err != nil {
            return err
        }

        intNumbers = append(intNumbers, intNumber)
    }

    // after
    intNumbersSlice, err := slice.CollectCtrl(numbers, func(number string) (int, error) {
        intNumber, err = strconv.Atoi(number)
        if err != nil {
            return -1, err
        }

        return intNumber, nil
    })

    if err !=  nil {
        return err
    }

    intNumbers = intNumbersSlice.Slice()
}
```

3. group slice
```golang
import "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"

func groupSlice() {
    type Person struct {
        Name string
        Sex string
    }

    var people = []Person {
        {
            Name: "a",
            Sex: "male",
        },
        {
            Name: "b",
            Sex: "male",
        },
        {
            Name: "c",
            Sex: "female",
        },
        {
            Name: "d",
            Sex: "female",
        },
        {
            Name: "e",
            Sex: "unknown",
        },
        {
            Name: "f",
            Sex: "unknown",
        },
    }

    // before
    var peopleGroupWithSex = make(map[string][]Person)
    for _, person := range people {
        group, ok := peopleGroupWithSex(person.Sex)
        if !ok {
            peopleGroupWithSex(person.Sex) = []Person{person}
            continue
        }
        group = append(group, person)
        peopleGroupWithSex(person.Sex) = group
    }

    // after
    peopleGroupWithSex = slice.Group(people, func(person Person) (string, Person) { return Person.Sex, Person}).Map()
}
```

#### slice Methods
```golang
import "cmp"
import "github.com/CrazyThursdayV50/gotils/pkg/builtin/api/slice"

func sliceMethods() {
    var before = []string{"1", "two", "III"}
    var after = slice.From(before...)

    // Len
    len(before)
    after.Len()

    // Swap
    before[i], before[j] = before[j], before[i]
    after.Swap(i, j)

    // Less
    var less = before[i] < before[j]
    after.WithLessFunc(cmp.Less)
    var less = after.Less(i, j)

    // Append
    before = append(before, "4️⃣")
    after.Append("4️⃣")

    // Cut
    var cut = before[1:2]
    var cut = after.Cut(1,2)

    // Clear
    clear(before)
    after.Clear()

    // Index
    var index int
    var target = "III"
    for i, e := range before {
        if e == target {
            index = i
            break
        }
    }

    after.Index(target, func(i, j string) bool { return i == j })

    // Get
    before[i] // maybe panic: out of range

    after.Get(i).Unwrap()
    if wrapper != nil {
        wrapper.Unwrap()
    }

    // Set
    before[0] = "1"
    after.Set(0, "1")

    // Iter
    for i, element := range before {
        // ....
    }

    breakAtIndex := after.IterOkay(func(index int, element string) bool { // ...})
    breakAtIndex, breakError := after.IterError(func(index int, element string) error { // ...})
    errors := after.IterFully(func(index int, element string) error { // ... })

    // 
}
```
