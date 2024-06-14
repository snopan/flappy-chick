use crate::components::*;
use bevy::prelude::*;

pub fn velocity(
    time: Res<Time>,
    mut query: Query<(&Velocity, &mut Transform)>
) {
    for (velocity, mut transform) in &mut query {
        let delta = time.delta().as_secs_f32();
        transform.translation.x += velocity.x * delta;
        transform.translation.y += velocity.y * delta;
    }
}