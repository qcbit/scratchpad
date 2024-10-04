// Validating integer inputs
using System;

string? readInput;
int input = 0;
bool valid = false;
do {
  Console.Write("Enter an integer between 5 and 10: ");
  readInput = Console.ReadLine();
  if (readInput is not null) {
    valid = int.TryParse(readInput, out input);
    if (valid && input < 5 || input > 10) {
      valid = false;
    }
  }
} while (! valid);

Console.WriteLine($"You input of {input} is accepted.");