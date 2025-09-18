use std::path::PathBuf;

pub struct Target {
    enable: bool,
    config_location: PathBuf,
}
pub enum TargetKind {
    Helix,
    Hypr,
}
