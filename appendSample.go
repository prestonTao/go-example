删除切片b

a = append(a,b...)

复制

//因为a类似于C里的数组，直接复制是浅复制
b = make([]T,len(a))
copy(b,a)


删除[i:j]

a = append(a[:i],a[j:]...)

删除第i个元素

a = append(a[:i],a[i+1:]...)

扩展j个空元素

a = append(a,make([]T,j))

插入j个空元素

a = append(a[:i],append(make([]T,j),a[i:...])...)

插入元素x

a = append(a[:i],append([]T{x},a[i:]...)...)

插入切片b

a = append(a[:i],append(b,a[i:]...)...)

弹出最后一个元素

x ,a = a[len(a)-1],a[:len(a)-1]

压入x

a = append(a,x)