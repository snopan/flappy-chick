use bevy::prelude::*;
use bevy::asset::AssetServer;

use crate::entities::create_player;

pub fn setup(
    mut commands: Commands,
    asset_server: Res<AssetServer>,
) {
    create_player(&mut commands, &asset_server);
}