mod systems;
mod entities;

use crate::systems::*;
use bevy::prelude::*;

fn main() {
    println!("Hello, world!");
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup)
        .run()
}
