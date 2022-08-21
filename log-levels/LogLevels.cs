using System;
using System.Collections.Generic;

static class LogLine
{
    public static string Message(string logLine)
    {
        return logLine.Split(":")[1].Trim();
    }

    public static string LogLevel(string logLine)
    {
        var logLevel = logLine.Split(":")[0].Trim();
        new List<char>() {'[', ']'}.ForEach(c => 
            logLevel = logLevel.Replace(c.ToString(), String.Empty));
        return logLevel.ToLower();
    }

    public static string Reformat(string logLine)
    {
        return $"{Message(logLine)} ({LogLevel(logLine)})";
    }
}
