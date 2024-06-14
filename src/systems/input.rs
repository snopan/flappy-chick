use crate::components::*;
use bevy::prelude::*;

pub fn input(
    keyboard_input: Res<ButtonInput<KeyCode>>,
    mut query: Query<(&mut Animation, &mut UpdateAnimation), With<Player>>
) {
    let (mut animation, mut update_animation) = match query.get_single_mut() {
        Ok(i) => i,
        Err(_) => return
    };

    if keyboard_input.just_pressed(KeyCode::Space) {
        *animation = Animation::PlayerFly;
        update_animation.0 = true;
    }
}