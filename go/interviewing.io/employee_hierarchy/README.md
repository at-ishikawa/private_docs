# Medium: [Employee Hierarchy](https://interviewing.io/questions/employee-hierarchy)

Given an array of employee IDs including who they report to, write a function to calculate the score for a given employee. The employee score for an employee equals "Total number of direct and indirect employees report to that employee, then plus 1."

The “plus one" here means adding the employee itself as self-reporting.

An employee without any other employees reporting to it, will have employee score 1.

Each employee has a unique eid (employee_id). Given a direct report map, where key is an eid, value is an array of eids who directly report to key. This map should contain all employees. The map could contain cycles.

Here is an example of direct report map: {123: [234, 345], 234: [456, 789], 345:[], 456:[], 789:[]}

Your solution should have better runtime when called multiple times.
