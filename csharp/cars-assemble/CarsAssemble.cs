using System;

internal static class AssemblyLine
{
    private const int CarsPerHour = 221;

    public static double SuccessRate(int speed)
    {
        return speed switch
        {
            0 => 0,
            >= 1 and <= 4 => 1,
            >= 5 and <= 8 => 0.9,
            9 => 0.8,
            10 => 0.77,
            _ => throw new ArgumentOutOfRangeException(nameof(speed), speed, null)
        };
    }

    public static double ProductionRatePerHour(int speed)
    {
        return CarsPerHour * speed * SuccessRate(speed);
    }

    public static int WorkingItemsPerMinute(int speed)
    {
        return (int)ProductionRatePerHour(speed) / 60;
    }
}