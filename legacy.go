import etcd
class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.left = None
        self.right = None
class TTree:
    def __init__(self):
        self.root = Node(None, None)
    def insert(self, key, value):
        self.root = self._insert(self.root, key, value)
    def _insert(self, node, key, value):
        if node.key is None:
            node.key = key
            node.value = value
        elif key < node.key:
            if node.left is None:
                node.left = Node(None, None)
            self._insert(node.left, key, value)
        else:
            if node.right is None:
                node.right = Node(None, None)
            self._insert(node.right, key, value)
        return node

    def search(self, key):
        return self._search(self.root, key)

    def _search(self, node, key):
        if node.key is None:
            return None
        elif key == node.key:
            return node.value
        elif key < node.key:
            if node.left is None:
                return None
            return self._search(node.left, key)
        else:
            if node.right is None:
                return None
            return self._search(node.right, key)

    def delete(self, key):
        self.root = self._delete(self.root, key)
    def _delete(self, node, key):
        if node is None:
            return node
        if key < node.key:
            node.left = self._delete(node.left, key)
        elif key > node.key:
            node.right = self._delete(node.right, key)
        else:
            if node.left is None:
                return node.right
            elif node.right is None:
                return node.left
            else:
                min_node = self._find_min(node.right)
                node.key = min_node.key
                node.value = min_node.value
                node.right = self._delete(node.right, min_node.key)
        return node

    def _find_min(self, node):
        while node.left is not None:
            node = node.left
        return node

class TTreeETCDClient:
    def __init__(self, etcd_client):
        self.etcd_client = etcd_client
        self.t_tree = TTree()

    def put(self, key, value):
        self.t_tree.insert(key, value)
        self.etcd_client.put(key, value)

    def get(self, key):
        value = self.t_tree.search(key)
        if value is not None:
            return value
        return self.etcd_client.get(key)

    def delete(self, key):
        self.t_tree.delete(key)
        self.etcd_client.delete(key)
etcd_client = etcd.Client(host='localhost', port=2379)
t_tree_etcd_client = TTreeETCDClient(etcd_client)
t_tree_etcd_client.put('key1', 'value1')
print(t_tree_etcd_client.get('key1'))  # Output: value1
t_tree_etcd_client.delete('key1')
print(t_tree_etcd_client.get('key1')) 
