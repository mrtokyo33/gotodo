# **GoTodo CLI**

shibuya-todo CLI will be an application where you can manage a todo list that is a list of tasks for you to complete

## **Commands:**
### **1. *gotodo add***
Adds a new task to the list

### **2. *gotodo list***
Displays all tasks including completed and pending ones

### **3. *gotodo do/undo***
Marks a task as completed or uncompleted

### **4. *gotodo rm***
Removes a task from the list

## **Architecture**
### **1. *Models ( Data Structures )***
This is the simplest package. It contains the structs that represent the application's data.

### **2. *Repositories ( Persistence )***
This is the most important layer for the architecture. Its purpose is to abstract how data is saved

- **Interface *( repository.go )*:** It dictates which data operations exist, regardless of how they are saved
- **Implementation *( json.go )*:** A struct that implements this interface, containing the logic for reading and writing the JSON file

### **3. *Services ( Business Rules )***
This layer contains all the logic of the application and is the only one that interacts directly with the repository

- **Dependency Injection:** The service struct should receive the repository interface
