use crate::components::*;
use bevy::prelude::*;

const GRAVITY: f32 = 9.8;
const MASS: f32 = 10.0;

pub fn gravity(
    time: Res<Time>,
    mut query: Query<&mut Velocity, With<Player>>
) {
    for mut velocity in &mut query {
        velocity.y = velocity.y - GRAVITY * MASS * time.delta().as_secs_f32();
    }
}