use crate::entities::*;
use crate::resources::*;
use bevy::prelude::*;

pub fn setup(
    mut commands: Commands,
    mut animations: ResMut<Animations>,
    mut texture_atlas_layouts: ResMut<Assets<TextureAtlasLayout>>,
    asset_server: Res<AssetServer>,
    windows: Query<&Window>,
) {
    let window = windows.single();
    commands.spawn(Camera2dBundle::default());

    animations.load(&asset_server, &mut texture_atlas_layouts);
    create_player(&mut commands);
    create_ground(&mut commands, &mut texture_atlas_layouts, &asset_server);
    create_border(&mut commands, window);
}
