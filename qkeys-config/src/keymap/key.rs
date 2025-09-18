use std::{
    fmt::{Display, Write},
    str::FromStr,
};

use anyhow::bail;
use crossterm::event::{KeyCode, KeyEvent, KeyModifiers};

pub struct Key {
    pub code: KeyCode,
    pub shift: bool,
    pub ctrl: bool,
    pub alt: bool,
}

impl Default for Key {
    fn default() -> Self {
        Self {
            code: KeyCode::Null,
            shift: false,
            ctrl: false,
            alt: false,
        }
    }
}

impl From<KeyEvent> for Key {
    fn from(value: KeyEvent) -> Self {
        let shift = match (value.code, value.modifiers) {
            (KeyCode::Char(c), _) => c.is_ascii_uppercase(),
            (KeyCode::BackTab, _) => false,
            (_, m) => m.contains(KeyModifiers::SHIFT),
        };

        Self {
            code: value.code,
            shift,
            ctrl: value.modifiers.contains(KeyModifiers::CONTROL),
            alt: value.modifiers.contains(KeyModifiers::ALT),
        }
    }
}

impl FromStr for Key {
    type Err = anyhow::Error;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        if s.is_empty() {
            bail!("empty key")
        }

        let mut key = Self::default();

        let mut split = s.split_inclusive('-').peekable();
        while let Some(next) = split.next() {
            match next.to_lowercase().as_str() {
                "s-" => key.shift = true,
                "c-" => key.shift = true,
                "a-" => key.shift = true,

                _ => {
                    key.code = match next {
                        "space" => KeyCode::Char(' '),
                        "backspace" | "bksp" => KeyCode::Backspace,
                        "enter" | "return" | "ret" => KeyCode::Enter,
                        "left" => KeyCode::Left,
                        "right" => KeyCode::Right,
                        "up" => KeyCode::Up,
                        "down" => KeyCode::Down,
                        "home" => KeyCode::Home,
                        "pageup" | "pgup" => KeyCode::PageUp,
                        "pagedown" | "pgdn" => KeyCode::PageDown,
                        "tab" => KeyCode::Tab,
                        "backtab" | "btab" => KeyCode::BackTab,
                        "delete" | "del" => KeyCode::Delete,
                        "insert" | "ins" => KeyCode::Insert,
                        "escape" | "esc" => KeyCode::Esc,
                        "f1" => KeyCode::F(1),
                        "f2" => KeyCode::F(2),
                        "f3" => KeyCode::F(3),
                        "f4" => KeyCode::F(4),
                        "f5" => KeyCode::F(5),
                        "f6" => KeyCode::F(6),
                        "f7" => KeyCode::F(7),
                        "f8" => KeyCode::F(8),
                        "f9" => KeyCode::F(9),
                        "f10" => KeyCode::F(10),
                        "f11" => KeyCode::F(11),
                        "f12" => KeyCode::F(12),
                        _ => match next {
                            s if split.peek().is_none() => {
                                let c = s.chars().next().unwrap();
                                KeyCode::Char(c)
                            }
                            s => bail!("unknown key: {s}"),
                        },
                    }
                }
            }
        }

        if key.code == KeyCode::Null {
            bail!("empty key")
        }
        Ok(key)
    }
}

impl Display for Key {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        if self.ctrl {
            write!(f, "C-")?;
        }

        if self.alt {
            write!(f, "A-")?;
        }

        if self.shift {
            write!(f, "S-")?;
        }

        let code = match self.code {
            KeyCode::Backspace => "backspace",
            KeyCode::Enter => "enter",
            KeyCode::Left => "left",
            KeyCode::Right => "right",
            KeyCode::Up => "up",
            KeyCode::Down => "down",
            KeyCode::PageUp => "pageup",
            KeyCode::PageDown => "pagedown",
            KeyCode::Tab => "tab",
            KeyCode::BackTab => "backtab",
            KeyCode::Delete => "delete",
            KeyCode::Insert => "insert",
            KeyCode::Esc => "esc",
            KeyCode::F(num) => {
                write!(f, "F{num}")?;
                ""
            }
            KeyCode::Char(' ') => "space",
            KeyCode::Char(c) => {
                f.write_char(c)?;
                ""
            }
            _ => "unknown",
        };
        write!(f, "{code}")
    }
}
