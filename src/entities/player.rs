use crate::components::*;
use bevy::prelude::*;

pub fn create_player(commands: &mut Commands) {
    commands.spawn((
        SpriteSheetBundle {
            transform: Transform::default(),
            ..Default::default()
        },
        Animator {
            ..Default::default()
        },
        AnimationState::UpdateTo(Animation::PlayerFall),
        Velocity { x: 0.0, y: 0.0 },
        Player,
    ));
}
