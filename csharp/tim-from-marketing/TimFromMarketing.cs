using System;

static class Badge
{
    public static string Print(int? id, string name, string? department)
    {
        var idString = id != null ? $"[{id}] - " : "";
        return $"{idString}{name} - {(department != null? department.ToUpper() : "OWNER")}";
    }
}
