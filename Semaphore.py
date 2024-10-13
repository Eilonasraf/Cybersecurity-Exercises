import threading
import time

# Define a binary semaphore
semaphore = threading.Semaphore(1)

def critical_section(id):
    print(f"Thread {id} attempting to acquire the semaphore.")
    # Acquire the semaphore (decrement the counter)
    semaphore.acquire()
    print(f"Thread {id} entered the critical section.")
    # Simulate some work in the critical section
    time.sleep(2)
    print(f"Thread {id} exited the critical section.")
    # Release the semaphore (increment the counter)
    semaphore.release()
    print(f"Thread {id} released the semaphore.")

# Create two threads
t1 = threading.Thread(target=critical_section, args=(1,)) # tuple
t2 = threading.Thread(target=critical_section, args=(2,)) # tuple

# Start the threads
t1.start()
t2.start()

# Wait for both threads to complete
t1.join()
t2.join()