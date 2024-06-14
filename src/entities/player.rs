use crate::components::*;
use bevy::prelude::*;
use bevy::asset::AssetServer;

pub fn create_player(
    commands: &mut Commands,
) {
    commands.spawn((
        SpriteSheetBundle { ..Default::default() },
        Animator { ..Default::default() },
        Animation::PlayerFall,
        UpdateAnimation(true)
    ));
}