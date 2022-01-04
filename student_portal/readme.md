//first yo need to create a schema name student_details and give your db username and password //CreateStudent creates student POST REQUEST
http://localhost:8090/create 

//GetAllStudent to get all student data GET REQUEST
http://localhost:8090/get 

//GetStudentByID returns Student with specific ID GET REQUEST
http://localhost:8090/get/give id 

//GetStudentByYearOfJoining returns Student with specific students (still not working) GET REQUEST
http://localhost:8090/update/year of joining

//UpdateMobileNumberByID will updates student Mobile Number with respective ID PUT REQUEST in body(changed mobile number) 
http://localhost:8090/update/give id

//DeleteStudentByID will delete's student with specific ID DELETE REQUEST
http://localhost:8090/delete/give id

//BODY INPUT
{
"id":1, "firstName": "Ram",
"lastName": "Kumar", 
"mobileNumber": 7683074449,
"yearOfJoining": 2017,
"emailId" :"ram.kumar@quest-global.com" 
}
