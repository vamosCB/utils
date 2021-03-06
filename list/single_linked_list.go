package list

//Object 数据实体对象
type Object interface{}

//Node 链表节点
type Node struct {
	Data Object
	Next *Node
}

//SingleLinkedList 链表
type SingleLinkedList struct {
	Length int   //链表长度
	Head   *Node //头节点
}

//NewSingleLinkedList new..
func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{Length: 0, Head: nil}
}

//Empty 链表是否为空
func (list *SingleLinkedList) Empty() bool {
	return list.Head == nil
}

//Len length of this list
func (list *SingleLinkedList) Len() (length int) {
	return list.Length
}

//Append 添加节点
func (list *SingleLinkedList) Append(data Object) {
	node := &Node{Data: data}
	if list.Empty() {
		list.Head = node
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	list.Length++
}

//HeadAdd 头部插入
func (list *SingleLinkedList) HeadAdd(data Object) {
	node := &Node{Data: data}
	node.Next = list.Head
	list.Head = node
	list.Length++
}

//InsertByIndex 在特定位置插入
func (list *SingleLinkedList) InsertByIndex(index int, data Object) {
	if index < 0 {
		list.HeadAdd(data)
	} else if index > list.Length {
		list.Append(data)
	} else {
		pre := list.Head
		count := 0
		for count < index-1 {
			pre = pre.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
	list.Length++
}

//DeleteByIndex 删除下标元素
func (list *SingleLinkedList) DeleteByIndex(index int) {
	pre := list.Head
	if index <= 0 {
		list.Head.Next = pre.Next
	} else if index >= list.Length {
		//包装自己的log包
		return
	} else {
		count := 0
		for count < index-1 {
			pre = pre.Next
			count++
		}
		pre.Next = pre.Next.Next
	}
	list.Length--
}

//DeleteByData 根据元素删除链表节点
func (list *SingleLinkedList) DeleteByData(data Object) {
	if list.Empty() {
		//日志输出
		return
	}
	pre := list.Head
	if pre.Data == data {
		pre = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Data == data {
				pre.Next = pre.Next.Next
			}
			pre = pre.Next
		}
	}
	list.Length--
}

//Contain 是否包含固定元素
func (list *SingleLinkedList) Contain(data Object) bool {
	pre := list.Head
	for pre != nil {
		if pre.Data == data {
			return true
		}
		pre = pre.Next
	}
	return false
}
