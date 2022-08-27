using System;
using System.Linq;

public static class Bob
{
    public static string Response(string statement)
    {
        statement = statement.Trim();
        return statement switch
        {
            var s when HasLetters(s) && s.EndsWith("?") && s.Equals(s.ToUpper()) => "Calm down, I know what I'm doing!",
            var s when s.EndsWith("?") => "Sure.",
            var s when string.IsNullOrWhiteSpace(s) => "Fine. Be that way!",
            var s when HasLetters(s) && s.Equals(s.ToUpper()) => "Whoa, chill out!",
            _ => "Whatever."
        };
    }

    public static bool HasLetters(string str) => str.Any(char.IsLetter);
}