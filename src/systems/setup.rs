use crate::entities::*;
use bevy::prelude::*;

pub fn setup(
    mut commands: Commands,
    asset_server: Res<AssetServer>,
    mut texture_atlas_layouts: ResMut<Assets<TextureAtlasLayout>>
) {

    commands.spawn(Camera2dBundle::default());
    create_player(&mut commands, &asset_server, &mut texture_atlas_layouts);
}