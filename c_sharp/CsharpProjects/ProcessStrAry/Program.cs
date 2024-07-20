using System;

string[] myStrings = new string[2] { "I like pizza. I like roast chicken. I like salad", "I like all three of the menu choices" };

int periodLocation = 0;
for(int ndx = 0; ndx < myStrings.Length; ndx++) {
  string str = myStrings[ndx];
  periodLocation = str.IndexOf('.');
  while (periodLocation > 0) {
    Console.WriteLine(str.Substring(0, periodLocation));
    str = str.Remove(0, periodLocation + 1);
    str = str.TrimStart(); // Removes leading whitespace
    periodLocation = str.IndexOf('.');
  }
  Console.WriteLine(str);
}