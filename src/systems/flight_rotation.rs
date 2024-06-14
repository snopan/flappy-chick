use std::cmp;

use crate::components::*;
use bevy::prelude::*;

const MAX_VELOCITY: f32 = 100.0;
const MAX_ANGLE: f32 = 0.3;

pub fn flight_rotation(
    mut query: Query<(&Velocity, &mut Transform), With<Player>>
) {
    for (velocity, mut transform) in &mut query {
        let mut vertical_velocity = velocity.y;
        if vertical_velocity > MAX_VELOCITY {
            vertical_velocity = MAX_VELOCITY;
        }
        if vertical_velocity < -MAX_VELOCITY {
            vertical_velocity = -MAX_VELOCITY
        }
       
       let angle = MAX_ANGLE/MAX_VELOCITY * vertical_velocity;
       transform.rotation.z = angle;
    }
}