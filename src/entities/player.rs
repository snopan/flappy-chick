use crate::components::*;
use bevy::prelude::*;

pub fn create_player(
    commands: &mut Commands,
) {
    commands.spawn((
        SpriteSheetBundle {
            transform: Transform{
                translation: Vec3{
                    x: 0.0,
                    y: 150.0,
                    z: 0.0
                },
                ..Default::default()
            },
            ..Default::default()
        },
        Animator { ..Default::default() },
        Animation::PlayerFall,
        UpdateAnimation(true),
        Velocity{
            x: 0.0,
            y: 0.0,
        },
        Player
    ));
}