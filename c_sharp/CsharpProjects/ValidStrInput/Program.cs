using System;

string? input;
bool valid = false;
do {
  Console.Write("Enter a role [Administrator, Manager, User]: ");
  input = Console.ReadLine();
  if (input is not null) {
    valid = input.ToLower().Trim() == "administrator" || input.ToLower().Trim() == "manager" || input.ToLower().Trim() == "user";
    if (! valid) {
      Console.WriteLine($"The role name that you entered, \"{input}\" is not valid.");
    }
  }
} while (! valid);

Console.Write($"The input value ({input}) is accepted.");