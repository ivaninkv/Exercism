using System;

public class Player
{
    public int RollDie()
    {
        var rnd = new Random();
        return rnd.Next(1, 19);
    }

    public double GenerateSpellStrength()
    {
        var rnd = new Random();
        return rnd.NextDouble() * 100.0;
    }
}
