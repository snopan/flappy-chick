use crate::components::*;
use bevy::prelude::*;

const GRAVITY: f32 = 9.8;
const MASS: f32 = 10.0;

pub fn gravity(
    time: Res<Time>,
    mut query: Query<(&mut Velocity, &mut AnimationState), With<Player>>,
) {
    for (mut velocity, mut animation_state) in &mut query {
        let old_velocity = velocity.y;
        let new_velocity = old_velocity - GRAVITY * MASS * time.delta().as_secs_f32();

        if old_velocity > 0.0 && new_velocity <= 0.0 {
            *animation_state = AnimationState::UpdateTo(Animation::PlayerFall);
        }

        velocity.y = new_velocity;
    }
}
