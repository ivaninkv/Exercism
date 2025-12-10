use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mut all_minutes = hours * 60 + minutes;
        while all_minutes < 0 {
            all_minutes += 24 * 60;
        }
        let h = all_minutes / 60 % 24;
        let m = all_minutes % 60;
        Clock {
            hours: h,
            minutes: m,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(0, self.hours * 60 + self.minutes + minutes)
    }
}
