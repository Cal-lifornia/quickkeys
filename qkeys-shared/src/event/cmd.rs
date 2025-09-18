use std::any::Any;

use hashbrown::HashMap;

use crate::{
    SStr,
    event::{Data, DataKey},
};

pub struct Cmd {
    pub name: SStr,
    pub args: HashMap<DataKey, Data>,
}

impl Cmd {
    pub fn new<N>(name: N) -> Self
    where
        N: Into<SStr>,
    {
        let cow: SStr = name.into();
        Self {
            name: cow,
            args: Default::default(),
        }
    }

    pub fn with(mut self, name: impl Into<DataKey>, value: impl Into<Data>) -> Self {
        self.args.insert(name.into(), value.into());
        self
    }

    pub fn with_option(mut self, name: impl Into<DataKey>, value: Option<impl Into<Data>>) -> Self {
        if let Some(v) = value {
            self.args.insert(name.into(), v.into());
        }
        self
    }

    pub fn with_any(mut self, name: impl Into<DataKey>, data: impl Any + Send + Sync) -> Self {
        self.args.insert(name.into(), Data::Any(Box::new(data)));
        self
    }

    pub fn get(&self, name: impl Into<DataKey>) -> Option<&Data> {
        self.args.get(&name.into())
    }

    pub fn str(&self, name: impl Into<DataKey>) -> Option<&str> {
        self.get(name)?.as_str()
    }

    pub fn take(&mut self, name: impl Into<DataKey>) -> Option<Data> {
        self.args.remove(&name.into())
    }

    pub fn take_first(&mut self) -> Option<Data> {
        self.take(0)
    }
}
