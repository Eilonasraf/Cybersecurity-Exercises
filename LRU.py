from collections import OrderedDict

class LRUCache:
    def __init__(self, capacity: int):
        self.cache = OrderedDict()  # This will hold the key-value pairs
        self.capacity = capacity
    
    def get(self, key: int):
       if key not in self.cache:
           return -1
       else:
           self.cache.move_to_end(key) # Most frequently used - MRU
           return self.cache[key]
        
    def put(self, key:int, value:int) -> None:
        if key in self.cache:
            self.cache.move_to_end(key) # Move key to end if it already exists
        self.cache[key] = value  # Insert or update the key-value pair
        
        if len(self.cache) > self.capacity:
            self.cache.popitem(last=False) # Pop the first item (least recently used)

obj = LRUCache(2)
print(obj.put(1, 1), obj.put(2, 2), obj.get(1), obj.put(3, 3), obj.get(2))


