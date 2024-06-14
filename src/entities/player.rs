use crate::components::*;
use bevy::prelude::*;
use bevy::asset::AssetServer;

pub fn create_player(
    commands: &mut Commands,
    asset_server: &Res<AssetServer>,
    texture_atlas_layouts: &mut ResMut<Assets<TextureAtlasLayout>>
) {
    let texture = asset_server.load("images/chicken_fall.png");
    let layout = TextureAtlasLayout::from_grid(Vec2::new(40.0, 40.0), 2, 1, None, None);
    let texture_atlas_layout = texture_atlas_layouts.add(layout);

    commands.spawn((
        SpriteSheetBundle {
            texture,
            atlas: TextureAtlas {
                layout: texture_atlas_layout,
                index: 0,
            },
            ..Default::default()
        },
        Animator {
            timer: Timer::from_seconds(0.1, TimerMode::Repeating),
            num_frames: 2,
        }
    ));

}