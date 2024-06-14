mod systems;
mod entities;
mod resources;
mod components;

use crate::systems::*;
use crate::resources::*;
use bevy::prelude::*;

fn main() {
    println!("Hello, world!");
    App::new()
        .init_resource::<Animations>()
        .add_plugins(DefaultPlugins.set(ImagePlugin::default_nearest()))
        .add_systems(Startup, setup)
        .add_systems(Update, input)
        .add_systems(Update, animation_update)
        .add_systems(Update, animate_sprite)
        .add_systems(Update, gravity)
        .add_systems(Update, velocity)
        .add_systems(Update, flight_rotation)
        .run()
}
