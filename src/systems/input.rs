use crate::components::*;
use bevy::prelude::*;

pub fn input(
    keyboard_input: Res<ButtonInput<KeyCode>>,
    mut query: Query<(&mut Animation, &mut UpdateAnimation, &mut Velocity), With<Player>>
) {
    for (mut animation, mut update_animation, mut velocity) in &mut query {
        if keyboard_input.just_pressed(KeyCode::Space) {
            *animation = Animation::PlayerFly;
            update_animation.0 = true;
            velocity.y = 100.0
        }
    }
}