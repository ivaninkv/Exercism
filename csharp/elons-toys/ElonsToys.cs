using System;

class RemoteControlCar
{
    private int distance;
    private int battery;

    public RemoteControlCar()
    {
        distance = 0;
        battery = 100;
    }

    public static RemoteControlCar Buy()
    {
        return new RemoteControlCar();
    }

    public string DistanceDisplay()
    {
        return $"Driven {distance} meters";
    }

    public string BatteryDisplay()
    {
        return battery > 0 ? $"Battery at {battery}%" : "Battery empty";
    }

    public void Drive()
    {
        if (battery > 0)
        {
            distance += 20;
            battery -= 1;
        }
    }
}
