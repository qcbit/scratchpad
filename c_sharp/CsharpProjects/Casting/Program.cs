int first = 2;
string second = "4";
// int result = first + second; // Produces CsharpProjects/Casting/Program.cs(3,14): error CS0029: Cannot implicitly convert type 'string' to 'int'
string result = first + second; // 24
Console.WriteLine(result);

// Information loss?
int myInt = 3;
Console.WriteLine($"Int: {myInt}"); // 3

decimal myDecimal = myInt; // widening conversion, implicit
Console.WriteLine($"decimal: {myDecimal}"); // 3

// 2nd example
decimal myDecimal2 = 3.14m;
int myInt2 = (int)myDecimal2; // Narrowing conversion, explicit
Console.WriteLine($"int: {myInt2}"); // 3, Loss of data

// 3rd example
decimal myDecimal3 = 1.23456789m;
float myFloat = (float)myDecimal3; // Narrowing conversion
Console.WriteLine($"Decimal: {myDecimal3}");
Console.WriteLine($"Float: {myFloat}");

// Conversions
int first2 = 5;
int second2 = 7;
string message = first2.ToString() + second2.ToString();
Console.WriteLine(message); // 57

string first3 = "5";
string second3 = "7";
int result2 = int.Parse(first3) + int.Parse(second3);
Console.WriteLine(result2); // 12

string value1 = "5";
string value2 = "7";
int result3 = Convert.ToInt32(value1) * Convert.ToInt32(value2);
Console.WriteLine(result3); // 35