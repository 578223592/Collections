# Collections
Collections：一个用反射过滤切片和转换切片类型的库

## 缘起

业务特点：
- 性能非“强”敏感：强字打上引号的指的是业务中性能瓶颈更多的在于外部系统：数据库、第三方接口等而非业务代码，所以内部耗时稍微增加也关系不大，当然也不能太过分。
- 切片类型转换和筛选：在许多场景，需要将一个切片转换为另一个切片，例如：将一个结构体类型的切片转换为另一个结构体类型的切片。
- 切片的筛选：在需要场景下，需要将一个切片中的元素进行筛选，只选择出符合条件的元素。
实际上，业务形态更多情况下like：
```go
//  some codes
slice1 := []Struct1{1, 2, 3, 4, 5}
slice2 := make([]Struct2, 0)
for _, item := range slice1 {
//对item进行条件过滤
    if (!condition(item)){
        continue
    }
	var item2 Struct2 = transform(item)
	slice2 = append(slice2, item2)
}
//do some thing use slice2
```
实际上如果转换链路比较短的话，还不会那么难受，但是转换链路比较长或者过滤次数比较多或者混合起来的话，整个转换的过程可能就会充斥着以`tempXXX`命名的变量。
这样的变量会给上下文的代码理解带来困难，更会造成需要上下文翻找这个变量实际在什么地方使用到了？是使用到了一次？还是两次？还是？？？

现有的库有一些是针对slice设计的，但是他们的设计初衷都是增加一个函数，用以快捷的操作slice，比如增加`insert`方法等等。
- https://github.com/elliotchance/pie
- https://github.com/jianfengye/collection

这确实非常有用，但是在我们的业务场景中，需求并不太一样。

我们想做出一个怎样的形态，可以见`collections_test.go`。

### 设计理念

- 易于使用.为了加速业务开发，我们不希望引入过多的理解成本，因此目前设计只设计了`where`函数与`transfer`函数。

- 任意类型转换的支持.希望做到slice中的元素类型可以任意转换而避免`for loop`+ `append`的组合方式。

- Type safety. 实际上，由于go对范型的限制（不支持范形方法），我们内部用到了反射进行处理，反射看起来和类型安全是一组反义词，然而，我们希望保证在运行的过程中不发生运行时错误。
- 链式调用.为了避免`tempXXX`命名的变量到处出现，我们希望采用类似于`jQuery`和[collect](https://github.com/tighten/collect)的链式调用，以简化代码。

- Performance. 如上，反射看起来和性能也不沾边，我们希望的是尽量减少这个库相比于原生实现的性能差异。

- Nil-safe. All of the functions will happily accept nil and treat them as empty slices. Apart from less possible panics, it makes it easier to chain.

- Immutable. Functions never modify inputs (except in cases where it would be illogical), unlike some built-ins such as sort.Strings
   todo  ???现在能做到Immutable吗？

## quick start

## 缘落
缘起终有缘落时。

### roadmap
- [ ] add `where` function
  - [ ] 添加where入参类型检查🧐，避免runtime error
- [ ] 添加where函数可以串联起来的优化
- [ ] 添加多goruntine以优化性能
- [ ] add more test
- [ ] add more 入参 limit ，保证安全 ，like gorm
- [ ] 考虑where和tranfer的返回参数增加一个error用以支持类型转换过程中的错误之类的

### 贡献指南


···todo···
