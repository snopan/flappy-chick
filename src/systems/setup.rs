use crate::entities::*;
use crate::resources::*;
use bevy::prelude::*;

pub fn setup(
    mut commands: Commands,
    mut animations: ResMut<Animations>,
    asset_server: Res<AssetServer>,
    mut texture_atlas_layouts: ResMut<Assets<TextureAtlasLayout>>
) {
    commands.spawn(Camera2dBundle::default());

    animations.load(&asset_server, &mut texture_atlas_layouts);
    create_player(&mut commands);
}