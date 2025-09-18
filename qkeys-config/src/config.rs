use crate::Target;

pub struct Config {
    log_level: LogLevel,
    targets: Vec<Target>,
}

pub enum LogLevel {
    Error,
    Warning,
    Info,
    Debug,
}
