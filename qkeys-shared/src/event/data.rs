use std::any::Any;

use bytes::Bytes;
use hashbrown::HashMap;
use serde::{Deserialize, Serialize, de};

use crate::SStr;

#[derive(Debug, Deserialize, Serialize)]
#[serde(untagged)]
pub enum Data {
    Nil,
    Boolean(bool),
    Integer(i64),
    Number(f64),
    String(SStr),
    List(Vec<Data>),
    #[serde(skip)]
    Dict(HashMap<DataKey, Data>),
    #[serde(skip)]
    Bytes(Bytes),
    #[serde(skip)]
    Any(Box<dyn Any + Send + Sync>),
}

impl Data {
    pub fn as_bool(&self) -> Option<bool> {
        match self {
            Self::Boolean(val) => Some(*val),
            Self::String(s) if s == "no" => Some(false),
            Self::String(s) if s == "yes" => Some(true),
            _ => None,
        }
    }

    pub fn as_str(&self) -> Option<&str> {
        if let Self::String(s) = self {
            Some(s)
        } else {
            None
        }
    }

    pub fn into_string(self) -> Option<SStr> {
        if let Self::String(s) = self {
            Some(s)
        } else {
            None
        }
    }

    pub fn into_dict(self) -> Option<HashMap<DataKey, Data>> {
        if let Self::Dict(map) = self {
            Some(map)
        } else {
            None
        }
    }

    pub fn into_any<T: 'static>(self) -> Option<T> {
        if let Self::Any(b) = self {
            b.downcast::<T>().ok().map(|b| *b)
        } else {
            None
        }
    }

    pub fn as_f64(&self) -> Option<f64> {
        match self {
            Self::Integer(i) => Some(*i as f64),
            Self::Number(n) => Some(*n),
            Self::String(s) => s.parse().ok(),
            _ => None,
        }
    }
}

impl From<bool> for Data {
    fn from(value: bool) -> Self {
        Self::Boolean(value)
    }
}

impl From<i32> for Data {
    fn from(value: i32) -> Self {
        Self::Integer(value as i64)
    }
}

impl From<i64> for Data {
    fn from(value: i64) -> Self {
        Self::Integer(value)
    }
}

impl From<f64> for Data {
    fn from(value: f64) -> Self {
        Self::Number(value)
    }
}

impl From<usize> for Data {
    fn from(value: usize) -> Self {
        Self::Integer(value as i64)
    }
}

impl From<SStr> for Data {
    fn from(value: SStr) -> Self {
        Self::String(value)
    }
}

impl From<&str> for Data {
    fn from(value: &str) -> Self {
        Self::String(std::borrow::Cow::Owned(value.to_owned()))
    }
}

impl PartialEq<bool> for Data {
    fn eq(&self, other: &bool) -> bool {
        self.as_bool() == Some(*other)
    }
}

#[derive(Debug, Hash, PartialEq, Eq, Deserialize, Serialize)]
#[serde(untagged)]
pub enum DataKey {
    Nil,
    Boolean(bool),
    #[serde(deserialize_with = "Self::deserialize_integer")]
    Integer(i64),
    String(SStr),
    #[serde(skip)]
    Bytes(Bytes),
}

impl DataKey {
    pub fn as_bool(&self) -> Option<bool> {
        match self {
            Self::Boolean(val) => Some(*val),
            Self::String(s) if s == "no" => Some(false),
            Self::String(s) if s == "yes" => Some(true),
            _ => None,
        }
    }

    pub fn as_str(&self) -> Option<&str> {
        if let Self::String(s) = self {
            Some(s)
        } else {
            None
        }
    }

    pub fn into_string(self) -> Option<SStr> {
        if let Self::String(s) = self {
            Some(s)
        } else {
            None
        }
    }

    pub fn is_integer(&self) -> bool {
        matches!(self, Self::Integer(_))
    }

    fn deserialize_integer<'de, D>(deserializer: D) -> Result<i64, D::Error>
    where
        D: de::Deserializer<'de>,
    {
        struct Visitor;

        impl<'de> de::Visitor<'de> for Visitor {
            type Value = i64;
            fn expecting(&self, formatter: &mut std::fmt::Formatter) -> std::fmt::Result {
                write!(formatter, "an interger or a string of an integer")
            }
            fn visit_i64<E>(self, v: i64) -> Result<Self::Value, E>
            where
                E: de::Error,
            {
                Ok(v)
            }

            fn visit_str<E>(self, v: &str) -> Result<Self::Value, E>
            where
                E: de::Error,
            {
                v.parse().map_err(serde::de::Error::custom)
            }
        }

        deserializer.deserialize_any(Visitor)
    }
}

impl From<usize> for DataKey {
    fn from(value: usize) -> Self {
        Self::Integer(value as i64)
    }
}

impl From<&'static str> for DataKey {
    fn from(value: &'static str) -> Self {
        Self::String(std::borrow::Cow::Borrowed(value))
    }
}

impl From<String> for DataKey {
    fn from(value: String) -> Self {
        Self::String(std::borrow::Cow::Owned(value))
    }
}

macro_rules! impl_as_integer {
    ($t:ty, $name:ident) => {
        impl Data {
            pub fn $name(&self) -> Option<$t> {
                match self {
                    Self::Integer(i) => <$t>::try_from(*i).ok(),
                    Self::String(s) => s.parse().ok(),
                    _ => None,
                }
            }
        }
    };
}

impl_as_integer!(usize, as_usize);
impl_as_integer!(isize, as_isize);
impl_as_integer!(i16, as_i16);
impl_as_integer!(i32, as_i32);
impl_as_integer!(i64, as_i64);
