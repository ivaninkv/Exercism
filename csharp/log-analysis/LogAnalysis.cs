public static class LogAnalysis
{
    public static string SubstringAfter(this string str, string delimiter)
    {
        return str[(str.IndexOf(delimiter) + delimiter.Length)..];
    }

    public static string SubstringBetween(this string str, string startStr, string endStr)
    {
        return str[(str.IndexOf(startStr) + startStr.Length)..str.IndexOf(endStr)];
    }

    public static string Message(this string str)
    {
        return str.SubstringAfter("]: ");
    }

    public static string LogLevel(this string str)
    {
        return str.SubstringBetween("[", "]");
    }
}