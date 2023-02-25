using System;

public static class Darts
{
    private const int outerRadius = 10;
    private const int middleRadius = 5;
    private const int innerRadius = 1;
    public static int Score(double x, double y)
    {
        var squareSum = x * x + y * y;

        var points = squareSum switch
        {
            <= innerRadius * innerRadius => 10,
            <= middleRadius * middleRadius => 5,
            <= outerRadius * outerRadius => 1,
            _ => 0
        };

        return points;
    }
}
