class Node:
    def __init__(self, key, val):
        self.key = key
        self.val = val
        self.prev = None
        self.next = None
        
class LRUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.cache = {} # map key to node
        
        # left = LRU, right = MRU
        self.left = Node(0, 0)
        self.right = Node(0, 0)
        self.left.next = self.right
        self.right.prev = self.left
    
    # remove node from list
    def remove(self, node):
        prev = node.prev 
        nxt = node.next
        prev.next = nxt
        nxt.prev = prev
        
    # insert node at right    
    def insert(self, node):
        prev = self.right.prev
        next = self.right
        prev.next = node
        next.prev = node
        node.next = next
        node.prev = prev
        
    def get(self, key: int):
        if key in self.cache:
            self.remove(self.cache[key])
            self.insert(self.cache[key])
            return self.cache[key].val
        return -1
            
    def put(self, key:int, value:int) -> None:
        if key in self.cache:
            self.remove(self.cache[key])
        self.cache[key] = Node(key, value)
        self.insert(self.cache[key])
        
        if len(self.cache) > self.cap:
            # remove from the list and delete the LRU from the hashmap
            lru = self.left.next
            self.remove(lru)
            del self.cache[lru.key]
            

obj = LRUCache(2)
print(obj.put(1, 1), obj.put(2, 2), obj.get(1), obj.put(3, 3), obj.get(2))

