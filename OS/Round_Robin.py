def round_robin(tasks, quantum):
    while tasks:
        task, time = tasks.pop(0) # from the start
        if time > quantum:
            print(f"{task} executed for {quantum} units")
            tasks.append((task, time - quantum))  # Remaining time
        else:
            print(f"{task} finished")

tasks = [('Task1', 5), ('Task2', 3), ('Task3', 8)]
quantum = 2
round_robin(tasks, quantum)
