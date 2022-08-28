using System;
using System.Linq;

public static class Bob
{
    public static string Response(string statement)
    {
        statement = statement.Trim();
        return statement switch
        {
            _ when IsYellQuestion(statement) => "Calm down, I know what I'm doing!",
            _ when IsYell(statement) => "Whoa, chill out!",
            _ when IsQuestion(statement) => "Sure.",
            _ when IsIdleTalk(statement) => "Fine. Be that way!",
            _ => "Whatever."
        };
    }

    private static bool HasLetters(string str) => str.Any(char.IsLetter);

    private static bool IsQuestion(string str) => str.EndsWith("?");

    private static bool IsYell(string str) => HasLetters(str) && str.Equals(str.ToUpper());

    private static bool IsYellQuestion(string str) => IsQuestion(str) && IsYell(str);

    private static bool IsIdleTalk(string str) => string.IsNullOrWhiteSpace(str);
}