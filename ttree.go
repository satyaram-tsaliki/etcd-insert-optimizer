class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.children = []
        self.index = None

class TTree:
    def __init__(self, degree):
        self.degree = degree
        self.root = Node(None, None)

    def insert(self, key, value):
        self.root = self._insert(self.root, key, value)

    def _insert(self, node, key, value):
        if node.key is None:
            node.key = key
            node.value = value
        elif key < node.key:
            if len(node.children) == 0:
                node.children.append(Node(None, None))
            self._insert(node.children[0], key, value)
        elif key > node.key:
            if len(node.children) == self.degree:
                # Split the node into two nodes
                mid = len(node.children) // 2
                new_node = Node(node.children[mid].key, node.children[mid].value)
                new_node.children = node.children[mid+1:]
                node.children = node.children[:mid]
                node.children.append(new_node)
            self._insert(node.children[-1], key, value)
        else:
            node.value = value
        return node

    def search(self, key):
        return self._search(self.root, key)

    def _search(self, node, key):
        if node.key is None:
            return None
        elif key == node.key:
            return node.value
        elif key < node.key:
            if len(node.children) == 0:
                return None
            # Use the index to quickly find the child node
            index = self._find_index(node.children, key)
            if index < len(node.children):
                return self._search(node.children[index], key)
        else:
            if len(node.children) == self.degree:
                # Use the index to quickly find the child node
                index = self._find_index(node.children, key)
                if index < len(node.children):
                    return self._search(node.children[index], key)
            return self._search(node.children[-1], key)

    def _find_index(self, children, key):
        left, right = 0, len(children)
        while left < right:
            mid = (left + right) // 2
            if children[mid].key < key:
                left = mid + 1
            else:
                right = mid
        return left

    def delete(self, key):
        self.root = self._delete(self.root, key)

    def _delete(self, node, key):
        if node is None:
            return node
        if key < node.key:
            node.children[0] = self._delete(node.children[0], key)
        elif key > node.key:
            node.children[-1] = self._delete(node.children[-1], key)
        else:
            if len(node.children) == 0:
                return None
            elif len(node.children) == 1:
                return node.children[0]
            else:
                min_node = self._find_min(node.children[1])
                node.key = min_node.key
                node.value = min_node.value
                node.children[1] = self._delete(node.children[1], min_node.key)
        return node

    def _find_min(self, node):
        while len(node.children) > 0:
            node = node.children[0]
        return node

t_tree = TTree(3)
t_tree.insert('key1', 'value1')
t_tree.insert('key2', 'value2')
t_tree.insert('key3', 'value3')

print(t_tree.search('key1'))  
print(t_tree.search('key2'))  
print(t_tree.search('key3'))  

t_tree.delete('key2')
print(t_tree.search('key2'))  
