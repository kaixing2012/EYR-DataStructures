

class Node(object):

    def __init__(self, data=None):
        self.data = data
        self.next = None

    # Represent data at this node
    def __repr__(self):
        return (
            repr(self.data)  # return this
            if self.data != None  # if
            else ""  # otherwise
        )


class LinkedList(object):

    def __init__(self):
        self.head = Node()

    # Length of linked list
    def __len__(self):
        size = 0
        for index, node in enumerate(self):
            # return 0 if no data within
            if node.data is None:
                return 0

            size = index + 1
        return size

    # Make linked list iterable

    def __iter__(self):
        node = self.head
        while node is not None:
            yield node
            node = node.next

    def insert(self, data):
        cur_node = self.head

        # traverse the list
        for node in self:
            # add to the first node
            # if the list is empty
            if node.data is None:
                self.head = Node(data)
                return
            elif node.next is None:
                cur_node = node
        # add to the last node
        # which is none
        cur_node.next = Node(data)

    def regain(self, data):
        result_set = []
        for index, node in enumerate(self):
            if node.data == data:
                found = {"index": index, "data": node}
                result_set.append(found)
        if len(result_set) > 0:
            return result_set
        else:
            raise Exception('Data cannot be found!')

    def alter(self, index, data):
        last_node = None
        for idx, node in enumerate(self):
            new_node = Node(data)
            if idx == index:
                if index == 0:
                    self.head = new_node
                    new_node.next = node.next
                    return
                else:
                    last_node.next = new_node
                    new_node.next = node.next
                    return
            last_node = node

    def delete(self, index):
        last_node = None
        for idx, node in enumerate(self):
            if idx == index:
                if index == 0:
                    self.head = node.next
                    return
                else:
                    last_node.next = node.next
                    return
            last_node = node


linked_list = LinkedList()
print(f"Init list: {list(linked_list)}\n")

linked_list.insert("Ed")
linked_list.insert("Ed")
linked_list.insert("Roy")
linked_list.insert("John")

print(f"After insertion: {list(linked_list)}\n")

data_set = linked_list.regain("Ed")
print(f"Regain data by given Ed var: {data_set}\n")

linked_list.alter(data_set[0]["index"], "Edward")
print(f"After alteration: {list(linked_list)}\n")

linked_list.delete(data_set[0]["index"])
print(f"After deletion of Ed: {list(linked_list)}\n")
