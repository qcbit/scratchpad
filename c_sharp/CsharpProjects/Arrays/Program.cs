string value = "abc123";
Console.WriteLine(new string((value.ToCharArray()).Reverse().ToArray()));
Console.WriteLine(string.Join(",", value.ToCharArray().Reverse().ToArray()));
