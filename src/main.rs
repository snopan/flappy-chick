mod components;
mod entities;
mod resources;
mod systems;

use crate::resources::*;
use crate::systems::*;
use bevy::prelude::*;
use bevy::window;

fn main() {
    println!("Hello, world!");
    App::new()
        .init_resource::<Animations>()
        .add_plugins(
            DefaultPlugins
                .set(ImagePlugin::default_nearest())
                .set(WindowPlugin {
                    primary_window: Some(Window {
                        resolution: (300.0, 600.0).into(),
                        title: "Game of Life".to_string(),
                        ..Default::default()
                    }),
                    ..Default::default()
                }),
        )
        .add_systems(Startup, setup)
        .add_systems(Update, border)
        .add_systems(Update, input)
        .add_systems(Update, animation_update)
        .add_systems(Update, animate_sprite)
        .add_systems(Update, gravity)
        .add_systems(Update, velocity)
        .add_systems(Update, flight_rotation)
        .run()
}
