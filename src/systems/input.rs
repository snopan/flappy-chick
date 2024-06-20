use crate::components::*;
use bevy::prelude::*;

pub fn input(
    keyboard_input: Res<ButtonInput<KeyCode>>,
    mut query: Query<(&mut Velocity, &mut AnimationState), With<Player>>,
) {
    for (mut velocity, mut animation_state) in &mut query {
        if keyboard_input.just_pressed(KeyCode::Space) {
            velocity.y = 100.0;
            *animation_state = AnimationState::UpdateTo(Animation::PlayerFly)
        }
    }
}
