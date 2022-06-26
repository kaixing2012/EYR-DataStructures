package data_structures

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (std *LinkedList) Insert(val interface{}) {
	if std.head == nil {
		std.head = &Node{val, nil}

		return
	}

	if std.head.next == nil {
		std.head.next = &Node{val, nil}

		return
	}

	next := std.head.next

	for next.next != nil {
		next = next.next
	}

	next.next = &Node{val, nil}
}

func (std *LinkedList) Retrieve(val interface{}) interface{} {
	var resultSet map[string]interface{}

	if std.head != nil {
		index := 0

		if std.head.value == val {
			resultSet = map[string]interface{}{
				"idx": index,
				"val": std.head.value,
			}
		}

		if resultSet == nil {
			next := std.head.next

			for next != nil {
				index++

				if next.value == val {
					resultSet = map[string]interface{}{
						"idx": index,
						"val": next.value,
					}
				}

				next = next.next
			}
		}
	}

	return resultSet
}

func (std *LinkedList) FilterByVal(val interface{}) interface{} {
	var resultSet []interface{}

	if std.head != nil {
		index := 0

		if std.head.value == val {
			resultSet = append(resultSet, map[string]interface{}{
				"idx": index,
				"val": std.head.value,
			})
		}

		next := std.head.next

		for next != nil {
			index++

			if next.value == val {
				resultSet = append(resultSet, map[string]interface{}{
					"idx": index,
					"val": next.value,
				})
			}

			next = next.next
		}

	}

	return resultSet
}

func (std *LinkedList) Update(oldVal interface{}, newVal interface{}) bool {
	result := false

	if std.head != nil {

		if std.head.value == oldVal {
			std.head.value = newVal

			result = std.head.value == newVal
		}

		if !result {
			next := std.head.next

			for next != nil {
				if next.value == oldVal {
					next.value = newVal

					result = next.value == newVal
					break
				}

				next = next.next
			}
		}

	}

	return result
}

func (std *LinkedList) Delete(val interface{}) bool {
	result := false

	if std.head != nil {

		if std.head.value == val {
			if std.head.next != nil {
				std.head = std.head.next
			} else {
				std.head = nil
			}

			result = true
		}

		if !result {
			prev := std.head
			next := std.head.next

			for next != nil {

				if next.value == val {
					if next.next != nil {
						prev.next = next.next
					} else {
						prev.next = nil
					}

					result = true
					break
				}

				prev = next
				next = next.next
			}
		}
	}

	return result
}

func (std *LinkedList) Size() int {
	var count []interface{}

	if std.head != nil {
		count = append(count, std.head)

		next := std.head.next

		for next != nil {
			count = append(count, std.head.value)
			next = next.next
		}
	}

	return len(count)
}
