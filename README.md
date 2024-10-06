**Objective:** To Build a Simple CRUD API with Go.

**Use Case – Student Records Management API**

This is a simple API that allows us to perform CRUD (Create, Read, Update, Delete) operations on a student record database. We can create, update, delete, and retrieve student details using POST or CURL commands.

-   **Create:** Adds a new student record by entering the name, email and course.
-   **Read:**  List all students’ records that have been created or search for a specific student by its ID.
-   **Update:** Update the student records.
-   **Delete:** Delete a student’s record from the database.

HTTP Server Status:

![HTTP Server Status](https://github.com/user-attachments/assets/bc18437b-3a9e-4258-ab74-cb3d4913b23a)

Used Postman to test APIs:

1. **POST**: Add a student record.
<img width="470" alt="POST" src="https://github.com/user-attachments/assets/da444fc8-b60c-4fbe-b37a-aa4073b29119">

2. To list all the students’ records.

**GET:** http://localhost:7881/students

<img width="470" alt="Get All" src="https://github.com/user-attachments/assets/09244a4a-d993-4604-8ad9-34f568e34d78">

List of all student records.

<img width="470" alt="get localhost" src="https://github.com/user-attachments/assets/fae8839e-1551-44d8-8358-06e27344c506">


3. List student details by ID.

**GET**: http://localhost:7881/students/4

<img width="470" alt="get id" src="https://github.com/user-attachments/assets/53ce55d1-f975-4600-b01d-8895819ce501">

Student details for ID-4.

<img width="470" alt="get id localhost" src="https://github.com/user-attachments/assets/e02b9385-27be-4c6e-8e1f-b3859ca780af">

4. Update student record. Corrected (updated) the name and course of the student.

**PUT:** [http://localhost:7881/students/2](http://localhost:7881/students/2)

<img width="470" alt="put" src="https://github.com/user-attachments/assets/5628f6fc-bc80-4b84-97bc-7d9e2cb457ef">

Student's name and course for ID-2 updated.

<img width="470" alt="put localhost" src="https://github.com/user-attachments/assets/97d6d696-17f2-4eea-9919-c72d4e8a3d5e">


**5.** Delete a student record.

**DELETE:** [http://localhost:7881/students/4](http://localhost:7881/students/4)

![delete](https://github.com/user-attachments/assets/33aa596b-dec9-4fa4-b022-4abeae106b3c)

Deleted the student record for ID-4.

<img width="470" alt="delete localhost" src="https://github.com/user-attachments/assets/2842314c-b441-4fce-9848-8601c48945a7">

List of updated student records.

<img width="470" alt="get all final" src="https://github.com/user-attachments/assets/3580983a-fa72-4921-b665-82f6ef1002cb">
